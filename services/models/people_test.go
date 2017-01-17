package models_test

import(
	"testing"
	"github.com/wilsontamarozzi/panda-api/services/models"
)

func TestPessoaFisicaValido(t *testing.T) {

	amountErrorsExpected := 0

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

		errorValidate := person.Validate()

		if len(errorValidate) != amountErrorsExpected {
			t.Errorf("[%s] Quantidade de erros esperado %d, atual %d", person.Name, amountErrorsExpected, len(errorValidate))
			
			for _, err := range errorValidate {
				t.Errorf(err)
			}
		}
	}
}

func TestPessoaFisicaSemCampoTipoDePessoa(t *testing.T) {
	
	// Erros esperados
	// - Tipo de pessoa é obrigatório
	// - Tipo de pessoa deve ser F (Fisica) ou J (Juridica)
	amountErrorsExpected := 2

	person := models.Person{
		Name : "Pessoa",
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
	}
	
	errorValidate := person.Validate()

	if len(errorValidate) != amountErrorsExpected {
		t.Errorf("[%s] Quantidade de erros esperado %d, atual %d", person.Name, amountErrorsExpected, len(errorValidate))
		
		for _, err := range errorValidate {
			t.Errorf(err)
		}
	}
}

func TestPessoaFisicaCampoTipoDePessoaInvalido(t *testing.T) {
	
	// Erros esperados
	// - Tipo de pessoa deve ser F (Fisica) ou J (Juridica)
	amountErrorsExpected := 1

	person := models.Person{
		Type : "X",
		Name : "Pessoa",
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
	}
	
	errorValidate := person.Validate()

	if len(errorValidate) != amountErrorsExpected {
		t.Errorf("[%s] Quantidade de erros esperado %d, atual %d", person.Name, amountErrorsExpected, len(errorValidate))
		
		for _, err := range errorValidate {
			t.Errorf(err)
		}
	}
}

func TestPessoaFisicaSemCampoSexo(t *testing.T) {
	
	// Erros esperados
	// - Campo sexo é obrigatório
	amountErrorsExpected := 1

	person := models.Person{
		Type : "F",
		Name : "Pessoa",
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
	}
	
	errorValidate := person.Validate()

	if len(errorValidate) != amountErrorsExpected {
		t.Errorf("[%s] Quantidade de erros esperado %d, atual %d", person.Name, amountErrorsExpected, len(errorValidate))
		
		for _, err := range errorValidate {
			t.Errorf(err)
		}
	}
}

func TestPessoaFisicaCampoSexoInvalido(t *testing.T) {
	
	// Erros esperados
	// - Genero deve ser M (Masculino) ou F (Femenino)
	amountErrorsExpected := 1

	person := models.Person{
		Type : "F",
		Name : "Pessoa",
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
	}
	
	errorValidate := person.Validate()

	if len(errorValidate) != amountErrorsExpected {
		t.Errorf("[%s] Quantidade de erros esperado %d, atual %d", person.Name, amountErrorsExpected, len(errorValidate))
		
		for _, err := range errorValidate {
			t.Errorf(err)
		}
	}
}

func TestPessoaFisicaCampoCPFInvalido(t *testing.T) {
	
	// Erros esperados
	// - CPF inválido
	amountErrorsExpected := 1

	person := models.Person{
		Type : "F",
		Name : "Pessoa",
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
	}
	
	errorValidate := person.Validate()

	if len(errorValidate) != amountErrorsExpected {
		t.Errorf("[%s] Quantidade de erros esperado %d, atual %d", person.Name, amountErrorsExpected, len(errorValidate))
		
		for _, err := range errorValidate {
			t.Errorf(err)
		}
	}
}

func TestPessoaJuridicaValido(t *testing.T) {
	
	amountErrorsExpected := 0

	person := models.Person{
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
	}
	
	errorValidate := person.Validate()

	if len(errorValidate) != amountErrorsExpected {
		t.Errorf("[%s] Quantidade de erros esperado %d, atual %d", person.Name, amountErrorsExpected, len(errorValidate))
		
		for _, err := range errorValidate {
			t.Errorf(err)
		}
	}
}

