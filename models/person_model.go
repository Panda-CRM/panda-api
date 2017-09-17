package models

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	"github.com/Panda-CRM/panda-api/helpers"
	"time"
)

var (
	ErrInvalidType   = errors.New("Tipo de pessoa deve ser F (Fisica) ou J (Juridica)")
	ErrEmptyGender   = errors.New("Campo sexo é obrigatório")
	ErrInvalidGender = errors.New("Genero deve ser M (Masculino) ou F (Feminino)")
	TYPE_PERSON      = "F"
	TYPE_COMPANY     = "J"
	GENDER_MALE      = "M"
	GENDER_FEMALE    = "F"
)

type Person struct {
	ModelBase
	IdCVC            *int    `json:"-" sql:"type:integer;unique"`
	Code             int     `json:"code,omitempty" sql:"auto_increment; primary_key; unique"`
	Type             string  `json:"type,omitempty" sql:"type:varchar(1); not null" valid:"required~Tipo de pessoa é obrigatório,length(1|1)~Tamanho do tipo de pessoa deve ser 1"`
	Name             string  `json:"name,omitempty" sql:"type:varchar(100); not null" valid:"required~Nome é obrigatório,length(2|100)~Nome deve ter minimo 2 e maximo 100 caracter"`
	CityName         string  `json:"city_name,omitempty" sql:"type:varchar(50)" valid:"length(0|50)~Cidade deve ter no maximo 50 caracter"`
	CompanyName      string  `json:"company_name,omitempty" sql:"type:varchar(100)" valid:"length(2|100)~Nome deve ter minimo 1 e maximo 100 caracter"`
	Address          string  `json:"address,omitempty" sql:"type:varchar(50)" valid:"length(0|50)~Endereço deve ter no maximo 50 caracter"`
	Number           string  `json:"number,omitempty" sql:"type:varchar(7)" valid:"length(0|7)~Numero deve ter no maximo 7 caracter"`
	Complement       string     `json:"complement,omitempty" sql:"type:varchar(50)" valid:"length(0|50)~Complemento deve ter no maximo 50 caracter"`
	District         string     `json:"district,omitempty" sql:"type:varchar(50)" valid:"length(0|50)~Bairro deve ter no maximo 50 caracter"`
	Zip              string     `json:"zip,omitempty" sql:"type:varchar(10)" valid:"length(0|10)~CEP deve ter no maximo 10 caracter"`
	BirthDate        *time.Time `json:"birth_date,omitempty" sql:"type:timestamp without time zone; default:null"`
	Cpf              *string    `json:"cpf,omitempty" sql:"type:varchar(14);unique" valid:"length(0|14)~CPF deve ter no maximo 14 caracter"`
	Rg               *string    `json:"rg,omitempty" sql:"type:varchar(20);unique" valid:"length(0|20)~RG deve ter no maximo 20 caracter"`
	Gender           string     `json:"gender,omitempty" sql:"type:varchar(1)"`
	BusinessPhone    string     `json:"business_phone,omitempty" sql:"type:varchar(20)" valid:"length(0|20)~Telefone Comercial deve ter no maximo 20 caracter"`
	HomePhone        string  `json:"home_phone,omitempty" sql:"type:varchar(20)" valid:"length(0|20)~Telefone Residencial deve ter no maximo 20 caracter"`
	MobilePhone      string  `json:"mobile_phone,omitempty" sql:"type:varchar(20)" valid:"length(0|20)~Telefone Celular deve ter no maximo 20 caracter"`
	Cnpj             *string `json:"cnpj,omitempty" sql:"type:varchar(18);unique" valid:"length(0|18)~CNPJ deve ter no maximo 18 caracter"`
	StateInscription string  `json:"state_inscription,omitempty" sql:"type:varchar(20)" valid:"length(0|20)~Inscrição Estadual deve ter no maximo 20 caracter"`
	Phone            string  `json:"phone,omitempty" sql:"type:varchar(20)" valid:"length(0|20)~Telefone deve ter no maximo 20 caracter"`
	Fax              string  `json:"fax,omitempty" sql:"type:varchar(20)" valid:"length(0|20)~FAX deve ter no maximo 20 caracter"`
	Email            string  `json:"email,omitempty" sql:"type:varchar(255)" valid:"length(0|255)~E-mail deve ter no maximo 255 caracter"`
	Website          string  `json:"website,omitempty" sql:"type:varchar(50)" valid:"length(0|50)~Website deve ter no maximo 50 caracter"`
	Observations     string  `json:"observations,omitempty" sql:"type:text"`
	User             bool    `json:"-" sql:"type:boolean"`
	Username         string  `json:"-" sql:"type:varchar(50); unique"`
	Password         string  `json:"-" sql:"type:varchar(50)"`
	RoleUUID         string  `json:"-" sql:"type:uuid"`
	Role             Role    `json:"role,omitempty"`
}

