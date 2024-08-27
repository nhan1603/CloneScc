package alert

import "errors"

var (
	// ErrNotFound means the request record was not found in db
	ErrNotFound = errors.New("alert model not found")
)
