package alert

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/nhan1603/CloneScc/api/internal/appconfig/httpserver"
)

// Alert represents the alert respond
type Alert struct {
	ID              string    `json:"id"`
	Type            string    `json:"type"`
	PremiseName     string    `json:"premiseName"`
	PremiseLocation string    `json:"premiseLocation"`
	CCTVDevice      string    `json:"cctvDevice"`
	CCTVDeviceFloor int       `json:"cctvDeviceFloor"`
	IsAcknowledged  bool      `json:"isAcknowledged"`
	IncidentAt      time.Time `json:"incidentAt"`
}

// GetAlertDetail retrieves a alert detail
func (h Handler) GetAlertDetail() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		log.Println("[GetAlertDetail] Starting get request detail ...")
		ctx := r.Context()

		alertID, err := strconv.ParseInt(chi.URLParam(r, "alertID"), 10, 64)
		if err != nil {
			log.Printf("[GetAlertDetail] failed to get param. Err: %+v\n", err.Error())
			return webErrUnableToGetAlertIDParam
		}

		rs, err := h.alertCtrl.Detail(ctx, alertID)
		if err != nil {
			log.Printf("[GetAlertDetail] failed to get alert detail. Err: %+v\n", err.Error())
			return convertCtrlErr(err)
		}

		resp := Alert{
			ID:              strconv.FormatInt(rs.ID, 10),
			Type:            rs.Type,
			PremiseName:     rs.PremiseName,
			PremiseLocation: rs.PremiseLocation,
			CCTVDevice:      rs.CCTVDevice,
			CCTVDeviceFloor: rs.CCTVDeviceFloor,
			IsAcknowledged:  rs.IsAcknowledged,
			IncidentAt:      rs.IncidentAt,
		}

		log.Println("[GetAlertDetail] Get alert detail success")
		httpserver.RespondJSON(w, resp)

		return nil
	})
}
