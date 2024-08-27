package request

import (
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/nhan1603/CloneScc/api/internal/appconfig/httpserver"
	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/volatiletech/null/v8"
)

// VerificationRequestResp represents the requests page
type VerificationRequestResp struct {
	Items      []RequestResp    `json:"items"`
	Pagination model.Pagination `json:"pagination"`
}

// RequestResp represents the request respond
type RequestResp struct {
	ID              string    `json:"id"`
	AlertID         string    `json:"alertID"`
	Alert           string    `json:"alert"`
	AlertType       string    `json:"alertType"`
	PremiseName     string    `json:"premiseName"`
	PremiseLocation string    `json:"premiseLocation"`
	Author          string    `json:"author"`
	Assignee        string    `json:"assignee"`
	Status          string    `json:"status"`
	StartTime       time.Time `json:"startTime"`
	VerifiedAt      null.Time `json:"verifiedAt"`
}

// GetRequests retrieves a page of request
func (h Handler) GetRequests() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		log.Println("[GetRequests] START processing requests")
		ctx := r.Context()

		inp, err := validateAndConvertInput(r.URL.Query())
		if err != nil {
			log.Printf("[GetRequests] failed to validate. Err: %+v\n", err.Error())
			return err
		}

		requests, totalCount, err := h.requestCtrl.List(ctx, inp)
		if err != nil {
			log.Printf("[GetRequests] failed to get requests. Err: %+v\n", err.Error())
			return err
		}
		var requestsPage VerificationRequestResp
		for _, item := range requests {
			requestsPage.Items = append(requestsPage.Items, RequestResp{
				ID:              strconv.FormatInt(item.ID, 10),
				AlertID:         strconv.FormatInt(item.AlertID, 10),
				Alert:           item.Alert,
				AlertType:       item.AlertType,
				PremiseName:     item.PremiseName,
				PremiseLocation: item.PremiseLocation,
				Author:          item.Author,
				Assignee:        item.Assignee,
				Status:          item.Status,
				StartTime:       item.StartTime,
				VerifiedAt:      item.VerifiedAt,
			})
		}

		requestsPage.Pagination = model.Pagination{
			TotalCount:  totalCount,
			CurrentPage: inp.Page,
			Limit:       inp.Limit,
		}

		httpserver.RespondJSON(w, requestsPage)
		return nil
	})
}

func validateAndConvertInput(req url.Values) (model.GetRequestsInput, error) {
	var inp model.GetRequestsInput

	reqPremiseID := strings.TrimSpace(req.Get("premiseID"))
	if len(reqPremiseID) > 0 {
		premiseID, err := strconv.ParseInt(reqPremiseID, 10, 64)
		if err != nil {
			return model.GetRequestsInput{}, webErrGetRequestsDueToInvalidPremiseID
		}
		inp.PremiseID = premiseID
	}

	reqAssigneeID := strings.TrimSpace(req.Get("assigneeID"))
	if len(reqAssigneeID) > 0 {
		assigneeID, err := strconv.ParseInt(reqAssigneeID, 10, 64)
		if err != nil {
			return model.GetRequestsInput{}, webErrGetRequestsDueToInvalidAssigneeID
		}
		inp.AssigneeID = assigneeID
	}

	reqLimit := strings.TrimSpace(req.Get("limit"))
	if len(reqLimit) > 0 {
		limit, err := strconv.Atoi(reqLimit)
		if err != nil {
			return model.GetRequestsInput{}, webErrGetRequestsDueToInvalidLimit
		}

		if limit <= 0 {
			return model.GetRequestsInput{}, webErrLimitMustBeGreaterThanZero
		}

		if limit > model.PaginationMaxLimit {
			return model.GetRequestsInput{}, webErrMaxLimitMustBeLessThanNumber
		}

		inp.Limit = limit
	}

	reqPage := strings.TrimSpace(req.Get("page"))
	if len(reqPage) > 0 {
		page, err := strconv.Atoi(reqPage)
		if err != nil {
			return model.GetRequestsInput{}, webErrGetRequestsDueToInvalidPage
		}
		if page <= 0 {
			return model.GetRequestsInput{}, webErrPageMustBeGreaterThanZero
		}
		inp.Page = page
	}

	if inp.Limit == 0 {
		inp.Limit = model.PaginationDefaultLimit
	}

	if inp.Page == 0 {
		inp.Page = model.PaginationDefaultPage
	}

	return inp, nil
}
