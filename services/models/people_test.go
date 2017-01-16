package models_test

import(
	"testing"
	"github.com/wilsontamarozzi/panda-api/services/models"
)

func TestPersonValidatePass(t *testing.T) {

	people := models.People{
		{
			Type : "F",
			Name : "Wilson",
			CityName : "Santa Barbara d'Oeste",
			Address : "Alfredo Claus",
			Number : "431",
			Complement : "Casa",
			District : "Conjunto Habitacional dos Trabalhadores",
			Zip : "13.453-514",
			Cpf : "416.781.718-75",
			Rg : "48.829.849-0",
			Gender : "M",
			BusinessPhone : "(99) 9999-9999",
			HomePhone : "(99) 9999-9999",
			MobilePhone : "(99) 9 9999-9999",
			Email : "wilson@example.com",
			Observations : "Observações",
		},
		{
			Type : "F",
			Name : "Leonice",
			CityName : "Santa Barbara d'Oeste",
			Address : "Alfredo Claus",
			Number : "431",
			Complement : "Casa",
			District : "Conjunto Habitacional dos Trabalhadores",
			Zip : "13.453-514",
			Cpf : "738.156.648-61",
			Rg : "23.468.339-9",
			Gender : "M",
			BusinessPhone : "(99) 9999-9999",
			HomePhone : "(99) 9999-9999",
			MobilePhone : "(99) 9 9999-9999",
			Email : "leonice@example.com",
			Observations : "Observações",
		},
	}

	for _, person := range people {
		if err := person.Validate(); err != nil {
			t.Errorf("[%s] Erros esperado 0, atual %d", person.Name, len(err))
			
			for _, errorValidate := range err {
				t.Errorf(errorValidate)
			}
		}
	}
}

func TestPersonValidateFail(t *testing.T) {
	
	amountErrorExpected := []int{
		2, // sem o campo tipo de pessoa
		1, // tipo de pessoa inválido
		1, // sem o campo sexo
		1, // sexo inválido
		1, // CPF inválido
	}

	people := models.People{
		{ // sem o campo tipo de pessoa
			Name : "Pessoa 1",
			CityName : "Santa Barbara d'Oeste",
			Address : "Alfredo Claus",
			Number : "431",
			Complement : "Casa",
			District : "Conjunto Habitacional dos Trabalhadores",
			Zip : "13.453-514",
			Cpf : "416.781.718-75",
			Rg : "48.829.849-0",
			Gender : "M",
			BusinessPhone : "(99) 9999-9999",
			HomePhone : "(99) 9999-9999",
			MobilePhone : "(99) 9 9999-9999",
			Email : "wilson@example.com",
			Observations : "Observações",
		},
		{ // tipo de pessoa inválido
			Type : "X",
			Name : "Pessoa 2",
			CityName : "Santa Barbara d'Oeste",
			Address : "Alfredo Claus",
			Number : "431",
			Complement : "Casa",
			District : "Conjunto Habitacional dos Trabalhadores",
			Zip : "13.453-514",
			Cpf : "416.781.718-75",
			Rg : "48.829.849-0",
			Gender : "M",
			BusinessPhone : "(99) 9999-9999",
			HomePhone : "(99) 9999-9999",
			MobilePhone : "(99) 9 9999-9999",
			Email : "wilson@example.com",
			Observations : "Observações",
		},
		{ // sem o campo sexo
			Type : "F",
			Name : "Pessoa 3",
			CityName : "Santa Barbara d'Oeste",
			Address : "Alfredo Claus",
			Number : "431",
			Complement : "Casa",
			District : "Conjunto Habitacional dos Trabalhadores",
			Zip : "13.453-514",
			Cpf : "416.781.718-75",
			Rg : "48.829.849-0",
			BusinessPhone : "(99) 9999-9999",
			HomePhone : "(99) 9999-9999",
			MobilePhone : "(99) 9 9999-9999",
			Email : "wilson@example.com",
			Observations : "Observações",
		},
		{ // sexo inválido
			Type : "F",
			Name : "Pessoa 4",
			CityName : "Santa Barbara d'Oeste",
			Address : "Alfredo Claus",
			Number : "431",
			Complement : "Casa",
			District : "Conjunto Habitacional dos Trabalhadores",
			Zip : "13.453-514",
			Cpf : "416.781.718-75",
			Rg : "48.829.849-0",
			Gender : "X",
			BusinessPhone : "(99) 9999-9999",
			HomePhone : "(99) 9999-9999",
			MobilePhone : "(99) 9 9999-9999",
			Email : "wilson@example.com",
			Observations : "Observações",
		},
		{ // CPF inválido
			Type : "F",
			Name : "Pessoa 5",
			CityName : "Santa Barbara d'Oeste",
			Address : "Alfredo Claus",
			Number : "431",
			Complement : "Casa",
			District : "Conjunto Habitacional dos Trabalhadores",
			Zip : "13.453-514",
			Cpf : "111.111.111-11",
			Rg : "48.829.849-0",
			Gender : "M",
			BusinessPhone : "(99) 9999-9999",
			HomePhone : "(99) 9999-9999",
			MobilePhone : "(99) 9 9999-9999",
			Email : "wilson@example.com",
			Observations : "Observações",
		},
	}

	for i, person := range people {
		if err := person.Validate(); len(err) != amountErrorExpected[i] {
			t.Errorf("[%s] Quantidade de erros esperado %d, atual %d", person.Name, amountErrorExpected[i], len(err))
			
			for _, errorValidate := range err {
				t.Errorf(errorValidate)
			}
		}
	}
}

