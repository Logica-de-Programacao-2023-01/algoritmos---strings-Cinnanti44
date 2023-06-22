package main

import "fmt"

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	// Converte a string em um slice de bytes
	bytes := []byte(input)

	// Inverte a ordem dos bytes
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}

	// Converte o slice de bytes de volta para uma string
	reversed := string(bytes)

	fmt.Println("A string invertida é:", reversed)

}
