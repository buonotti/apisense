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

// Where returns all elements in a slice that match the given predicate
func Where[T any](s []T, f func(T) bool) []T {
	res := make([]T, 0)
	for _, v := range s {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

// All returns whether all elements in the given slice match the given predicate
func All[T any](s []T, f func(T) bool) bool {
	if s == nil || len(s) == 0 {
		return false
	}
	for _, v := range s {
		if !f(v) {
			return false
		}
	}
	return true
}

// Any returns whether any element in the given slice matches the given predicate
func Any[T any](s []T, f func(T) bool) bool {
	if s == nil || len(s) == 0 {
		return false
	}
	for _, v := range s {
		if f(v) {
			return true
		}
	}
	return false
}

// Map maps the sequence of items to another using the given transformer function
func Map[TIn any, TOut any](s []TIn, f func(TIn) TOut) []TOut {
	res := make([]TOut, 0)
	for _, v := range s {
		res = append(res, f(v))
	}
	return res
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

// Max returns the max in a slice of integers
func Max(arr []int) int {
	if len(arr) == 0 {
		panic("cannot find max of empty slice")
	}
	maxVal := arr[0]
	for _, elem := range arr[1:] {
		if elem > maxVal {
			maxVal = elem
		}
	}
	return maxVal
}

// FirstOrDefault returns the first value in arr if it has items else fallback
func FirstOrDefault[T any](arr []T, fallback T) T {
	if len(arr) == 0 {
		return fallback
	}
	return arr[0]
}
