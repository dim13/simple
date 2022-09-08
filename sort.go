package simple

import "golang.org/x/exp/constraints"

func Sort[T constraints.Ordered](v []T) {
	for i := 1; i < len(v); i++ {
		for j := 0; j < i; j++ {
			if v[i] < v[j] {
				v[i], v[j] = v[j], v[i]
			}
		}
	}
}
