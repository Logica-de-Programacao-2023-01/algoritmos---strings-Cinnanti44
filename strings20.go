package main

import (
	"fmt"
	"unicode"
)

func main() {
	// Solicita ao usuário que insira uma string
	var str string
	fmt.Print("Digite uma string em camelCase: ")
	fmt.Scanln(&str)

	// Verifica se a string está em camelCase
	isCamelCase := true
	for i, c := range str {
		if i == 0 {
			if !unicode.IsUpper(c) {
				isCamelCase = false
				break
			}
		} else if unicode.IsUpper(c) {
			if i == len(str)-1 || unicode.IsUpper(rune(str[i+1])) {
				isCamelCase = false
				break
			}
		}
	}

	// Conta o número de palavras na string
	numPalavras := 1
	for _, c := range str {
		if unicode.IsUpper(c) {
			numPalavras++
		}
	}

	// Imprime o resultado
	if isCamelCase {
		fmt.Println("A string está em camelCase e possui", numPalavras, "palavras.")
	} else {
		fmt.Println("A string não está em camelCase.")
	}
}
