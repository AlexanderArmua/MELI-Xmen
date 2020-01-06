package main

import (
	"./lib"
	"fmt"
)

func main() {
	mutante := []string {
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTG",
	}

	esMutante, error := lib.IsMutant(mutante, 4)

	if error == nil {
		fmt.Printf("Es mutante: %v\n", esMutante)
	} else {
		fmt.Println(error.Error())
	}
}


