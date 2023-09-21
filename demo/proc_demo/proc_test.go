package main

import (
	"math/rand"
	"testing"
)

func scalarMultiply(values []int, multiplicator int) []int {
	for i := range values {
		values[i] *= multiplicator
	}

	return values
}

const numCPU = 8

func parallelScalarMultiply(values []int, multiplicator int) []int {
	for i := 0; i <= len(values)-numCPU; i += numCPU {
		values[i] *= multiplicator
		values[i+1] *= multiplicator
		values[i+2] *= multiplicator
		values[i+3] *= multiplicator
		values[i+4] *= multiplicator
		values[i+5] *= multiplicator
		values[i+6] *= multiplicator
		values[i+7] *= multiplicator
	}

	for i := len(values) - numCPU + 1; i < len(values); i += 1 {
		values[i] *= multiplicator
	}

	return values
}

func prepareArgs(size int) ([]int, int) {
	res := make([]int, size)
	for i := range res {
		res[i] = int(rand.Int() % (2 << 15))
	}

	return res, rand.Int() % 2 << 7
}

func BenchmarkScalarMultiply(b *testing.B) {
	var res []int
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		values, multiplicator := prepareArgs(10000)
		b.StartTimer()

		res = scalarMultiply(values, multiplicator)
	}

	_ = res
}

func BenchmarkParallelScalarMultiply(b *testing.B) {
	var res []int
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		values, multiplicator := prepareArgs(1000)
		b.StartTimer()

		res = parallelScalarMultiply(values, multiplicator)
	}

	_ = res
}
