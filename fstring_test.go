package fstring

import (
	"log"
	"testing"

	"github.com/eliothedeman/randutil"
)

func init() {
	log.SetFlags(log.Llongfile)
}

func TestFlagLen(t *testing.T) {
	t.Run("smallEnough", func(t *testing.T) {
		f := toLen(10)
		if lenInFlags(f) != 10 {
			t.Error(lenInFlags(f))
		}
	})
	t.Run("tooBig", func(t *testing.T) {
		f := toLen(255)
		if lenInFlags(f) == 255 {
			t.Error(lenInFlags(f))
		}
	})
}

func TestToFromString(t *testing.T) {
	t.Run("small", func(t *testing.T) {
		s := fromString("hello")
		if toString(s) != "hello" {
			t.Error(len(toString(s)))
		}

	})
}

func TestLargeStringEq(t *testing.T) {
	a := fromString("abcdefghijklmnopqrstuvwxyz")
	b := fromString("abcdefghijklmnopqrstuvwxyz")

	if !eq(a, b) {
		t.Fail()
	}

}

func genSmallStrings(n, l int) []stringStructSmall {
	x := make([]stringStructSmall, n)
	for i := range x {
		x[i] = fromString(randutil.AlphaString(l))
	}

	return x
}

func genStrings(n, l int) []string {
	x := make([]string, n)
	for i := range x {
		x[i] = randutil.AlphaString(l)
	}

	return x
}

func copySmallStrings(s []stringStructSmall) []stringStructSmall {
	x := make([]stringStructSmall, len(s))
	for i := range x {
		y := []byte(toString(s[i]))
		z := make([]byte, len(y))
		copy(z, y)
		x[i] = fromString(string(z))
	}
	return x
}
func copyStrings(s []string) []string {
	x := make([]string, len(s))
	for i := range x {
		y := []byte(s[i])
		x[i] = string(y)
	}
	return x
}

func BenchmarkEq(b *testing.B) {
	n := 1000000
	small := genSmallStrings(n, 15)
	old := genStrings(n, 15)
	old2 := copyStrings(old)
	small2 := copySmallStrings(small)

	b.Run("smallFalse", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if eq(small[i%n], small[(i+1)%n]) {
				b.Fatal("Shouldn't be equal")
			}
		}
	})
	b.Run("smallFalseBuiltin", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if eqs(old[i%n], old[(i+1)%n]) {
				b.Fatal("Shouldn't be equal")
			}
		}
	})
	b.Run("smallTrue", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if !eq(small[i%n], small2[i%n]) {
				b.Fatal("Should be equal")
			}
		}
	})
	b.Run("smallTrueBuiltin", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if !eqs(old[i%n], old2[i%n]) {
				b.Fatal("Should be equal")
			}
		}
	})

	small = genSmallStrings(n, 50)
	old = genStrings(n, 50)
	old2 = copyStrings(old)
	small2 = copySmallStrings(small)

	b.Run("largeFalse", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if eq(small[i%n], small[(i+1)%n]) {
				b.Fatal("Shouldn't be equal")
			}
		}
	})
	b.Run("largeFalseBuiltin", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if eqs(old[i%n], old[(i+1)%n]) {
				b.Fatal("Shouldn't be equal")
			}
		}
	})
	b.Run("largeTrue", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if !eq(small[i%n], small2[i%n]) {
				b.Fatal("Should be equal", small[i%n], small2[i%n])
			}
		}
	})
	b.Run("largeTrueBuiltin", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			if !eqs(old[i%n], old2[i%n]) {
				b.Fatal("Should be equal")
			}
		}
	})

	b.Run("bigTrue", func(b *testing.B) {
		x := fromString("abcdefghijklmnopqrstuvwxyz")
		y := fromString("abcdefghijklmnopqrstuvwxyz")
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			eq(x, y)
		}
	})
	b.Run("bigFalse", func(b *testing.B) {
		x := fromString("abcdefghijklmnopqrstu4vwxy")
		y := fromString("abcdefghijklmnopqrstuvwxyz")
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			eq(x, y)
		}
	})
}
