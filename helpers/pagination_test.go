package helpers

import (
	"testing"
)

var (
	TEST_AMOUNT_ITEMS  = 5
	TEST_ITEM_PER_PAGE = 2
)

func TestMakePaginationPage1(t *testing.T) {
	currentPage := 1
	pagination := MakePagination(TEST_AMOUNT_ITEMS, currentPage, TEST_ITEM_PER_PAGE)

	if pagination.StartIndex != 0 {
		t.Errorf("Esperado %s", 0)
		t.Errorf("Recebido %s", pagination.StartIndex)
	}

	if pagination.TotalPages != 3 {
		t.Errorf("Esperado %s", 3)
		t.Errorf("Recebido %s", pagination.TotalPages)
	}
}

func TestMakePaginationPage2(t *testing.T) {
	currentPage := 2
	pagination := MakePagination(TEST_AMOUNT_ITEMS, currentPage, TEST_ITEM_PER_PAGE)

	if pagination.StartIndex != 2 {
		t.Errorf("Esperado %s", 2)
		t.Errorf("Recebido %s", pagination.StartIndex)
	}

	if pagination.TotalPages != 3 {
		t.Errorf("Esperado %s", 3)
		t.Errorf("Recebido %s", pagination.TotalPages)
	}
}
