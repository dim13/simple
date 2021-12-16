package simple

import (
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	a := []int{1, 4, 9, 13, 2, 5, 8, 7}
	Sort(sort.IntSlice(a))
	t.Log(a)
}

func BenchmarkSort(b *testing.B) {
	b.Run("std", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a := []int{1, 4, 9, 13, 2, 5, 8, 7}
			sort.Sort(sort.IntSlice(a))
		}
	})
	b.Run("simple", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a := []int{1, 4, 9, 13, 2, 5, 8, 7}
			Sort(sort.IntSlice(a))
		}
	})
}
