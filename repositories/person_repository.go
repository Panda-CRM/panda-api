package repositories

import (
	"github.com/wilsontamarozzi/panda-api/database"
	"github.com/wilsontamarozzi/panda-api/helpers"
	"github.com/wilsontamarozzi/panda-api/models"
	"log"
	"net/url"
)

type PersonRepository interface {
	List(q url.Values) models.PersonList
	Get(id string) models.Person
	GetByCPF(cpf string) models.Person
	GetByIdCVC(idCVC int) models.Person
	Delete(id string) error
	Create(*models.Person) error
	Update(*models.Person) error
	CountRows() int
}

type personRepository struct{}

func NewPersonRepository() *personRepository {
	return new(personRepository)
}

func (r personRepository) List(q url.Values) models.PersonList {
	db := database.GetInstance()
	db = db.WhereWithoutNull("name iLIKE ? OR company_name iLIKE ?", "%"+q.Get("filter")+"%", "%"+q.Get("filter")+"%")
	db = db.WhereWithoutNull("code = ?", q.Get("code"))
	db = db.WhereWithoutNull("name iLIKE ?", "%"+q.Get("name")+"%")
	db = db.WhereWithoutNull("company_name iLIKE ?", "%"+q.Get("company_name")+"%")
	db = db.WhereWithoutNull("gender = ?", q.Get("gender"))
	db = db.WhereWithoutNull("type = ?", q.Get("type"))
	db = db.WhereWithoutNull("is_user = ?", q.Get("only_users"))
	db = db.WhereWithoutNull("cpf = ?", q.Get("cpf"))

	pageParams := helpers.MakePagination(r.CountRows(), q.Get("page"), q.Get("per_page"))
	var people models.PersonList
	people.Pages = pageParams

	db.Debug().Limit(pageParams.ItemPerPage).
		Offset(pageParams.StartIndex).
		Order("registered_at desc").
		Find(&people.People)

	return people
}

func (r personRepository) Get(id string) models.Person {
	db := database.GetInstance()
	var person models.Person
	db.Where("uuid = ?", id).
		First(&person)

	return person
}

func (r personRepository) GetByCPF(cpf string) models.Person {
	db := database.GetInstance()
	var person models.Person
	db.Where("cpf = ?", cpf).
		First(&person)

	return person
}

func (r personRepository) GetByIdCVC(idCVC int) models.Person {
	db := database.GetInstance()
	var person models.Person
	db.Where("id_cvc = ?", idCVC).
		First(&person)

	return person
}

func (r personRepository) Delete(id string) error {
	db := database.GetInstance()
	err := db.Where("uuid = ?", id).Delete(&models.Person{}).Error
	if err != nil {
		log.Print(err.Error())
	}

	return err
}

func (r personRepository) Create(p *models.Person) error {
	db := database.GetInstance()
	err := db.Create(&p).Error
	if err != nil {
		log.Print(err.Error())
	}

	return err
}

func (r personRepository) Update(p *models.Person) error {
	db := database.GetInstance()
	err := db.Model(&p).
		Omit("type", "code", "uuid", "registered_at", "registered_uuid").
		Updates(&p).Error

	if err != nil {
		log.Print(err.Error())
	}

	return err
}

func (r personRepository) CountRows() int {
	db := database.GetInstance()
	var count int
	db.Model(&models.Person{}).Count(&count)

	return count
}
