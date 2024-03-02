package errors

import "errors"

var ()

func ErrMissingRequestParameter(msg string) error {
	return errors.New(msg)
}
