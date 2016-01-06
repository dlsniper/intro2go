package main

import (
	"fmt"
	"math"
	"os"
	"sync"
	"time"
)

func powersOf(what float64, wait *sync.WaitGroup, result chan float64) {
	for i := float64(0); i <= 10; i++ {
		result <- math.Pow(what, i)
		time.Sleep(1 * time.Millisecond)
	}
	wait.Done()
}

func main() {
	wait := &sync.WaitGroup{}
	result := make(chan float64)
	for i := 1; i < 1000; i++ {
		wait.Add(1)
		go powersOf(float64(i), wait, result)
	}

	var timeout <-chan time.Time
	go func(timeout *<-chan time.Time) {
		*timeout = time.After(20 * time.Millisecond)
	}(&timeout)

	progress := time.Tick(1 * time.Millisecond)

	var sum float64
	for {
		select {
		case value := <-result:
			sum += value
		case <-progress:
			fmt.Printf("sum so far: %.f\n", sum)
		case <-timeout:
			goto timeOut
		}
	}
	wait.Wait()

timeOut:
	fmt.Printf("Sum is: %.f\n", sum)
	os.Exit(0)
}
