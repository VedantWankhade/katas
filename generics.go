package katas

import (
	"fmt"
	"reflect"
)

type Num interface {
	int | float64
}

func addNums[N int | float64](n1, n2 N) N {
	return n1 + n2
}

// same as above function
func addNum2[N Num](n1, n2 N) N {
	return n1 + n2
}

func generics() {
	a1 := addNums(1, 2.44)
	fmt.Println(reflect.TypeOf(a1))
	fmt.Println(a1)

	// a2 := addNums[int](2, 2.44) // wont work because we are restricting the generic to int
	a2 := addNums[int](2, 3)
	fmt.Println(a2)

	b1 := addNum2(1, 2.44)
	fmt.Println(reflect.TypeOf(b1))
	fmt.Println(b1)

	// a2 := addNum2[int](2, 2.44) // wont work because we are restricting the generic to int
	b2 := addNum2[int](2, 3)
	fmt.Println(b2)
}

// TODO:
// ~T syntax - underlying type generic
// comparable generic
// generic struct
// other stuff
