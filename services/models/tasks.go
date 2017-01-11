package models

import (
	"strings"
	"time"
	"errors"
	"github.com/asaskevich/govalidator"
)

var (
	ErrEmptyCategory = errors.New("ID da categoria não pode ser vázio")
	ErrInvalidCategory = errors.New("ID da categoria inválida")
	ErrInvalidPerson = errors.New("ID da pessoa inválida")
	ErrEmptyAssignee = errors.New("ID da responsável não pode ser vázio")
	ErrInvalidAssignee = errors.New("ID da responsável inválida")
)

type Task struct {
	UUID 				string			`json:"id,omitempty" sql:"type:uuid; primary_key; default:uuid_generate_v4();unique"`
	Code 				int 			`json:"code" sql:"auto_increment; primary_key"`
	Title				string 			`json:"title" sql:"type:varchar(100); not null" valid:"length(2|100)~Título deve ter minimo 2 e maximo 100 caracter"`
	Due 				time.Time  		`json:"due" sql:"type:timestamp without time zone; default:NOW()"`
	Visualized 			bool 			`json:"visualized" sql:"boolean"`
	CompletedAt			*time.Time 		`json:"completed_at" sql:"type:timestamp without time zone; default:null"`
	RegisteredAt 		time.Time 		`json:"registered_at" sql:"type:timestamp without time zone; default:NOW()"`

	CategoryUUID 		string 			`json:"-" sql:"type:uuid; not null"`
	Category 			TaskCategory	`json:"category" gorm:"many2many:category_uuid;"`

	RegisteredByUUID	string 			`json:"-" sql:"type:uuid; not null"`
	RegisteredBy		Person			`json:"registered_by" gorm:"many2many:registered_by_uuid;"`

	PersonUUID			string 			`json:"-" sql:"type:uuid; default:uuid_nil()"`
	Person 				Person			`json:"person" gorm:"many2many:person_uuid;"`

	AssigneeUUID		string 			`json:"-" sql:"type:uuid; not null"`
	Assignee 			Person			`json:"assignee" gorm:"many2many:assignee_uuid;"`

	TaskHistorics 		TaskHistorics 	`json:"task_historics" gorm:"ForeignKey:task_uuid;"`
}

type Tasks []Task

type TaskHistoric struct {
	UUID 				string 			`json:"id,omitempty" sql:"type:uuid; primary_key; default:uuid_generate_v4();unique"`
	TaskUUID 			string 			`json:"-" sql:"type:uuid; not null"`
	Comment 			string 			`json:"comment" sql:"type:text"`
	RegisteredAt 		time.Time 		`json:"registered_at" sql:"type:timestamp without time zone; default:NOW()"`

	RegisteredByUUID	string 			`json:"-" sql:"type:uuid; not null"`
	RegisteredBy		Person			`json:"registered_by" gorm:"many2many:registered_by_uuid;"`
}

type TaskHistorics []TaskHistoric


type TaskRequest struct {
	Task Task `json:"task"`
}

func (t Task) Validate() []string {

	var errors []string

	if govalidator.IsNull(t.Category.UUID) {
		errors = append(errors, ErrEmptyCategory.Error())
	} else if !govalidator.IsUUIDv4(t.Category.UUID) {
		errors = append(errors, ErrInvalidCategory.Error())
	}

	if govalidator.IsNull(t.Assignee.UUID) {
		errors = append(errors, ErrEmptyAssignee.Error())
	} else if !govalidator.IsUUIDv4(t.Assignee.UUID) {
		errors = append(errors, ErrInvalidAssignee.Error())
	}

	if !govalidator.IsNull(t.Person.UUID) && !govalidator.IsUUIDv4(t.Person.UUID) {
		errors = append(errors, ErrInvalidPerson.Error())
	}

	// Valida a estrutura pelas tags
	if _, err := govalidator.ValidateStruct(t); err != nil {	
		for _, element := range strings.Split(err.Error(), ";") {
			errors = append(errors, element)
		}
	}

	return errors
}