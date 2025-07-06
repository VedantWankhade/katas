package katas

import "fmt"

func maps() {
	m1 := make(map[string]int)
	m1["one"] = 1
	m1["two"] = 2

	m1["three"] = 3
	fmt.Println(m1, len(m1))

	delete(m1, "two")
	fmt.Println(m1, len(m1))

	for k, v := range m1 {
		fmt.Println(k, v)
	}
}
