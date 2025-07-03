package katas

import "fmt"

func ifElse() {
	if true {
		fmt.Println("hello true")
	}

	if 7%2 == 0 {
		fmt.Println("7%2")
	}

	if num := 2; num == 2 {
		fmt.Println(num)
	}
}
