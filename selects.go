package katas

import (
	"fmt"
	"time"
)

func selects() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "two"
	}()
	go func() {
		time.Sleep(time.Second)
		c1 <- "one"
	}()

	for range 2 {
		select {

		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
