package testcalc

import (
	"structures/calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	result := calculator.Add(34.67, 78.6)
	if result != 34.67+78.6 {
		t.Errorf("Expected %f but got %f", 34.67+78.6, result)
	}
}
