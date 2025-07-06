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

	twoAdder := xAddGenerator(2)
	fourAdder := xAddGenerator(4)

	fmt.Println(twoAdder(5))
	fmt.Println(fourAdder(5))
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

func xAddGenerator(x int) func(int) int {
	return func(i int) int {
		return i + x
	}
}
