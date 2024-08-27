package alert

import (
	"net/http"

	"github.com/nhan1603/CloneScc/api/internal/appconfig/httpserver"
	"github.com/nhan1603/CloneScc/api/internal/controller/alerts"
)

const (
	// ErrCodeValidationFailed represents the error code for a failed validation
	ErrCodeValidationFailed = "validation_failed"
)

var (
	webErrGetAlertsDueToInvalidPremiseID = &httpserver.Error{Status: http.StatusBadRequest, Code: ErrCodeValidationFailed, Desc: "invalid premise id"}
	webErrGetAlertsDueToInvalidLimit     = &httpserver.Error{Status: http.StatusBadRequest, Code: ErrCodeValidationFailed, Desc: "invalid limit"}
	webErrGetAlertsDueToInvalidPage      = &httpserver.Error{Status: http.StatusBadRequest, Code: ErrCodeValidationFailed, Desc: "invalid page"}
	webErrLimitMustBeGreaterThanZero     = &httpserver.Error{Status: http.StatusBadRequest, Code: ErrCodeValidationFailed, Desc: "limit must be greater than 0"}
	webErrPageMustBeGreaterThanZero      = &httpserver.Error{Status: http.StatusBadRequest, Code: ErrCodeValidationFailed, Desc: "page must be greater than 0"}
	webErrMaxLimitMustBeLessThanNumber   = &httpserver.Error{Status: http.StatusBadRequest, Code: ErrCodeValidationFailed, Desc: "limit must be less than 1000"}
	webErrUnableToGetAlertIDParam        = &httpserver.Error{Status: http.StatusBadRequest, Code: "invalid_request_body", Desc: "unable to get alert id parameter"}
	webErrGetAlertDetailNotFound         = &httpserver.Error{Status: http.StatusBadRequest, Code: "alert_not_found", Desc: "the alert not found"}
)

func convertCtrlErr(err error) error {
	switch err {
	case alerts.ErrNotFound:
		return webErrGetAlertDetailNotFound
	default:
		return err
	}
}
