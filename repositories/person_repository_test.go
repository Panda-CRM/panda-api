package repositories

import (
	"github.com/wilsontamarozzi/panda-api/database"
	"github.com/wilsontamarozzi/panda-api/models"
	"net/url"
	"testing"
)

var (
	TEST_PERSON_PARAMS_LIST url.Values
	TEST_PERSON_MODEL       models.Person
)

func init() {
	TEST_PERSON_PARAMS_LIST = url.Values{}
	TEST_PERSON_PARAMS_LIST.Add("page", "1")
	TEST_PERSON_PARAMS_LIST.Add("per_page", "50")

	cpf := "09515274010"
	TEST_PERSON_MODEL.Type = models.TYPE_PERSON
	TEST_PERSON_MODEL.Name = "Admin"
	TEST_PERSON_MODEL.Cpf = &cpf
	TEST_PERSON_MODEL.UUID = "ce7405d8-3b78-4de7-8b58-6b32ac913701"
	TEST_PERSON_MODEL.CreatedByUUID = "ce7405d8-3b78-4de7-8b58-6b32ac913701"

	database.RebuildDataBase(false)
}

func TestPersonCreate(t *testing.T) {
	err := NewPersonRepository().Create(&TEST_PERSON_MODEL)
	if err != nil {
		t.Errorf("Não foi possivel cadastrar a pessoa")
	}
}

func TestPessoaGet(t *testing.T) {
	person := NewPersonRepository().Get(TEST_PERSON_MODEL.UUID)
	if person.UUID != TEST_PERSON_MODEL.UUID {
		t.Errorf("Usuário não encontrado")
	}
}

func TestPessoaGetByCPF(t *testing.T) {
	person := NewPersonRepository().GetByCPF(*TEST_PERSON_MODEL.Cpf)
	if person.UUID != TEST_PERSON_MODEL.UUID {
		t.Errorf("Usuário não encontrado")
	}
}

func TestPessoaLista(t *testing.T) {
	personList := NewPersonRepository().List(TEST_PERSON_PARAMS_LIST)
	people := personList.People
	if people[0].UUID != TEST_PERSON_MODEL.UUID {
		t.Errorf("Usuários não listados")
	}
}
