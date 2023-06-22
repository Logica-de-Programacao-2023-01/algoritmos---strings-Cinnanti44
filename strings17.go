package main

import "fmt"

func main() {
	// Solicita ao usuário que insira uma string
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)

	// Cria um mapa para armazenar a contagem de cada letra na string
	counts := make(map[rune]int)
	for _, c := range str {
		counts[c]++
	}

	// Imprime as letras únicas na ordem em que aparecem na string
	fmt.Println("Letras únicas na string:")
	for _, c := range str {
		if counts[c] == 1 {
			fmt.Printf("%c ", c)
		}
	}
	fmt.Println()
}
