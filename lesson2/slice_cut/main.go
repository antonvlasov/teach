package main

import "fmt"

func printSliceProperties(sl []int8) {
	fmt.Printf("ptr: %d; len: %d; cap: %d; values: %v\n", &sl[0], len(sl), cap(sl), sl)
}

func main() {
	sl := make([]int8, 3, 6)
	for i := 0; i < len(sl); i++ {
		sl[i] = int8(i)
	}
	printSliceProperties(sl)

	sl = append(sl, 3)
	printSliceProperties(sl)

	sl = sl[:2]
	printSliceProperties(sl)

	sl = sl[:4]
	printSliceProperties(sl)

	sl = sl[:6]
	printSliceProperties(sl)

	sl = append(sl, 9)
	printSliceProperties(sl)

	sl = sl[3:7]
	printSliceProperties(sl)

	sl = append(sl, []int8{11, 12, 13, 14, 15}...)
	printSliceProperties(sl)
}
