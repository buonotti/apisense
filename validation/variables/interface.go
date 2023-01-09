package variables

type EndpointParameter interface {
	Value(index int) any
}
