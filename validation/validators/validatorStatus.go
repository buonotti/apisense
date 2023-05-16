package validators

type ValidatorStatus string

const (
	ValidatorStatusUnknown ValidatorStatus = "unknown"
	ValidatorStatusSuccess ValidatorStatus = "success"
	ValidatorStatusSkipped ValidatorStatus = "skipped"
	ValidatorStatusFail    ValidatorStatus = "fail"
)
