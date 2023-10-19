package main

import (
	"fmt"
	"sync"
)

type NeedToSync struct {
	mu       sync.Mutex
	counters map[string]int
}

func (r *NeedToSync) IncConcurrent(name string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.counters[name]++
}

func (r *NeedToSync) Inc(name string) {
	r.counters[name]++
}

func main() {
	c := NeedToSync{
		counters: map[string]int{
			"$uicideboy$":         0,
			"My bloody valentine": 0,
		},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.IncConcurrent(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("$uicideboy$", 10000)
	go doIncrement("$uicideboy$", 10000)
	go doIncrement("My bloody valentine", 10000)

	wg.Wait()

	for key, value := range c.counters {
		fmt.Println(key, ":", value)
	}
}
