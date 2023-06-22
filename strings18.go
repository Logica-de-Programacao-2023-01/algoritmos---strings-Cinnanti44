package main

import (
	"fmt"
	"unicode"
)

func main() {
	// Solicita ao usuário que insira uma string
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)

	// Verifica se a string contém somente números
	onlyNumbers := true
	for _, c := range str {
		if !unicode.IsDigit(c) {
			onlyNumbers = false
			break
		}
	}

	// Imprime o resultado
	if onlyNumbers {
		fmt.Println("A string contém somente números.")
	} else {
		fmt.Println("A string não contém somente números.")
	}
}
