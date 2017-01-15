package database

import (
    "os"
	"github.com/jinzhu/gorm"
    "github.com/wilsontamarozzi/panda-api/services/models"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/asaskevich/govalidator"
)

const(
    ENV_DB_DATABASE = "DB_DATABASE"
    ENV_DB_HOST = "DB_HOST"
    ENV_DB_NAME = "DB_NAME"
    ENV_DB_USER = "DB_USER"
    ENV_DB_PASSWORD = "DB_PASSWORD"
    ENV_DB_SSL_MODE = "DB_SSL_MODE"
    ENV_DB_MAX_CONNECTION = "DB_MAX_CONNECTION"
    ENV_DB_LOG_MODE = "DB_LOG_MODE"
)

var(
    DB_DATABASE string = "postgres"
    DB_HOST string = "localhost"
    DB_NAME string = "panda"
    DB_USER string = "pandaapi"
    DB_PASSWORD string = "1234"
    DB_SSL_MODE string = "disable" // disable | require
    DB_MAX_CONNECTION int = 1
    DB_LOG_MODE bool = true
)

func init() {
    getEnvDatabaseConfig()
}

func getEnvDatabaseConfig() {
    dbDatabase      := os.Getenv(ENV_DB_DATABASE)
    dbHost          := os.Getenv(ENV_DB_HOST)
    dbName          := os.Getenv(ENV_DB_NAME)
    dbUser          := os.Getenv(ENV_DB_USER)
    dbPassword      := os.Getenv(ENV_DB_PASSWORD)
    dbSslMode       := os.Getenv(ENV_DB_SSL_MODE)
    dbMaxConnection := os.Getenv(ENV_DB_MAX_CONNECTION)
    dbLogMode       := os.Getenv(ENV_DB_LOG_MODE)

    maxConnection, err1 := govalidator.ToInt(dbMaxConnection)
    logMode, err2 := govalidator.ToBoolean(dbLogMode)
    
    if len(dbDatabase) > 0  { DB_DATABASE         = dbDatabase          }
    if len(dbHost) > 0      { DB_HOST             = dbHost              }
    if len(dbName) > 0      { DB_NAME             = dbName              }
    if len(dbUser) > 0      { DB_USER             = dbUser              }
    if len(dbPassword) > 0  { DB_PASSWORD         = dbPassword          }
    if len(dbSslMode) > 0   { DB_SSL_MODE         = dbSslMode           }
    if err1 == nil          { DB_MAX_CONNECTION   = int(maxConnection)  }
    if err2 == nil          { DB_LOG_MODE         = logMode             }
}

func GetConnection() *gorm.DB {	
    db, err := gorm.Open(DB_DATABASE, "host=" + DB_HOST + " user=" + DB_USER + " dbname=" + DB_NAME + " sslmode=" + DB_SSL_MODE + " password=" + DB_PASSWORD)

    if err != nil {
        panic(err)
    }

    //Ativa log de todas as saidas da conexão (SQL)
    db.LogMode(DB_LOG_MODE)
    //Seta o maximo de conexões
    db.DB().SetMaxIdleConns(DB_MAX_CONNECTION)
    db.DB().SetMaxOpenConns(DB_MAX_CONNECTION)

    DropTablesIfExists(db)
    AutoMigrate(db)
    AutoPopulate(db)
    AddForeignKeys(db)

 	return db
}

func DropTablesIfExists(db *gorm.DB) {
    db.Exec("DROP TABLE users, people, tasks, task_categories, task_historics CASCADE;")
}

func AutoMigrate(db *gorm.DB) {
    db.AutoMigrate(
        &models.Person{}, 
        &models.User{}, 
        &models.TaskCategory{}, 
        &models.Task{}, 
        &models.TaskHistoric{},
    )
}

func AddForeignKeys(db *gorm.DB) {
    db.Model(&models.Person{}).AddForeignKey("registered_by_uuid", "people(uuid)", "RESTRICT", "RESTRICT")
    db.Model(&models.Users{}).AddForeignKey("person_uuid", "people(uuid)", "RESTRICT", "RESTRICT")
    db.Model(&models.Task{}).AddForeignKey("category_uuid", "task_categories(uuid)", "RESTRICT", "RESTRICT")
    db.Model(&models.Task{}).AddForeignKey("person_uuid", "people(uuid)", "RESTRICT", "RESTRICT")
    db.Model(&models.Task{}).AddForeignKey("assignee_uuid", "people(uuid)", "RESTRICT", "RESTRICT")
    db.Model(&models.Task{}).AddForeignKey("registered_by_uuid", "people(uuid)", "RESTRICT", "RESTRICT")
    db.Model(&models.TaskHistoric{}).AddForeignKey("task_uuid", "tasks(uuid)", "CASCADE", "CASCADE")
    db.Model(&models.TaskHistoric{}).AddForeignKey("registered_by_uuid", "people(uuid)", "RESTRICT", "RESTRICT")
}

func AutoPopulate(db *gorm.DB) {
    PopulatePerson(db)
    PopulateUser(db)
    PopulateTaskCategory(db)
}

func PopulatePerson(db *gorm.DB) {
    db.Create(&models.Person{
        UUID : "ce7405d8-3b78-4de7-8b58-6b32ac913701",
        Code : 0,
        Name : "Admin",
        Type : "F",
        RegisteredByUUID : "ce7405d8-3b78-4de7-8b58-6b32ac913701",
        IsUser : true,
    })
}

func PopulateUser(db *gorm.DB) {
    db.Create(&models.User{
        UUID : "5b4e149d-4799-4192-b95e-aa2b57d99465",
        PersonUUID : "ce7405d8-3b78-4de7-8b58-6b32ac913701",
        Active : true,
        Username : "admin",
        Password : "202cb962ac59075b964b07152d234b70",
    })
}

func PopulateTaskCategory(db *gorm.DB) {
    db.Create(&models.TaskCategory{
        UUID : "756524a2-9555-4ae5-9a6c-b2232de896af",
        Description : "Geral",
    })
}