package lib

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"strings"
)

func GenerateHash(dna []string) uint32 {
	allRowAsOne := strings.Join(dna, "")
	h := fnv.New32a()
	h.Write([]byte(allRowAsOne))

	return h.Sum32()
}

func EqualString(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func ConverByteToResultado(mutantByte []byte) (Resultado, error) {
	res := &Resultado{}
	err := json.Unmarshal(mutantByte, &res)

	return *res, err
}

func ConvertResultadoToByte(dna []string, isMutant bool) []byte {
	result := Resultado{dna,isMutant}
	mutanteAsJson, _ := json.Marshal(result)

	return mutanteAsJson
}

func GetValueAsByte(hash interface{}) []byte{
	return []byte(fmt.Sprint(hash))
}