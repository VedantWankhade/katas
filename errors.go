package katas

import (
	"errors"
	"fmt"
)

func errorInGo() {
	_, e1 := f1(true)
	if e1 != nil {
		fmt.Println("error:", e1.Error())
	}
	_, e2 := f1(false)
	if e2 != nil {
		fmt.Println("error:", e2.Error())
	}

	e3 := f2(true)
	if e3 != nil && errors.Is(e3, ErrPassedFalse) {
		fmt.Println("FALSE error: %w", e3.Error())
	} else if e3 != nil && errors.Is(e3, ErrPassedTrue) {
		fmt.Println("TRUE error: %w", e3.Error())
	}
}

func f1(b bool) (int, error) {
	if b {
		return 1, nil
	}
	return 0, errors.New("f1 was passed false")
}

var ErrPassedFalse error = fmt.Errorf("passed false")
var ErrPassedTrue = errors.New("passed true")

func f2(b bool) error {
	if b {
		return ErrPassedTrue
	}
	return ErrPassedFalse
}
