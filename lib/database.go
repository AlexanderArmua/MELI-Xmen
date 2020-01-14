package lib

import (
	"fmt"
	"github.com/prologic/bitcask"
	"hash/fnv"
	"strconv"
	"strings"
)

type Resultado struct {
	Dna      []string
	IsMutant bool
}

type Stats struct {
	CountMutantDna float32
	CountHumanDna  float32
}

func (stats Stats) GetRatio() float32 {
	return stats.CountMutantDna / stats.CountHumanDna
}

var nombreDB = "./databases/XMenDatabase"

func GetResultado(dna []string) (bool, error) {
	hash := GenerateHash(dna)

	isMutant, error := getIsMutantFromBD(hash)

	return isMutant, error
}

func SaveResult(dna []string, isMutant bool) {
	hash := GenerateHash(dna)

	database := getDatabase()
	database.Put(getValueAsByte(hash), getValueAsByte(isMutant))

	defer database.Close()
}

func CalculateStats() Stats {
	stats := Stats{0,0}

	database := getDatabase()
	defer database.Close()

	chKeys := database.Keys()

	for key := range chKeys {
		value, _ := database.Get(key)

		isMutant, _ := strconv.ParseBool(string(value))

		addNewStat(isMutant, &stats)
	}


	return stats
}

func GenerateHash(dna []string) uint32 {
	allRowAsOne := strings.Join(dna, "")
	h := fnv.New32a()
	h.Write([]byte(allRowAsOne))

	return h.Sum32()
}

func getIsMutantFromBD(hash uint32) (bool, error) {
	database := getDatabase()

	mutantByte, error := database.Get(getValueAsByte(hash))
	defer database.Close()

	isMutante, _ := strconv.ParseBool(string(mutantByte))

	return isMutante, error
}

func getDatabase() *bitcask.Bitcask {
	database, _ := bitcask.Open(nombreDB)

	return database
}

func getValueAsByte(hash interface{}) []byte{
	return []byte(fmt.Sprint(hash))
}

func addNewStat(isMutant bool, stats *Stats) {
	stats.CountHumanDna++
	if isMutant {
		stats.CountMutantDna++
	}
}
