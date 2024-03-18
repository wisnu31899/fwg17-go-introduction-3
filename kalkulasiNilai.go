package main

import (
	"fmt"
	"sync"
)

func calculateGrade(student Student, resultCh chan<- string, mu *sync.Mutex) {
	defer wg.Done() // Menggunakan variabel wg yang didefinisikan di luar fungsi main()

	var grade string
	switch {
	case student.Score >= 80:
		grade = "A"
	case student.Score >= 70:
		grade = "B"
	case student.Score >= 60:
		grade = "C"
	case student.Score >= 50:
		grade = "D"
	default:
		grade = "E"
	}

	mu.Lock() // Mengunci akses ke variabel shared
	resultCh <- fmt.Sprintf("%s - Score: %d, Grade: %s", student.Name, student.Score, grade)
	mu.Unlock() // Membuka kunci akses ke variabel shared
}
