package simple

import "constraints"

func Sort[T constraints.Ordered](v []T) {
	n := len(v)
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if v[i] < v[j] {
				v[i], v[j] = v[j], v[i]
			}
		}
	}
}
