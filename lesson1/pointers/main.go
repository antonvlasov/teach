package main

import "fmt"

//go:noinline
func max(x *int, y *int, res *int) {
	var (
		xDeref int = *x
		yDeref int = *y
	)

	if xDeref > yDeref {
		*res = xDeref
		return
	}

	*res = yDeref
}

func main() {
	var (
		a   int = 42
		b   int = 13
		res int
	)

	var (
		aPtr   *int = &a
		bPtr   *int = &b
		resPtr *int = &res
	)

	max(aPtr, bPtr, resPtr)

	fmt.Println(res)
}
