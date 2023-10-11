package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	N := 4 // играться здесь, ставьте N воркеров
  fiboNumber := 45
  
	jobs := make(chan int, N)
	var wg sync.WaitGroup
	startTime := time.Now()

	for i := 0; i < N; i++ {
		wg.Add(1)
		go worker(&wg, jobs, i)
	}

	for i := 1; i < fiboNumber; i++ {
		jobs <- i
	}

	close(jobs)
	wg.Wait()
	endTime := time.Now()
	fmt.Printf("Общее время выполнения: %v\n", endTime.Sub(startTime))

}

func worker(wg *sync.WaitGroup, jobs <-chan int, workerId int) {
	defer wg.Done()

	for n := range jobs {
		fmt.Printf(strconv.Itoa(n)+": "+strconv.Itoa(fib(n))+" from worker %v\n", workerId)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
