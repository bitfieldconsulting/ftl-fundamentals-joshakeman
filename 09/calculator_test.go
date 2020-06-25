package calculator_test

import (
	"calculator"
	"testing"
)

func TestSubtract(t *testing.T) {
	testCases := []struct {
		a, b int
		want int
	}{
		{a: 6, b: 7, want: -1},
		{a: 9, b: 3, want: 6},
		{a: 0, b: 0, want: 0},
		{a: -6, b: -7, want: 1},
		{a: -2, b: 5, want: -7},
	}
	for _, testCase := range testCases {
		got := calculator.Subtract(testCase.a, testCase.b)
		if got != testCase.want {
			t.Errorf("Subtracted (%v,%v): Wanted %v, got %v", testCase.a, testCase.b, testCase.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	testCases := []struct {
		a, b int
		want int
	}{
		{a: 2, b: calculator.Multiply(2, 2), want: 8},
		{a: 9, b: 0, want: 0},
		{a: 3, b: 3, want: 9},
		{a: 1, b: -10, want: -10},
		{a: -5, b: -5, want: 25},
	}

	for _, testCase := range testCases {
		got := calculator.Multiply(testCase.a, testCase.b)
		if got != testCase.want {
			t.Errorf("Multiplied (%v,%v): Wanted %v, got %v", testCase.a, testCase.b, testCase.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	testCases := []struct {
		a, b int
		want int
	}{
		{a: 14, b: 4, want: 3},
		{a: 4, b: -2, want: -2},
		{a: 12, b: 4, want: 3},
		{a: 2, b: 3, want: 0},
		{a: 4, b: 1, want: 4},
	}
	for _, testCase := range testCases {
		got := calculator.Divide(testCase.a, testCase.b)
		if got != testCase.want {
			t.Errorf("Divided (%v,%v): Wanted %v, got %v", testCase.a, testCase.b, testCase.want, got)
		}
	}
}
