package response

// SchemaEntry is a field definition of the response
type SchemaEntry struct {
	Name       string        `yaml:"name"`     // Name is the name of the field
	Type       string        `yaml:"type"`     // Type is the type of the field
	Minimum    interface{}   `yaml:"min"`      // Minimum is the minimum allowed value of the field
	Maximum    interface{}   `yaml:"max"`      // Maximum is the maximum allowed value of the field
	IsRequired bool          `yaml:"required"` // Required is true if the field is required (not null or not empty in case of an array)
	Fields     []SchemaEntry `yaml:"fields"`   // Fields describe the children of this field if the field is an object or array
}
