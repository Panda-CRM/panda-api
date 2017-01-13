package controllers_test

import (
    "fmt"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
    "panda-api/routers"
    "io/ioutil"
    "encoding/json"
)

type Token struct {
	Token 	string `json:"token"`
	UserId 	string `json:"user_id"`
}

var (
    server 	*httptest.Server
    token 	Token
)

func init() {
    server = httptest.NewServer(routers.InitRoutes())
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