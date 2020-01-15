package lib

import "testing"

func TestGenerateHash(t *testing.T) {
	mutante := []string {"ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG",}

	t.Run("Generate Hash", func(t *testing.T) {
		got := GenerateHash(mutante)
		var want uint32 = 3100903834

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func BenchmarkGenerateHash(b *testing.B) {
	mutante := []string {"ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG",}

	b.Run("Bench Generar Hash ", func(b *testing.B) {
		got := GenerateHash(mutante)
		var want uint32 = 3100903834

		if got != want {
			b.Errorf("got %v want %v", got, want)
		}
	})

	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				GenerateHash(mutante)
			}
		},
	)

}

func TestEqualString(t *testing.T) {
	string1 := getMutantDNAString()
	string2 := getMutantDNAString()
	string3 := []string {"AAGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG",	}

	t.Run("Compare two equal DNA", func(t *testing.T) {
		got := EqualString(string1, string2)
		want := true

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Compare two distinct DNA", func(t *testing.T) {
		got := EqualString(string1, string3)
		want := false

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func BenchmarkEqualString(b *testing.B) {
	string1 := getMutantDNAString()
	string2 := getMutantDNAString()

	b.Run("Bench Compare two DNA ", func(b *testing.B) {
		got := EqualString(string1, string2)
		want := true

		if got != want {
			b.Errorf("got %v want %v", got, want)
		}
	})

	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				EqualString(string1, string2)
			}
		},
	)
}

func TestConverByteToResultado(t *testing.T) {
	mutante := getMutantDNA()
	adnMutante := getMutantDNAString()

	t.Run("Convert []byte to Struct", func(t *testing.T) {
		got, err := ConverByteToResultado(mutante)
		want := Resultado{Dna:adnMutante, IsMutant:true}

		if err != nil || got.IsMutant != want.IsMutant || !EqualString(got.Dna, want.Dna){
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func BenchmarkConverByteToResultado(b *testing.B) {
	mutante := getMutantDNA()
	adnMutante := getMutantDNAString()

	b.Run("Bench Convert []byte struct", func(b *testing.B) {
		got, err := ConverByteToResultado(mutante)
		want := Resultado{Dna:adnMutante, IsMutant:true}

		if err != nil || got.IsMutant != want.IsMutant || !EqualString(got.Dna, want.Dna){
			b.Errorf("got %v want %v", got, want)
		}
	})

	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				ConverByteToResultado(mutante)
			}
		},
	)
}

func TestConvertResultadoToByte(t *testing.T) {
	mutante := getMutantDNAString()

	t.Run("Convert Struct to []byte", func(t *testing.T) {
		got := ConvertResultadoToByte(mutante, true)
		want := getMutantDNA()

		if !equalBytes(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func BenchmarkConvertResultadoToByte(b *testing.B) {
	mutante := getMutantDNAString()

	b.Run("Bench Convert Struct to []byte ", func(b *testing.B) {
		got := ConvertResultadoToByte(mutante, true)
		want := getMutantDNA()

		if !equalBytes(got, want) {
			b.Errorf("got %v want %v", got, want)
		}
	})

	b.RunParallel(
		func(pb *testing.PB) {
			for pb.Next() {
				ConvertResultadoToByte(mutante, true)
			}
		},
	)
}

func equalBytes(a, b []byte) bool {
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

func getMutantDNA() []byte {
	return []byte{123, 34, 68, 110, 97, 34, 58, 91, 34, 65, 65, 65, 65, 71, 65, 34, 44, 34, 67, 67, 67, 67, 71, 67, 34,
		44, 34, 84, 84, 65, 84, 71, 84, 34, 44, 34, 65, 71, 65, 65, 71, 71, 34, 44, 34, 67, 67, 67, 67, 84, 65,
		34, 44, 34, 84, 67, 65, 67, 84, 71, 34, 93, 44, 34, 73, 115, 77, 117, 116, 97, 110, 116, 34, 58, 116,
		114, 117, 101, 125}
}

func getMutantDNAString() []string {
	return []string {"AAAAGA","CCCCGC","TTATGT","AGAAGG","CCCCTA","TCACTG",}
}
