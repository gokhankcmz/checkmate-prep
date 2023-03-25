package validation

import "cmDeneme/checkmate/settings"

type IValidation interface {
	Validate() []error
	DefaultMessages() *settings.DefaultErrorMessages
}
