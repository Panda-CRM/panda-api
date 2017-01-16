package controllers_test

import (
    "fmt"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
    "io/ioutil"
    "encoding/json"
    "github.com/wilsontamarozzi/panda-api/routers"
    "github.com/wilsontamarozzi/panda-api/services/models"
)

type Token struct {
	Token 	string `json:"token"`
	UserId 	string `json:"user_id"`
}

type PeopleResponse struct {
    People []models.Person `json:"people"`
}

type PersonResponse struct {
    Person models.Person `json:"person"`
}

var (
    SERVER     *httptest.Server
    token      Token
    PEOPLE_URL string
)

/**
 * Sequencia dos testes
 * @C - {Testa o cadastro de pessoas}
 * @R (ALL) - {Testa a listagem de pessoas}
 * @R - {Testa a busca de pessoa por ID}
 * @U - {Testa a alteração de pessoa por ID}
 * @D - {Testa a exclusão de pessoa por ID}
 */
func init() {
    SERVER = httptest.NewServer(routers.InitRoutes())
    PEOPLE_URL = fmt.Sprintf("%s/api/v1/people", SERVER.URL)
}

/**
 * @Autor Wilson
 * @Cenario Recebe o Token do JWT para iniciar os testes
 */
func TestAuthToken(t *testing.T) {
    url := fmt.Sprintf("%s/api/v1/auth/auth_token", SERVER.URL)
	
	userJson := `{"username": "admin", "password": "123"}`

    payload := strings.NewReader(userJson)

    req, _ := http.NewRequest("POST", url, payload)

    res, _ := http.DefaultClient.Do(req)

    defer res.Body.Close()
    body, _ := ioutil.ReadAll(res.Body)

    json.Unmarshal([]byte(string(body)), &token)

    if token.Token == "" || token.UserId == "" {
    	t.Error("Token ou Usuário ID vázio")
    }

    if res.StatusCode != http.StatusOK {
        t.Errorf("HTTP STATUS esperado %d, atual %d", http.StatusOK, res.StatusCode)
    }
}

/**
 * @Autor Wilson
 * @Cenario Testa o cadastro de pessoas
 */ 
func TestCreatePerson(t *testing.T) {
    
    expectedStatus := []int{
        201, // Cadastro Pessoa Fisica completo sem erros
        201, // Cadastro Pessoa Jurídica completo sem erros
    }

    people := []models.PersonRequest{
        // Cadastro Pessoa Fisica completo sem erros
        models.PersonRequest{
            models.Person{
                Type          : "F",
                Name          : "Wilson",
                CityName      : "Santa Barbara d'Oeste",
                Address       : "Alfredo Claus",
                Number        : "431",
                Complement    : "Casa",
                District      : "Conjunto Habitacional dos Trabalhadores",
                Zip           : "13.453-514",
                Cpf           : "416.781.718-75",
                Rg            : "488298490",
                Gender        : "M",
                BusinessPhone : "(99) 9999-9999",
                HomePhone     : "(99) 9999-9999",
                MobilePhone   : "(99) 9 9999-9999",
                Email         : "wilson.tamarozzi@gmail.com",
                Observations  : "Observações",
            },
        },
        // Cadastro Pessoa Jurídica completo sem erros
        models.PersonRequest{
            models.Person{
                Type             : "J",
                Name             : "Monde",
                CityName         : "Americana",
                CompanyName      : "Monde Sistemas De Informacao Ltda",
                Address          : "R Pernambuco",
                Number           : "1466",
                Complement       : "Sala 05",
                District         : "Jardim Nossa Senhora de Fatima",
                Zip              : "13.478-570",
                Cnpj             : "10.795.027/0001-71",
                StateInscription : "Isento",
                Phone            : "(99) 9999-9999",
                Fax              : "(99) 9999-9999",
                Email            : "contato@monde.com.br",
                Website          : "https://www.monde.com.br",
                Observations     : "Observações",
            },
        },
    }

    for i, person := range people {
        // converte a struct para json
        personJson, _ := json.Marshal(person)
        // converte o json para io.Reader
        payload := strings.NewReader(string(personJson))
        // prepara o request com Method, URL e Payload
        req, _ := http.NewRequest("POST", PEOPLE_URL, payload)
        // adiciona o header com o token
        req.Header.Add("Authorization", "Bearer " + token.Token)
        // executa a requisição ao servidor
        res, _ := http.DefaultClient.Do(req)
        // fecha conexão
        defer res.Body.Close()
        // lê a saida
        body, _ := ioutil.ReadAll(res.Body)
        // cria variavel pra pegar o retorno
        var personResponse PersonResponse
        // convert o json de retorno em objeto
        json.Unmarshal([]byte(string(body)), &personResponse)
        // analise se há o UUID no response de pessoa        
        if personResponse.Person.UUID == "" {
            t.Errorf("Esperado o UUID do response do cadastro de pessoa")
        }
        // analisa status de retorno com o esperado
        if res.StatusCode != expectedStatus[i] {
            t.Errorf("HTTP STATUS esperado %d, atual %d", http.StatusCreated, res.StatusCode)
            // lê o erro retornado pelo servidor
            body, _ := ioutil.ReadAll(res.Body)
            // exibe o erro de retorno
            t.Errorf(string(body))
        }        
    }
}

