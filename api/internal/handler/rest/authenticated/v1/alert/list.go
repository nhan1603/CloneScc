package alert

import (
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/nhan1603/CloneScc/api/internal/appconfig/httpserver"
	"github.com/nhan1603/CloneScc/api/internal/model"
)

// AlertsPage represents the alerts page respond
type AlertsPage struct {
	Items      []Alert          `json:"items"`
	Pagination model.Pagination `json:"pagination"`
}

// GetAlerts retrieves a list of alert
func (h Handler) GetAlerts() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		log.Println("[GetAlerts] Starting get alerts ...")
		ctx := r.Context()

		inp, err := validateAndConvertInput(r.URL.Query())
		if err != nil {
			log.Printf("[GetAlerts] failed to validate. Err: %+v\n", err.Error())
			return err
		}

		alerts, totalCount, err := h.alertCtrl.List(ctx, inp)
		if err != nil {
			log.Printf("[GetAlerts] failed to get alerts. Err: %+v\n", err.Error())
			return err
		}
		var alertsPage AlertsPage
		for _, item := range alerts {
			alertsPage.Items = append(alertsPage.Items, Alert{
				ID:              strconv.FormatInt(item.ID, 10),
				Type:            item.Type,
				PremiseName:     item.PremiseName,
				PremiseLocation: item.PremiseLocation,
				CCTVDevice:      item.CCTVDevice,
				CCTVDeviceFloor: item.CCTVDeviceFloor,
				IsAcknowledged:  item.IsAcknowledged,
				IncidentAt:      item.IncidentAt,
			})
		}

		alertsPage.Pagination = model.Pagination{
			TotalCount:  totalCount,
			CurrentPage: inp.Page,
			Limit:       inp.Limit,
		}

		httpserver.RespondJSON(w, alertsPage)
		return nil
	})
}

func validateAndConvertInput(req url.Values) (model.GetAlertsInput, error) {
	var inp model.GetAlertsInput

	reqPremiseID := strings.TrimSpace(req.Get("premiseID"))
	if len(reqPremiseID) > 0 {
		premiseID, err := strconv.Atoi(reqPremiseID)
		if err != nil {
			return model.GetAlertsInput{}, webErrGetAlertsDueToInvalidPremiseID
		}
		inp.PremiseID = int64(premiseID)
	}

	reqLimit := strings.TrimSpace(req.Get("limit"))
	if len(reqLimit) > 0 {
		limit, err := strconv.Atoi(reqLimit)
		if err != nil {
			return model.GetAlertsInput{}, webErrGetAlertsDueToInvalidLimit
		}

		if limit <= 0 {
			return model.GetAlertsInput{}, webErrLimitMustBeGreaterThanZero
		}

		if limit > model.PaginationMaxLimit {
			return model.GetAlertsInput{}, webErrMaxLimitMustBeLessThanNumber
		}

		inp.Limit = limit
	}

	reqPage := strings.TrimSpace(req.Get("page"))
	if len(reqPage) > 0 {
		page, err := strconv.Atoi(reqPage)
		if err != nil {
			return model.GetAlertsInput{}, webErrGetAlertsDueToInvalidPage
		}
		if page <= 0 {
			return model.GetAlertsInput{}, webErrPageMustBeGreaterThanZero
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
