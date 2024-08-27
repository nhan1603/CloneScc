package model

import (
	"time"
)

// UserRole represents the role of the user
type UserRole string

const (
	// UserRoleOperationUser defines role as operation user
	UserRoleOperationUser UserRole = "OPERATION_USER"
	// UserRoleSecurityGuard defines role as security guard
	UserRoleSecurityGuard UserRole = "SECURITY_GUARD"
)

// String converts to string value
func (r UserRole) String() string {
	return string(r)
}

// User represents the user
type User struct {
	ID          int64
	DisplayName string
	Email       string
	Password    string
	Role        UserRole
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
