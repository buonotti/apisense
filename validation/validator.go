package validation

// Validator is an interface that all validators in the pipeline must implement
type Validator interface {
	Name() string                         // Name returns the name of the validator
	Validate(item PipelineTestCase) error // Validate validates the given item and return nil on success or an error on failure
	IsFatal() bool                        // IsFatal returns true if the validator is fatal and the pipeline should stop on failure
}
