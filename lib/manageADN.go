package lib

import (
	"errors"
	"strings"
	"sync"
)

func IsMutant(matrix []string, sizeWord int) (bool, error) {
	nMatrix := len(matrix)
	wordsFinded := 0

	for row := 0; row < nMatrix; row++ { // Arriba a Abajo
		if isInValidRow(matrix[row], sizeWord, nMatrix) {
			return false, errors.New("Tamaño De Matriz Inválido")
		}

		for col := 0; col < nMatrix; col++ { // Izquierda a Derecha
			if isInvalidrChar(matrix[row][col:col+1]) {
				return false, errors.New("Caracter de ADN Inválido")
			}

			var wg sync.WaitGroup
			wg.Add(4)

			go func(){
				if nMatrix - col >= sizeWord {
					if SearchHorizontalWord(matrix, row, col, sizeWord) {
						wordsFinded++
					}
				}
				wg.Done()
			}()

			go func(){
				if nMatrix - row >= sizeWord {
					if SearchVerticalWord(matrix, row, col, sizeWord) {
						wordsFinded++
					}
				}
				wg.Done()
			}()

			go func() {
				if nMatrix - col >= sizeWord && nMatrix - row >= sizeWord {
					if SearchDiagonalDownWord(matrix, row, col, sizeWord) {
						wordsFinded++
					}
				}
				wg.Done()
			}()

			go func() {
				if nMatrix - col >= sizeWord && row + 1 >= sizeWord {
					if SearchDiagonalUpWord(matrix, row, col, sizeWord) {
						wordsFinded++
					}
				}
				wg.Done()
			}()

			wg.Wait()

			// Mas de una secuencia de cuatro letras iguales
			if wordsFinded >= 2 {
				return true, nil
			}
		}
	}

	return false, nil
}

func isInValidRow(row string, sizeWord, nMatrix int) bool {
	lenRow := len(row)
	return lenRow != nMatrix || lenRow < sizeWord
}

func isInvalidrChar(char string) bool {
	return !strings.Contains("ATCG", char)
}

func SearchHorizontalWord(matrix []string, row, col int, sizeWord int) bool {
	return searchWord(matrix, row, col, sizeWord, nextHorizontalChar(matrix))
}

func SearchVerticalWord(matrix []string, row, col int, sizeWord int) bool {
	return searchWord(matrix, row, col, sizeWord, nextVerticalChar(matrix))
}

func SearchDiagonalDownWord(matrix []string, row, col int, sizeWord int) bool {
	return searchWord(matrix, row, col, sizeWord, nextDiagonalDownChar(matrix))
}

func SearchDiagonalUpWord(matrix []string, row, col int, sizeWord int) bool {
	return searchWord(matrix, row, col, sizeWord, nextDiagonalUpChar(matrix))
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
