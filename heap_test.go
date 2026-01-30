package dsa

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	h := NewHeap[int](func(t1, t2 int) int {
		return t1 - t2
	})

	fmt.Println(h)
	h.add(5)
	fmt.Println(h)
	h.add(2)
	fmt.Println(h)
	h.add(3)
	fmt.Println(h)
	h.add(8)
	fmt.Println(h)
	h.add(6)
	fmt.Println(h)
	fmt.Println(h.peek())
	fmt.Println("-------------")
	fmt.Println(*h.get())
	fmt.Println(*h.get())
	fmt.Println(*h.get())
	fmt.Println(*h.get())
	fmt.Println(*h.get())

}
