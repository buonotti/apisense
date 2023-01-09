package variables

func NewVariableEndpointParameter(values []string) VariableEndpointParameter {
	return VariableEndpointParameter{values: values}
}

type VariableEndpointParameter struct {
	values []string
}

func (p VariableEndpointParameter) Value(index int) any {
	return p.values[index]
}
