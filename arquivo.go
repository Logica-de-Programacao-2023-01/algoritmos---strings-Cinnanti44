package main

import (
	"errors"
	"fmt"
)

func main() {
	var pesodaMelancia int
	fmt.Print("digite o peso da melancia:")
	fmt.Scanln(&pesodaMelancia)

	podeDividir, err := dividirMelancia(pesodaMelancia)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("É possível:", podeDividir)
	}
}
func dividirMelancia(peso int) (bool, error) {
	if peso <= 0 {
		return false, errors.New("Peso da melancia deve ser maior que 0")
	}

	if peso%2 == 0 && peso > 2 {
		return true, nil
	}

	return false, nil
}
