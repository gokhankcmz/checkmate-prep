package checkmate

import (
	settings2 "cmDeneme/checkmate/settings"
	"cmDeneme/checkmate/stringValidation"
	"cmDeneme/checkmate/validation"
)

type validator struct {
	settings    *settings2.Settings
	validations []validation.IValidation
}

func New(settings *settings2.Settings) *validator {
	return &validator{
		settings: settings,
	}
}

func (v *validator) ValidateString(value string, defaultMessage ...string) *stringValidation.Validation {
	sCase := &stringValidation.Validation{
		Subject:  value,
		Settings: v.settings,
	}
	if len(defaultMessage) > 0 {
		sCase.DefaultMessage = defaultMessage[0]
	}
	v.addValidation(sCase)
	return sCase
}

func (v *validator) ValidateStringPtr(value *string, defaultMessage ...string) *stringValidation.Validation {
	if value == nil {
		return &stringValidation.Validation{}
	}
	return v.ValidateString(*value, defaultMessage...)
}

func (v *validator) addValidation(validation validation.IValidation) {
	v.validations = append(v.validations, validation)
}

func (v *validator) Validate() []error {
	errors := make([]error, 0)
	for _, vr := range v.validations {
		errors = append(errors, vr.Validate()...)

	}
	return errors
}