func TestPessoaJuridicaSemCampoTipoDePessoa(t *testing.T) {
	
	// Erros esperados
	// - Tipo de pessoa é obrigatório
	// - Tipo de pessoa deve ser F (Fisica) ou J (Juridica)
	amountErrorsExpected := 2

	person := models.Person{
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
	}
	
	errorValidate := person.Validate()

	if len(errorValidate) != amountErrorsExpected {
		t.Errorf("[%s] Quantidade de erros esperado %d, atual %d", person.Name, amountErrorsExpected, len(errorValidate))
		
		for _, err := range errorValidate {
			t.Errorf(err)
		}
	}
}

func TestPessoaJuridicaCampoTipoDePessoaInvalido(t *testing.T) {
	
	// Erros esperados
	// - Tipo de pessoa deve ser F (Fisica) ou J (Juridica)
	amountErrorsExpected := 1

	person := models.Person{
		Type : "X",
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
	}
	
	errorValidate := person.Validate()

	if len(errorValidate) != amountErrorsExpected {
		t.Errorf("[%s] Quantidade de erros esperado %d, atual %d", person.Name, amountErrorsExpected, len(errorValidate))
		
		for _, err := range errorValidate {
			t.Errorf(err)
		}
	}
}

func TestPessoaJuridicaCampoCNPJInvalido(t *testing.T) {
	
	// Erros esperados
	// - CNPJ inválido
	amountErrorsExpected := 1

	person := models.Person{
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
	}
	
	errorValidate := person.Validate()

	if len(errorValidate) != amountErrorsExpected {
		t.Errorf("[%s] Quantidade de erros esperado %d, atual %d", person.Name, amountErrorsExpected, len(errorValidate))
		
		for _, err := range errorValidate {
			t.Errorf(err)
		}
	}
}

func TestPessoaTamanhoMaximoDosCampos(t *testing.T) {
	
	// Erros esperados
	// - Tipo de pessoa deve ser F (Fisica) ou J (Juridica)
	// - Tamanho do tipo de pessoa deve ser 1
	// - Nome deve ter minimo 2 e maximo 100 caracter
	// - Cidade deve ter no maximo 50 caracter
	// - Endereço deve ter no maximo 50 caracter
	// - Numero deve ter no maximo 7 caracter
	// - Complemento deve ter no maximo 50 caracter
	// - Bairro deve ter no maximo 50 caracter
	// - CEP deve ter no maximo 10 caracter
	// - CPF deve ter no maximo 14 caracter
	// - RG deve ter no maximo 20 caracter
	// - Telefone Comercial deve ter no maximo 20 caracter
	// - Telefone Residencial deve ter no maximo 20 caracter
	// - Telefone Celular deve ter no maximo 20 caracter
	// - E-mail deve ter no maximo 255 caracter
	amountErrorsExpected := 15

	person := models.Person{
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
	}
	
	errorValidate := person.Validate()

	if len(errorValidate) != amountErrorsExpected {
		t.Errorf("[%s] Quantidade de erros esperado %d, atual %d", person.Name, amountErrorsExpected, len(errorValidate))
		
		for _, err := range errorValidate {
			t.Errorf(err)
		}
	}
}

func TestPessoaSemCamposObrigatorios(t *testing.T) {
	
	// Erros esperados
	// - Tipo de pessoa deve ser F (Fisica) ou J (Juridica)
	// - Tipo de pessoa é obrigatório
	// - Nome é obrigatório
	amountErrorsExpected := 3

	person := models.Person{
		CityName : "Americana",
	}
	
	errorValidate := person.Validate()

	if len(errorValidate) != amountErrorsExpected {
		t.Errorf("[%s] Quantidade de erros esperado %d, atual %d", person.Name, amountErrorsExpected, len(errorValidate))
		
		for _, err := range errorValidate {
			t.Errorf(err)
		}
	}
}