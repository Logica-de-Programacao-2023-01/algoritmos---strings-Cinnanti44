package main

import (
	"fmt"
	"strings"
)

func main() {
	// Solicita ao usuário que insira uma string
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)

	// Substitui as vogais por '*'
	str = strings.ReplaceAll(str, "a", "*")
	str = strings.ReplaceAll(str, "e", "*")
	str = strings.ReplaceAll(str, "i", "*")
	str = strings.ReplaceAll(str, "o", "*")
	str = strings.ReplaceAll(str, "u", "*")

	// Imprime a string modificada
	fmt.Printf("A string modificada é: %s\n", str)
}
