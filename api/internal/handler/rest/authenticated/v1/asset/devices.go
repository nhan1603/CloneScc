package asset

import (
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/nhan1603/CloneScc/api/internal/appconfig/httpserver"
	assetCtrl "github.com/nhan1603/CloneScc/api/internal/controller/asset"
	"github.com/nhan1603/CloneScc/api/internal/model"
)

// Devices is the response for each item of devices
type Devices struct {
	ID              int64  `json:"id"`
	PremiseID       int64  `json:"premiseID"`
	PremiseName     string `json:"premiseName"`
	PremiseLocation string `json:"premiseLocation"`
	DeviceName      string `json:"deviceName"`
	DeviceCode      string `json:"deviceCode"`
	IsActive        bool   `json:"isActive"`
	FloorNumber     int    `json:"floorNumber"`
	DeviceURL       string `json:"deviceURL"`
}

// DevicesResponse is the response for GetDevices
type DevicesResponse struct {
	Items      []Devices        `json:"items"`
	Pagination model.Pagination `json:"pagination"`
}

// GetDevices retrieves a list of devices
func (h Handler) GetDevices() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		log.Println("[GetDevices] START processing requests")
		input, err := validateAndConvertInput(r.URL.Query())
		if err != nil {
			return err
		}

		devices, total, err := h.assetCtrl.GetDevices(r.Context(), input)
		if err != nil {
			return err
		}

		items := make([]Devices, len(devices))
		for idx, item := range devices {
			items[idx] = Devices{
				ID:              item.ID,
				PremiseID:       item.PremiseID,
				PremiseName:     item.PremiseName,
				PremiseLocation: item.PremiseLocation,
				DeviceName:      item.DeviceName,
				DeviceCode:      item.DeviceCode,
				IsActive:        item.IsActive,
				FloorNumber:     item.FloorNumber,
				DeviceURL:       item.DeviceURL,
			}
		}

		httpserver.RespondJSON(w, DevicesResponse{
			Items: items,
			Pagination: model.Pagination{
				TotalCount:  total,
				CurrentPage: input.Page,
				Limit:       input.Limit,
			},
		})
		return nil
	})
}

func validateAndConvertInput(req url.Values) (assetCtrl.GetDevicesInput, error) {
	output := assetCtrl.GetDevicesInput{
		Name: strings.TrimSpace(req.Get("name")),
	}

	reqID := strings.TrimSpace(req.Get("premiseID"))
	if len(reqID) > 0 {
		premiseID, err := strconv.Atoi(reqID)
		if err != nil {
			return assetCtrl.GetDevicesInput{}, webErrInvalidID
		}
		output.PremiseID = int64(premiseID)
	}

	reqLimit := strings.TrimSpace(req.Get("limit"))
	if len(reqLimit) > 0 {
		limit, err := strconv.Atoi(reqLimit)
		if err != nil {
			return assetCtrl.GetDevicesInput{}, webErrInvalidNumber
		}

		if limit <= 0 {
			return assetCtrl.GetDevicesInput{}, webErrLimitMustBeGreaterThanZero
		}

		if limit > model.PaginationMaxLimit {
			return assetCtrl.GetDevicesInput{}, webErrMaxLimitMustBeLessThanNumber
		}
		output.Limit = limit
	}

	if output.Limit == 0 {
		output.Limit = model.PaginationDefaultLimit
	}

	reqPage := strings.TrimSpace(req.Get("page"))
	if len(reqPage) > 0 {
		page, err := strconv.Atoi(reqPage)
		if err != nil {
			return assetCtrl.GetDevicesInput{}, webErrInvalidNumber
		}
		if page < 0 {
			return assetCtrl.GetDevicesInput{}, webErrPageMustBeGreaterThanZero
		}
		output.Page = page
	}

	if output.Page == 0 {
		output.Page = model.PaginationDefaultPage
	}

	return output, nil
}
