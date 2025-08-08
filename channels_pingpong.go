package katas

import (
	"fmt"
	"sync"
)

func pingpong() {
	var wg sync.WaitGroup
	wg.Add(2)
	c := make(chan bool)
	go ping(&wg, c)
	go pong(&wg, c)
	wg.Wait()
}
func pong(wg *sync.WaitGroup, c chan bool) {
	c <- true
	for range 5 {
		if <-c {
			fmt.Println("pong")
			c <- true
		}
	}
	wg.Done()
}
func ping(wg *sync.WaitGroup, c chan bool) {
	for range 5 {
		if <-c {
			fmt.Println("ping")
			c <- true
		}
	}
	<-c
	wg.Done()
}
