package models

import(
	"testing"
	"net/url"
)

var TaskCategoryMock TaskCategory
var TaskCategoryRepositoryMock TaskCategoryRepository

func init() {
	TaskCategoryRepositoryMock = NewTaskCategoryRepository()
	TaskCategoryMock = TaskCategory{
		Description : "Categoria 1",
	}
}

func TestTarefaCategoriaValida(t *testing.T) {

	amountErrorsExpected := 0

	category := TaskCategory{
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

	category := TaskCategory{}

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

	category := TaskCategory{
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

	category := TaskCategory{
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

func TestCadastrarNovaCategoriaDeTarefa(t *testing.T) {
	category := TaskCategoryMock
	expectedAmountItems := TaskCategoryRepositoryMock.CountRows() + 1

	err := TaskCategoryRepositoryMock.Create(&category)
	if err != nil {
		t.Errorf("Não era esperado o erro: %s", err)
	}

	if category.UUID == "" {
		t.Error("Era esperado o ID da categoria")	
	}

	amountItems := TaskCategoryRepositoryMock.CountRows()
	if amountItems < expectedAmountItems {
		t.Errorf("Quantidade de registros esperados %d, recebido %d", expectedAmountItems, amountItems)	
	}
}

func TestBuscaTodasCategoriaDeTarefa(t *testing.T) {
	params := url.Values{}
	params.Add("page", "1")
	params.Add("per_page", "50")

	categories, _ := TaskCategoryRepositoryMock.GetAll(params)
	if len(categories) <= 0 {
		t.Errorf("Esperado mais de %d, recebido %d", 0, len(categories))
	}
}

func TestBuscaDeCategoriaDeTarefaPorId(t *testing.T) {
	category := TaskCategoryMock

	err1 := TaskCategoryRepositoryMock.Create(&category)
	if err1 != nil {
		t.Errorf("Não era esperado o erro: %s", err1)
	}

	result := TaskCategoryRepositoryMock.Get(category.UUID)
	if result != category {
		t.Errorf("Esperado %s, recebido %s", category, result)
	}
}

func TestAlterarCategoriaDeTarefa(t *testing.T) {
	category := TaskCategoryMock

	err1 := TaskCategoryRepositoryMock.Create(&category)
	if err1 != nil {
		t.Errorf("Não era esperado o erro: %s", err1)
	}

	category.Description = "Alterada a Categoria"

	err2 := TaskCategoryRepositoryMock.Update(&category)
	if err2 != nil {
		t.Errorf("Não era esperado o erro: %s", err1)
	}
}

func TestExclusaoCategoriaDeTarefa(t *testing.T) {
	category := TaskCategoryMock
	expectedAmountItems := TaskCategoryRepositoryMock.CountRows()

	err1 := TaskCategoryRepositoryMock.Create(&category)
	if err1 != nil {
		t.Errorf("Não era esperado o erro: %s", err1)
	}

	err2 := TaskCategoryRepositoryMock.Delete(category.UUID)
	if err2 != nil {
		t.Errorf("Não era esperado o erro: %s", err1)
	}

	amountItems := TaskCategoryRepositoryMock.CountRows()
	if amountItems != expectedAmountItems {
		t.Error("Quantidade de registros esperados %s, recebido %s", expectedAmountItems, amountItems)	
	}
}