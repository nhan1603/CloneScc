package asset

import "errors"

var (
	ErrCctvNotFound     = errors.New("cctv not found")
	ErrVerificationData = errors.New("Cannot retrieve verification request data")
	// ErrCanNotLoadPremisesTable means cannot load premises table
	ErrCanNotLoadPremisesTable = errors.New("cannot load premises table")
)
