package lib

import (
	"github.com/prologic/bitcask"
	"github.com/spf13/cast"
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
	inicializarDatabase()
}

func inicializarDatabase() {
	resultados = make(map[uint32]Resultado)
}

//func GetResultado(dna []string) (Resultado, bool) {
//	hash := GenerateHash(dna)
//
//	item, ok := resultados[hash]
//
//	return item, ok
//}

func GetResultado(dna []string) (bool, error) {
	hash := GenerateHash(dna)

	db, _ := bitcask.Open("./databases/XMenDatabase")
	defer db.Close()

	//db.Put([]byte("hash"), []byte("1/[ATGCG,ASDQWE,QWEQWEQWEQWEQWEQWEQDSFDS]"))

	isMutant, error := db.Get([]byte(cast.ToString(hash)))

	//item, ok := resultados[hash]

	return cast.ToBool(isMutant), error
}

//func SaveResult(dna []string, isMutant bool) {
//	hash := GenerateHash(dna)
//	resultados[hash] = Rfesultado{dna, isMutant}
//}

func SaveResult(dna []string, isMutant bool) {
	hash := GenerateHash(dna)

	db, _ := bitcask.Open("./databases/XMenDatabase")
	defer db.Close()

	db.Put([]byte(cast.ToString(hash)), []byte(cast.ToString(isMutant)))
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