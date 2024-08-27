package alerts

import "errors"

var (
	// ErrNotFound means item not found in db
	ErrNotFound = errors.New("alert not found")
	// ErrCreateAlert means cannot create alert instance in the database
	ErrCreateAlert = errors.New("error create alert")
	// ErrCCTVNotFound means cctv item not found in db
	ErrCCTVNotFound = errors.New("cctv not found")
)
