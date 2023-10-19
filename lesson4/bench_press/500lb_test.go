package benchpress

import (
	"math/rand"
	"testing"
)

const sliceLen = 5

func BenchmarkFindSum(b *testing.B) {
	list := make([]int, 0)
	for i := 0; i < sliceLen; i++ {
		list = append(list, rand.Intn(sliceLen))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindSum(list)
	}
}
func BenchmarkFindSumConc(b *testing.B) {
	list := make([]int, 0)
	for i := 0; i < sliceLen; i++ {
		list = append(list, rand.Intn(sliceLen))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindSumConc(list)
	}
}
