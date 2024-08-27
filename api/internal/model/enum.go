package model

// VerificationRequestStatus is enum for verification_requests.status
type VerificationRequestStatus string

const (
	// VerificationRequestStatusNew is an enum representing a new request
	VerificationRequestStatusNew VerificationRequestStatus = "New"
	// VerificationRequestStatusResolved is an enum representing a resolved request
	VerificationRequestStatusResolved VerificationRequestStatus = "Resolved"
)

// ToString convert enum to string
func (vr VerificationRequestStatus) ToString() string {
	return string(vr)
}
