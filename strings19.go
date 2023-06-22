package main

import "fmt"

func main() {
	// Solicita ao usuário que insira uma string
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)

	// Inverte a string
	reversed := ""
	for _, c := range str {
		reversed = string(c) + reversed
	}

	// Imprime a string invertida
	fmt.Println("String invertida:", reversed)
}
