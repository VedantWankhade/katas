package test

import "testing"

type TestCase[I any, E any] struct {
	Input    I
	Expected E
}

type TwoInputTestCase[I1, I2, E any] struct {
	Input1   I1
	Input2   I2
	Expected E
}

func Test[I any, E comparable](t *testing.T, fn func(I) E, tests []TestCase[I, E]) {
	for _, test := range tests {
		actual := fn(test.Input)
		if actual != test.Expected {
			t.Logf("Input: %v\tExpected: %v\tActual: %v", test.Input, test.Expected, actual)
			t.Fail()
		}
	}
}

func TestTwoInputsWithComparator[I1, I2 any, E any](t *testing.T, fn func(I1, I2) E, comp func(E, E) bool, tests []TwoInputTestCase[I1, I2, E]) {
	for _, test := range tests {
		actual := fn(test.Input1, test.Input2)
		if !comp(actual, test.Expected) {
			t.Logf("Input: %v, %v\tExpected: %v\tActual: %v", test.Input1, test.Input2, test.Expected, actual)
			t.Fail()
		}
	}
}
