package model

// Devices represents an instance of devices internally in the system
type Devices struct {
	ID              int64
	PremiseID       int64
	PremiseName     string
	PremiseLocation string
	DeviceName      string
	DeviceCode      string
	IsActive        bool
	FloorNumber     int
	DeviceURL       string
}
