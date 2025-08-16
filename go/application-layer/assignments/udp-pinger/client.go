package udppinger

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// I feel like I made this programm unnecessarilly complicated, but its still cool I think
func client() {
	numPackets := 10
	var wg sync.WaitGroup
	numRet := 0
	var x sync.Mutex
	ret := make(chan int, 10)

	for i := range numPackets {
		wg.Add(1)
		go func(*sync.Mutex, *sync.WaitGroup, chan<- int) {
			defer wg.Done()

			start := time.Now()
			conn, err := net.Dial("udp", "127.0.0.1:12000")
			if err != nil {
				return
			}
			fmt.Printf("SENT packet %d\n", i)
			_, err = fmt.Fprintf(conn, "%d", i)
			if err != nil {
				return
			}

			conn.Close()
			retChan := make(chan time.Time)

			go func(chan<- time.Time, time.Time) {
				lis, err := net.ListenPacket(conn.LocalAddr().Network(), conn.LocalAddr().String())
				if err != nil {
					return
				}
				resMsg := make([]byte, 1024)
				_, _, err = lis.ReadFrom(resMsg)
				if err != nil {
					return
				}
				retChan <- time.Now()
			}(retChan, start)

			select {
			case t := <-retChan:
				delta := t.Nanosecond() - start.Nanosecond()
				fmt.Printf("RECEIVED packet %d in %d nanoseconds\n", i, delta)
				x.Lock()
				defer x.Unlock()
				numRet++
				ret <- delta
			case <-time.After(4 * time.Second):
				fmt.Printf("TIMEOUT packet %d\n", i)
				return
			}
		}(&x, &wg, ret)
	}

	wg.Wait()
	fmt.Println("------------------------------------")
	fmt.Printf("Packet Loss: %f%%\n", (float64(numRet)/float64(numPackets))*100)
	sum := 0
	for range numRet {
		sum += <-ret
	}
	fmt.Printf("Average RTT: %d nanoseconds\n", sum/numRet)
	fmt.Println("------------------------------------")
}
