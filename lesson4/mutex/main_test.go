package main

import (
	"fmt"
	"sync"
	"testing"
)

// func TestIncForRace(t *testing.T) {
// 	c := NeedToSync{
// 		counters: map[string]int{
// 			"chief keef": 0,
// 			"gorillaz":   0,
// 		},
// 	}

// 	var wg sync.WaitGroup

// 	doIncrement := func(name string, n int) {
// 		for i := 0; i < n; i++ {
// 			c.Inc(name)
// 		}
// 		wg.Done()
// 	}

// 	wg.Add(3)
// 	go doIncrement("chief keef", 10000)
// 	go doIncrement("chief keef", 10000)
// 	go doIncrement("gorillaz", 1000)

// 	wg.Wait()

// 	for key, value := range c.counters {
// 		fmt.Println(key, ":", value)
// 	}
// }

func TestIncConcurrentForRace(t *testing.T) {
	c := NeedToSync{
		counters: map[string]int{
			"the cabs":            0,
			"The Notorious B.I.G": 0,
		},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		defer wg.Done()

		for i := 0; i < n; i++ {
			c.IncConcurrent(name)
		}
	}

	wg.Add(3)
	go doIncrement("the cabs", 10000)
	go doIncrement("the cabs", 10000)
	go doIncrement("The Notorious B.I.G", 10000)

	wg.Wait()

	for key, value := range c.counters {
		fmt.Println(key, ":", value)
	}
}
