package simple

import "constraints"

func Sort[T constraints.Ordered](v []T) {
	n := len(v)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if v[i] < v[j] {
				v[i], v[j] = v[j], v[i]
			}
		}
	}
}
