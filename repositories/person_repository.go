package repositories

import (
	"github.com/wilsontamarozzi/panda-api/database"
	"github.com/wilsontamarozzi/panda-api/helpers"
	"github.com/wilsontamarozzi/panda-api/models"
	"log"
	"net/url"
	"strconv"
)

type PersonRepository interface {
	List(q url.Values) models.PersonList
	Get(id string) models.Person
	GetByCPF(cpf string) models.Person
	GetByIdCVC(idCVC int) models.Person
	Delete(id string) error
	Create(p *models.Person) error
	Update(p *models.Person) error
	CountRows() int
}

type personRepository struct{}

func NewPersonRepository() *personRepository {
	return new(personRepository)
}

func (repository personRepository) List(q url.Values) models.PersonList {
	db := database.GetInstance()

	currentPage, _ := strconv.Atoi(q.Get("page"))
	itemPerPage, _ := strconv.Atoi(q.Get("per_page"))
	pagination := helpers.MakePagination(repository.CountRows(), currentPage, itemPerPage)

	if q.Get("filter") != "" {
		db = db.Where("name iLIKE ?", "%"+q.Get("filter")+"%").
			Or("company_name iLIKE ?", "%"+q.Get("filter")+"%")
	}

	if q.Get("code") != "" {
		db = db.Where("code = ?", q.Get("code"))
	}

	if q.Get("name") != "" {
		db = db.Where("name iLIKE ?", "%"+q.Get("name")+"%")
	}

	if q.Get("company_name") != "" {
		db = db.Where("company_name iLIKE ?", "%"+q.Get("company_name")+"%")
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

	if q.Get("cpf") != "" {
		db = db.Where("cpf = ?", q.Get("cpf"))
	}

	var people models.PersonList
	people.Meta.Pagination = pagination

	db.Limit(pagination.ItemPerPage).
		Offset(pagination.StartIndex).
		Order("registered_at desc").
		Find(&people.People)

	return people
}

func (repository personRepository) Get(id string) models.Person {
	db := database.GetInstance()

	var person models.Person
	db.Where("uuid = ?", id).
		First(&person)

	return person
}

func (repository personRepository) GetByCPF(cpf string) models.Person {
	db := database.GetInstance()

	var person models.Person
	db.Where("cpf = ?", cpf).
		First(&person)

	return person
}

func (repository personRepository) GetByIdCVC(idCVC int) models.Person {
	db := database.GetInstance()

	var person models.Person
	db.Where("id_cvc = ?", idCVC).
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

	err := db.Create(&p).Error
	if err != nil {
		log.Print(err.Error())
	}

	return err
}

func (repository personRepository) Update(p *models.Person) error {
	db := database.GetInstance()

	err := db.Model(&p).
		Omit("type", "code", "uuid", "registered_at", "registered_uuid").
		Updates(&p).Error

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
