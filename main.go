package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup // Deklarasi variabel WaitGroup di luar fungsi main()

type Student struct {
	Name  string
	Score int
}

func main() {
	// //soal 1
	// var wg sync.WaitGroup
	// wg.Add(4)

	// deret := deretBilangan{limit: 40}
	// go deret.prima(&wg)
	// go deret.ganjil(&wg)
	// go deret.genap(&wg)
	// go deret.fibonacci(&wg)

	// wg.Wait()

	// //soal 2
	// ch := make(chan int)
	// var wg sync.WaitGroup

	// wg.Add(2)
	// go Fibonacci(ch, &wg)
	// go GanjilGenap(ch, &wg)

	// wg.Wait()

	//soal 3
	students := []Student{
		{Name: "Alice", Score: 85},
		{Name: "Bob", Score: 65},
		{Name: "Charlie", Score: 75},
		{Name: "David", Score: 55},
		{Name: "Emma", Score: 95},
	}

	resultCh := make(chan string) // Hapus buffer channel karena tidak diperlukan

	var mu sync.Mutex // Mutex untuk mengamankan operasi pada variabel shared

	wg.Add(len(students))
	for _, student := range students {
		go calculateGrade(student, resultCh, &mu)
	}

	go func() {
		wg.Wait() // menunggu semua goroutine selesai
		close(resultCh)
	}()

	for result := range resultCh {
		fmt.Println(result)
	}
}
