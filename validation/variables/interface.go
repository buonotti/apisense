package variables

// EndpointParameter is an interface that defines a parameter that will be interpolated as variable in an endpoint request
type EndpointParameter interface {
	Value(index int) any // Value returns the value of the parameter at the given index
}
