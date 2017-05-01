package repositories

import (
	"github.com/wilsontamarozzi/panda-api/database"
	"github.com/wilsontamarozzi/panda-api/models"
	"github.com/wilsontamarozzi/panda-api/helpers"
	"strconv"
	"time"
	"net/url"
	"log"
)

type PersonRepositoryInterface interface{
	GetAll(q url.Values) models.People
	Get(id string) models.Person
	Delete(id string) error
	Create(p *models.Person) error
	Update(p *models.Person) error
	CountRows() int
}

type personRepository struct{}

func NewPersonRepository() *personRepository {
	return new(personRepository)
}

func (repository personRepository) GetAll(q url.Values) models.People {
	db := database.GetInstance()

	currentPage, _ := strconv.Atoi(q.Get("page"))
	itemPerPage, _ := strconv.Atoi(q.Get("per_page"))
	pagination := helpers.MakePagination(repository.CountRows(), currentPage, itemPerPage)

	if q.Get("filter") != "" {
		db = db.Where("name iLIKE ?", "%" + q.Get("filter") + "%").
			Or("company_name iLIKE ?", "%" + q.Get("filter") + "%")
	}

	if q.Get("code") != "" {
		db = db.Where("code = ?", q.Get("code"))
	}

	if q.Get("name") != "" {
		db = db.Where("name iLIKE ?", "%" + q.Get("name") + "%")
	}

	if q.Get("company_name") != "" {
		db = db.Where("company_name iLIKE ?", "%" + q.Get("company_name") + "%")
	}

	if q.Get("gender") != "" {
		db = db.Where("gender = ?", q.Get("gender"))
	}

	if q.Get("type") != "" {
		db = db.Where("type = ?", q.Get("type"))
	}

	if q.Get("only_users") != "" {
		db = db.Where("is_user = ?", q.Get("only_users"))
	}

	var people models.People
	people.Meta.Pagination = pagination

	db.Limit(pagination.ItemPerPage).
		Offset(pagination.StartIndex).
		Order("registered_at desc").
		Find(&people)

	return people
}

func (repository personRepository) Get(id string) models.Person {
	db := database.GetInstance()

	var person models.Person
	db.Where("uuid = ?", id).
		First(&person)

	return person
}

func (repository personRepository) Delete(id string) error {
	db := database.GetInstance()

	err := db.Where("uuid = ?", id).Delete(&models.Person{}).Error
	if err != nil {
		log.Print(err.Error())
	}

	return err
}

func (repository personRepository) Create(p *models.Person) error {
	db := database.GetInstance()

	record := models.Person{
		Type 				: p.Type,
		Name 				: p.Name,
		CityName 			: p.CityName,
		CompanyName 		: p.CompanyName,
		Address 			: p.Address,
		Number 				: p.Number,
		Complement 			: p.Complement,
		District 			: p.District,
		Zip 				: p.Zip,
		BirthDate 			: p.BirthDate,
		Cpf 				: p.Cpf,
		Rg 					: p.Rg,
		Gender 				: p.Gender,
		BusinessPhone 		: p.BusinessPhone,
		HomePhone 			: p.HomePhone,
		MobilePhone 		: p.MobilePhone,
		Cnpj 				: p.Cnpj,
		StateInscription 	: p.StateInscription,
		Phone 				: p.Phone,
		Fax 				: p.Fax,
		Email 				: p.Email,
		Website 			: p.Website,
		Observations 		: p.Observations,
		RegisteredAt 		: time.Now(),
		RegisteredByUUID	: p.RegisteredByUUID,
	}

	err := db.Create(&record).Error
	if err != nil {
		log.Print(err.Error())
	}

	*(p) = record
	return err
}

func (repository personRepository) Update(p *models.Person) error {
	db := database.GetInstance()

	record := models.Person{
		Name 				: p.Name,
		CityName 			: p.CityName,
		CompanyName 		: p.CompanyName,
		Address 			: p.Address,
		Number 				: p.Number,
		Complement 			: p.Complement,
		District 			: p.District,
		Zip 				: p.Zip,
		BirthDate 			: p.BirthDate,
		Cpf 				: p.Cpf,
		Rg 					: p.Rg,
		Gender 				: p.Gender,
		BusinessPhone 		: p.BusinessPhone,
		HomePhone 			: p.HomePhone,
		MobilePhone 		: p.MobilePhone,
		Cnpj 				: p.Cnpj,
		StateInscription 	: p.StateInscription,
		Phone 				: p.Phone,
		Fax 				: p.Fax,
		Email 				: p.Email,
		Website 			: p.Website,
		Observations 		: p.Observations,
	}

	err := db.Model(&models.Person{}).
		Where("uuid = ?", p.UUID).
		Updates(&record).Error

	if err != nil {
		log.Print(err.Error())
	}

	return err
}

func (repository personRepository) CountRows() int {
	db := database.GetInstance()
	var count int
	db.Model(&models.Person{}).Count(&count)

	return count
}