package cctv

import (
	"net/http"

	"github.com/nhan1603/CloneScc/api/internal/appconfig/httpserver"
)

const (
	// ErrCodeValidationFailed represents the error code for a failed validation
	ErrCodeValidationFailed = "validation_failed"
)

// Web errors
var (
	webErrorDeviceCodeIsRequired = &httpserver.Error{Status: http.StatusBadRequest, Code: ErrCodeValidationFailed, Desc: "device_code is required"}
	webErrorNotFoundCam          = &httpserver.Error{Status: http.StatusBadRequest, Code: ErrCodeValidationFailed, Desc: "not found cctv for this device"}
)
