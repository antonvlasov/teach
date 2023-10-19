package benchpress

import "sync"

func FindSum(list []int) int {
	sum := 0

	for _, number := range list {
		sum += number
	}

	return sum
}

func FindSumConc(list []int) int {
	sum := 0

	var rwm sync.RWMutex

	var wg sync.WaitGroup

	wg.Add(len(list))

	for _, number := range list {
		go func(number int) {
			defer wg.Done()
			rwm.Lock()
			defer rwm.Unlock()
			sum += number
		}(number)
	}

	wg.Wait()

	return sum
}
