package helpers_test

import (
	"testing"
	"github.com/wilsontamarozzi/panda-api/helpers"
)

type CPF struct {
	Cpf string
	Expected bool
}

type CNPJ struct {
	Cnpj string
	Expected bool
}

func TestValidateCPF(t *testing.T) {
	items := []CPF{
		{Cpf: "41678171875", Expected: true},
		{Cpf: "416.781.718-75", Expected: true},
		{Cpf: "22222222222", Expected: false},
		{Cpf: "222.222.222-22", Expected: false},
		{Cpf: "XXXXXXXXXXX", Expected: false},
		{Cpf: "XXX.XXX.XXX-XX", Expected: false},
		{Cpf: "416.X81.718-75", Expected: false},
		{Cpf: "123456789123456", Expected: false},
		{Cpf: "", Expected: false},
	}

	for _, item := range items {
		actual := helpers.ValidateCPF(item.Cpf) == nil

		if actual != item.Expected {
			t.Errorf("CPF(%s) esperado %t, atual %t", item.Cpf, item.Expected, actual)
		}
	}
}

func TestValidateCNPJ(t *testing.T) {
	items := []CNPJ{
		{Cnpj: "12345678000195", Expected: true},
		{Cnpj: "12.345.678/0001-95", Expected: true},
		{Cnpj: "11111111111111", Expected: false},
		{Cnpj: "11.111.111/1111-11", Expected: false},
		{Cnpj: "123456780001950", Expected: false},
		{Cnpj: "12.345.678/0001-950", Expected: false},
		{Cnpj: "12345678912345X", Expected: false},
		{Cnpj: "12.345.678/9123-45X", Expected: false},
		{Cnpj: "XXXXXXXXXXXXXX", Expected: false},
		{Cnpj: "XX.XXX.XXX/XXXX-XX", Expected: false},
	}

	for _, item := range items {
		actual := helpers.ValidateCNPJ(item.Cnpj) == nil

		if actual != item.Expected {
			t.Errorf("CNPJ(%s) esperado %t, atual %t", item.Cnpj, item.Expected, actual)
		}
	}
}