package errors

import "errors"

var (
	ErrUnableToConvertMapToProto = errors.New("ErrUnableToConvertMapToProto: some map fields are incompatible with the proto type")
)
