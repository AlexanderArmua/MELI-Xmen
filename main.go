package main

import (
	"./lib"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Persona struct {
	DNA []string `json:"dna" binding:"required"`
}

var stats *lib.Stats

func main() {
	r := gin.Default()

	r.POST("/mutant/", func(c *gin.Context) {
		var persona Persona
		c.BindJSON(&persona)

		if isMutant(persona.DNA) {
			c.Done()
		} else {
			c.String(http.StatusForbidden, "Forbidden")
		}
	})

	r.GET("/stats", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"count_mutant_dna": stats.CountMutantDna,
			"count_human_dna": stats.CountHumanDna,
			"ratio": stats.GetRatio(),
		})
	})

	r.Run()
}

func init() {
	generateFirstCache()
}

func generateFirstCache() {
	go func() {
		stats = lib.CalculateStats()
	}()
}

func isMutant(dna []string) bool {
	persona, error := lib.GetResultadoFromBD(dna)

	if error == nil {
		return persona.IsMutant
	}

	esMutante, error := lib.IsMutant(dna)

	if error != nil {
		return false
	}

	defer lib.SaveResult(dna, esMutante, stats)

	return esMutante
}


