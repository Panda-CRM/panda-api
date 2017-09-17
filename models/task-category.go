package models

import (
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	"github.com/wilsontamarozzi/panda-api/helpers"
)

type TaskCategory struct {
	UUID        string `json:"id,omitempty" sql:"type:uuid; primary_key; default:uuid_generate_v4();unique"`
	Description string `json:"description,omitempty" sql:"type:varchar(25); not null;" valid:"required~Descrição é obrigatório,length(2|25)~Descrição deve ter minimo 2 e maximo 25 caracter"`
}

type TaskCategoryList struct {
	TaskCategories []TaskCategory     `json:"task_categories"`
	Pages          helpers.PageParams `json:"pages"`
}

func (c TaskCategory) IsEmpty() bool {
	return c == TaskCategory{}
}

func (c TaskCategory) Validate() []string {
	var errs []string

	if _, err := govalidator.ValidateStruct(c); err != nil {
		errsV := err.(govalidator.Errors).Errors()
		for _, element := range errsV {
			errs = append(errs, element.Error())
		}
	}

	return errs
}

func (c TaskCategory) PopulateDefault(db *gorm.DB) {
	c.UUID = "756524a2-9555-4ae5-9a6c-b2232de896af"
	c.Description = "Geral"
	db.Create(&c)
}

func (c *TaskCategory) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("uuid", nil)
	return nil
}
