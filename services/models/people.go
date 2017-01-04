package models

import (
	"time"
	"errors"
	"panda-api/helpers"
	"github.com/asaskevich/govalidator"
)

var (
	ErrEmptyName = errors.New("Nome não pode ser vázio")
	ErrEmptyType = errors.New("Tipo de pessoa não definido")
	ErrInvalidType = errors.New("Tipo de pessoa inválido")
	ErrEmptyGender = errors.New("O campo sexo é obrigatório")
	ErrInvalidGender = errors.New("Campo sexo inválido")
)

type Person struct {
	UUID 				string		`json:"id,omitempty" sql:"type:uuid; primary_key; default:uuid_generate_v4()"`
	Code 				int 		`json:"code" sql:"auto_increment; primary_key"`
	Type				string		`json:"type" sql:"type:varchar(1); not null" valid:"length(1|1)"`
	Name				string 		`json:"name" sql:"type:varchar(100); not null" valid:"length(2|100)"`
	CityName 			string		`json:"city_name" sql:"type:varchar(50)" valid:"length(0|50)"`
	CompanyName 		string		`json:"company_name" sql:"type:varchar(100)" valid:"length(2|100)"`
	Address 			string 		`json:"address" sql:"type:varchar(50)" valid:"length(0|50)"`
	Number 				string	 	`json:"number" sql:"type:varchar(7)" valid:"length(0|7)"`
	Complement 			string	 	`json:"complement" sql:"type:varchar(50)" valid:"length(0|50)"`
	District 			string	 	`json:"district" sql:"type:varchar(50)" valid:"length(0|50)"`
	Zip 				string	 	`json:"zip" sql:"type:varchar(9)" valid:"length(0|9)"`
	BirthDate 			*time.Time  `json:"birth_date" sql:"type:timestamp without time zone; default:null"`
	Cpf 				string	 	`json:"cpf" sql:"type:varchar(14)" valid:"length(0|14)"`
	Rg 					string	 	`json:"rg" sql:"type:varchar(20)" valid:"length(0|20)"`
	Gender 				string	 	`json:"gender" sql:"type:varchar(1)" valid:"length(1|1)"`
	BusinessPhone 		string	 	`json:"business_phone" sql:"type:varchar(20)" valid:"length(0|20)"`
	HomePhone 			string	 	`json:"home_phone" sql:"type:varchar(20)" valid:"length(0|20)"`
	MobilePhone 		string	 	`json:"mobile_phone" sql:"type:varchar(20)" valid:"length(0|20)"`
	Cnpj 				string	 	`json:"cnpj" sql:"type:varchar(18)" valid:"length(0|18)"`
	StateInscription 	string 		`json:"state_inscription" sql:"type:varchar(20)" valid:"length(0|20)"`
	Phone 				string	 	`json:"phone" sql:"type:varchar(20)" valid:"length(0|20)"`
	Fax 				string	 	`json:"fax" sql:"type:varchar(20)" valid:"length(0|20)"`
	Email 				string	 	`json:"email" sql:"type:varchar(255)" valid:"length(0|255)"`
	Website 			string	 	`json:"website" sql:"type:varchar(50)" valid:"length(0|50)"`
	Observations 		string	 	`json:"observations" sql:"type:text"`
	RegisteredAt		time.Time 	`json:"registered_at" sql:"type:timestamp without time zone; default:NOW()"`
	RegisteredByUUID	string 		`json:"registered_by" sql:"type:uuid"`
}

type People []Person

type PersonRequest struct {
	Person Person 		`json:"person"`
}

func (p Person) Validate() []string {

	var errors []string

	if govalidator.IsNull(p.Name) {
		errors = append(errors, ErrEmptyName.Error())
	}

	if govalidator.IsNull(p.Type) {
		errors = append(errors, ErrEmptyType.Error())
	}

	if p.Type != "F" && p.Type != "J" {
		errors = append(errors, ErrInvalidType.Error())
	}

	if p.Type == "F" {
		if govalidator.IsNull(p.Gender) {
			errors = append(errors, ErrEmptyGender.Error())
		} else if p.Gender != "M" && p.Gender != "F" {
			errors = append(errors, ErrInvalidGender.Error())
		}

		if !govalidator.IsNull(p.Cpf) {
			if err := helpers.ValidateCPF(p.Cpf); err != nil {
				for _, element := range err {
					errors = append(errors, element.Error())
				}
			}
		}
	}

	if p.Type == "J" {
		if !govalidator.IsNull(p.Cnpj) {
			if err := helpers.ValidateCNPJ(p.Cnpj); err != nil {
				for _, element := range err {
					errors = append(errors, element.Error())
				}
			}
		}
	}

	if _, err := govalidator.ValidateStruct(p); err != nil {
		errors = append(errors, err.Error())
	}

	return errors
}