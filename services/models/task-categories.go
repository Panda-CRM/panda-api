package models

import (
	"strings"
	"github.com/asaskevich/govalidator"
)

type TaskCategory struct {
	UUID		string	`json:"id" sql:"type:uuid; primary_key; default:uuid_generate_v4()"`
	Description	string 	`json:"description" sql:"type:varchar(25); not null; unique" valid:"required~Descrição é obrigatório,length(2|25)~Descrição deve ter minimo 2 e maximo 25 caracter"`
}

type TaskCategories []TaskCategory

type TaskCategoryRequest struct {
	TaskCategory TaskCategory `json:"task_category"`
}

func (c TaskCategory) Validate() []string {

	var errors []string

	// Valida a estrutura pelas tags
	if _, err := govalidator.ValidateStruct(c); err != nil {	
		for _, element := range strings.Split(err.Error(), ";") {
			errors = append(errors, element)
		}
	}

	return errors
}