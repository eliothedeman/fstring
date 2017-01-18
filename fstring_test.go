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
		t.Error(s.cat(x).String(), a+b)
	}
}
