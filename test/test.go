package test

import (
	"testing"
)

type TestCase[I any, E any] struct {
	Input    I
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
