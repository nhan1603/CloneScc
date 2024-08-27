package model

import (
	"time"
)

// GetAlertsInput represents for input to get alerts
type GetAlertsInput struct {
	PremiseID int64
	Limit     int
	Page      int
}

// Alert represents the alert
type Alert struct {
	ID              int64
	CCTVDeviceID    int64
	Type            string
	Description     string
	PremiseName     string
	PremiseLocation string
	CCTVDevice      string
	CCTVDeviceFloor int
	IsAcknowledged  bool
	IncidentAt      time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// AlertMessage represents alert info
type AlertMessage struct {
	ID          string    `json:"id"`
	CCTVName    string    `json:"cctvName"`
	FloorNumber string    `json:"floorNumber"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	IncidentAt  time.Time `json:"incidentAt"`
}

// AlertType is enum for alert.type
type AlertType string

const (
	// AlertTypeUnauthorizedAccess is an enum representing an unauthorized access
	AlertTypeUnauthorizedAccess AlertType = "Unauthorized Access"
	// AlertTypeSuspiciousActivities is an enum representing an suspicious activities
	AlertTypeSuspiciousActivities AlertType = "Suspicious Activities"
	// AlertTypePropertyDamage is an enum representing an property damage
	AlertTypePropertyDamage AlertType = "Property Damage"
)

// ToString convert enum to string
func (vr AlertType) ToString() string {
	return string(vr)
}
