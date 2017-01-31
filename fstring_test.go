package fstring

import (
	"log"
	"testing"
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

func BenchmarkEq(b *testing.B) {

	b.Run("smallTrue", func(b *testing.B) {
		x := fromString("world")
		y := fromString("world")
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			eq(x, y)
		}
	})
	b.Run("smallTrueBuiltin", func(b *testing.B) {
		x := "world"
		y := "world"
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			eqs(x, y)
		}
	})
	b.Run("smallFalse", func(b *testing.B) {
		x := fromString("rld")
		y := fromString("world")
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			eq(x, y)
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
