package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Digite uma string: ")
	input, _ := reader.ReadString('\n')

	fmt.Print("Digite o caractere a ser substituído: ")
	oldChar, _ := reader.ReadString('\n')
	oldChar = strings.TrimSpace(oldChar)

	fmt.Print("Digite o caractere de substituição: ")
	newChar, _ := reader.ReadString('\n')
	newChar = strings.TrimSpace(newChar)

	input = strings.ReplaceAll(input, oldChar, newChar)

	fmt.Println("Resultado:", input)

}
