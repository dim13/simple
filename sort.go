package simple

import "golang.org/x/exp/constraints"

func Sort[S ~[]E, E constraints.Ordered](x S) {
	for i := 1; i < len(x); i++ {
		for j := 0; j < i; j++ {
			if x[i] < x[j] {
				x[i], x[j] = x[j], x[i]
			}
		}
	}
}
