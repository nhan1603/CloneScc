package asset

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/nhan1603/CloneScc/api/internal/appconfig/httpserver"
	"github.com/nhan1603/CloneScc/api/internal/controller/asset"
)

// updateDeviceTokenRequest represents a payload struct for UpdateDeviceToken
type updateDeviceTokenRequest struct {
	UserID      int64  `json:"userID"`
	DeviceToken string `json:"deviceToken"`
	Platform    string `json:"platform"`
}

// UpdateDeviceToken is an api used to update token of devices
func (h Handler) UpdateDeviceToken() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		log.Println("[UpdateDeviceToken] START processing payload requests")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("[UpdateDeviceToken] read body error %v", err.Error())
			return webErrInvalidRequestBody
		}

		var req updateDeviceTokenRequest
		if err := json.Unmarshal(body, &req); err != nil {
			log.Printf("[UpdateDeviceToken] unmarshal error %v", err.Error())
			return err
		}

		input, err := req.validateAndMapUpdateDeviceTokenInput()
		if err != nil {
			return err
		}

		if err := h.assetCtrl.UpdateDeviceToken(r.Context(), input); err != nil {
			return err
		}

		httpserver.RespondJSON(w, httpserver.Success{
			Message: true,
		})

		return nil
	})
}

func (r updateDeviceTokenRequest) validateAndMapUpdateDeviceTokenInput() (asset.UpdateDeviceTokenInput, error) {
	if r.UserID < 1 {
		return asset.UpdateDeviceTokenInput{}, webErrInvalidUserID
	}

	deviceToken := strings.TrimSpace(r.DeviceToken)
	if deviceToken == "" {
		return asset.UpdateDeviceTokenInput{}, webErrDeviceTokenIsRequired
	}

	platform := strings.TrimSpace(r.Platform)
	if platform == "" {
		return asset.UpdateDeviceTokenInput{}, webErrPlatformIsRequired
	}

	return asset.UpdateDeviceTokenInput{
		UserID:      r.UserID,
		DeviceToken: deviceToken,
		Platform:    platform,
	}, nil
}
