package erring

import "errors"

var (
	ErrNoData         = errors.New("no values")
	ErrDuplicateValue = errors.New("duplicate value")
)
