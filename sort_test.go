package simple

import (
	"math/rand/v2"
	"slices"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	a := []int{1, 4, 9, 13, 2, 5, 8, 7}
	b := []int{1, 2, 4, 5, 7, 8, 9, 13}
	Sort(a)
	if !slices.Equal(a, b) {
		t.Error("got", a)
	}
}

func BenchmarkSort(b *testing.B) {
	sample := rand.Perm(1000)
	b.Run("sort.IntSlice", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a := slices.Clone(sample)
			sort.Sort(sort.IntSlice(a))
		}
	})
	b.Run("sort.Ints", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a := slices.Clone(sample)
			sort.Ints(a)
		}
	})
	b.Run("simple.Sort", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a := slices.Clone(sample)
			Sort(a)
		}
	})
	b.Run("slices.Sort", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			a := slices.Clone(sample)
			slices.Sort(a)
		}
	})
}

func FuzzSort(f *testing.F) {
	f.Fuzz(func(t *testing.T, v []byte) {
		sorted := slices.Clone(v)
		slices.Sort(sorted)
		Sort(v)
		if !slices.Equal(sorted, v) {
			t.Error("got", v)
		}
	})
}
