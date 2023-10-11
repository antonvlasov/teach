package main

import (
	"fmt"
	"time"
)

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	n := 43

	start := time.Now()
	for i := 0; i <= n; i++ {
		fibValue := fib(i)
		fmt.Printf("%d = %d\n", i, fibValue)
	}
	end := time.Now()

	fmt.Printf("Общее время выполнения: %v\n", end.Sub(start))
}
