package database

import (
    "os"
	"github.com/jinzhu/gorm"
    "panda-api/services/models"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/asaskevich/govalidator"
)

const DB_DATABASE = "postgres"
const DB_HOST = "ec2-54-235-173-161.compute-1.amazonaws.com"
const DB_NAME = "d380btbdjq6o8q"
const DB_USER = "bocfuxgbikaxkq"
const DB_PASSWORD = "8e913f6d5081b277484f6a5739f99b7ab2d4f38086a43715c27ad9bfb77b0731"
const DB_MAX_CONNECTION = 1
const DB_SSL_MODE = "require"
const DB_LOG_MODE = true

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
    db.Model(&models.Task{}).AddForeignKey("category_uuid", "task_categories(uuid)", "RESTRICT", "RESTRICT")
    db.Model(&models.Task{}).AddForeignKey("person_uuid", "people(uuid)", "RESTRICT", "RESTRICT")
    db.Model(&models.Task{}).AddForeignKey("assignee_uuid", "people(uuid)", "RESTRICT", "RESTRICT")
    db.Model(&models.Task{}).AddForeignKey("registered_by_uuid", "people(uuid)", "RESTRICT", "RESTRICT")
    db.Model(&models.TaskHistoric{}).AddForeignKey("task_uuid", "tasks(uuid)", "CASCADE", "CASCADE")
}

func AutoPopulate(db *gorm.DB) {
    PopulatePerson(db)
    PopulateUser(db)
    PopulateTaskCategory(db)
}

func PopulatePerson(db *gorm.DB) {
    var person models.Person

    person.UUID = "ce7405d8-3b78-4de7-8b58-6b32ac913701"
    person.Code = 0
    person.Name = "Admin"
    person.Type = "F"
    person.RegisteredByUUID = "ce7405d8-3b78-4de7-8b58-6b32ac913701"
    person.IsUser = true

    db.Set("gorm:save_associations", false).Create(&person)
}

func PopulateUser(db *gorm.DB) {
    var user models.User

    user.UUID = "5b4e149d-4799-4192-b95e-aa2b57d99465"
    user.PersonUUID = "ce7405d8-3b78-4de7-8b58-6b32ac913701"
    user.Active = true
    user.Username = "admin"
    user.Password = "202cb962ac59075b964b07152d234b70"

    db.Set("gorm:save_associations", false).Create(&user)
}

func PopulateTaskCategory(db *gorm.DB) {
    var taskCategory models.TaskCategory

    taskCategory.Description = "Geral"

    db.Set("gorm:save_associations", false).Create(&taskCategory)
}