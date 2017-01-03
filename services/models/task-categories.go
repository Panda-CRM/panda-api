package models

import (
	"errors"
	"github.com/asaskevich/govalidator"
)

var (
	ErrEmptyDescription = errors.New("Descrição não pode ser vázio")
	ErrLenghtDescription = errors.New("Descrição deve ter mais que 1 caracteres")
)

type TaskCategory struct {
	UUID		string	`json:"id" sql:"type:uuid; primary_key; default:uuid_generate_v4()"`
	Description	string 	`json:"description" sql:"type:varchar(25); not null; unique"`
}

type TaskCategories []TaskCategory

type TaskCategoryRequest struct {
	TaskCategory TaskCategory `json:"task_category"`
}

func (c TaskCategory) Validate() []string {

	var errors []string

	if govalidator.IsNull(c.Description) {
		errors = append(errors, ErrEmptyDescription.Error())
	} else if len(c.Description) < 2 {
		errors = append(errors, ErrLenghtDescription.Error())
	}

	return errors
}