package models

import (
	"github.com/asaskevich/govalidator"
	"github.com/wilsontamarozzi/panda-api/helpers"
)

type TaskCategory struct {
	UUID string 		`json:"id" sql:"type:uuid; primary_key; default:uuid_generate_v4();unique"`
	Description	string 	`json:"description" sql:"type:varchar(25); not null;" valid:"required~Descrição é obrigatório,length(2|25)~Descrição deve ter minimo 2 e maximo 25 caracter"`
}

type TaskCategories struct {
	TaskCategories []TaskCategory
	Meta helpers.Meta
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