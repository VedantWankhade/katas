package katas

import "fmt"

var z int
var y int = 99
var x = "hello global"

// xx := "nope" // wont work

func variables() {
	var a int
	fmt.Println(a)
	var b string
	fmt.Println(b)

	var c = 22
	fmt.Println(c)

	d := "hello"
	fmt.Println(d)

	var (
		m string = "m"
		n bool
	)

	fmt.Println(m)
	fmt.Println(n)

	fmt.Println(z)

	fmt.Println(y)
	fmt.Println(x)
}
