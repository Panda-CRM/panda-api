package services

import (
	"panda-api/services/models"
)

func AuthenticationUser(username string, password string) models.Person {

	var user models.Person

	Con.Joins("JOIN users ON users.person_uuid = people.uuid").
		Where("users.username = ? AND users.password = ?", username, password).
		First(&user)

	return user
}