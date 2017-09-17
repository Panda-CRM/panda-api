package models

import (
	"testing"
	"time"
)

var (
	TEST_TASK_PERSON    Person
	TEST_TASK_CATEGORY  TaskCategory
	TEST_TASK_HISTORICS TaskHistorics
)

func init() {
	TEST_TASK_PERSON.UUID = "ce7405d8-3b78-4de7-8b58-6b32ac913701"
	TEST_TASK_PERSON.Name = "admin"

	TEST_TASK_CATEGORY.UUID = "756524a2-9555-4ae5-9a6c-b2232de896af"
	TEST_TASK_CATEGORY.Description = "Geral"

	TEST_TASK_HISTORICS = TaskHistorics{
		TaskHistoric{Comment: "Primeiro comentário"},
		TaskHistoric{Comment: "Segundo comentário"},
	}
}

func TestTarefaValida(t *testing.T) {
	amountErrorsExpected := 0
	data := time.Now()

	task := Task{
		Title:         "Geral",
		Due:           time.Now(),
		CompletedAt:   &data,
		Category:      TEST_TASK_CATEGORY,
		Person:        TEST_TASK_PERSON,
		Assignee:      TEST_TASK_PERSON,
		TaskHistorics: TEST_TASK_HISTORICS,
	}

	errorValidate := task.Validate()
	if len(errorValidate) != amountErrorsExpected {
		t.Errorf("[%s] Quantidade de erros esperado %d, atual %d", task.Title, amountErrorsExpected, len(errorValidate))
		for _, err := range errorValidate {
			t.Errorf(err)
		}
	}
}

func TestTarefaSemCampoCategoria(t *testing.T) {
	amountErrorsExpected := 1
	data := time.Now()

	task := Task{
		Title:         "Geral",
		Due:           time.Now(),
		CompletedAt:   &data,
		Person:        TEST_TASK_PERSON,
		Assignee:      TEST_TASK_PERSON,
		TaskHistorics: TEST_TASK_HISTORICS,
	}

	errorValidate := task.Validate()
	if len(errorValidate) != amountErrorsExpected {
		t.Errorf("[%s] Quantidade de erros esperado %d, atual %d", task.Title, amountErrorsExpected, len(errorValidate))
		for _, err := range errorValidate {
			t.Errorf(err)
		}
	}
}

func TestTarefaSemCampoPessoa(t *testing.T) {
	amountErrorsExpected := 1
	data := time.Now()

	task := Task{
		Title:         "Geral",
		Due:           time.Now(),
		CompletedAt:   &data,
		Category:      TEST_TASK_CATEGORY,
		Assignee:      TEST_TASK_PERSON,
		TaskHistorics: TEST_TASK_HISTORICS,
	}

	errorValidate := task.Validate()
	if len(errorValidate) != amountErrorsExpected {
		t.Errorf("[%s] Quantidade de erros esperado %d, atual %d", task.Title, amountErrorsExpected, len(errorValidate))
		for _, err := range errorValidate {
			t.Errorf(err)
		}
	}
}

func TestTarefaSemCampoResponsavel(t *testing.T) {
	amountErrorsExpected := 1
	data := time.Now()

	task := Task{
		Title:         "Geral",
		Due:           time.Now(),
		CompletedAt:   &data,
		Category:      TEST_TASK_CATEGORY,
		Person:        TEST_TASK_PERSON,
		TaskHistorics: TEST_TASK_HISTORICS,
	}

	errorValidate := task.Validate()
	if len(errorValidate) != amountErrorsExpected {
		t.Errorf("[%s] Quantidade de erros esperado %d, atual %d", task.Title, amountErrorsExpected, len(errorValidate))
		for _, err := range errorValidate {
			t.Errorf(err)
		}
	}
}
