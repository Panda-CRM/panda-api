package services

import (
	"github.com/jinzhu/gorm"
	"github.com/wilsontamarozzi/panda-api/database"
)

var Con *gorm.DB

func init() {
	Con = database.GetConnection()
}