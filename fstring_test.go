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
	str := "abcdefghijklmnopqrstuvwxyz"
	a := fromString(str)
	b := fromString(str)

	if !eq(a, b) {
		t.Fail()
	}

}

func genStrings(n, l int) []string {
	x := make([]string, n)
	for i := range x {
		x[i] = randutil.AlphaString(l)
	}

	return x
}

func copySmallStrings(s []string) []stringStructSmall {
	x := make([]stringStructSmall, len(s))
	for i := range x {
		x[i] = fromString(s[i])
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
	old := genStrings(n, 15)
	small := copySmallStrings(old)
	old2 := copyStrings(old)
	small2 := copySmallStrings(old)

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

	old = genStrings(n, 50)
	small = copySmallStrings(old)
	old2 = copyStrings(old)
	small2 = copySmallStrings(old)

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
