package comparer

type excludesComparer struct{}

func (excludesComparer) Compare(a any, b any) bool {
	return !containsComparer{}.Compare(a, b)
}
