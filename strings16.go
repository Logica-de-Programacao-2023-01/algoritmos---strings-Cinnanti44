package main

import (
	"fmt"
	"strings"
)

func main() {
	// Solicita ao usuário que insira duas strings
	var str1, str2 string
	fmt.Print("Digite a primeira string: ")
	fmt.Scanln(&str1)
	fmt.Print("Digite a segunda string: ")
	fmt.Scanln(&str2)

	// Verifica se a segunda string é uma substring da primeira
	if strings.Contains(str1, str2) {
		fmt.Printf("'%s' é uma substring de '%s'\n", str2, str1)
	} else {
		fmt.Printf("'%s' não é uma substring de '%s'\n", str2, str1)
	}
}
