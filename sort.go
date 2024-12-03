package simple

import "cmp"

func Sort[S ~[]E, E cmp.Ordered](x S) {
	for i := range len(x) {
		for j := range i {
			if x[i] < x[j] {
				x[i], x[j] = x[j], x[i]
			}
		}
	}
}
