package database

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/Panda-CRM/panda-api/models"
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
	//RebuildDataBase(true)
}

func getEnvDatabaseConfig() {
	log.Print("[CONFIG] Lendo configurações do banco de dados")
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

	log.Print("[DATABASE] Conexão com o banco realizada com sucesso")
	//Ativa log de todas as saidas da conexão (SQL)
	db.LogMode(DB_LOG_MODE)
	//Seta o maximo de conexões
	db.DB().SetMaxIdleConns(DB_MAX_CONNECTION)
	db.DB().SetMaxOpenConns(DB_MAX_CONNECTION)
	return db
}

func RebuildDataBase(pupulate bool) {
	DropTablesIfExists()
	AutoMigrate()
	if pupulate {
		AutoPopulate()
	}
	//AddForeignKeys()
	//CreateTriggers()
}

func AutoPopulate() {
	db := GetInstance()
	models.Role{}.PopulateDefault(db)
	models.Person{}.PopulateDefault(db)
	models.TaskCategory{}.PopulateDefault(db)
}

func DropTablesIfExists() {
	GetInstance().Exec(`DROP TABLE IF EXISTS
		people,
		tasks,
		task_categories,
		task_historics,
		products,
		sales,
		sale_products,
		role_permissions,
		roles,
	CASCADE;`)
}

func AutoMigrate() {
	db := GetInstance()
	db.AutoMigrate(&models.Person{})
	db.AutoMigrate(&models.TaskCategory{})
	db.AutoMigrate(&models.Task{})
	db.AutoMigrate(&models.TaskHistoric{})
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Sale{})
	db.AutoMigrate(&models.SaleProduct{})
	db.AutoMigrate(&models.RolePermission{})
	db.AutoMigrate(&models.Role{})
}

func AddForeignKeys() {
	con := GetInstance()
	// Person Table
	con.Model(&models.Person{}).AddForeignKey("registered_by_uuid", "people(uuid)", "RESTRICT", "RESTRICT")
	// Task Table
	con.Model(&models.Task{}).AddForeignKey("category_uuid", "task_categories(uuid)", "RESTRICT", "RESTRICT")
	con.Model(&models.Task{}).AddForeignKey("person_uuid", "people(uuid)", "RESTRICT", "RESTRICT")
	con.Model(&models.Task{}).AddForeignKey("assignee_uuid", "people(uuid)", "RESTRICT", "RESTRICT")
	con.Model(&models.Task{}).AddForeignKey("registered_by_uuid", "people(uuid)", "RESTRICT", "RESTRICT")
	// Task Historic Table
	con.Model(&models.TaskHistoric{}).AddForeignKey("task_uuid", "tasks(uuid)", "CASCADE", "CASCADE")
	con.Model(&models.TaskHistoric{}).AddForeignKey("registered_by_uuid", "people(uuid)", "RESTRICT", "RESTRICT")
	// Sale Product Table
	con.Model(&models.SaleProduct{}).AddForeignKey("sale_uuid", "sales(uuid)", "RESTRICT", "RESTRICT")
}

func CreateTriggers() {
	createTriggerSaleProduct()
}

func createTriggerSaleProduct() {
	con := GetInstance()
	con.Exec(`
	CREATE OR REPLACE FUNCTION calculate_product() RETURNS trigger AS
	$$
	BEGIN
		IF NEW.product_value > 0 AND NEW.commission_value > 0 THEN
			NEW.commission_percentage = (NEW.commission_value*100) / NEW.product_value;
		ELSE
			NEW.commission_percentage = 0;
		END IF;
		RETURN NEW;
	END;
	$$
	LANGUAGE plpgsql;`)

	con.Exec(`
	CREATE TRIGGER calculate_product
		BEFORE INSERT OR UPDATE ON sale_products
		FOR EACH ROW
		EXECUTE PROCEDURE calculate_product();`)
}