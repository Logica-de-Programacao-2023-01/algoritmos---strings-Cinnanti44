package main

import (
	"fmt"
	"strings"
)

func isPalindrome(str string) bool {
	// Remove espaços em branco e converte para minúsculas
	str = strings.ToLower(strings.ReplaceAll(str, " ", ""))

	// Verifica se a string é igual à sua versão invertida
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	// Solicita ao usuário que insira uma string
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)

	// Verifica se a string é um palíndromo
	if isPalindrome(str) {
		fmt.Printf("A string '%s' é um palíndromo.\n", str)
	} else {
		fmt.Printf("A string '%s' não é um palíndromo.\n", str)
	}
}
