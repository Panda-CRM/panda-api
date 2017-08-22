package models

import (
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	"github.com/wilsontamarozzi/panda-api/helpers"
)

type Product struct {
	UUID        string `json:"id" sql:"type:uuid; primary_key; default:uuid_generate_v4();unique"`
	IdCVC       int    `json:"-" sql:"type:integer;unique"`
	Description string `json:"description" sql:"type:varchar(50);unique;not null" valid:"length(2|50)~Descrição deve ter minimo 2 e maximo 50 caracter"`
	Active      bool   `sql:"type:boolean; default:true"`
}

type ProductList struct {
	Products []Product    `json:"products"`
	Meta     helpers.Meta `json:"meta"`
}

func (p Product) IsEmpty() bool {
	return p == Product{}
}

func (p Product) Validate() []string {
	var errs []string
	if _, err := govalidator.ValidateStruct(p); err != nil {
		errsV := err.(govalidator.Errors).Errors()
		for _, element := range errsV {
			errs = append(errs, element.Error())
		}
	}
	return errs
}

func (p *Product) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("uuid", nil)
	return nil
}
