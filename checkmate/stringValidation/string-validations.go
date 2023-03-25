package stringValidation

import (
	"cmDeneme/checkmate/helper"
	"cmDeneme/checkmate/settings"
	"net/url"
	"strings"
	"unicode"

	"github.com/google/uuid"
)

type StringValidation struct {
	Field     string
	FieldName string
	Settings  *settings.Settings
	Rules     []func() error
}

func (v *StringValidation) Validate() []error {
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

func (v *StringValidation) DefaultMessages() *settings.DefaultErrorMessages {
	return &v.Settings.DefaultErrorMessages
}

func (v *StringValidation) NotEmpty(message ...string) *StringValidation {
	v.Rules = append(v.Rules, func() error {
		if v.Field == "" {
			return helper.ErrorBuilder(
				v.DefaultMessages().NotEmpty,
				strings.Join(message, " "),
				map[string]interface{}{
					"fieldName": v.FieldName,
				})
		}
		return nil
	})
	return v
}

func (v *StringValidation) MinLength(n int, message ...string) *StringValidation {
	v.Rules = append(v.Rules, func() error {
		if len(v.Field) <= n {
			return helper.ErrorBuilder(
				v.DefaultMessages().MinLength,
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

func (v *StringValidation) MaxLength(n int, message ...string) *StringValidation {
	v.Rules = append(v.Rules, func() error {
		if len(v.Field) > n {
			return helper.ErrorBuilder(
				v.DefaultMessages().MaxLength,
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

func (v *StringValidation) ExactLength(n int, message ...string) *StringValidation {
	v.Rules = append(v.Rules, func() error {
		if len(v.Field) != n {
			return helper.ErrorBuilder(
				v.DefaultMessages().ExactLength,
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

func (v *StringValidation) InRange(min, max int, message ...string) *StringValidation {
	v.Rules = append(v.Rules, func() error {
		if len(v.Field) < min || len(v.Field) > max {
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

func (v *StringValidation) Contain(substr string, message ...string) *StringValidation {
	v.Rules = append(v.Rules, func() error {
		if !strings.Contains(v.Field, substr) {
			return helper.ErrorBuilder(
				v.DefaultMessages().Contain,
				strings.Join(message, " "),
				map[string]interface{}{
					"fieldName": v.FieldName,
					"substr":    substr,
				})
		}
		return nil
	})
	return v
}

func (v *StringValidation) Alphabetic(message ...string) *StringValidation {
	v.Rules = append(v.Rules, func() error {
		runes := []rune(strings.ReplaceAll(v.Field, " ", ""))
		for _, r := range runes {
			if !unicode.IsLetter(r) {
				return helper.ErrorBuilder(
					v.DefaultMessages().Alphabetic,
					strings.Join(message, " "),
					map[string]interface{}{
						"fieldName": v.FieldName,
					})
			}
		}
		return nil
	})
	return v
}

func (v *StringValidation) Numeric(message ...string) *StringValidation {
	v.Rules = append(v.Rules, func() error {
		runes := []rune(v.Field)
		for _, r := range runes {
			if !unicode.IsDigit(r) {
				return helper.ErrorBuilder(
					v.DefaultMessages().Numeric,
					strings.Join(message, " "),
					map[string]interface{}{
						"fieldName": v.FieldName,
					})
			}
		}
		return nil
	})
	return v
}

func (v *StringValidation) NoWhiteSpace(message ...string) *StringValidation {
	v.Rules = append(v.Rules, func() error {
		if strings.Contains(v.Field, " ") {
			return helper.ErrorBuilder(
				v.DefaultMessages().NoWhiteSpace,
				strings.Join(message, " "),
				map[string]interface{}{
					"fieldName": v.FieldName,
				})
		}
		return nil
	})
	return v
}

func (v *StringValidation) URL(message ...string) *StringValidation {
	v.Rules = append(v.Rules, func() error {
		_, err := url.ParseRequestURI(v.Field)
		if err != nil {
			return helper.ErrorBuilder(
				v.DefaultMessages().URL,
				strings.Join(message, " "),
				map[string]interface{}{
					"fieldName": v.FieldName,
					"err":       err.Error(),
				})
		}
		return nil
	})
	return v
}

func (v *StringValidation) UUID(message ...string) *StringValidation {
	v.Rules = append(v.Rules, func() error {
		_, err := uuid.Parse(v.Field)
		if err != nil {
			return helper.ErrorBuilder(
				v.DefaultMessages().UUID,
				strings.Join(message, " "),
				map[string]interface{}{
					"fieldName": v.FieldName,
					"err":       err.Error(),
				})
		}
		return nil
	})
	return v
}
