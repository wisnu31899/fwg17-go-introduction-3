package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

func Fibonacci(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Generating Fibonacci numbers...")

	a, b := 0, 1
	for i := 0; i < 10; i++ {
		mutex.Lock()
		ch <- a
		mutex.Unlock()
		a, b = b, a+b
	}
	close(ch)
}

func GanjilGenap(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Filtering odd and even numbers...")

	for num := range ch {
		if num%2 == 0 {
			fmt.Println("Even:", num)
		} else {
			fmt.Println("Odd:", num)
		}
	}
}
