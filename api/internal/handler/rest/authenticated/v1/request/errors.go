package request

import (
	"net/http"

	"github.com/nhan1603/CloneScc/api/internal/appconfig/httpserver"
	"github.com/nhan1603/CloneScc/api/internal/controller/requests"
)

const (
	errCodeValidationFailed   = "validation_failed"
	errCodeInvalidRequestBody = "invalid_request_body"
)

var (
	webErrUnableToGetRequestIDParam         = &httpserver.Error{Status: http.StatusBadRequest, Code: errCodeInvalidRequestBody, Desc: "unable to get request id parameter"}
	webErrGetRequestDetailNotFound          = &httpserver.Error{Status: http.StatusBadRequest, Code: "request_not_found", Desc: "the request not found"}
	webErrGetRequestsDueToInvalidPremiseID  = &httpserver.Error{Status: http.StatusBadRequest, Code: errCodeValidationFailed, Desc: "invalid premise id"}
	webErrGetRequestsDueToInvalidAssigneeID = &httpserver.Error{Status: http.StatusBadRequest, Code: errCodeValidationFailed, Desc: "invalid assignee id"}
	webErrGetRequestsDueToInvalidLimit      = &httpserver.Error{Status: http.StatusBadRequest, Code: errCodeValidationFailed, Desc: "invalid limit"}
	webErrGetRequestsDueToInvalidPage       = &httpserver.Error{Status: http.StatusBadRequest, Code: errCodeValidationFailed, Desc: "invalid page"}
	webErrLimitMustBeGreaterThanZero        = &httpserver.Error{Status: http.StatusBadRequest, Code: errCodeValidationFailed, Desc: "limit must be greater than 0"}
	webErrPageMustBeGreaterThanZero         = &httpserver.Error{Status: http.StatusBadRequest, Code: errCodeValidationFailed, Desc: "page must be greater than 0"}
	webErrMaxLimitMustBeLessThanNumber      = &httpserver.Error{Status: http.StatusBadRequest, Code: errCodeValidationFailed, Desc: "limit must be less than 1000"}

	webErrInvalidAlertID          = &httpserver.Error{Status: http.StatusBadRequest, Code: errCodeValidationFailed, Desc: "invalid alertID"}
	webErrInvalidRequestBy        = &httpserver.Error{Status: http.StatusBadRequest, Code: errCodeValidationFailed, Desc: "invalid requestBy"}
	webErrInvalidAssignedUserID   = &httpserver.Error{Status: http.StatusBadRequest, Code: errCodeValidationFailed, Desc: "invalid assignedUserId"}
	webErrInvalidRequestAndAssign = &httpserver.Error{Status: http.StatusBadRequest, Code: errCodeValidationFailed, Desc: "requester and assigned user should be different"}
	webErrInvalidRequestBody      = &httpserver.Error{Status: http.StatusBadRequest, Code: errCodeInvalidRequestBody, Desc: "invalid request body"}
	webErrExceedMaxMemoryAllowed  = &httpserver.Error{Status: http.StatusBadRequest, Code: errCodeInvalidRequestBody, Desc: "max memory allowed is 100MB"}
	webErrExceedMaxMediaAllowed   = &httpserver.Error{Status: http.StatusBadRequest, Code: errCodeInvalidRequestBody, Desc: "max media allowed is 20"}
	webErrInvalidRequestID        = &httpserver.Error{Status: http.StatusBadRequest, Code: errCodeValidationFailed, Desc: "invalid requestID"}
	webErrMessageIsRequired       = &httpserver.Error{Status: http.StatusBadRequest, Code: errCodeValidationFailed, Desc: "message is required"}
	webErrRequestAlreadyResolved  = &httpserver.Error{Status: http.StatusBadRequest, Code: errCodeValidationFailed, Desc: "this request is already resolved"}
)

func convertCtrlErr(err error) error {
	switch err {
	case requests.ErrNotFound:
		return webErrGetRequestDetailNotFound
	case requests.ErrRequestResolved:
		return webErrRequestAlreadyResolved
	default:
		return err
	}
}
