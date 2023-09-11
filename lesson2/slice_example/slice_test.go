package main

import (
	"math/rand"
	"testing"
)

//go:noinline
func simpleEscapes(size int) int {
	sl := make([]int, size)

	for i := 0; i < len(sl); i++ {
		sl[i] = int(rand.Int63())
	}

	return sl[len(sl)/2]
}

//go:noinline
func simpleNoExcape() int {
	sl := make([]int, 10)

	for i := 0; i < len(sl); i++ {
		sl[i] = int(rand.Int63())
	}

	return sl[len(sl)/2]
}

func BenchmarkSimpleEscape(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = simpleEscapes(10)
	}
}

func BenchmarkSimpleNoEscape(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = simpleNoExcape()
	}
}
