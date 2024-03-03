package messages

import (
	"fmt"
	"strings"
)

var (
	MessageNoRequestBodyProvided = func(funcName string) string {
		return fmt.Sprintf("%s: no request body provided", funcName)
	}

	MessageProvideAtleastOneOfTheFollowing = func(funcName string, propNames []string) string {
		vars := strings.Join(propNames, ", ")

		return fmt.Sprintf("%s: provide at least one of the following `%s`", funcName, vars)
	}

	MessageNoTargetProvided = func(funcName string) string {
		return fmt.Sprintf("%s: no target provided", funcName)
	}

	MessageRequiredFieldNotProvided = func(funcName string, fieldName string) string {
		return fmt.Sprintf("%s: required field `%s` not provided", funcName, fieldName)
	}
)
