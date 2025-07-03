package katas

import "fmt"

func arrays() {
	var arr1 [5]int

	// size := 5
	// var arr2 [size]int // doesnt work - need compile time constant for size of array

	const size2 = 5
	var arr3 [size2]int // works due to const
	fmt.Println(arr1, len(arr1))
	fmt.Println(arr3, len(arr3))

	b := [4]int{1, 2}
	fmt.Println(b, len(b))

	c := [...]string{"a", "b", "c"}
	fmt.Println(c, len(c))

	d := [...]int{1, 4: 4, 5}
	fmt.Println(d, len(d))

	twoD := [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)

}
