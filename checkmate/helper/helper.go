package helper

import (
	"errors"
	"strings"
)

func FixErrorMessage(err error, defaultMessage, fieldMessage string, ruleMessage ...string) error {
	text := strings.Join(ruleMessage, " ")

	if text == "" {
		text = fieldMessage
	}

	if text == "" {
		text = defaultMessage
	}

	if err == nil {
		return errors.New(strings.ReplaceAll(text, "{err}", ""))
	}
	return errors.New(strings.ReplaceAll(text, "{err}", err.Error()))
}
