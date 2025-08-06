package katas

import (
	"errors"
	"fmt"
)

type myError struct {
	msg string
}

func (e myError) Error() string {
	return e.msg
}

func customErrors() {
	e := f4()
	var err myError
	if e != nil && errors.As(e, &err) {
		fmt.Println("go the error type:", e.Error())
	} else {
		fmt.Println("No error or no error type matched")
	}
}

func f4() error {
	return myError{
		"just an error bro",
	}
}
