package katas

import (
	"fmt"
	"sync"
	"time"
)

func channelsSync() {
	done := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(2)
	go task1(done, &wg)
	go task2(done, &wg)
	wg.Wait()
}

func task1(done chan bool, wg *sync.WaitGroup) {
	fmt.Println("Performing prerequite tasks - before task 2 finishes")
	_ = <-done
	fmt.Println("Performing tasks after task 2 finished")
	fmt.Println("Task 1 finished")
	wg.Done()
}

func task2(done chan bool, wg *sync.WaitGroup) {
	fmt.Println("Performing task 2")
	time.Sleep(2 * time.Second)
	fmt.Println("Task 2 finished")
	done <- true
	wg.Done()
}
