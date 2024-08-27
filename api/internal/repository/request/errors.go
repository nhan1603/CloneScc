package request

import "errors"

var (
	// ErrNotFound means the request record was not found in db
	ErrNotFound = errors.New("request model not found")
	// ErrCanNotLoadRelatedTablesData means cannot load related tables data
	ErrCanNotLoadRelatedTablesData = errors.New("cannot load related tables data")
)
