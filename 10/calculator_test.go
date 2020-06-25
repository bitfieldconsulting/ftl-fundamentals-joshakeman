package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		a, b int
		want int
	}{
		{a: 1, b: 1, want: 2},
		{a: 2, b: 2, want: 4},
		{a: 6, b: 7, want: 13},
		{a: -8, b: -3, want: -11},
		{a: -2, b: 0, want: -2},
		{a: 0, b: 0, want: 0},
		{a: 3, b: 5 / 5, want: 4},
		{a: 3 * 10, b: 6 / 3, want: 32},
	}
	for _, testCase := range testCases {
		got := calculator.Add(testCase.a, testCase.b)
		if testCase.want != got {
			t.Errorf("Add(%d, %d): want %d, got %d", testCase.a, testCase.b, testCase.want, got)
		}
	}
}

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
		a, b        int
		want        int
		errExpected bool
	}{
		{a: 14, b: 4, want: 3, errExpected: false},
		{a: 4, b: -2, want: -2, errExpected: false},
		{a: 12, b: 4, want: 3, errExpected: false},
		{a: 2, b: 0, want: 0, errExpected: true},
		{a: 4, b: 1, want: 4, errExpected: false},
	}
	for _, testCase := range testCases {
		got, err := calculator.Divide(testCase.a, testCase.b)
		if testCase.errExpected == true && err == nil {
			t.Errorf("You divided %v by %v. This test expected a divide by zero error but got no error", testCase.a, testCase.b)
		}
		if got != testCase.want {
			t.Errorf("Divided (%v,%v): Wanted %v, got %v", testCase.a, testCase.b, testCase.want, got)
		}
	}
}
