package katas

import (
	"fmt"
	"time"
)

func switchConstruct() {
	x := 2
	switch x {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	day := time.Now().Weekday()

	switch day {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	// type switch - switches on type of an interface variable
	var y interface{}
	y = "hello"

	switch y.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	}

	var yy any
	yy = 22
	switch yy.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	}
}
