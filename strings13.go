package main

import (
	"fmt"
	"strconv"
)

func isNumericSequence(str string) bool {
	// Verifica se a string contém apenas dígitos
	if _, err := strconv.Atoi(str); err != nil {
		return false
	}

	// Converte a string em uma lista de inteiros
	var nums []int
	for _, digit := range str {
		num, _ := strconv.Atoi(string(digit))
		nums = append(nums, num)
	}

	// Verifica se os números formam uma sequência crescente
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			return false
		}
	}
	return true
}

func main() {
	// Solicita ao usuário que insira uma string
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)

	// Verifica se a string é uma sequência numérica crescente
	if isNumericSequence(str) {
		fmt.Printf("A string '%s' é uma sequência numérica crescente.\n", str)
	} else {
		fmt.Printf("A string '%s' não é uma sequência numérica crescente.\n", str)
	}
}
