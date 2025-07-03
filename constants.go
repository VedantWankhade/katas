package katas

import (
	"fmt"
	"reflect"
)

const s string = "global constant string bro"

/*
Golang constants are created at compile time - even the local ones in functions.
*/
func constants() {
	fmt.Println(s)

	fmt.Println(reflect.TypeOf(s))

	const a int = 22 // typed constant
	fmt.Println(reflect.TypeOf(a))

	const b = 99 // untyped constant
	fmt.Println(reflect.TypeOf(b))

	var x int = 1

	fmt.Println(x + a)
	fmt.Println(x + b)
}
