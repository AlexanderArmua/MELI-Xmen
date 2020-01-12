package lib

import (
	"hash/fnv"
	"strings"
)

type Resultado struct {
	Dna      []string
	IsMutant bool
}

type Stats struct {
	CountMutantDna int
	CountHumanDna int
	Ratio float32
}

var resultados map[uint32]Resultado

func init() {
	resultados = make(map[uint32]Resultado)
}

func GetResultado(dna []string) (Resultado, bool) {
	hash := GenerateHash(dna)

	item, ok := resultados[hash]

	return item, ok
}

func SaveResult(dna []string, isMutant bool) {
	hash := GenerateHash(dna)
	resultados[hash] = Resultado{dna, isMutant}
}

func GenerateHash(dna []string) uint32 {
	allRowAsOne := strings.Join(dna, "")

	h := fnv.New32a()
	h.Write([]byte(allRowAsOne))

	return h.Sum32()
}

func GetStats() Stats {
	stats := Stats{0,0,0.0}

	for _, value := range resultados {
		if value.IsMutant {
			stats.CountMutantDna++
		} else {
			stats.CountHumanDna++
		}

		if stats.CountHumanDna == 0 {
			stats.Ratio = float32(stats.CountMutantDna) / 1
		} else {
			stats.Ratio = float32(stats.CountMutantDna) / float32(stats.CountHumanDna)
		}

	}

	return stats
}