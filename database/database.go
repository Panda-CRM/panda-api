package database

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/wilsontamarozzi/panda-api/models"
	"log"
	"os"
	"sync"
)

const (
	ENV_DB_DRIVER         = "DB_DRIVER"
	ENV_DB_HOST           = "DB_HOST"
	ENV_DB_NAME           = "DB_NAME"
	ENV_DB_USER           = "DB_USER"
	ENV_DB_PASSWORD       = "DB_PASSWORD"
	ENV_DB_SSL_MODE       = "DB_SSL_MODE"
	ENV_DB_MAX_CONNECTION = "DB_MAX_CONNECTION"
	ENV_DB_LOG_MODE       = "DB_LOG_MODE"
)

var (
	DB_DRIVER         string = "postgres"
	DB_HOST           string = "localhost"
	DB_NAME           string = "panda"
	DB_USER           string = "pandaapi"
	DB_PASSWORD       string = "1234"
	DB_SSL_MODE       string = "disable" // disable | require
	DB_MAX_CONNECTION int    = 1
	DB_LOG_MODE       bool   = false
)

var ONCE sync.Once
var DBSession *gorm.DB

func init() {
	getEnvDatabaseConfig()
	GetInstance()
	//RebuildDataBase()
}

func getEnvDatabaseConfig() {
	dbDriver := os.Getenv(ENV_DB_DRIVER)
	dbHost := os.Getenv(ENV_DB_HOST)
	dbName := os.Getenv(ENV_DB_NAME)
	dbUser := os.Getenv(ENV_DB_USER)
	dbPassword := os.Getenv(ENV_DB_PASSWORD)
	dbSslMode := os.Getenv(ENV_DB_SSL_MODE)
	dbMaxConnection := os.Getenv(ENV_DB_MAX_CONNECTION)
	dbLogMode := os.Getenv(ENV_DB_LOG_MODE)

	maxConnection, err1 := govalidator.ToInt(dbMaxConnection)
	logMode, err2 := govalidator.ToBoolean(dbLogMode)

	if len(dbDriver) > 0 {
		DB_DRIVER = dbDriver
	}
	if len(dbHost) > 0 {
		DB_HOST = dbHost
	}
	if len(dbName) > 0 {
		DB_NAME = dbName
	}
	if len(dbUser) > 0 {
		DB_USER = dbUser
	}
	if len(dbPassword) > 0 {
		DB_PASSWORD = dbPassword
	}
	if len(dbSslMode) > 0 {
		DB_SSL_MODE = dbSslMode
	}
	if err1 == nil {
		DB_MAX_CONNECTION = int(maxConnection)
	}
	if err2 == nil {
		DB_LOG_MODE = logMode
	}
}

func GetInstance() *gorm.DB {
	ONCE.Do(func() {
		DBSession = buildConnection()
	})
	return DBSession
}

func buildConnection() *gorm.DB {
	strConnection := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s", DB_HOST, DB_USER, DB_NAME, DB_SSL_MODE, DB_PASSWORD)
	db, err := gorm.Open(DB_DRIVER, strConnection)
	if err != nil {
		panic(err)
	}

	log.Print("Conexão com o banco realizada com sucesso")
	//Ativa log de todas as saidas da conexão (SQL)
	db.LogMode(DB_LOG_MODE)
	//Seta o maximo de conexões
	db.DB().SetMaxIdleConns(DB_MAX_CONNECTION)
	db.DB().SetMaxOpenConns(DB_MAX_CONNECTION)

	return db
}

func RebuildDataBase() {
	DropTablesIfExists()
	AutoMigrate()
	AutoPopulate()
	AddForeignKeys()
}

func AutoPopulate() {
	PopulatePerson()
	PopulateUser()
	PopulateTaskCategory()
}

func DropTablesIfExists() {
	GetInstance().Exec("DROP TABLE IF EXISTS users, people, tasks, task_categories, task_historics, products, sales, sale_products CASCADE;")
}

func TruncateTables() {
	con := GetInstance()
	con.Exec("TRUNCATE TABLE products CASCADE;")
	con.Exec("TRUNCATE TABLE tasks CASCADE;")
	con.Exec("TRUNCATE TABLE sales CASCADE;")
}

func AutoMigrate() {
	GetInstance().AutoMigrate(
		&models.Person{},
		&models.User{},
		&models.TaskCategory{},
		&models.Task{},
		&models.TaskHistoric{},
		&models.Product{},
		&models.Sale{},
		&models.SaleProduct{},
	)
}

func AddForeignKeys() {
	con := GetInstance()
	/* Person Table */
	con.Model(&models.Person{}).AddForeignKey("registered_by_uuid", "people(uuid)", "RESTRICT", "RESTRICT")
	/* User Table */
	con.Model(&models.Users{}).AddForeignKey("person_uuid", "people(uuid)", "RESTRICT", "RESTRICT")
	/* Task Table */
	con.Model(&models.Task{}).AddForeignKey("category_uuid", "task_categories(uuid)", "RESTRICT", "RESTRICT")
	con.Model(&models.Task{}).AddForeignKey("person_uuid", "people(uuid)", "RESTRICT", "RESTRICT")
	con.Model(&models.Task{}).AddForeignKey("assignee_uuid", "people(uuid)", "RESTRICT", "RESTRICT")
	con.Model(&models.Task{}).AddForeignKey("registered_by_uuid", "people(uuid)", "RESTRICT", "RESTRICT")
	/* Task Historic Table */
	con.Model(&models.TaskHistoric{}).AddForeignKey("task_uuid", "tasks(uuid)", "CASCADE", "CASCADE")
	con.Model(&models.TaskHistoric{}).AddForeignKey("registered_by_uuid", "people(uuid)", "RESTRICT", "RESTRICT")
	/* Sale Table */
	//con.Model(&models.Sale{}).AddForeignKey("buyer_uuid", "people(uuid)", "RESTRICT", "RESTRICT")
	//con.Model(&models.Sale{}).AddForeignKey("seller_uuid", "people(uuid)", "RESTRICT", "RESTRICT")
	/* Sale Product Table */
	con.Model(&models.SaleProduct{}).AddForeignKey("sale_uuid", "sales(uuid)", "RESTRICT", "RESTRICT")
}

func PopulatePerson() {
	GetInstance().Create(&models.Person{
		UUID:             "ce7405d8-3b78-4de7-8b58-6b32ac913701",
		Code:             0,
		Name:             "Admin",
		Type:             "F",
		RegisteredByUUID: "ce7405d8-3b78-4de7-8b58-6b32ac913701",
		IsUser:           true,
	})
}

func PopulateUser() {
	GetInstance().Create(&models.User{
		UUID:       "5b4e149d-4799-4192-b95e-aa2b57d99465",
		PersonUUID: "ce7405d8-3b78-4de7-8b58-6b32ac913701",
		Active:     true,
		Username:   "admin",
		Password:   "202cb962ac59075b964b07152d234b70",
	})
}

func PopulateTaskCategory() {
	GetInstance().Create(&models.TaskCategory{
		UUID:        "756524a2-9555-4ae5-9a6c-b2232de896af",
		Description: "Geral",
	})
}
