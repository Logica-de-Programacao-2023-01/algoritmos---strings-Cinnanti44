package main

import (
	"fmt"
)

func main() {
	// Solicita ao usuário que insira uma string
	var str string
	fmt.Print("Digite uma string numérica: ")
	fmt.Scanln(&str)

	// Verifica se a string é uma sequência numérica decrescente
	isDescending := true
	for i := 1; i < len(str); i++ {
		if str[i-1] <= str[i] {
			isDescending = false
			break
		}
	}

	// Imprime o resultado
	if isDescending {
		fmt.Println("A string é uma sequência numérica decrescente.")
	} else {
		fmt.Println("A string não é uma sequência numérica decrescente.")
	}
}
