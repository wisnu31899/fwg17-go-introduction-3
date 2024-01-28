package main

import (
	"fmt"
	"sync"
)

type deretBilangan struct {
	limit int
}

// Fungsi untuk mengecek apakah suatu bilangan adalah bilangan prima
func isPrima(num int) bool {
	if num <= 1 {
		return false
	}
	if num <= 3 {
		return true
	}
	if num%2 == 0 || num%3 == 0 {
		return false
	}
	for i := 5; i*i <= num; i = i + 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}
	return true
}

// Fungsi untuk mencetak deret bilangan prima hingga batas limit
func (d deretBilangan) prima(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Deret bilangan prima:")
	for i := 2; i <= d.limit; i++ {
		if isPrima(i) {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

// Fungsi untuk mencetak deret bilangan ganjil hingga batas limit
func (d deretBilangan) ganjil(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Deret bilangan ganjil:")
	for i := 1; i <= d.limit; i += 2 {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

// Fungsi untuk mencetak deret bilangan genap hingga batas limit
func (d deretBilangan) genap(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Deret bilangan genap:")
	for i := 2; i <= d.limit; i += 2 {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

// Fungsi untuk mencetak deret bilangan Fibonacci hingga batas limit
func (d deretBilangan) fibonacci(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Deret bilangan Fibonacci:")
	a, b := 0, 1
	for a <= d.limit {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
	fmt.Println()
}
