package util

func FindFirst[T any](s []T, f func(T) bool) T {
	if s == nil || len(s) == 0 {
		panic("util: FindFirst called on nil or empty slice")
	}
	for _, v := range s {
		if f(v) {
			return v
		}
	}
	panic("util: FindFirst called on slice with no matching element")
}
