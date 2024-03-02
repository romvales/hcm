package errors

import "errors"

var (
	ErrInvalidClientFromRequestUnimplemented = errors.New("ErrInvalidClientFromRequest: unimplemented")
	ErrIndexItemNotFound                     = errors.New("ErrIndexNotFound: unable to find such item in the database")
)
