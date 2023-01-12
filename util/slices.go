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

// Transpose returns the matrix transposed from (nxm) to (mxn)
func Transpose[T any](slice [][]T) [][]T {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]T, xl)
	for i := range result {
		result[i] = make([]T, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}