/**
 * @Autor Wilson
 * @Cenario Teste as validações do cadastro de pessoa
 */
func TestCreatePersonValidate(t *testing.T) {
    expectedStatus := []int{
        422, // Cadastro faltando o campo Sexo
        422, // Cadastro com campo Sexo inválido
        422, // Cadastro faltando o campo Tipo de Pessoa
        422, // Cadastro com campo Tipo de Pessoa inválido
        422, // Cadastro com campo CPF inválido
        422, // Cadastro com campo CNPJ inválido
    }

    people := []models.PersonRequest{
        // Cadastro faltando o campo Sexo
        models.PersonRequest{
            models.Person{
                Type          : "F",
                Name          : "Wilson",
                CityName      : "Santa Barbara d'Oeste",
                Address       : "Alfredo Claus",
                Number        : "431",
                Complement    : "Casa",
                District      : "Conjunto Habitacional dos Trabalhadores",
                Zip           : "13.453-514",
                Cpf           : "416.781.718-75",
                Rg            : "488298490",
                BusinessPhone : "(99) 9999-9999",
                HomePhone     : "(99) 9999-9999",
                MobilePhone   : "(99) 9 9999-9999",
                Email         : "wilson.tamarozzi@gmail.com",
                Observations  : "Observações",
            },
        },
        // Cadastro com campo Sexo inválido
        models.PersonRequest{
            models.Person{
                Type          : "F",
                Name          : "Wilson",
                CityName      : "Santa Barbara d'Oeste",
                Address       : "Alfredo Claus",
                Number        : "431",
                Complement    : "Casa",
                District      : "Conjunto Habitacional dos Trabalhadores",
                Zip           : "13.453-514",
                Cpf           : "416.781.718-75",
                Rg            : "488298490",
                Gender        : "X",
                BusinessPhone : "(99) 9999-9999",
                HomePhone     : "(99) 9999-9999",
                MobilePhone   : "(99) 9 9999-9999",
                Email         : "wilson.tamarozzi@gmail.com",
                Observations  : "Observações",
            },
        },
        // Cadastro faltando o campo Tipo de Pessoa
        models.PersonRequest{
            models.Person{
                Name          : "Wilson",
                CityName      : "Santa Barbara d'Oeste",
                Address       : "Alfredo Claus",
                Number        : "431",
                Complement    : "Casa",
                District      : "Conjunto Habitacional dos Trabalhadores",
                Zip           : "13.453-514",
                Cpf           : "416.781.718-75",
                Rg            : "488298490",
                Gender        : "M",
                BusinessPhone : "(99) 9999-9999",
                HomePhone     : "(99) 9999-9999",
                MobilePhone   : "(99) 9 9999-9999",
                Email         : "wilson.tamarozzi@gmail.com",
                Observations  : "Observações",
            },
        },
        // Cadastro com campo Tipo de Pessoa inválido
        models.PersonRequest{
            models.Person{
                Type          : "X",
                Name          : "Wilson",
                CityName      : "Santa Barbara d'Oeste",
                Address       : "Alfredo Claus",
                Number        : "431",
                Complement    : "Casa",
                District      : "Conjunto Habitacional dos Trabalhadores",
                Zip           : "13.453-514",
                Cpf           : "416.781.718-75",
                Rg            : "488298490",
                Gender        : "M",
                BusinessPhone : "(99) 9999-9999",
                HomePhone     : "(99) 9999-9999",
                MobilePhone   : "(99) 9 9999-9999",
                Email         : "wilson.tamarozzi@gmail.com",
                Observations  : "Observações",
            },
        },
        // Cadastro com campo CPF inválido
        models.PersonRequest{
            models.Person{
                Type          : "F",
                Name          : "Wilson",
                CityName      : "Santa Barbara d'Oeste",
                Address       : "Alfredo Claus",
                Number        : "431",
                Complement    : "Casa",
                District      : "Conjunto Habitacional dos Trabalhadores",
                Zip           : "13.453-514",
                Cpf           : "111.111.111-11",
                Rg            : "488298490",
                Gender        : "M",
                BusinessPhone : "(99) 9999-9999",
                HomePhone     : "(99) 9999-9999",
                MobilePhone   : "(99) 9 9999-9999",
                Email         : "wilson.tamarozzi@gmail.com",
                Observations  : "Observações",
            },
        },        
        // Cadastro com campo CNPJ inválido
        models.PersonRequest{
            models.Person{
                Type             : "J",
                Name             : "Monde",
                CityName         : "Americana",
                CompanyName      : "Monde Sistemas De Informacao Ltda",
                Address          : "R Pernambuco",
                Number           : "1466",
                Complement       : "Sala 05",
                District         : "Jardim Nossa Senhora de Fatima",
                Zip              : "13.478-570",
                Cnpj             : "11.111.111/1111-11",
                StateInscription : "Isento",
                Phone            : "(99) 9999-9999",
                Fax              : "(99) 9999-9999",
                Email            : "contato@monde.com.br",
                Website          : "https://www.monde.com.br",
                Observations     : "Observações",
            },
        },
    }

    for i, person := range people {
        // converte a struct para json
        personJson, _ := json.Marshal(person)
        // converte o json para io.Reader
        payload := strings.NewReader(string(personJson))
        // cria a requisição usando Method, URL e Payload
        req, _ := http.NewRequest("POST", PEOPLE_URL, payload)
        // adiciona o header com o token JWT
        req.Header.Add("Authorization", "Bearer " + token.Token)
        // executa a requisição
        res, _ := http.DefaultClient.Do(req)
        // fecha a conexão
        defer res.Body.Close()
        // analise o http status com o esperado
        if res.StatusCode != expectedStatus[i] {
            t.Errorf("HTTP STATUS esperado %d, atual %d", http.StatusCreated, res.StatusCode)
            // lê o erro retornado do servidor
            body, _ := ioutil.ReadAll(res.Body)
            // exibie o erro do retorno
            t.Errorf(string(body))
        }        
    }
}

