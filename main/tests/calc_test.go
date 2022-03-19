package tests

// go test files always are written like this:   name_test.go

import (
	"first_goland_project/cal"
	"testing"
)

func TestAdd(t *testing.T) {
	a := 50.0
	b := 20.0
	resultado := cal.Add(a, b)
	if resultado != a+b {
		t.Errorf("got %f, want %f", resultado, a+b)
	}
}

func TestSubtract(t *testing.T) {
	a := 50.0
	b := 20.0
	resultado := cal.Subtract(a, b)
	if resultado != a-b {
		t.Errorf("got %f, want %f", resultado, a-b)
	}
}

func TestMultiply(t *testing.T) {
	a := 50.0
	b := 20.0
	resultado := cal.Multiply(a, b)
	if resultado != a*b {
		t.Errorf("got %f, want %f", resultado, a*b)
	}
}

func TestDivide(t *testing.T) {
	a := 50.0
	b := 20.0
	resultado := cal.Divide(a, b)
	if resultado != a/b {
		t.Errorf("got %f, want %f", resultado, a/b)
	}
}
