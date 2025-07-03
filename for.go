package katas

import "fmt"

func forLoop() {
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 1; j < 4; j++ {
		fmt.Println(j)
	}

	for x := range 5 {
		fmt.Println(x)
	}

	for {
		// infinite
		fmt.Println("bro")
		break
	}
}
