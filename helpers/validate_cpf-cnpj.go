package helpers

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var DEBBUG = false

var (
	ErrEmptyCPF           = errors.New("CPF não pode ser vázio")
	ErrInvalidCPF         = errors.New("CPF inválido")
	ErrInvalidCaracterCPF = errors.New("CPF com caracter inválido")

	ErrEmptyCNPJ           = errors.New("CNPJ não pode ser vázio")
	ErrInvalidCNPJ         = errors.New("CNPJ inválido")
	ErrInvalidCaracterCNPJ = errors.New("CNPJ com caracter inválido")
)

/*	@autor: Wilson T.J.

	Método responsável por validar um CPF

	CPF: string (14)
	Aceita: ###.###.###-## | ###########
*/
func ValidateCPF(cpf string) []error {

	var errors []error

	if DEBBUG == true {
		fmt.Println("CPF: ", cpf)
	}

	// Analise se foi passado um CPF
	if cpf != "" {
		var cpfValidate string

		// Remove pontos e traços
		cpfValidate = strings.Replace(cpf, ".", "", -1)
		cpfValidate = strings.Replace(cpfValidate, "-", "", -1)

		if DEBBUG == true {
			fmt.Println("Removido pontos: ", cpfValidate)
		}

		// Verifica se tem a quantidade certa de caracters
		if len(cpfValidate) == 11 {

			// Extrai os digitos do CPF já convertendo para inteiro
			num1, _ := strconv.Atoi(cpfValidate[0:1])
			num2, _ := strconv.Atoi(cpfValidate[1:2])
			num3, _ := strconv.Atoi(cpfValidate[2:3])
			num4, _ := strconv.Atoi(cpfValidate[3:4])
			num5, _ := strconv.Atoi(cpfValidate[4:5])
			num6, _ := strconv.Atoi(cpfValidate[5:6])
			num7, _ := strconv.Atoi(cpfValidate[6:7])
			num8, _ := strconv.Atoi(cpfValidate[7:8])
			num9, _ := strconv.Atoi(cpfValidate[8:9])
			num10, _ := strconv.Atoi(cpfValidate[9:10])
			num11, _ := strconv.Atoi(cpfValidate[10:11])

			// Valiação dos CPFs inválidos conhecidos
			if (num1 == num2) &&
				(num2 == num3) &&
				(num3 == num4) &&
				(num4 == num5) &&
				(num6 == num7) &&
				(num8 == num9) &&
				(num9 == num10) &&
				(num10 == num11) {

				errors = append(errors, ErrInvalidCPF)
			} else {
				// Soma os valores dos 9 primeiros número fazendo a multiplicação seguindo a regra
				sum1 := num1*10 + num2*9 + num3*8 + num4*7 + num5*6 + num6*5 + num7*4 + num8*3 + num9*2

				if DEBBUG == true {
					fmt.Println("Soma1: ", sum1)
				}

				// Faz a divisão da soma por 11 para pegar o resto
				rest1 := (sum1 * 10) % 11
				if DEBBUG == true {
					fmt.Println("Resto1: ", rest1)
				}

				if rest1 == 10 {
					rest1 = 0
				}

				// Soma os valores dos 9 primeiros número mais o primeiro digito verificado, fazendo a multiplicação seguindo a regra
				sum2 := num1*11 + num2*10 + num3*9 + num4*8 + num5*7 + num6*6 + num7*5 + num8*4 + num9*3 + num10*2
				if DEBBUG == true {
					fmt.Println("Soma2: ", sum2)
				}

				// Faz a divisão da soma por 11 para pegar o resto
				rest2 := (sum2 * 10) % 11
				if DEBBUG == true {
					fmt.Println("Resto2: ", rest2)
				}

				if rest2 == 10 {
					rest2 = 0
				}

				// Verifica se o valor do rest1 é igual ao penultimo digito do CPF
				// Verifica se o valor do rest2 é igual ao ultimo digito do CPF
				if (rest1 == num10) && (rest2 == num11) {
					if DEBBUG == true {
						fmt.Println("Passou na validação!")
					}
					return nil
				} else {
					errors = append(errors, ErrInvalidCPF)
				}
			}
		} else {
			errors = append(errors, ErrInvalidCaracterCPF)
		}
	} else {
		errors = append(errors, ErrEmptyCPF)
	}

	return errors
}

