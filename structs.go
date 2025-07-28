package katas

import (
	"fmt"
	"reflect"
)

func structs() {
	type person struct {
		name string
		age  int
	}

	var a person // initialize field members to their zero values
	fmt.Println(a)
	a = person{
		name: "a",
		age:  22,
	}

	fmt.Println(a)

	var b person
	fmt.Println(b)
	b.name = "b"
	b.age = 99
	fmt.Println(b)

	fmt.Println(a == b)
	b.age = 22
	b.name = "a"

	fmt.Println(a == b)
	c := person{
		"c",
		100,
	}
	fmt.Println(c)

	d := struct {
		name string
		age  int
	}{
		"d",
		66,
	}

	fmt.Println(d)

	fmt.Println(c == d) // type matches even though they are technically different
	d.age = 100
	d.name = "c"

	fmt.Println(c == d)
	fmt.Println(reflect.TypeOf(c))
	fmt.Println(reflect.TypeOf(d))

	e := struct {
		name string
	}{
		"e",
	}
	fmt.Println(e)
	// fmt.Println(e == d) // type mismatch
	// fmt.Println(e == a) // type mismatch
}
