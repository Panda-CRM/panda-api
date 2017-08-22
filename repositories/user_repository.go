package repositories

import (
	"github.com/wilsontamarozzi/panda-api/database"
	"github.com/wilsontamarozzi/panda-api/models"
)

type UserRepository interface {
	Authentication(username string, password string) models.Person
}

type userRepository struct{}

func NewUserRepository() *userRepository {
	return new(userRepository)
}

func (repository userRepository) Authentication(username string, password string) models.Person {
	db := database.GetInstance()

	var user models.Person
	db.Joins("JOIN users ON users.person_uuid = people.uuid").
		Where("users.username = ? AND users.password = ?", username, password).
		First(&user)

	return user
}
