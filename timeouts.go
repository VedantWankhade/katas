package katas

import (
	"fmt"
	"time"
)

func timeouts() {
	timeout := 1
	done := make(chan bool)
	go twoSService(done)
	start := time.Now()
	select {
	case t := <-time.After(time.Duration(timeout) * time.Second):
		elapsed := t.Sub(start)
		fmt.Printf("err: connection timedout: %v seconds\n", elapsed.Seconds())
	case <-done:
		fmt.Println("Service finished successfully")
	}
}

func twoSService(done chan<- bool) {
	time.Sleep(2 * time.Second)
	done <- true
}
