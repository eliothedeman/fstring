package fstring

import "testing"

func TestRawLength(t *testing.T) {
	s := FString{
		len: 10,
	}

	if s.toRaw().len() != 10 {
		t.Error(s.toRaw().len())
	}
}

func TestFromString(t *testing.T) {
	s := fromString("Hello world")

	if s.String() != "Hello world" {
		t.Fail()
	}
}

func TestCat(t *testing.T) {
	a := "Hlaskdjfl"
	b := "x"
	s := fromString(a)
	x := fromString(b)

	if s.cat(x).String() != a+b {
		// t.Error(s.cat(x).String(), a+b)
	}
}

func BenchmarkCatSmall(b *testing.B) {
	a := "Hlaskdjfl"
	y := "x"
	s := fromString(a)
	x := fromString(y)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.cat(x)
	}
}

func BenchmarkCatSmallBuiltin(b *testing.B) {
	a := "Hlaskdjfl"
	y := "x"
	x := ""
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x = a + y
	}
	b.Log(x)
}
