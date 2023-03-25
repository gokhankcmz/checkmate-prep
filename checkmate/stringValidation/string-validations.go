package stringValidation

import (
	"cmDeneme/checkmate/helper"
	"cmDeneme/checkmate/validation"
	"net/url"
	"strings"
	"unicode"

	"github.com/google/uuid"
)

type StringValidation struct {
	Field string
	validation.Validation
}

func (v *StringValidation) NotEmpty(message ...string) *StringValidation {
	v.AddRule(func() error {
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
	v.AddRule(func() error {
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
	v.AddRule(func() error {
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
	v.AddRule(func() error {
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
	v.AddRule(func() error {
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
	v.AddRule(func() error {
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
	v.AddRule(func() error {
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
	v.AddRule(func() error {
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
	v.AddRule(func() error {
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
	v.AddRule(func() error {
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
	v.AddRule(func() error {
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

func (v *StringValidation) Email(message ...string) *StringValidation {
	//TODO: Regex maybe?
	v.AddRule(func() error {
		emailParts := strings.Split(v.Field, "@")
		if len(emailParts) != 2 {
			return helper.ErrorBuilder(
				v.DefaultMessages().Email,
				strings.Join(message, " "),
				map[string]interface{}{
					"fieldName": v.FieldName,
				})
		}
		localPart := emailParts[0]
		domainPart := emailParts[1]
		if localPart == "" || domainPart == "" {
			return helper.ErrorBuilder(
				v.DefaultMessages().Email,
				strings.Join(message, " "),
				map[string]interface{}{
					"fieldName": v.FieldName,
				})
		}
		if strings.Contains(localPart, " ") || strings.Contains(domainPart, " ") {
			return helper.ErrorBuilder(
				v.DefaultMessages().Email,
				strings.Join(message, " "),
				map[string]interface{}{
					"fieldName": v.FieldName,
				})
		}
		return nil
	})
	return v
}

func (v *StringValidation) Custom(f func() error) *StringValidation {
	v.AddRule(f)
	return v
}
