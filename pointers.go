package katas

import "fmt"

func pointers() {
	v1 := 22
	fmt.Println(v1)
	zeroval(v1)
	fmt.Println(v1)

	v2 := 99
	fmt.Println(v2)
	zeroptr(&v2)
	fmt.Println(v2)

	var v3 *int
	v3 = &v1
	fmt.Println(v3, *v3)
}

func zeroval(val int) {
	val = 0
}

func zeroptr(val *int) {
	*val = 0
}
