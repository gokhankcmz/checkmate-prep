package helper

import (
	"errors"
	"fmt"
	"strings"
)

func ErrorBuilder(defaultMessage string, message string, args map[string]interface{}) error {
	if message == "" {
		message = defaultMessage
	}

	for k, v := range args {
		message = strings.ReplaceAll(message, fmt.Sprintf("{%s}", k), fmt.Sprint(v))
	}

	return errors.New(RemoveSubstrings(message))
}

func RemoveSubstrings(s string) string {
	var result string
	var inBrackets bool

	for _, c := range s {
		if c == '{' {
			inBrackets = true
			continue
		}

		if c == '}' {
			inBrackets = false
			continue
		}

		if !inBrackets {
			result += string(c)
		}
	}

	return strings.TrimSpace(result)
}
