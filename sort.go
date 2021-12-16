package simple

import "sort"

func Sort(v sort.Interface) {
	n := v.Len()
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if v.Less(i, j) {
				v.Swap(i, j)
			}
		}
	}
}
