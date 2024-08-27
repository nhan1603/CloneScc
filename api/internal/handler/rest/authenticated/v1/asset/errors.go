// Package asset includes the neccessary api for retrieving asset-related data
package asset

import (
	"net/http"

	"github.com/nhan1603/CloneScc/api/internal/appconfig/httpserver"
)

const (
	errCodeValidationFailed   = "validation_failed"
	errCodeInvalidRequestBody = "invalid_request_body"
	errCodeInternal           = "internal_error"
)

// Error codes
var (
	// 4xx
	webErrInvalidID                    = &httpserver.Error{Status: http.StatusBadRequest, Code: errCodeValidationFailed, Desc: "invalid id"}
	webErrInvalidUserID                = &httpserver.Error{Status: http.StatusBadRequest, Code: errCodeValidationFailed, Desc: "invalid alertID"}
	webErrInvalidNumber                = &httpserver.Error{Status: http.StatusBadRequest, Code: errCodeValidationFailed, Desc: "invalid number"}
	webErrLimitMustBeGreaterThanZero   = &httpserver.Error{Status: http.StatusBadRequest, Code: errCodeValidationFailed, Desc: "limit must be greater than 0"}
	webErrPageMustBeGreaterThanZero    = &httpserver.Error{Status: http.StatusBadRequest, Code: errCodeValidationFailed, Desc: "page must be greater than 0"}
	webErrMaxLimitMustBeLessThanNumber = &httpserver.Error{Status: http.StatusBadRequest, Code: errCodeValidationFailed, Desc: "limit must be less than 1000"}
	webErrDeviceTokenIsRequired        = &httpserver.Error{Status: http.StatusBadRequest, Code: errCodeValidationFailed, Desc: "deviceToken is required"}
	webErrPlatformIsRequired           = &httpserver.Error{Status: http.StatusBadRequest, Code: errCodeValidationFailed, Desc: "platform is required"}
	webErrInvalidLimit                 = &httpserver.Error{Status: http.StatusBadRequest, Code: errCodeValidationFailed, Desc: "limit is in invalid format"}
	webErrInvalidPage                  = &httpserver.Error{Status: http.StatusBadRequest, Code: errCodeValidationFailed, Desc: "page is in invalid format"}
	webErrInvalidRequestBody           = &httpserver.Error{Status: http.StatusBadRequest, Code: errCodeInvalidRequestBody, Desc: "invalid request body"}

	// 5xx
	webErrInternalError = &httpserver.Error{Status: http.StatusInternalServerError, Code: errCodeInternal, Desc: "Something went wrong"}
)
