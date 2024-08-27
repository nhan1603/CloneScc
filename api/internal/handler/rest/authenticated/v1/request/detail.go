package request

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/nhan1603/CloneScc/api/internal/appconfig/httpserver"
	"github.com/volatiletech/sqlboiler/v4/types"
)

// DetailResp represents the request detail respond
type DetailResp struct {
	ID          string           `json:"id"`
	Title       string           `json:"title"`
	Author      string           `json:"author"`
	Assignee    string           `json:"assignee"`
	Message     string           `json:"message"`
	StartTime   time.Time        `json:"startTime"`
	AlertDetail AlertDetailsResp `json:"alertDetail"`
	Respond     *Respond         `json:"respond"`
}

// AlertDetailsResp represents the alert detail response
type AlertDetailsResp struct {
	ID              string    `json:"id"`
	Type            string    `json:"type"`
	PremiseName     string    `json:"premiseName"`
	PremiseLocation string    `json:"premiseLocation"`
	CCTVDevice      string    `json:"cctvDevice"`
	CCTVDeviceFloor int       `json:"cctvDeviceFloor"`
	IsAcknowledged  bool      `json:"isAcknowledged"`
	IncidentAt      time.Time `json:"incidentAt"`
}

// Respond represents the request respond data
type Respond struct {
	ID         string     `json:"id"`
	User       string     `json:"user"`
	Message    string     `json:"message"`
	MediaData  types.JSON `json:"mediaData"`
	VerifiedAt time.Time  `json:"verifiedAt"`
}

// GetRequestDetail retrieves a request detail
func (h Handler) GetRequestDetail() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		log.Println("[GetRequestDetail] START processing requests")
		ctx := r.Context()

		reqID, err := strconv.ParseInt(chi.URLParam(r, "requestID"), 10, 64)
		if err != nil {
			log.Printf("[GetRequestDetail] failed to get param. Err: %+v\n", err.Error())
			return webErrUnableToGetRequestIDParam
		}

		rs, err := h.requestCtrl.Detail(ctx, reqID)
		if err != nil {
			log.Printf("[GetRequestDetail] failed to get request detail. Err: %+v\n", err.Error())
			return convertCtrlErr(err)
		}

		resp := DetailResp{
			ID:        strconv.FormatInt(rs.ID, 10),
			Title:     rs.Title,
			Author:    rs.Author,
			Assignee:  rs.Assignee,
			Message:   rs.Message,
			StartTime: rs.StartTime,
			AlertDetail: AlertDetailsResp{
				ID:              strconv.FormatInt(rs.AlertDetail.ID, 10),
				Type:            rs.AlertDetail.Type,
				PremiseName:     rs.AlertDetail.PremiseName,
				PremiseLocation: rs.AlertDetail.PremiseLocation,
				CCTVDevice:      rs.AlertDetail.CCTVDevice,
				CCTVDeviceFloor: rs.AlertDetail.CCTVDeviceFloor,
				IsAcknowledged:  rs.AlertDetail.IsAcknowledged,
				IncidentAt:      rs.AlertDetail.IncidentAt,
			},
		}

		if rs.Respond != nil {
			resp.Respond = &Respond{
				ID:         strconv.FormatInt(rs.Respond.ID, 10),
				User:       rs.Respond.User,
				Message:    rs.Respond.Message,
				MediaData:  rs.Respond.MediaData,
				VerifiedAt: rs.Respond.VerifiedAt,
			}
		}

		httpserver.RespondJSON(w, &resp)

		return nil
	})
}
