package variables

// NewVariableEndpointParameter returns a new EndpointParameter with the given values
func NewVariableEndpointParameter(values []string) EndpointParameter {
	return VariableEndpointParameter{values: values}
}

// VariableEndpointParameter is a parameter that returns a different value based on the index in the given collection
type VariableEndpointParameter struct {
	values []string // values is the collection of values that will be returned
}

// Value returns the value in the initial collection at the given index
func (p VariableEndpointParameter) Value(index int) any {
	return p.values[index]
}
