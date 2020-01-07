package main

import (
	"github.com/AlexanderArmua/EjercicioMutantesMELI/lib"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Persona struct {
	DNA []string `json:"dna" binding:"required"`
}

func main() {
	// Default returns an Engine instance with the Logger and Recovery middleware already attached
	r := gin.Default()

	// GET is a shortcut for router.Handle("GET", path, handle)
	r.POST("/mutant/", func(c *gin.Context) {
		var persona Persona
		c.BindJSON(&persona)

		esMutante, error := lib.IsMutant(persona.DNA, 4)

		if error == nil && esMutante{
			c.Done()
		} else {
			c.String(http.StatusForbidden, "Forbidden")
		}
	})

	r.GET("/stats", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"count_mutant_dna": 40,
			"count_human_dna": 100,
			"ratio": 0.4,
		})
	})

	r.Run()
}


