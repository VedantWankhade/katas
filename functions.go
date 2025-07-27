package katas

import "fmt"

func functions() {
	a, b := addAndSubtract(5, 2)
	fmt.Println(a, b)
	fmt.Println(add(1, 2, 3, 4))
	nums := []int{1, 2, 3}
	fmt.Println(add(nums...))

	nums2 := [3]int{9, 8}
	// fmt.Println(add(nums2...)) // wont work cause nums2 is not a slice
	fmt.Println(add(nums2[:]...)) // do this instead

	sequenceOf2Adder := sequenceOfXAdder(2)
	sequenceOf3Adder := sequenceOfXAdder(3)

	fmt.Printf("%s\t%s\n", "Sequence of 2", "Sequence of 3")
	for range 5 {
		fmt.Printf("%d\t%d\n", sequenceOf2Adder(), sequenceOf3Adder())
	}

	fmt.Println(fib(6))
}

func sequenceOfXAdder(x int) func() int {
	y := 0
	inc := x
	return func() int {
		y = y + x
		x = x + inc
		return y
	}
}

func addAndSubtract(num1, num2 int) (int, int) {
	return num1 + num2, num1 - num2
}

func add(nums ...int) int {
	sum := 0
	for _, x := range nums {
		sum += x
	}
	return sum
}

func fib(n int) int {
	if n > 1 {
		return fib(n-2) + fib(n-1)
	} else {
		return n
	}
}
