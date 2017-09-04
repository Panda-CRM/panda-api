package models

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	"github.com/wilsontamarozzi/panda-api/helpers"
	"time"
)

var (
	ErrEmptyCategory   = errors.New("ID da categoria não pode ser vázio")
	ErrInvalidCategory = errors.New("ID da categoria inválida")
	ErrEmptyPerson     = errors.New("ID da pessoa não pode ser vázio")
	ErrInvalidPerson   = errors.New("ID da pessoa inválida")
	ErrEmptyAssignee   = errors.New("ID da responsável não pode ser vázio")
	ErrInvalidAssignee = errors.New("ID da responsável inválida")
)

type Task struct {
	UUID             string        `json:"id,omitempty" sql:"type:uuid; primary_key; default:uuid_generate_v4(); unique"`
	Code             int           `json:"code,omitempty" sql:"auto_increment; primary_key; unique"`
	Title            string        `json:"title,omitempty" sql:"type:varchar(100); not null" valid:"length(2|100)~Título deve ter minimo 2 e maximo 100 caracter"`
	Due              time.Time     `json:"due,omitempty" sql:"type:timestamp without time zone; default:NOW()"`
	Visualized       bool          `json:"visualized" sql:"boolean"`
	CompletedAt      *time.Time    `json:"completed_at,omitempty" sql:"type:timestamp without time zone; default:null"`
	RegisteredAt     *time.Time    `json:"registered_at,omitempty" sql:"type:timestamp without time zone; default:NOW()"`
	CategoryUUID     string        `json:"-" sql:"type:uuid; not null"`
	Category         TaskCategory  `json:"category"`
	RegisteredByUUID string        `json:"-" sql:"type:uuid; not null"`
	RegisteredBy     Person        `json:"registered_by"`
	PersonUUID       string        `json:"-" sql:"type:uuid; not null"`
	Person           Person        `json:"person"`
	AssigneeUUID     string        `json:"-" sql:"type:uuid; not null"`
	Assignee         Person        `json:"assignee"`
	TaskHistorics    TaskHistorics `json:"task_historics"`
}

type TaskList struct {
	Tasks []Task       `json:"tasks"`
	Pages   helpers.PageParams `json:"pages"`
}

func (t Task) IsEmpty() bool {
	return t.UUID == ""
}

func (t Task) Validate() []string {
	var errs []string

	if govalidator.IsNull(t.Category.UUID) {
		errs = append(errs, ErrEmptyCategory.Error())
	} else if !govalidator.IsUUIDv4(t.Category.UUID) {
		errs = append(errs, ErrInvalidCategory.Error())
	}

	if govalidator.IsNull(t.Assignee.UUID) {
		errs = append(errs, ErrEmptyAssignee.Error())
	} else if !govalidator.IsUUIDv4(t.Assignee.UUID) {
		errs = append(errs, ErrInvalidAssignee.Error())
	}

	if govalidator.IsNull(t.Person.UUID) {
		errs = append(errs, ErrEmptyPerson.Error())
	} else if !govalidator.IsUUIDv4(t.Person.UUID) {
		errs = append(errs, ErrInvalidPerson.Error())
	}

	if _, err := govalidator.ValidateStruct(t); err != nil {
		errsV := err.(govalidator.Errors).Errors()
		for _, element := range errsV {
			errs = append(errs, element.Error())
		}
	}

	return errs
}

func (t *Task) BeforeCreate(scope *gorm.Scope) error {
	dateTime := time.Now()
	scope.SetColumn("uuid", nil)
	scope.SetColumn("code", nil)
	scope.SetColumn("visualized", false)
	scope.SetColumn("registered_at", dateTime)
	scope.SetColumn("category_uuid", t.Category.UUID)
	scope.SetColumn("person_uuid", t.Person.UUID)
	scope.SetColumn("assignee_uuid", t.Assignee.UUID)
	return nil
}

func (t *Task) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("category_uuid", t.Category.UUID)
	scope.SetColumn("person_uuid", t.Person.UUID)
	scope.SetColumn("assignee_uuid", t.Assignee.UUID)
	return nil
}

type TaskHistoric struct {
	UUID             string     `json:"id,omitempty" sql:"type:uuid; primary_key; default:uuid_generate_v4();unique"`
	TaskUUID         string     `json:"-" sql:"type:uuid; not null"`
	Comment          string     `json:"comment,omitempty" sql:"type:text"`
	RegisteredAt     *time.Time `json:"registered_at,omitempty" sql:"type:timestamp without time zone; default:NOW()"`
	RegisteredByUUID string     `json:"-" sql:"type:uuid; not null"`
	RegisteredBy     Person     `json:"registered_by"`
}

type TaskHistorics []TaskHistoric

func (th *TaskHistoric) BeforeCreate(scope *gorm.Scope) error {
	dateTime := time.Now()
	scope.SetColumn("uuid", nil)
	scope.SetColumn("registered_at", dateTime)
	scope.SetColumn("code", nil)
	return nil
}

type ReportGeneral struct {
	Total        int `json:"total"`
	Completed    int `json:"completed"`
	NotCompleted int `json:"not_completed"`
	Overdue      int `json:"overdue"`
}

type ReportCategory struct {
	CategoryUUID string `json:"category_uuid"`
	Description  string `json:"description"`
	ReportGeneral
}

type ReportCategories []ReportCategory

type ReportAssignee struct {
	AssigneeUUID string `json:"assignee_uuid"`
	Name         string `json:"name"`
	ReportGeneral
}

type ReportAssignees []ReportAssignee

type ReportAssigneeAndCategory struct {
	AssigneeUUID string           `json:"assignee_uuid"`
	Name         string           `json:"name"`
	Categories   []ReportCategory `json:"category"`
}

type ReportAssigneesAndCategory []ReportAssigneeAndCategory