/**
 * @Autor Wilson
 * @Cenario Testa a listagem de pessoas
 */
func TestGetPeople(t *testing.T) {
    // prepara a requisição usando Method, URL e Payload	
    req, _ := http.NewRequest("GET", PEOPLE_URL, nil)
    // adiciona o header com o token JWT
    req.Header.Add("Authorization", "Bearer " + token.Token)
    // executa a requisição
    res, _ := http.DefaultClient.Do(req)
    // fecha a conexão
    defer res.Body.Close()
    // lê o response do servidor
    body, _ := ioutil.ReadAll(res.Body)
    // cria a struct de objetos
    var listPeople PeopleResponse
    // converte o json em uma lista de objetos
    json.Unmarshal([]byte(string(body)), &listPeople)
    // analise se retornou um número diferente de 3 pessoas (2 do cadastro mais o Admin)
    if len(listPeople.People) != 3 {
        t.Errorf("Esperado 3 registros, atual %d", len(listPeople.People))
    }
    // analise o http status com o esperado
    if res.StatusCode != http.StatusOK {
        t.Errorf("HTTP STATUS esperado %d, atual %d", http.StatusOK, res.StatusCode)
    }	
}

/**
 * @Autor Wilson
 * @Cenario Testa a busca de pessoa por ID
 */
func TestGetPerson(t *testing.T) {
    // prepara a requisição usando Method, URL e Payload    
    req, _ := http.NewRequest("GET", PEOPLE_URL, nil)
    // adiciona o header com o token JWT
    req.Header.Add("Authorization", "Bearer " + token.Token)
    // executa a requisição
    res, _ := http.DefaultClient.Do(req)
    // fecha a conexão
    defer res.Body.Close()
    // lê o response do servidor
    body, _ := ioutil.ReadAll(res.Body)
    // cria a struct de objetos
    var listPeople PeopleResponse
    // converte o json em uma lista de objetos
    json.Unmarshal([]byte(string(body)), &listPeople)

    // após realizar a busca de todos os registros de uma vez
    // é feito a busca todos individualmente
    for _, person := range listPeople.People {
        // prepara a requisição
        req, _ := http.NewRequest("GET", PEOPLE_URL + "/" + person.UUID, nil)
        // adiciona o header com o token JWT
        req.Header.Add("Authorization", "Bearer " + token.Token)
        // executa a requisição
        res, _ := http.DefaultClient.Do(req)
        // fecha a conexão
        defer res.Body.Close()
        // analisa o http status com o esperado
        if res.StatusCode != http.StatusOK {
            t.Errorf("HTTP STATUS esperado %d, atual %d", http.StatusOK, res.StatusCode)
        }
    }
}

/**
 * @Autor Wilson
 * @Cenario Testa a alteração de pessoa pelo ID
 */
