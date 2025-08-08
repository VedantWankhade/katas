package katas

import (
	"fmt"
	"time"
)

func bufferedChannels() {
	messages := make(chan string, 3)
	go func() {
		fmt.Println(<-messages)
	}()
	messages <- "1"
	fmt.Println("message 1 sent")
	// we can send upto 3 messages even though only one of them is used in the goroutine
	messages <- "2"
	fmt.Println("message 2 sent")
	messages <- "3"
	fmt.Println("message 3 sent")
	messages <- "4"
	fmt.Println("message 4 sent")
	// the following will block
	messages <- "5"
	fmt.Println("message 5 sent")
	time.Sleep(2 * time.Second)
}
