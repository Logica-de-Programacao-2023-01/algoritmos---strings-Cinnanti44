package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	// Remove todas as vogais da string
	output := removeVowels(input)

	fmt.Println("A string sem vogais é:", output)
}

// Função auxiliar que remove todas as vogais de uma string
func removeVowels(s string) string {
	vowels := "aeiouAEIOU"
	result := ""

	for _, c := range s {
		if !strings.ContainsAny(string(c), vowels) {
			result += string(c)
		}
	}

	return result

}
