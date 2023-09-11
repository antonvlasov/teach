package main

import "fmt"

//go:noinline
func add(x int, y int) int {
	retVal := x + y

	return retVal
}

func main() {
	var (
		a   int = 42
		b   int = 13
		res int
	)

	res = add(a, b)

	fmt.Println(res)
}
