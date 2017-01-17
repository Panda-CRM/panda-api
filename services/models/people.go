package models

import (
	"strings"
	"time"
	"errors"
	"github.com/wilsontamarozzi/panda-api/helpers"
	"github.com/asaskevich/govalidator"
)

var (
	ErrInvalidType = errors.New("Tipo de pessoa deve ser F (Fisica) ou J (Juridica)")
	ErrEmptyGender = errors.New("Campo sexo é obrigatório")
	ErrInvalidGender = errors.New("Genero deve ser M (Masculino) ou F (Femenino)")
)

type Person struct {
	UUID string 				`json:"id,omitempty" sql:"type:uuid; primary_key; default:uuid_generate_v4();unique"`
	Code int 					`json:"code" sql:"auto_increment; primary_key"`
	Type string					`json:"type" sql:"type:varchar(1); not null" valid:"required~Tipo de pessoa é obrigatório,length(1|1)~Tamanho do tipo de pessoa deve ser 1"`
	Name string 				`json:"name" sql:"type:varchar(100); not null" valid:"required~Nome é obrigatório,length(2|100)~Nome deve ter minimo 2 e maximo 100 caracter"`
	CityName string				`json:"city_name" sql:"type:varchar(50)" valid:"length(0|50)~Cidade deve ter no maximo 50 caracter"`
	CompanyName string			`json:"company_name" sql:"type:varchar(100)" valid:"length(2|100)~Nome deve ter minimo 1 e maximo 100 caracter"`
	Address string 				`json:"address" sql:"type:varchar(50)" valid:"length(0|50)~Endereço deve ter no maximo 50 caracter"`
	Number string	 			`json:"number" sql:"type:varchar(7)" valid:"length(0|7)~Numero deve ter no maximo 7 caracter"`
	Complement string	 		`json:"complement" sql:"type:varchar(50)" valid:"length(0|50)~Complemento deve ter no maximo 50 caracter"`
	District string	 			`json:"district" sql:"type:varchar(50)" valid:"length(0|50)~Bairro deve ter no maximo 50 caracter"`
	Zip string	 				`json:"zip" sql:"type:varchar(10)" valid:"length(0|10)~CEP deve ter no maximo 10 caracter"`
	BirthDate *time.Time  		`json:"birth_date" sql:"type:timestamp without time zone; default:null"`
	Cpf string	 				`json:"cpf" sql:"type:varchar(14)" valid:"length(0|14)~CPF deve ter no maximo 14 caracter"`
	Rg string	 				`json:"rg" sql:"type:varchar(20)" valid:"length(0|20)~RG deve ter no maximo 20 caracter"`
	Gender string	 			`json:"gender" sql:"type:varchar(1)"`
	BusinessPhone string 		`json:"business_phone" sql:"type:varchar(20)" valid:"length(0|20)~Telefone Comercial deve ter no maximo 20 caracter"`
	HomePhone string	 		`json:"home_phone" sql:"type:varchar(20)" valid:"length(0|20)~Telefone Residencial deve ter no maximo 20 caracter"`
	MobilePhone string	 		`json:"mobile_phone" sql:"type:varchar(20)" valid:"length(0|20)~Telefone Celular deve ter no maximo 20 caracter"`
	Cnpj string	 				`json:"cnpj" sql:"type:varchar(18)" valid:"length(0|18)~CNPJ deve ter no maximo 18 caracter"`
	StateInscription string 	`json:"state_inscription" sql:"type:varchar(20)" valid:"length(0|20)~Inscrição Estadual deve ter no maximo 20 caracter"`
	Phone string	 			`json:"phone" sql:"type:varchar(20)" valid:"length(0|20)~Telefone deve ter no maximo 20 caracter"`
	Fax string	 				`json:"fax" sql:"type:varchar(20)" valid:"length(0|20)~FAX deve ter no maximo 20 caracter"`
	Email string	 			`json:"email" sql:"type:varchar(255)" valid:"length(0|255)~E-mail deve ter no maximo 255 caracter"`
	Website string	 			`json:"website" sql:"type:varchar(50)" valid:"length(0|50)~Website deve ter no maximo 50 caracter"`
	Observations string	 		`json:"observations" sql:"type:text"`
	RegisteredAt time.Time 		`json:"registered_at" sql:"type:timestamp without time zone; default:NOW()"`
	RegisteredByUUID string 	`json:"registered_by" sql:"type:uuid"`
	IsUser bool 				`json:"-" sql:"type:boolean"`
}

type People []Person

type PersonRequest struct {
	Person Person `json:"person"`
}

func (p Person) Validate() []string {

	var errors []string

	if p.Type != "F" && p.Type != "J" {
		errors = append(errors, ErrInvalidType.Error())
	}

	// Valida pessoa física
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

	// Valida pessoa jurídica
	if p.Type == "J" {
		if !govalidator.IsNull(p.Cnpj) {
			if err := helpers.ValidateCNPJ(p.Cnpj); err != nil {
				for _, element := range err {
					errors = append(errors, element.Error())
				}
			}
		}
	}

	// Valida a estrutura pelas tags
	if _, err := govalidator.ValidateStruct(p); err != nil {
		// splita erros por ponto e vigular
		errParse := strings.Split(err.Error(), ";")
		// remove o ultimo indice que vem vázio
		removeLastEmpty := errParse[:len(errParse)-1]

		for _, element := range removeLastEmpty {
			errors = append(errors, element)
		}
	}

	return errors
}