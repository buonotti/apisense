package util

import (
	"strings"
)

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

func Join(arr []string, joiner string) string {
	res := strings.Builder{}
	for _, elem := range arr[:len(arr)-1] {
		res.Write([]byte(elem))
		res.Write([]byte(joiner))
	}
	res.Write([]byte(arr[len(arr)-1]))
	return res.String()
}
