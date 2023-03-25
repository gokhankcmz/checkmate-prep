package validation

import "cmDeneme/checkmate/settings"

type IValidation interface {
	Validate() []error
}

type Validation struct {
	FieldName string
	Settings  *settings.Settings
	Rules     []func() error
}

func (v *Validation) Validate() []error {
	var errors = make([]error, 0)
	for _, r := range v.Rules {
		if err := r(); err != nil {
			errors = append(errors, err)
			if v.Settings.StopAtFirstError {
				return errors
			}
		}
	}
	return errors
}

func (v *Validation) AddRule(rule func() error) {
	v.Rules = append(v.Rules, rule)
}
func (v *Validation) DefaultMessages() *settings.DefaultErrorMessages {
	return &v.Settings.DefaultErrorMessages
}
