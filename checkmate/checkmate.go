package checkmate

import (
	"cmDeneme/checkmate/customValidation"
	"cmDeneme/checkmate/intValidations"
	"cmDeneme/checkmate/settings"
	"cmDeneme/checkmate/stringValidation"
	"cmDeneme/checkmate/validation"
)

type validator struct {
	settings    *settings.Settings
	validations []validation.IValidation
}

var vs *settings.Settings

func Init(settings *settings.Settings) {
	settings.DefaultErrorMessages = settings.DefaultErrorMessages.WithDefaults()
	vs = settings
}

func New() *validator {
	if vs == nil {
		vs = &settings.Settings{
			StopAtFirstError:     false,
			DefaultErrorMessages: (&settings.DefaultErrorMessages{}).WithDefaults(),
		}
	}
	return &validator{
		settings: vs,
	}
}

func (v *validator) ValidateStruct(value validation.IValidation) {
	v.addValidation(value)
}

func (v *validator) Custom(value func() []error) {
	v.addValidation(&customValidation.CustomValidation{
		F: value,
	})
}
func (v *validator) ValidateInt(value *int, fieldName string) *intValidations.IntValidations {
	if value == nil {
		return &intValidations.IntValidations{}
	}
	vCase := &intValidations.IntValidations{
		Field: *value,
		Validation: validation.Validation{
			FieldName: fieldName,
			Settings:  v.settings,
		},
	}
	v.addValidation(vCase)
	return vCase
}

func (v *validator) ValidateString(value *string, fieldName string) *stringValidation.StringValidation {
	if value == nil {
		return &stringValidation.StringValidation{}
	}
	vCase := &stringValidation.StringValidation{
		Field: *value,
		Validation: validation.Validation{
			FieldName: fieldName,
			Settings:  v.settings,
		},
	}
	v.addValidation(vCase)
	return vCase
}

func (v *validator) addValidation(validation validation.IValidation) {
	v.validations = append(v.validations, validation)
}

func (v *validator) Validate() []error {
	errs := make([]error, 0)
	for _, vr := range v.validations {
		errs = append(errs, vr.Validate()...)
		if v.settings.StopAtFirstError {
			return errs
		}

	}
	return errs
}
