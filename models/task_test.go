package models

import(
	"testing"
	"time"
)

func TestTarefaValida(t *testing.T) {

	amountErrorsExpected := 0

	data := time.Now()

	task := Task{
		Title : "Geral",
		Due : time.Now(),
		CompletedAt : &data,
		Category : TaskCategory{
	        UUID : "756524a2-9555-4ae5-9a6c-b2232de896af",
	        Description : "Geral",
	    },
		Person : Person{
	        UUID : "ce7405d8-3b78-4de7-8b58-6b32ac913701",
	        Name : "Admin",
	    },
		Assignee : Person{
	        UUID : "ce7405d8-3b78-4de7-8b58-6b32ac913701",
	        Name : "Admin",
	    },
		TaskHistorics : TaskHistorics{
			{
				Comment : "Primeiro comentário",
			},
			{
				Comment : "Primeiro comentário",
			},
		},
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
		Title : "Geral",
		Due : time.Now(),
		CompletedAt : &data,
		Person : Person{
	        UUID : "ce7405d8-3b78-4de7-8b58-6b32ac913701",
	        Name : "Admin",
	    },
		Assignee : Person{
	        UUID : "ce7405d8-3b78-4de7-8b58-6b32ac913701",
	        Name : "Admin",
	    },
		TaskHistorics : TaskHistorics{
			{
				Comment : "Primeiro comentário",
			},
			{
				Comment : "Primeiro comentário",
			},
		},
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
		Title : "Geral",
		Due : time.Now(),
		CompletedAt : &data,
		Category : TaskCategory{
	        UUID : "756524a2-9555-4ae5-9a6c-b2232de896af",
	        Description : "Geral",
	    },
		Assignee : Person{
	        UUID : "ce7405d8-3b78-4de7-8b58-6b32ac913701",
	        Name : "Admin",
	    },
		TaskHistorics : TaskHistorics{
			{
				Comment : "Primeiro comentário",
			},
			{
				Comment : "Primeiro comentário",
			},
		},
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
		Title : "Geral",
		Due : time.Now(),
		CompletedAt : &data,
		Category : TaskCategory{
	        UUID : "756524a2-9555-4ae5-9a6c-b2232de896af",
	        Description : "Geral",
	    },
	    Person : Person{
	        UUID : "ce7405d8-3b78-4de7-8b58-6b32ac913701",
	        Name : "Admin",
	    },
		TaskHistorics : TaskHistorics{
			{
				Comment : "Primeiro comentário",
			},
			{
				Comment : "Primeiro comentário",
			},
		},
	}

	errorValidate := task.Validate()

	if len(errorValidate) != amountErrorsExpected {
		t.Errorf("[%s] Quantidade de erros esperado %d, atual %d", task.Title, amountErrorsExpected, len(errorValidate))
		
		for _, err := range errorValidate {
			t.Errorf(err)
		}
	}
}