func ValidateCNPJ(cnpj string) []error {

	var errors []error

	if DEBBUG == true {
		fmt.Println("CNPJ: ", cnpj)
	}

	// Analise se foi passado um CNPJ
	if cnpj != "" {
		var cnpjValidate string

		// Remove pontos, traços e barras
		cnpjValidate = strings.Replace(cnpj, ".", "", -1)
		cnpjValidate = strings.Replace(cnpjValidate, "-", "", -1)
		cnpjValidate = strings.Replace(cnpjValidate, "/", "", -1)

		if DEBBUG == true {
			fmt.Println("Removido pontos: ", cnpjValidate)
		}

		// Verifica se tem a quantidade certa de caracters
		if len(cnpjValidate) == 14 {

			// Extrai os digitos do CNPJ já convertendo para inteiro
			num1, _ := strconv.Atoi(cnpjValidate[0:1])
			num2, _ := strconv.Atoi(cnpjValidate[1:2])
			num3, _ := strconv.Atoi(cnpjValidate[2:3])
			num4, _ := strconv.Atoi(cnpjValidate[3:4])
			num5, _ := strconv.Atoi(cnpjValidate[4:5])
			num6, _ := strconv.Atoi(cnpjValidate[5:6])
			num7, _ := strconv.Atoi(cnpjValidate[6:7])
			num8, _ := strconv.Atoi(cnpjValidate[7:8])
			num9, _ := strconv.Atoi(cnpjValidate[8:9])
			num10, _ := strconv.Atoi(cnpjValidate[9:10])
			num11, _ := strconv.Atoi(cnpjValidate[10:11])
			num12, _ := strconv.Atoi(cnpjValidate[11:12])
			num13, _ := strconv.Atoi(cnpjValidate[12:13])
			num14, _ := strconv.Atoi(cnpjValidate[13:14])

			// Valiação dos CNPJs inválidos conhecidos
			if (num1 == num2) &&
				(num2 == num3) &&
				(num3 == num4) &&
				(num4 == num5) &&
				(num5 == num6) &&
				(num6 == num7) &&
				(num7 == num8) &&
				(num8 == num9) &&
				(num9 == num10) &&
				(num10 == num11) &&
				(num11 == num12) &&
				(num12 == num13) &&
				(num13 == num14) {

				errors = append(errors, ErrInvalidCNPJ)
			} else {
				// Soma os valores dos 12 primeiros número fazendo a multiplicação seguindo a regra
				sum1 := num1*5 + num2*4 + num3*3 + num4*2 + num5*9 + num6*8 + num7*7 + num8*6 + num9*5 + num10*4 + num11*3 + num12*2

				if DEBBUG == true {
					fmt.Println("Soma1: ", sum1)
				}

				// Faz a divisão da soma por 11 para pegar o resto
				rest1 := sum1 % 11
				if DEBBUG == true {
					fmt.Println("Resto1: ", rest1)
				}

				if rest1 < 2 {
					rest1 = 0
				} else {
					rest1 = 11 - rest1
				}

				// Soma os valores dos 12 primeiros número mais o primeiro digito verificado, fazendo a multiplicação seguindo a regra
				sum2 := num1*6 + num2*5 + num3*4 + num4*3 + num5*2 + num6*9 + num7*8 + num8*7 + num9*6 + num10*5 + num11*4 + num12*3 + num13*2
				if DEBBUG == true {
					fmt.Println("Soma2: ", sum2)
				}

				// Faz a divisão da soma por 11 para pegar o resto
				rest2 := sum2 % 11
				if DEBBUG == true {
					fmt.Println("Resto2: ", rest2)
				}

				if rest2 < 2 {
					rest2 = 0
				} else {
					rest2 = 11 - rest2
				}

				// Verifica se o valor do rest1 é igual ao penultimo digito do CNPJ
				// Verifica se o valor do rest2 é igual ao ultimo digito do CNPJ
				if (rest1 == num13) && (rest2 == num14) {
					if DEBBUG == true {
						fmt.Println("Passou na validação!")
					}
					return nil
				} else {
					errors = append(errors, ErrInvalidCNPJ)
				}
			}
		} else {
			errors = append(errors, ErrInvalidCaracterCNPJ)
		}
	} else {
		errors = append(errors, ErrEmptyCNPJ)
	}

	return errors
}
