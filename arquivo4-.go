package main

import (
	"errors"
	"fmt"
)

func classificaçãolista(prices []float64) (int, error) {
	n := len(prices)
	if n == 0 {
		return 0, errors.New("a lista está vazia")
	}
	if n == 1 {
		return 3, nil
	}
	Écrescente := true
	Édecrescente := true
	for i := 1; i < n; i++ {
		if prices[i] < prices[i-1] {
			Écrescente = false
		}
		if prices[i] > prices[i-1] {
			Édecrescente = false
		}
	}
	if Écrescente {
		return 1, nil
	}
	if Édecrescente {
		return 2, nil
	}
	return 3, nil
}
func main() {
	prices1 := []float64{5, 4, 3, 2, 1}
	result1, err1 := classificaçãolista(prices1)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(prices1, "=>", result1)
	}
}
