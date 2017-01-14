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
)

type Token struct {
	Token 	string `json:"token"`
	UserId 	string `json:"user_id"`
}

var (
    server 	*httptest.Server
    token 	Token
)

var PEOPLE_URL string

func init() {
    server = httptest.NewServer(routers.InitRoutes())
    PEOPLE_URL = fmt.Sprintf("%s/api/v1/people", server.URL)
}

func TestAuthToken(t *testing.T) {
    url := fmt.Sprintf("%s/api/v1/auth/auth_token", server.URL)
	
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

    if res.StatusCode != 200 {
        t.Errorf("Success expected: %d", res.StatusCode)
    }
}

/**
 * @Autor Wilson
 * @Cenario Testa o cadastro de varios tipo de pessoa
 * sendo eles com e sem algumas informações essenciais
 *
 * C / R (ALL) / R / U / D
 */ 
func TestCreatePerson(t *testing.T) {    
    person := []struct{
        Name string `json:"name"`
        Type string `json:"type"`
        Gender string `json:"gender"`
    }{
        {"Wilson", "F", "M"},
        {"Monde", "J", ""},
    }

    json, _ := json.Marshal(person)
    payload := strings.NewReader(string(json))

    req, _ := http.NewRequest("POST", PEOPLE_URL, payload)

    req.Header.Add("Authorization", "Bearer " + token.Token)

    res, _ := http.DefaultClient.Do(req)
    body, _ := ioutil.ReadAll(res.Body)

    fmt.Println(string(body))

    defer res.Body.Close()

    if res.StatusCode != 201 {
        t.Errorf("Success expected: %d", res.StatusCode)
    }
}

/**
 * @Autor Wilson
 * @Cenario Testa a listagem de pessoas
 */
func TestGetPeople(t * testing.T) {
	url := fmt.Sprintf("%s/api/v1/people", server.URL)
	
    req, _ := http.NewRequest("GET", url, nil)

    req.Header.Add("Authorization", "Bearer " + token.Token)

    res, _ := http.DefaultClient.Do(req)

    defer res.Body.Close()
    //body, _ := ioutil.ReadAll(res.Body)

    if res.StatusCode != 200 {
        t.Errorf("Success expected: %d", res.StatusCode)
    }	
}

/**
 * @Autor Wilson
 * @Cenario Testa a busca de pessoa por ID
 */
func TestGetPerson(t * testing.T) {
	url := fmt.Sprintf("%s/api/v1/people/ce7405d8-3b78-4de7-8b58-6b32ac913701", server.URL)
	
    req, _ := http.NewRequest("GET", url, nil)

    req.Header.Add("Authorization", "Bearer " + token.Token)

    res, _ := http.DefaultClient.Do(req)

    defer res.Body.Close()
    //body, _ := ioutil.ReadAll(res.Body)

    if res.StatusCode != 200 {
        t.Errorf("Success expected: %d", res.StatusCode)
    }	
}