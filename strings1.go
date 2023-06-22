package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	output := strings.ToUpper(input)

	fmt.Println("A string em maiúsculas é:", output)
}
