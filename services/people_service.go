package services

import (
	"time"
	"panda-api/services/models"
	"panda-api/helpers"
)

func GetPeople(pag helpers.Pagination, code string, name string) models.People {

	var people models.People

	db := Con

	if code != "" {
		db = db.Where("code = ?", code)
	}

	if name != "" {
		db = db.Where("name iLIKE ?", "%" + name + "%")	
	}
	
	db.Limit(pag.ItemPerPage).
		Offset(pag.StartIndex).
		Order("registered_at desc").
		Find(&people)

    return people
}

func GetPerson(personId string) models.Person {

	var person models.Person

	Con.Where("uuid = ?", personId).
		First(&person)

	return person
}

func DeletePerson(personId string) error {
	return Con.Where("uuid = ?", personId).Delete(&models.Person{}).Error
}

func CreatePerson(person models.Person) error {
	
	record := models.Person{
		Type 				: person.Type,
		Name 				: person.Name,
		CityName 			: person.CityName,
		CompanyName 		: person.CompanyName,
		Address 			: person.Address,
		Number 				: person.Number,
		Complement 			: person.Complement,
		District 			: person.District,
		Zip 				: person.Zip,
		BirthDate 			: person.BirthDate,
		Cpf 				: person.Cpf,
		Rg 					: person.Rg,
		Gender 				: person.Gender,
		BusinessPhone 		: person.BusinessPhone,
		HomePhone 			: person.HomePhone,
		MobilePhone 		: person.MobilePhone,
		Cnpj 				: person.Cnpj,
		StateInscription 	: person.StateInscription,
		Phone 				: person.Phone,
		Fax 				: person.Fax,
		Email 				: person.Email,
		Website 			: person.Website,
		Observations 		: person.Observations,
		RegisteredAt 		: time.Now(),
		RegisteredByUUID	: person.RegisteredByUUID,
	}

	return Con.Set("gorm:save_associations", false).
		Table("people").
		Create(&record).Error
}

func UpdatePerson(person models.Person) error {
	return Con.Set("gorm:save_associations", false).
		Table("people").
		Where("uuid = ?", person.UUID).
		Updates(models.Person{
			Name 				: person.Name,
			CityName 			: person.CityName,
			CompanyName 		: person.CompanyName,
			Address 			: person.Address,
			Number 				: person.Number,
			Complement 			: person.Complement,
			District 			: person.District,
			Zip 				: person.Zip,
			BirthDate 			: person.BirthDate,
			Cpf 				: person.Cpf,
			Rg 					: person.Rg,
			Gender 				: person.Gender,
			BusinessPhone 		: person.BusinessPhone,
			HomePhone 			: person.HomePhone,
			MobilePhone 		: person.MobilePhone,
			Cnpj 				: person.Cnpj,
			StateInscription 	: person.StateInscription,
			Phone 				: person.Phone,
			Fax 				: person.Fax,
			Email 				: person.Email,
			Website 			: person.Website,
			Observations 		: person.Observations,
		}).Error
}

func CountRowsPerson() int {
	var count int
	Con.Table("people").Count(&count)

	return count
}