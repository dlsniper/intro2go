package main

import (
	"fmt"
	"math"
	"sync"
)

var wait = &sync.WaitGroup{}

func powersOf(what float64) {
	for i := float64(0); i <= 10; i++ {
		fmt.Printf("%2.f ^ %2.f = %.0f\n", what, i, math.Pow(what, i))
	}
	wait.Done()
}

func main() {
	for i:=1; i<10; i++ {
		wait.Add(1)
		go func () {
			powersOf(float64(i))
		}()
	}

	wait.Wait()
}
