package lib

import "testing"

func TestGenerateHash(t *testing.T) {
	mutante := []string {
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTG",
	}

	t.Run("Generate Hash Mutant", func(t *testing.T) {
		got := GenerateHash(mutante)
		var want uint32 = 3100903834

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
