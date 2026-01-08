package sequence_test

import (
	"fmt"
	"testing"

	"github.com/vedantwankhade/katas/dsa/sequence"
)

func TestLinkedList(t *testing.T) {
	arr := sequence.NewLinkedList[int]()
	fmt.Println(arr, arr.Len())
	fmt.Println(arr.Get(2))

	arr.Add(1)
	fmt.Println(arr, arr.Len())
	fmt.Println(arr.Get(0))

	arr.Add(1)
	fmt.Println(arr, arr.Len())
	fmt.Println(arr.Get(1))

	arr.Add(2)
	fmt.Println(arr, arr.Len())
	fmt.Println(arr.Get(2))

	arr.Add(3)
	fmt.Println(arr, arr.Len())
	fmt.Println(arr.Get(3))

	arr.Add(4)
	fmt.Println(arr, arr.Len())
	fmt.Println(arr.Get(4))

	for i, t := range arr.Iter() {
		fmt.Println(i, t)
	}

	arr.AddAt(0, 99)
	fmt.Println(arr)
	arr.AddAt(2, 88)
	fmt.Println(arr)
	arr.AddAt(4, 2423)
	fmt.Println(arr)
	arr.AddAt(7, 112)

	fmt.Println(arr)
	arr.AddAt(9, 108080808)
	fmt.Println(arr)
}
