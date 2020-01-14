package lib

import (
	"encoding/json"
	"fmt"
	"github.com/prologic/bitcask"
	"hash/fnv"
	"strings"
)

type Resultado struct {
	Dna      []string `json:"Dna"`
	IsMutant bool	  `json:"IsMutant"`
}

type Stats struct {
	CountMutantDna float32
	CountHumanDna  float32
}

func (stats Stats) GetRatio() float32 {
	return stats.CountMutantDna / stats.CountHumanDna
}

const nombreDB = "./databases/XMenDatabase"

func GetResultado(dna []string) (Resultado, error) {
	hash := GenerateHash(dna)

	resultado, error := GetIsMutantFromBD(hash)

	return resultado, error
}

func GetIsMutantFromBD(hash uint32) (Resultado, error) {
	database, errorDB := getDatabase()

	mutantByte, errorGet := database.Get(getValueAsByte(hash))
	defer database.Close()

	// En caso de cualquier error o que no encuentre el dato, no explota y puede continuar
	if errorDB != nil || mutantByte == nil || errorGet != nil {
		return Resultado{}, errorGet
	}

	result, errorParse := converByteToResultado(mutantByte)

	return result, errorParse
}

func converByteToResultado(mutantByte []byte) (Resultado, error) {
	res := &Resultado{}
	err := json.Unmarshal(mutantByte, &res)

	return *res, err
}

func convertResultadoToByte(dna []string, isMutant bool) []byte {
	result := Resultado{dna,isMutant}
	mutanteAsJson, _ := json.Marshal(result)

	return mutanteAsJson
}

func SaveResult(dna []string, isMutant bool) {
	hash := GenerateHash(dna)

	database, _ := getDatabase()
	defer database.Close()

	database.Put(getValueAsByte(hash), convertResultadoToByte(dna, isMutant))
}

func GenerateHash(dna []string) uint32 {
	allRowAsOne := strings.Join(dna, "")
	h := fnv.New32a()
	h.Write([]byte(allRowAsOne))

	return h.Sum32()
}

func getDatabase() (*bitcask.Bitcask, error) {
	database, error := bitcask.Open(nombreDB)
	if error != nil {
		fmt.Printf("No se pudo abrir la database - Error: %v.\n", error)
	}
	return database, error
}

func getValueAsByte(hash interface{}) []byte{
	return []byte(fmt.Sprint(hash))
}

func CalculateStats() Stats {
	stats := Stats{0,0}

	database, _ := getDatabase()
	defer database.Close()

	chKeys := database.Keys()

	for key := range chKeys {
		mutantByte, _ := database.Get(key)

		mutante, _ := converByteToResultado(mutantByte)

		addNewStat(mutante.IsMutant, &stats)
	}

	return stats
}

func addNewStat(isMutant bool, stats *Stats) {
	stats.CountHumanDna++
	if isMutant {
		stats.CountMutantDna++
	}
}
