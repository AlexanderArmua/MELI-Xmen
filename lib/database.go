package lib

import (
	"errors"
	"fmt"
	"github.com/prologic/bitcask"
)

type Resultado struct {
	Dna      []string `json:"Dna"`
	IsMutant bool	  `json:"IsMutant"`
}

type Stats struct {
	CountMutantDna float32
	CountHumanDna  float32
}

func (stats *Stats) GetRatio() float32 {
	if stats.CountHumanDna == 0 {
		return 0
	}
	return stats.CountMutantDna / stats.CountHumanDna
}

func (stats *Stats) AddNewStat(isMutant bool) {
	stats.CountHumanDna++
	if isMutant {
		stats.CountMutantDna++
	}
}

const nombreDB = "./databases/XMenDatabase"

func GetResultado(dna []string) (Resultado, error) {
	hash := GenerateHash(dna)

	resultado, error := GetIsMutantFromBD(hash)

	// El algoritmo de Hasheo puede tener errores, con esta validacion aseguramos que sea el correcto.
	if EqualString(resultado.Dna, dna) {
		return resultado, error
	}

	return Resultado{}, errors.New("Crasheo de Hash")
}

func GetIsMutantFromBD(hash uint32) (Resultado, error) {
	database, errorDB := getDatabase()

	mutantByte, errorGet := database.Get(GetValueAsByte(hash))
	defer database.Close()

	// En caso de cualquier error o que no encuentre el dato, no explota y puede continuar
	if errorDB != nil || mutantByte == nil || errorGet != nil {
		return Resultado{}, errorGet
	}

	result, errorParse := ConverByteToResultado(mutantByte)

	return result, errorParse
}

func SaveResult(dna []string, isMutant bool, stats *Stats) {
	hash := GenerateHash(dna)

	database, _ := getDatabase()
	defer database.Close()

	database.Put(GetValueAsByte(hash), ConvertResultadoToByte(dna, isMutant))

	stats.AddNewStat(isMutant)
}

func getDatabase() (*bitcask.Bitcask, error) {
	database, error := bitcask.Open(nombreDB)
	if error != nil {
		fmt.Printf("No se pudo abrir la database - Error: %v.\n", error)
	}
	return database, error
}

func CalculateStats() *Stats {
	stats := new(Stats)

	database, _ := getDatabase()
	defer database.Close()

	chKeys := database.Keys()

	for key := range chKeys {
		mutantByte, _ := database.Get(key)

		mutante, _ := ConverByteToResultado(mutantByte)

		stats.AddNewStat(mutante.IsMutant)
	}

	return stats
}

