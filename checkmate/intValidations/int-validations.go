package intValidations

import (
	"cmDeneme/checkmate/helper"
	"cmDeneme/checkmate/settings"
	"strings"
)

type IntValidations struct {
	Field     int
	FieldName string
	Settings  *settings.Settings
	Rules     []func() error
}

func (v *IntValidations) Validate() []error {
	var errors = make([]error, 0)
	for _, r := range v.Rules {
		if err := r(); err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}

func (v *IntValidations) DefaultMessages() *settings.DefaultErrorMessages {
	return &v.Settings.DefaultErrorMessages
}

func (v *IntValidations) NotZero(message ...string) *IntValidations {
	v.Rules = append(v.Rules, func() error {
		if v.Field == 0 {
			return helper.ErrorBuilder(
				v.DefaultMessages().NotZero,
				strings.Join(message, " "),
				map[string]interface{}{
					"fieldName": v.FieldName,
				})
		}
		return nil
	})
	return v
}

func (v *IntValidations) Positive(message ...string) *IntValidations {
	v.Rules = append(v.Rules, func() error {
		if v.Field <= 0 {
			return helper.ErrorBuilder(
				v.DefaultMessages().Positive,
				strings.Join(message, " "),
				map[string]interface{}{
					"fieldName": v.FieldName,
				})
		}
		return nil
	})
	return v
}

func (v *IntValidations) Negative(message ...string) *IntValidations {
	v.Rules = append(v.Rules, func() error {
		if v.Field >= 0 {
			return helper.ErrorBuilder(
				v.DefaultMessages().Negative,
				strings.Join(message, " "),
				map[string]interface{}{
					"fieldName": v.FieldName,
				})
		}
		return nil
	})
	return v
}

func (v *IntValidations) BiggerThan(n int, message ...string) *IntValidations {
	v.Rules = append(v.Rules, func() error {
		if v.Field <= n {
			return helper.ErrorBuilder(
				v.DefaultMessages().BiggerThan,
				strings.Join(message, " "),
				map[string]interface{}{
					"fieldName": v.FieldName,
					"n":         n,
				})
		}
		return nil
	})
	return v
}

func (v *IntValidations) LessThan(n int, message ...string) *IntValidations {
	v.Rules = append(v.Rules, func() error {
		if v.Field >= n {
			return helper.ErrorBuilder(
				v.DefaultMessages().LessThan,
				strings.Join(message, " "),
				map[string]interface{}{
					"fieldName": v.FieldName,
					"n":         n,
				})
		}
		return nil
	})
	return v
}

func (v *IntValidations) InRange(min, max int, message ...string) *IntValidations {
	v.Rules = append(v.Rules, func() error {
		if v.Field < min || v.Field > max {
			return helper.ErrorBuilder(
				v.DefaultMessages().InRange,
				strings.Join(message, " "),
				map[string]interface{}{
					"fieldName": v.FieldName,
					"min":       min,
					"max":       max,
				})
		}
		return nil
	})
	return v
}

func (v *IntValidations) Exactly(n int, message ...string) *IntValidations {
	v.Rules = append(v.Rules, func() error {
		if v.Field != n {
			return helper.ErrorBuilder(
				v.DefaultMessages().Exactly,
				strings.Join(message, " "),
				map[string]interface{}{
					"fieldName": v.FieldName,
					"n":         n,
				})
		}
		return nil
	})
	return v
}

func (v *IntValidations) EvenNumber(message ...string) *IntValidations {
	v.Rules = append(v.Rules, func() error {
		if v.Field%2 != 0 {
			return helper.ErrorBuilder(
				v.DefaultMessages().EvenNumber,
				strings.Join(message, " "),
				map[string]interface{}{
					"fieldName": v.FieldName,
				})
		}
		return nil
	})
	return v
}

func (v *IntValidations) OddNumber(message ...string) *IntValidations {
	v.Rules = append(v.Rules, func() error {
		if v.Field%2 == 0 {
			return helper.ErrorBuilder(
				v.DefaultMessages().OddNumber,
				strings.Join(message, " "),
				map[string]interface{}{
					"fieldName": v.FieldName,
				})
		}
		return nil
	})
	return v
}
