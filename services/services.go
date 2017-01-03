package services

import (
	"github.com/jinzhu/gorm"
	"panda-api/database"
)

var Con *gorm.DB

func init() {
	Con = database.GetConnection()
}