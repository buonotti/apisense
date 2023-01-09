package variables

type ConstantEndpointParameter struct {
	value string
}

func NewConstantEndpointParameter(value string) ConstantEndpointParameter {
	return ConstantEndpointParameter{value: value}
}

func (p ConstantEndpointParameter) Value(int) any {
	return p.value
}
