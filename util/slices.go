package util

// FindFirst return a reference to the first element of the slice that matches
// the predicate. It returns nil if the slice is nil or empty or if no item
// matches the predicate.
func FindFirst[T any](s []T, f func(T) bool) *T {
	if s == nil || len(s) == 0 {
		return nil
	}
	for _, v := range s {
		if f(v) {
			return &v
		}
	}
	return nil
}

// Contains returns whether the given slice contains the given element
func Contains[T comparable](s []T, v T) bool {
	if s == nil || len(s) == 0 {
		return false
	}
	for _, e := range s {
		if e == v {
			return true
		}
	}
	return false
}
