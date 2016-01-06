package main

import (
	"fmt"
	"math"
	"sync"
)

func powersOf(what float64, wait *sync.WaitGroup, result chan float64) {
	for i := float64(0); i <= 100; i++ {
		result <- math.Pow(what, i)
	}
	wait.Done()
}

func main() {
	wait := &sync.WaitGroup{}
	result := make(chan float64, 10000000)
	for i := 1; i < 100; i++ {
		wait.Add(1)
		go powersOf(float64(i), wait, result)
	}
	wait.Wait()

	var sum float64
	for len(result) > 0 {
		sum += <-result
	}

	fmt.Printf("Sum is: %.0f\n", sum)
}
