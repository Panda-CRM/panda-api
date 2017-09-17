package repositories

import (
	"github.com/wilsontamarozzi/panda-api/database"
	"github.com/wilsontamarozzi/panda-api/models"
)

type UserRepository interface {
	Authenticator(username string, password string) models.Person
	Get(id string) models.Person
}

type userRepository struct{}

func NewUserRepository() *userRepository {
	return new(userRepository)
}

func (r userRepository) Authenticator(username string, password string) models.Person {
	db := database.GetInstance()
	var user models.Person
	db.Where("username = ? AND password = ?", username, password).
		Preload("Role.Permissions").First(&user)
	return user
}

func (r userRepository) Get(id string) models.Person {
	var user models.Person
	user.UUID = id
	return r.find(user)
}

func (r userRepository) find(u models.Person) models.Person {
	db := database.GetInstance()
	var person models.Person
	switch {
	case u.UUID != "":
		db.WhereWithoutNull("uuid = ?", u.UUID)
	}
	db.Preload("Role.Permissions").First(&person)
	return person
}
