package main

import "fmt"

func main() {
	var input string
	var oldChar, newChar byte

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	fmt.Print("Digite o caractere que deseja substituir: ")
	fmt.Scanf("%c\n", &oldChar)

	fmt.Print("Digite o caractere de substituição: ")
	fmt.Scanf("%c\n", &newChar)

	// Converte a string em um slice de bytes
	bytes := []byte(input)

	// Substitui todas as ocorrências de oldChar por newChar
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == oldChar {
			bytes[i] = newChar
		}
	}

	// Converte o slice de bytes de volta para uma string
	output := string(bytes)

	fmt.Println("A string modificada é:", output)

}
