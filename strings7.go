package main

import (
	"fmt"
	"unicode"
)

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	containsNumber := false
	for _, char := range input {
		if unicode.IsDigit(char) {
			containsNumber = true
			break
		}
	}

	if containsNumber {
		fmt.Println("A string contém pelo menos um número.")
	} else {
		fmt.Println("A string não contém nenhum número.")
	}
}
