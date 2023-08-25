package sync_pool

import (
	"math/rand"
	"testing"
	"time"
)

var (
	seed   = time.Now().UnixNano()
	random = rand.New(rand.NewSource(seed))
	n      = 1000000
)

func generateWithCap(n int) []int {
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, random.Int())
	}
	return nums
}

func generate(n int) []int {
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, random.Int())
	}
	return nums
}

func BenchmarkGenerateWithCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generateWithCap(n)
	}
}

func BenchmarkGenerate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		generate(n)
	}
}
