package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Print("digite:")
	fmt.Scanln(&input)

	// Deleta todas as vogais
	vowels := "AEIOUaeiou"
	for _, vowel := range vowels {
		input = strings.ReplaceAll(input, string(vowel), "")
	}

	// Insere um caractere "" antes de cada consoante
	var output string
	for _, char := range input {
		if strings.ContainsAny(string(char), "BCDFGHJKLMNPQRSTVWXYZbcdfghjklmnpqrstvwxyz") {
			output += "." + string(char)
		}
	}

	// Substitui todas as consoantes maiúsculas pelas correspondentes em minúsculas
	output = strings.ToLower(output)

	fmt.Println(output)
}