type PersonList struct {
	TotalCount int                `json:"total_count"`
	People     []Person           `json:"people"`
	Pages      helpers.PageParams `json:"pages"`
}

func (p Person) IsEmpty() bool {
	return p.UUID == ""
}

func (p Person) IsPerson() bool {
	return p.Type == TYPE_PERSON
}

func (p Person) IsCompany() bool {
	return p.Type == TYPE_COMPANY
}

func (p Person) IsMale() bool {
	return p.Gender == GENDER_MALE
}

func (p Person) IsFemale() bool {
	return p.Gender == GENDER_FEMALE
}

func (p Person) IsUser() bool {
	return p.User
}

func (p Person) ValidatePerson() []string {
	var errs []string
	if p.Gender == "" {
		errs = append(errs, ErrEmptyGender.Error())
	} else if !p.IsMale() && !p.IsFemale() {
		errs = append(errs, ErrInvalidGender.Error())
	}

	if p.Cpf != nil {
		if err := helpers.ValidateCPF(*p.Cpf); err != nil {
			for _, element := range err {
				errs = append(errs, element.Error())
			}
		}
	}
	return errs
}

func (p Person) ValidateCompany() []string {
	var errs []string
	if p.Cnpj != nil {
		if err := helpers.ValidateCNPJ(*p.Cnpj); err != nil {
			for _, element := range err {
				errs = append(errs, element.Error())
			}
		}
	}
	return errs
}

func (p Person) Validate() []string {
	var errs []string
	// Valida se é uma pessoa ou empresa
	if !p.IsPerson() && !p.IsCompany() {
		errs = append(errs, ErrInvalidType.Error())
	}
	// Valida pessoa física
	if p.IsPerson() {
		errs = append(errs, p.ValidatePerson()...)
	}
	// Valida pessoa jurídica
	if p.IsCompany() {
		errs = append(errs, p.ValidateCompany()...)
	}
	// Valida os campos da estrutura
	if _, err := govalidator.ValidateStruct(p); err != nil {
		errsV := err.(govalidator.Errors).Errors()
		for _, element := range errsV {
			errs = append(errs, element.Error())
		}
	}
	return errs
}

func (p Person) PopulateDefault(db *gorm.DB) {
	p.UUID = "ce7405d8-3b78-4de7-8b58-6b32ac913701"
	p.CreatedByUUID = "ce7405d8-3b78-4de7-8b58-6b32ac913701"
	p.Type = TYPE_PERSON
	p.Name = "Admin"
	p.User = true
	p.Username = "admin"
	p.Password = "202cb962ac59075b964b07152d234b70"
	p.RoleUUID = "899182db-4f57-4ec3-a263-3f83a4a66a6a"
	db.Create(&p)
}

func (p *Person) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("code", nil)
	scope.SetColumn("created_at", time.Now())
	if !p.IsUser() {
		scope.SetColumn("uuid", nil)
	}
	return nil
}
