package variables

// ConstantEndpointParameter is a parameter that always returns the same value
type ConstantEndpointParameter struct {
	value string // value is the value that is always returned
}

// NewConstantEndpointParameter returns a new EndpointParameter with the given value
func NewConstantEndpointParameter(value string) EndpointParameter {
	return ConstantEndpointParameter{value: value}
}

// Value always returns the same value
func (p ConstantEndpointParameter) Value(int) any {
	return p.value
}
