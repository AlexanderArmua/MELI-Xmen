package main

import (
	"fmt"
)

func main() {
	humano := []string {
		"ATGCGA",
		"CAGTGC",
		"TTATTT",
		"AGACGG",
		"GCGTCA",
		"TCACTG",
	}

	mutante := []string {
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTG",
	}
	
	fmt.Printf("Humano es mutante: %v\n", isMutant(humano))

	fmt.Printf("Mutante es mutante: %v\n", isMutant(mutante))
}

func isMutant(matrix []string) bool {
	nMatrix := len(matrix)
	sizeWord := 4
	palabrasEncontradas := 0

	for row := 0; row < nMatrix; row++ { // Arriba a Abajo
		for col := 0; col < nMatrix; col++ { // Izquierda a Derecha
			if isInValidRow(matrix[row], sizeWord, nMatrix) {
				panic("El tamaño de la matris es incorrecto.")
			}

			if nMatrix - col >= sizeWord {
				if searchWord(matrix, row, col, sizeWord, nextHorizontalChar(matrix)) {
					palabrasEncontradas++
				}
			}

			if nMatrix - row >= sizeWord {
				if searchWord(matrix, row, col, sizeWord, nextVerticalChar(matrix)) {
					palabrasEncontradas++
				}
			}

			if nMatrix - col >= sizeWord && nMatrix - row >= sizeWord {
				if searchWord(matrix, row, col, sizeWord, nextDiagonalDownChar(matrix)) {
					palabrasEncontradas++
				}
			}

			if nMatrix - col >= sizeWord && row + 1 >= sizeWord {
				if searchWord(matrix, row, col, sizeWord, nextDiagonalUpChar(matrix)) {
					palabrasEncontradas++
				}
			}

			if palabrasEncontradas >= 2 {
				return true
			}
		}
	}

	return false
}

func isInValidRow(row string, sizeWord, nMatrix int) bool {
	lenRow := len(row)
	return lenRow != nMatrix || lenRow < sizeWord
}


func searchWord(matrix []string, row, col int, sizeWord int, nextChar func(int, int, int) string) bool {
	firstLetter := matrix[row][col: col + 1]
	countChars := 1

	for i := 1; i < sizeWord; i++ {
		if nextChar(row, col, i) == firstLetter{
			countChars++
		} else {
			// Se encontro un caracter que cortaba la cadena, no se alcanzo la palabra.
			break
		}

		if countChars == sizeWord {
			return true
		}
	}

	return false
}

func nextHorizontalChar(matrix []string) func(int, int, int) string {
	return func(row, col, index int) string {
		return matrix[row][col + index: col + index + 1]
	}
}

func nextVerticalChar(matrix []string) func(int, int, int) string {
	return func(row, col, index int) string {
		return matrix[row + index][col: col + 1]
	}
}

func nextDiagonalDownChar(matrix []string) func(int, int, int) string {
	return func(row, col, index int) string {
		return matrix[row + index][col + index: col + index + 1]
	}
}

func nextDiagonalUpChar(matrix []string) func(int, int, int) string {
	return func(row, col, index int) string {
		return matrix[row - index][col + index: col + index + 1]
	}
}
