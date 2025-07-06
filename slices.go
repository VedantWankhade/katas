package katas

import (
	"fmt"
	"reflect"
	"slices"
)

func sliceS() {
	s1 := make([]string, 3)
	fmt.Println(s1, len(s1))
	fmt.Println(reflect.TypeOf(s1))

	s1[0] = "1"
	s1[1] = "2"
	s1[2] = "3"
	fmt.Println(s1)

	// s1[4] = "4"// wont work
	// fmt.Println(s1)
	s1 = append(s1, "4")
	fmt.Println(s1, len(s1))

	var s2 []int
	fmt.Println(s2, len(s2))
	fmt.Println(s2 == nil)
	s2 = append(s2, 1, 2, 3, 4)
	fmt.Println(s2, len(s2))

	var s3 []int = make([]int, 2)
	copy(s3, s2) // copy from s2 to s3 upto size of s3
	fmt.Println(s3, len(s3))

	var s4 = s2[0:2]
	fmt.Println(s4, len(s4))
	s4[0] = 99 // changes the underneath array - hence changes s2 as well
	fmt.Println(s4)
	fmt.Println(s2)

	for x := range s4 {
		fmt.Println(x)
	}

	fmt.Println(slices.Equal(s2, s3))
}
