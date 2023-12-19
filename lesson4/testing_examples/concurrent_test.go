package main

import (
	"sync"
	"testing"
)

func TestConcurrentRead(t *testing.T) {
	//
	// Each element is handled as distinct variable
	// so it is concurrent safe to read from slice.
	//

	v := make([]int, 0, 100000)

	for j := 0; j < 100000; j++ {
		v = append(v, j)
	}

	var wg sync.WaitGroup

	readSlice := func(ind int) {
		defer wg.Done()
		_ = v[ind]
	}

	for i := range v {
		wg.Add(1)
		go readSlice(i)
	}

	wg.Wait()
}

func TestSafeConcurrentWrite(t *testing.T) {
	//
	// Each element is handled as distinct variable
	// so it is concurrent safe to wtite to slice using indecies.
	//

	v := make([]int, 100000)

	var wg sync.WaitGroup

	readSlice := func(i, value int) {
		defer wg.Done()

		v[i] = value
	}

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go readSlice(i, i)
	}

	wg.Wait()
}

func TestBadConcurrentWrite(t *testing.T) {
	//
	// When appending we mess with whole memory
	// so it is NOT concurrent safe to append elements to slice.
	//

	v := make([]int, 0, 100000)

	var wg sync.WaitGroup

	readSlice := func(ind, value int) {
		defer wg.Done()

		v = append(v, value)
	}

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go readSlice(i, i)
	}

	wg.Wait()
}
