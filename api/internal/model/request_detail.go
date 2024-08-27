package model

import (
	"time"

	"github.com/volatiletech/sqlboiler/v4/types"
)

// RequestDetail represents the request detail
type RequestDetail struct {
	ID          int64
	Title       string
	Author      string
	Assignee    string
	Message     string
	StartTime   time.Time
	AlertDetail Alert
	Respond     *RequestRespond
}

// RequestRespond represents the request respond data
type RequestRespond struct {
	ID         int64
	User       string
	Message    string
	MediaData  types.JSON
	VerifiedAt time.Time
}