func TestUpdatePerson(t *testing.T) {
    // prepara a requisição
    req, _ := http.NewRequest("GET", PEOPLE_URL, nil)
    // adiciona o header com o token JWT
    req.Header.Add("Authorization", "Bearer " + token.Token)
    // executa a requisição
    res, _ := http.DefaultClient.Do(req)
    // fecha a conexão
    defer res.Body.Close()
    // lê o response do servidor
    body, _ := ioutil.ReadAll(res.Body)
    // cria a struct de objetos
    var listPeople PeopleResponse
    // converte o json em uma lista de objetos
    json.Unmarshal([]byte(string(body)), &listPeople)

    // após realizar a busca de todos os registros de uma vez
    // é feito a alteração todos individualmente (remove o ultimo indice que é o admin)
    for _, person := range listPeople.People[:len(listPeople.People)-1] {
        var p models.PersonRequest

        if person.Type == "F" {
            person.Type = "F"
            person.Name = "Wilson Tamarozzi"
            person.CityName = "Campinas"
            person.Address = "Avenida Coronel Silva Teles"
            person.Number = "675"
            person.Complement = ""
            person.District = "Cambuí"
            person.Zip = "13.024-001"
            person.Cpf = "416.781.718-75"
            person.Rg = "488298490"
            person.Gender = "M"
            person.BusinessPhone = "(99) 9999-9998"
            person.HomePhone = "(99) 9999-9998"
            person.MobilePhone = "(99) 9 9999-9998"
            person.Email = "wilson.tamarozzi@gmail.com"
            person.Observations = "Observações novas"
        }

        if person.Type == "J" {
            person.Type = "J"
            person.Name = "Panda"
            person.CityName = "São Paulo"
            person.CompanyName = "Panda Sistemas De Informacao Ltda"
            person.Address = "Av. Brasil"
            person.Number = "645"
            person.Complement = ""
            person.District = "Jardim Paulista"
            person.Zip = "13.478-570"
            person.Cnpj = "10.795.027/0001-71"
            person.StateInscription = "Isento"
            person.Phone = "(99) 9999-9998"
            person.Fax = "(99) 9999-9998"
            person.Email = "contato@panda.com.br"
            person.Website = "https://www.panda.com.br"
            person.Observations = "Observações novas"
        }
        
        p.Person = person
        // converte a struct para json
        personJson, _ := json.Marshal(p)
        // converte o json para io.Reader
        payload := strings.NewReader(string(personJson))
        // prepara a requisição
        req, _ := http.NewRequest("PUT", PEOPLE_URL + "/" + person.UUID, payload)
        // adiciona o header com o token JWT
        req.Header.Add("Authorization", "Bearer " + token.Token)
        // executa a requisição
        res, _ := http.DefaultClient.Do(req)
        // fecha conexão
        defer res.Body.Close()
        // analisa o http status com o esperado
        if res.StatusCode != 201 {
            t.Errorf("HTTP STATUS esperado %d, atual %d", 201, res.StatusCode)
            // lê o erro retornado do servidor
            body, _ := ioutil.ReadAll(res.Body)
            // exibie o erro do retorno
            t.Errorf(string(body))
        }
    }
}

/**
 * @Autor Wilson
 * @Cenario Testa a exclusão de pessoa por ID
 */
func TestDeletePerson(t *testing.T) {
    // prepara a requisição
    req, _ := http.NewRequest("GET", PEOPLE_URL, nil)
    // adiciona o header com o token JWT
    req.Header.Add("Authorization", "Bearer " + token.Token)
    // executa a requisição
    res, _ := http.DefaultClient.Do(req)
    // fecha a conexão
    defer res.Body.Close()
    // lê o response do servidor
    body, _ := ioutil.ReadAll(res.Body)
    // cria a struct de objetos
    var listPeople PeopleResponse
    // converte o json em uma lista de objetos
    json.Unmarshal([]byte(string(body)), &listPeople)

    // após realizar a busca de todos os registros de uma vez
    // é feito a exclusão todos individualmente (remove o ultimo indice que é o admin)
    for _, person := range listPeople.People {
        if person.UUID != "ce7405d8-3b78-4de7-8b58-6b32ac913701" {
            // prepara a requisição
            req, _ := http.NewRequest("DELETE", PEOPLE_URL + "/" + person.UUID, nil)
            // adiciona o header com o token JWT
            req.Header.Add("Authorization", "Bearer " + token.Token)
            // executa a requisição
            res, _ := http.DefaultClient.Do(req)
            // fecha conexão
            defer res.Body.Close()
            // analisa o http status com o esperado
            if res.StatusCode != 204 {
                t.Errorf("HTTP STATUS esperado %d, atual %d", 204, res.StatusCode)
            }
        }
    }
}