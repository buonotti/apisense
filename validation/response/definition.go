package response

// Schema describes how the response should look like
type Schema struct {
	Entries []SchemaEntry `toml:"entry"` // Entries are all the field definitions of the response
}

// SchemaEntry is a field definition of the response
type SchemaEntry struct {
	Name         string        `toml:"name"`     // Name is the name of the field
	Type         string        `toml:"type"`     // Type is the type of the field
	Minimum      interface{}   `toml:"min"`      // Minimum is the minimum allowed value of the field
	Maximum      interface{}   `toml:"max"`      // Maximum is the maximum allowed value of the field
	IsRequired   bool          `toml:"required"` // Required is true if the field is required (not null or not empty in case of an array)
	ChildEntries []SchemaEntry `toml:"fields"`   // ChildEntries describe the children of this field if the field is an object or array
}
