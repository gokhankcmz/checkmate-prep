package stringValidation

import (
	"cmDeneme/checkmate/helper"
	"cmDeneme/checkmate/settings"
	"fmt"
)

type Validation struct {
	Settings       *settings.Settings
	Subject        string
	DefaultMessage string
	Errors         []error
	Rules          []func() error
}

func (v *Validation) Validate() []error {
	for _, r := range v.Rules {
		if err := r(); err != nil {
			v.Errors = append(v.Errors, err)
		}
	}
	return v.Errors
}

func (v *Validation) NotEmpty(message ...string) *Validation {
	v.Rules = append(v.Rules, func() error {
		if v.Subject == "" {
			return helper.FixErrorMessage(nil, v.Settings.NotEmpty, v.DefaultMessage, message...)
		}
		return nil
	})
	return v
}

func (v *Validation) LongerThan(n int, message ...string) *Validation {
	v.Rules = append(v.Rules, func() error {
		if len(v.Subject) <= n {
			return helper.FixErrorMessage(nil, fmt.Sprint("field is not longer than ", n), v.DefaultMessage, message...)
		}
		return nil
	})
	return v
}
