package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("DIGITE UMA FRASE:")
	fmt.Scanln(&input)

	// Deleta todas as vogais
	input = strings.ReplaceAll(input, "A", "")
	input = strings.ReplaceAll(input, "a", "")
	input = strings.ReplaceAll(input, "E", "")
	input = strings.ReplaceAll(input, "e", "")
	input = strings.ReplaceAll(input, "I", "")
	input = strings.ReplaceAll(input, "i", "")
	input = strings.ReplaceAll(input, "O", "")
	input = strings.ReplaceAll(input, "o", "")
	input = strings.ReplaceAll(input, "U", "")
	input = strings.ReplaceAll(input, "u", "")

	// Insere um caractere "." antes de cada consoante
	var result strings.Builder
	for _, char := range input {
		if isConsonant(char) {
			result.WriteString(".")
		}
		result.WriteRune(char)
	}

	// Substitui todas as consoantes maiúsculas pelas correspondentes em minúsculas
	output := strings.ToLower(result.String())

	fmt.Println(output)
}

func isConsonant(char rune) bool {
	return char >= 'A' && char <= 'Z' && char != 'A' && char != 'E' && char != 'I' && char != 'O' && char != 'U' ||
		char >= 'a' && char <= 'z' && char != 'a' && char != 'e' && char != 'i' && char != 'o' && char != 'u'
}
