package validation

type IValidation interface {
	Validate() []error
}
