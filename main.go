package main

import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Add(4)

	deret := deretBilangan{limit: 40}
	go deret.prima(&wg)
	go deret.ganjil(&wg)
	go deret.genap(&wg)
	go deret.fibonacci(&wg)

	wg.Wait()
}
