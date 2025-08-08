package katas

import (
	"fmt"
	"time"
)

func goroutines() {
	f("hello")
	fmt.Println("------------------------------------")

	go f("bye")
	go f("ok")

	time.Sleep(time.Second)
	fmt.Println("done")
}

func f(msg string) {
	for i := range 5 {
		fmt.Println(msg, ":", i)
	}
}
