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

	// Remove espaços em branco extras da entrada do usuário
	input = strings.TrimSpace(input)

	// Conta o número de palavras na entrada do usuário
	words := strings.Fields(input)
	count := len(words)

	fmt.Println("A entrada contém palavra(s)", count)
}
