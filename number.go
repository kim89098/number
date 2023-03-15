package number

import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Bound[T constraints.Ordered](v, min, max T) T {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}
