package query

// Definition is a query parameter that should be added to the call
type Definition struct {
	Name  string `yaml:"name" json:"name"`   // Name is the name of the query parameter
	Value string `yaml:"value" json:"value"` // Value is the value of the query parameter
}
