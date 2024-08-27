// Package model includes the necessary data model for internal process
package model

import (
	"time"

	"github.com/volatiletech/null/v8"
)

// GetRequestsInput represents for input to get requests
type GetRequestsInput struct {
	PremiseID  int64
	AssigneeID int64
	Limit      int
	Page       int
}

// RequestSummary represent model summary for request
type RequestSummary struct {
	ID              int64
	AlertID         int64
	Alert           string
	AlertType       string
	PremiseName     string
	PremiseLocation string
	Author          string
	Assignee        string
	Message         string
	Status          string
	StartTime       time.Time
	EndTime         time.Time
	VerifiedAt      null.Time
}

// VerificationRequest represent model for verification_requests table
type VerificationRequest struct {
	ID             int64
	AlertID        int64
	RequestBy      int64
	AssignedUserID int64
	Message        string
	Status         VerificationRequestStatus
	StartTime      time.Time
	EndTime        time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// VerificationRequestResponses represent model for verification_request_responses table
type VerificationRequestResponses struct {
	ID                    int64
	VerificationRequestID int64
	Message               string
	MediaData             []MediaData
	VerifiedAt            time.Time
	CreatedAt             time.Time
	UpdatedAt             time.Time
}

// MediaData represent VerificationRequestDetail.MediaData
type MediaData struct {
	FileName      string
	FileExtension string
	URL           string
}

// ResponseMessage represents reponse info
type ResponseMessage struct {
	RequestID   string    `json:"requestId"`
	PremiseName string    `json:"premiseName"`
	CctvName    string    `json:"cctvName"`
	AlertID     string    `json:"alertId"`
	IncidentAt  time.Time `json:"incidentAt"`
}
