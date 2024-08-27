package requests

import "errors"

var (
	// ErrNotFound means item not found in db
	ErrNotFound = errors.New("request not found")
	// ErrRequestResolved means request is resolved
	ErrRequestResolved = errors.New("request resolved")
)
