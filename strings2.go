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

	input = strings.Replace(input, " ", "", -1)

	fmt.Println("Resultado:", input)

}
