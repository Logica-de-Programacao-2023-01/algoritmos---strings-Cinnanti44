package main

import (
	"errors"
	"fmt"
)

func maximoDominos(m, n int) (int, error) {
	if m <= 0 || n <= 0 {
		return 0, errors.New("o valor de M e N deve ser maior que zero")
	}

	// Calcula o número máximo de peças de dominó
	// que podem ser colocadas no tabuleiro
	max := (m * n) / 2

	return max, nil
}

func main() {
	// tabuleiro retangular de 3 x 3 //quadrados
	max1, err1 := maximoDominos(3, 3)
	if err1 != nil {
		fmt.Println("Erro:", err1)
	} else {
		fmt.Println("Máximo de peças de dominó:", max1)
	}
}
