package query

// Definition is a query parameter that should be added to the call
type Definition struct {
	Name  string `yaml:"name" json:"name" validate:"required"`   // Name is the name of the query parameter
	Value string `yaml:"value" json:"value" validate:"required"` // Value is the value of the query parameter
}
