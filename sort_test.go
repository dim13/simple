package simple

import (
	"math/rand/v2"
	"slices"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	array := []int{1, 4, 9, 13, 2, 5, 8, 7}
	want := []int{1, 2, 4, 5, 7, 8, 9, 13}
	Sort(array)
	if !slices.Equal(array, want) {
		t.Error("got", array)
	}
}

func BenchmarkSort(b *testing.B) {
	sample := rand.Perm(1000)
	b.Run("slices.Sort", func(b *testing.B) {
		for range b.N {
			a := slices.Clone(sample)
			slices.Sort(a)
		}
	})
	b.Run("sort.IntSlice", func(b *testing.B) {
		for range b.N {
			a := slices.Clone(sample)
			sort.IntSlice(a).Sort()
		}
	})
	b.Run("simple.Sort", func(b *testing.B) {
		for range b.N {
			a := slices.Clone(sample)
			Sort(a)
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
