package lib

import (
	"fmt"
	"testing"
)

func TestSearchHorizontalWord(t *testing.T) {
	humano := getADNHumano()
	mutante := getADNMutante()

	t.Run("Find Horizontal Word", func(t *testing.T) {
		sizeWord := 4
		got := SearchHorizontalWord(mutante, 5, 0, sizeWord)
		want := true

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Don't Find Horizontal Word", func(t *testing.T) {
		sizeWord := 4
		got := SearchHorizontalWord(humano, 5, 0, sizeWord)
		want := false

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSearchVerticalWord(t *testing.T) {
	humano := getADNHumano()
	mutante := getADNMutante()

	t.Run("Find Vertical Word", func(t *testing.T) {
		sizeWord := 4
		got := SearchVerticalWord(mutante, 0, 4, sizeWord)
		want := true

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Don't Find Vertical Word", func(t *testing.T) {
		sizeWord := 4
		got := SearchVerticalWord(humano, 4, 0, sizeWord)
		want := false

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSearchDiagonalDownWord(t *testing.T) {
	humano := getADNHumano()
	mutante := getADNMutante()

	t.Run("Find Diagonal Down Word", func(t *testing.T) {
		sizeWord := 4
		got := SearchDiagonalDownWord(mutante, 0, 0, sizeWord)
		want := true

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Don't Find Diagonal Down Word", func(t *testing.T) {
		sizeWord := 4
		got := SearchDiagonalDownWord(humano, 0, 0, sizeWord)
		want := false

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSearchDiagonalUpWord(t *testing.T) {
	humano := getADNHumano()
	mutante := getADNMutante()

	t.Run("Find Diagonal Up Word", func(t *testing.T) {
		sizeWord := 4
		got := SearchDiagonalUpWord(mutante, 5, 0, sizeWord)
		want := true

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Don't Find Diagonal Up Word", func(t *testing.T) {
		sizeWord := 4
		got := SearchDiagonalUpWord(humano, 5, 0, sizeWord)
		want := false

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestIsMutant(t *testing.T) {
	humano := getADNHumano()
	mutante := getADNMutante()

	t.Run("Is Mutant", func(t *testing.T) {
		got, error := IsMutant(mutante)
		want := true

		if error == nil {
			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		} else {
			t.Errorf(error.Error())
		}
	})

	t.Run("Is Human", func(t *testing.T) {
		got, error := IsMutant(humano)
		want := false

		if error == nil {
			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		} else {
			t.Errorf(error.Error())
		}
	})
}

func BenchmarkIsMutant(b *testing.B) {
	humano := getADNHumano()
	mutante := getADNMutante()

	b.Run(fmt.Sprintf("Humano no es Mutante "), func(b *testing.B) {
		got, _ := IsMutant(humano)
		want := false
		if got != want {
			b.Errorf("got %v want %v", got, want)
		}
	})

	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				IsMutant(humano)
			}
		},
	)

	b.Run(fmt.Sprintf("Mutante es Mutante "), func(b *testing.B) {
		got, _ := IsMutant(mutante)
		want := true
		if got != want {
			b.Errorf("got %v want %v", got, want)
		}
	})

	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				IsMutant(mutante)
			}
		},
	)

}

func getADNHumano() []string {
	return []string {
		"ATGCGA",
		"CAGTGC",
		"TTATTT",
		"AGACGG",
		"GCGTCA",
		"TCACTG",
	}
}

func getADNMutante() []string {
	return []string {
		"ATGCGA",
		"CAGTGC",
		"TTACGT",
		"AGCAGG",
		"CCCCTT",
		"CCCCTG",
	}
}