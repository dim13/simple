package simple

import "cmp"

func Sort[S ~[]E, E cmp.Ordered](s S) {
	for i := range len(s) {
		for j := range i {
			if s[i] < s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
}
