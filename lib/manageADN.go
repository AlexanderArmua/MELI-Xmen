package lib

import (
	"errors"
	"github.com/magiconair/properties"
	"strings"
	"sync"
)

type Props struct {
	sizeWord int        	`properties:"sizeWord,default=4"`
	minCountWords int 		`properties:"minCountWords,default=2"`
	acceptedChars string    `properties:"acceptedChars,layout=ATCG"`
}

var props Props

func init() {
	loadConfigFile()
}

func loadConfigFile() {
	p := properties.MustLoadFile("./mutantes.conf", properties.UTF8)
	props.sizeWord = p.GetInt("sizeWord", 4)
	props.minCountWords = p.GetInt("minCountWords", 2)
	props.acceptedChars = p.GetString("acceptedChars", "ATCG")
}

func IsMutant(matrix []string) (bool, error) {
	nMatrix := len(matrix)
	wordsFinded := 0

	for row := range matrix {
		if isInValidRow(matrix[row], props.sizeWord, nMatrix) {
			return false, errors.New("Tamaño De Matriz Inválido")
		}

		for col := range matrix {
			if isInvalidrChar(matrix[row][col:col+1]) {
				return false, errors.New("Caracter de ADN Inválido")
			}

			/**
				Al comenzar valida y busca palabras correctas.
				En el momento que las encuentra, sigue solo para validar toda la matris
			 */
			if wordsFinded < props.minCountWords {
				var wg sync.WaitGroup
				wg.Add(4)

				go func(){
					if nMatrix - col >= props.sizeWord {
						if SearchHorizontalWord(matrix, row, col, props.sizeWord) {
							wordsFinded++
						}
					}
					wg.Done()
				}()

				go func(){
					if nMatrix - row >= props.sizeWord {
						if SearchVerticalWord(matrix, row, col, props.sizeWord) {
							wordsFinded++
						}
					}
					wg.Done()
				}()

				go func() {
					if nMatrix - col >= props.sizeWord && nMatrix - row >= props.sizeWord {
						if SearchDiagonalDownWord(matrix, row, col, props.sizeWord) {
							wordsFinded++
						}
					}
					wg.Done()
				}()

				go func() {
					if nMatrix - col >= props.sizeWord && row + 1 >= props.sizeWord {
						if SearchDiagonalUpWord(matrix, row, col, props.sizeWord) {
							wordsFinded++
						}
					}
					wg.Done()
				}()

				wg.Wait()
			}
		}
	}

	return wordsFinded >= props.minCountWords, nil
}

func isInValidRow(row string, sizeWord, nMatrix int) bool {
	lenRow := len(row)
	return lenRow != nMatrix || lenRow < sizeWord
}

func isInvalidrChar(char string) bool {
	return !strings.Contains(props.acceptedChars, char)
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
			break
		}
	}

	return countChars == sizeWord
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
