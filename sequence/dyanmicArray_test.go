package sequence_test

import (
	"fmt"
	"testing"

	"github.com/vedantwankhade/katas/dsa/sequence"
)

func TestDynamicArray(t *testing.T) {
	arr := sequence.NewArray[int](4)
	fmt.Println(arr, arr.Len(), arr.Size())
	fmt.Println(arr.Get(2))

	arr.Add(1)
	fmt.Println(arr, arr.Len(), arr.Size())
	fmt.Println(arr.Get(0))

	arr.Add(1)
	fmt.Println(arr, arr.Len(), arr.Size())
	fmt.Println(arr.Get(1))

	arr.Add(2)
	fmt.Println(arr, arr.Len(), arr.Size())
	fmt.Println(arr.Get(2))

	arr.Add(3)
	fmt.Println(arr, arr.Len(), arr.Size())
	fmt.Println(arr.Get(3))

	arr.Add(4)
	fmt.Println(arr, arr.Len(), arr.Size())
	fmt.Println(arr.Get(4))

	for i, t := range arr.Iter() {
		fmt.Println(i, t)
	}
}
