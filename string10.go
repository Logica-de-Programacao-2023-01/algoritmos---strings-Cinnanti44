package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var s1, s2 string

	fmt.Print("Digite a primeira string: ")
	fmt.Scanln(&s1)

	fmt.Print("Digite a segunda string: ")
	fmt.Scanln(&s2)

	// Converte as strings para minúsculas e remove espaços em branco
	s1 = strings.ToLower(strings.ReplaceAll(s1, " ", ""))
	s2 = strings.ToLower(strings.ReplaceAll(s2, " ", ""))

	// Ordena as strings
	s1 = sortString(s1)
	s2 = sortString(s2)

	// Verifica se as strings são iguais
	if s1 == s2 {
		fmt.Println("As strings são anagramas!")
	} else {
		fmt.Println("As strings não são anagramas.")
	}
}

// Função auxiliar que ordena uma string
func sortString(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")

}
