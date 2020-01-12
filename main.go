package main

import (
	"./lib"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Persona struct {
	DNA []string `json:"dna" binding:"required"`
}

var stats lib.Stats

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
		//stats = lib.GetStats()
		c.JSON(http.StatusOK, gin.H{
			"count_mutant_dna": stats.CountMutantDna,
			"count_human_dna": stats.CountHumanDna,
			"ratio": stats.Ratio,
		})
	})

	r.Run()
}

func init() {
	generateCacheStatsEvery5Secs()
}

func generateCacheStatsEvery5Secs() {
	go func() {
		for {
			stats = lib.GetStats()
			time.Sleep(5 * time.Second)
		}
	}()
}

func isMutant(dna []string) bool {
	item, ok := lib.GetResultado(dna)

	if ok {
		return item.IsMutant
	}

	esMutante, error := lib.IsMutant(dna, 4)

	if error != nil {
		return false
	}

	defer lib.SaveResult(dna, esMutante)

	return esMutante
}