func TestCorporateValidatePass(t *testing.T) {

	people := models.People{
		{
			Type : "J",
			Name : "Panda",
			CityName : "Americana",
			CompanyName : "Panda System LDTA",
			Address : "Rua Pernambuco",
			Number : "1466",
			Complement : "Sala 05",
			District : "Jardim Nossa Senhora de Fatima",
			Zip : "13.453-514",
			Cnpj : "36.454.648/0001-85",
			StateInscription : "Isento",
			Phone : "(99) 9999-9999",
			Fax : "(99) 9999-9999",
			Email : "panda@example.com",
			Website : "http://www.example.com",
			Observations : "Observações",
		},
	}

	for _, person := range people {
		if err := person.Validate(); err != nil {
			t.Errorf("[%s] Erros esperado 0, atual %d", person.Name, len(err))
			
			for _, errorValidate := range err {
				t.Errorf(errorValidate)
			}
		}
	}
}

func TestCorporateValidateFail(t *testing.T) {
	
	amountErrorExpected := []int{
		2, // sem o campo tipo de pessoa
		1, // CNPJ inválido
	}

	people := models.People{
		{ // sem o campo tipo de pessoa
			Name : "Panda",
			CityName : "Americana",
			CompanyName : "Panda System LDTA",
			Address : "Rua Pernambuco",
			Number : "1466",
			Complement : "Sala 05",
			District : "Jardim Nossa Senhora de Fatima",
			Zip : "13.453-514",
			Cnpj : "36.454.648/0001-85",
			StateInscription : "Isento",
			Phone : "(99) 9999-9999",
			Fax : "(99) 9999-9999",
			Email : "panda@example.com",
			Website : "http://www.example.com",
			Observations : "Observações",
		},
		{ // CNPJ inválido
			Type : "J",
			Name : "Panda",
			CityName : "Americana",
			CompanyName : "Panda System LDTA",
			Address : "Rua Pernambuco",
			Number : "1466",
			Complement : "Sala 05",
			District : "Jardim Nossa Senhora de Fatima",
			Zip : "13.453-514",
			Cnpj : "11.111.111/1111-11",
			StateInscription : "Isento",
			Phone : "(99) 9999-9999",
			Fax : "(99) 9999-9999",
			Email : "panda@example.com",
			Website : "http://www.example.com",
			Observations : "Observações",
		},
	}

	for i, person := range people {
		if err := person.Validate(); len(err) != amountErrorExpected[i] {
			t.Errorf("[%s] Quantidade de erros esperado %d, atual %d", person.Name, amountErrorExpected[i], len(err))
			
			for _, errorValidate := range err {
				t.Errorf(errorValidate)
			}
		}
	}
}

func TestPersonValidateSizeMaxField(t *testing.T) {
	
	amountErrorExpected := []int{15, 3}

	people := models.People{
		// --- Erros esperados para o primeiro cenário (total: 15)
		// Tipo de pessoa deve ser F (Fisica) ou J (Juridica)
		// Tamanho do tipo de pessoa deve ser 1
		// Nome deve ter minimo 2 e maximo 100 caracter
		// Cidade deve ter no maximo 50 caracter
		// Endereço deve ter no maximo 50 caracter
		// Numero deve ter no maximo 7 caracter
		// Complemento deve ter no maximo 50 caracter
		// Bairro deve ter no maximo 50 caracter
		// CEP deve ter no maximo 10 caracter
		// CPF deve ter no maximo 14 caracter
		// RG deve ter no maximo 20 caracter
		// Telefone Comercial deve ter no maximo 20 caracter
		// Telefone Residencial deve ter no maximo 20 caracter
		// Telefone Celular deve ter no maximo 20 caracter
		// E-mail deve ter no maximo 255 caracter
		{
			Type : "FF",
			Gender : "MM",
			Name : "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
			CityName : "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
			Address : "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
			Number : "12345678",
			Complement : "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
			District : "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
			Zip : "11.111-1111",
			Cpf : "416.781.718-757",
			Rg : "AAAAAAAAAAAAAAAAAAAAA",
			BusinessPhone : "AAAAAAAAAAAAAAAAAAAAA",
			HomePhone : "AAAAAAAAAAAAAAAAAAAAA",
			MobilePhone : "AAAAAAAAAAAAAAAAAAAAA",
			Email : "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
		},
		// --- Erros esperado para esse cenário (total: 3)
		// Tipo de pessoa deve ser F (Fisica) ou J (Juridica)
		// Tipo de pessoa é obrigatório
		// Nome é obrigatório
		{
			CityName : "Americana",
		},
	}

	for i, person := range people {
		if err := person.Validate(); len(err) != amountErrorExpected[i] {
			t.Errorf("[%s] Quantidade de erros esperado %d, atual %d", person.Name, amountErrorExpected[i], len(err))

			for _, errorValidate := range err {
				t.Errorf(errorValidate)
			}
		}
	}
}