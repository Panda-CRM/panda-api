package database

import (
    "os"
	"github.com/jinzhu/gorm"
    "github.com/wilsontamarozzi/panda-api/services/models"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/asaskevich/govalidator"
)

/*
const DB_DATABASE = "postgres"
const DB_HOST = "ec2-54-235-173-161.compute-1.amazonaws.com"
const DB_NAME = "d380btbdjq6o8q"
const DB_USER = "bocfuxgbikaxkq"
const DB_PASSWORD = "8e913f6d5081b277484f6a5739f99b7ab2d4f38086a43715c27ad9bfb77b0731"
const DB_SSL_MODE = "require"
const DB_MAX_CONNECTION = 1
const DB_LOG_MODE = true
*/

const DB_DATABASE = "postgres"
const DB_HOST = "localhost"
const DB_NAME = "panda"
const DB_USER = "pandaapi"
const DB_PASSWORD = "1234"
const DB_SSL_MODE = "disable"
const DB_MAX_CONNECTION = 1
const DB_LOG_MODE = false


func GetENVLogMode() bool {
    env := os.Getenv("DB_LOG_MOD")

    logMode, err := govalidator.ToBoolean(env);

    if err != nil {
        return DB_LOG_MODE
    }
    
    return logMode
}

func GetConnection() *gorm.DB {	
    db, err := gorm.Open(DB_DATABASE, "host=" + DB_HOST + " user=" + DB_USER + " dbname=" + DB_NAME + " sslmode=" + DB_SSL_MODE + " password=" + DB_PASSWORD)

    if err != nil {
        panic(err)
    }

    //Ativa log de todas as saidas da conexão (SQL)
    db.LogMode(GetENVLogMode())
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