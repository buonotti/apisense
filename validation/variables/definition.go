package variables

// Definition describes a variable that should be interpolated in the base url and the query parameters
type Definition struct {
	Name       string   `toml:"name"`     // Name is the name of the variable
	IsConstant bool     `toml:"constant"` // IsConstant is true if the value of the variable is constant or else false
	Values     []string `toml:"values"`   // Values are all the possible values of the variable (only 1 in case of a constant)
}
