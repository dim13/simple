package simple

import (
	"math/rand/v2"
	"slices"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	a := []int{1, 4, 9, 13, 2, 5, 8, 7}
	Sort(a)
	t.Log(a)
}

func copySample(b *testing.B, s []int) []int {
	b.Helper()
	ret := make([]int, len(s))
	copy(ret, s)
	return ret
}

func newSample(b *testing.B, sz int) []int {
	sample := make([]int, sz)
	for i := range sample {
		sample[i] = rand.IntN(sz)
	}
	return sample
}

func BenchmarkSort(b *testing.B) {
	sample := newSample(b, 1000)
	b.Run("sort.IntSlice", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a := copySample(b, sample)
			sort.Sort(sort.IntSlice(a))
		}
	})
	b.Run("sort.Ints", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a := copySample(b, sample)
			sort.Ints(a)
		}
	})
	b.Run("simple.Sort", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a := copySample(b, sample)
			Sort(a)
		}
	})
	b.Run("slices.Sort", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a := copySample(b, sample)
			slices.Sort(a)
		}
	})
}
