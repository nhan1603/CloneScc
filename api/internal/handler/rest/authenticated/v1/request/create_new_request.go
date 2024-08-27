package request

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/nhan1603/CloneScc/api/internal/appconfig/httpserver"
	"github.com/nhan1603/CloneScc/api/internal/controller/requests"
)

type createNewRequestRequest struct {
	AlertID        string `json:"alertId"`
	RequestBy      string `json:"requestBy"`
	AssignedUserID string `json:"assignedUserId"`
	Content        string `json:"content"`
}

// ValidateAndMap validates the input and maps it to the CreateRequestInput type
func (r createNewRequestRequest) ValidateAndMap() (requests.CreateRequestInput, error) {
	alertID, err := strconv.ParseInt(r.AlertID, 10, 64)
	if err != nil || alertID <= 0 {
		return requests.CreateRequestInput{}, webErrInvalidAlertID
	}

	requestBy, err := strconv.ParseInt(r.RequestBy, 10, 64)
	if err != nil || requestBy <= 0 {
		return requests.CreateRequestInput{}, webErrInvalidRequestBy
	}

	assignedUserID, err := strconv.ParseInt(r.AssignedUserID, 10, 64)
	if err != nil || assignedUserID <= 0 {
		return requests.CreateRequestInput{}, webErrInvalidAssignedUserID
	}

	if r.RequestBy == r.AssignedUserID {
		return requests.CreateRequestInput{}, webErrInvalidRequestAndAssign
	}

	return requests.CreateRequestInput{
		AlertID:        alertID,
		RequestBy:      requestBy,
		AssignedUserID: assignedUserID,
		Content:        r.Content,
	}, nil
}

type successResponse struct {
	Success bool `json:"success"`
}

// CreateNewRequest create a new request for alert
func (h Handler) CreateNewRequest() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		log.Println("[CreateNewRequest] START processing requests")
		ctx := r.Context()

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("[CreateNewRequest] read body error %v", err.Error())
			return webErrInvalidRequestBody
		}

		var request createNewRequestRequest
		if err := json.Unmarshal(body, &request); err != nil {
			log.Printf("[CreateNewRequest] Unmarshal error %v", err.Error())
			return err
		}

		inp, err := request.ValidateAndMap()
		if err != nil {
			log.Printf("[CreateNewRequest] invalid request error %v", err.Error())
			return err
		}

		if err := h.requestCtrl.CreateRequest(ctx, inp); err != nil {
			return err
		}

		httpserver.RespondJSON(w, successResponse{Success: true})
		return nil
	})
}
