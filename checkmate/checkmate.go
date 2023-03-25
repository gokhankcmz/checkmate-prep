package checkmate

import (
	"cmDeneme/checkmate/intValidations"
	"cmDeneme/checkmate/settings"
	"cmDeneme/checkmate/stringValidation"
	"cmDeneme/checkmate/validation"
)

type validator struct {
	settings    *settings.Settings
	validations []validation.IValidation
}

var v *settings.Settings

func Init(settings *settings.Settings) {
	settings.DefaultErrorMessages = settings.DefaultErrorMessages.WithDefaults()
	v = settings
}

func New() *validator {
	return &validator{
		settings: v,
	}
}

func (v *validator) ValidateInt(value *int, fieldName string) *intValidations.IntValidations {
	if value == nil {
		return &intValidations.IntValidations{}
	}
	sCase := &intValidations.IntValidations{
		Field:     *value,
		FieldName: fieldName,
		Settings:  v.settings,
	}
	v.addValidation(sCase)
	return sCase
}

func (v *validator) ValidateString(value *string, fieldName string) *stringValidation.StringValidation {
	if value == nil {
		return &stringValidation.StringValidation{}
	}
	sCase := &stringValidation.StringValidation{
		Field:     *value,
		FieldName: fieldName,
		Settings:  v.settings,
	}
	v.addValidation(sCase)
	return sCase
}

func (v *validator) addValidation(validation validation.IValidation) {
	v.validations = append(v.validations, validation)
}

func (v *validator) Validate() []error {
	errors := make([]error, 0)
	for _, vr := range v.validations {
		errors = append(errors, vr.Validate()...)
		if v.settings.StopAtFirstError {
			return errors
		}

	}
	return errors
}
