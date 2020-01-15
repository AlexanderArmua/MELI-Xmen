package lib

import (
	"testing"
)

func TestGetResultado(t *testing.T) {
	mutante := getADNMutanteDatabase()
	want := Resultado{mutante, true}

	stats := new(Stats)
	SaveResult(want.Dna, want.IsMutant, stats)

	t.Run("Find Horizontal Word", func(t *testing.T) {
		got, err := GetResultado(mutante)

		if err != nil || want.IsMutant != true || !EqualString(got.Dna, want.Dna) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func BenchmarkGetResultado(b *testing.B) {
	mutante := getADNMutanteDatabase()
	want := Resultado{mutante, true}

	stats := new(Stats)
	SaveResult(want.Dna, want.IsMutant, stats)

	b.Run("Bench Get Resultados ", func(b *testing.B) {
		got, err := GetResultado(mutante)

		if err != nil || want.IsMutant != true || !EqualString(got.Dna, want.Dna) {
			b.Errorf("got %v want %v", got, want)
		}
	})

	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				GetResultado(mutante)
			}
		},
	)
}

func TestCalculateStats(t *testing.T) {
	mutante := Resultado{getADNMutanteDatabase(), true}
	humano := Resultado{getADNHumanoDatabase(), false}

	stats := new(Stats)
	SaveResult(mutante.Dna, mutante.IsMutant, stats)
	SaveResult(humano.Dna, humano.IsMutant, stats)

	t.Run("Calculate Stats", func(t *testing.T) {
		got := CalculateStats()
		want := Stats{CountMutantDna: 1, CountHumanDna:  2,}

		if got.CountHumanDna != want.CountHumanDna || got.CountMutantDna != want.CountMutantDna {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func BenchmarkCalculateStats(b *testing.B) {
	mutante := Resultado{getADNMutanteDatabase(), true}
	humano := Resultado{getADNHumanoDatabase(), false}

	stats := new(Stats)
	SaveResult(mutante.Dna, mutante.IsMutant, stats)
	SaveResult(humano.Dna, humano.IsMutant, stats)

	b.Run("Bench Calculate Stats ", func(b *testing.B) {
		got := CalculateStats()
		want := Stats{CountMutantDna: 1, CountHumanDna:  2,}

		if got.CountHumanDna != want.CountHumanDna || got.CountMutantDna != want.CountMutantDna {
			b.Errorf("got %v want %v", got, want)
		}
	})

	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				CalculateStats()
			}
		},
	)

}

func getADNHumanoDatabase() []string {
	return []string {
		"ATGCGA",
		"CAGTGC",
		"TTATTT",
		"AGACGG",
		"GCGTCA",
		"TCACTG",
	}
}

func getADNMutanteDatabase() []string {
	return []string {
		"ATGCGA",
		"CAGTGC",
		"TTTTGT",
		"AGTAGG",
		"CCCCTA",
		"TCACTG",
	}
}