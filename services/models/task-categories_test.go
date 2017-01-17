package models_test

import(
	"testing"
	"github.com/wilsontamarozzi/panda-api/services/models"
)

func TestTarefaCategoriaValida(t *testing.T) {

	amountErrorsExpected := 0

	category := models.TaskCategory{
		Description : "Geral",
	}

	errorValidate := category.Validate()

	if len(errorValidate) != amountErrorsExpected {
		t.Errorf("[%s] Quantidade de erros esperado %d, atual %d", category.Description, amountErrorsExpected, len(errorValidate))
		
		for _, err := range errorValidate {
			t.Errorf(err)
		}
	}
}

func TestTarefaCategoriaSemCampoDescricao(t *testing.T) {

	// Erros esperados
	// - Descrição é obrigatório
	amountErrorsExpected := 1

	category := models.TaskCategory{}

	errorValidate := category.Validate()

	if len(errorValidate) != amountErrorsExpected {
		t.Errorf("[%s] Quantidade de erros esperado %d, atual %d", category.Description, amountErrorsExpected, len(errorValidate))
		
		for _, err := range errorValidate {
			t.Errorf(err)
		}
	}
}

func TestTarefaCategoriaTamanhoMaximoDosCampos(t *testing.T) {

	// Erros esperados
	// - Descrição deve ter minimo 2 e maximo 25 caracter
	amountErrorsExpected := 1

	category := models.TaskCategory{
		Description : "AAAAAAAAAAAAAAAAAAAAAAAAAA",
	}

	errorValidate := category.Validate()

	if len(errorValidate) != amountErrorsExpected {
		t.Errorf("[%s] Quantidade de erros esperado %d, atual %d", category.Description, amountErrorsExpected, len(errorValidate))
		
		for _, err := range errorValidate {
			t.Errorf(err)
		}
	}
}

func TestTarefaCategoriaTamanhoMinimoDosCampos(t *testing.T) {

	// Erros esperados
	// - Descrição deve ter minimo 2 e maximo 25 caracter
	amountErrorsExpected := 1

	category := models.TaskCategory{
		Description : "A",
	}

	errorValidate := category.Validate()

	if len(errorValidate) != amountErrorsExpected {
		t.Errorf("[%s] Quantidade de erros esperado %d, atual %d", category.Description, amountErrorsExpected, len(errorValidate))
		
		for _, err := range errorValidate {
			t.Errorf(err)
		}
	}
}