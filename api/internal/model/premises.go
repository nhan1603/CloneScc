package model

// Premises represents an instance of premises internally in the system
type Premises struct {
	ID           int64
	Name         string
	Location     string
	PremisesCode string
	Description  string
	CCTVCount    int
}
