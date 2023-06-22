package main

import "fmt"

func main() {
	problemas := [][]bool{
		{true, true, false},
		{true, false, false},
		{true, true, true},
		{false, false, true},
		{true, true, false},
		{true, true, false},
	}

	problemasresolvidos := 0
	for _, p := range problemas {
		numerosTrue := 0
		for _, opinion := range p {
			if opinion {
				numerosTrue++
			}
		}
		if numerosTrue >= 2 {
			problemasresolvidos++
		}
	}

	fmt.Println(problemasresolvidos)
}
