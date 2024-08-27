package alerts

import (
	"context"
	"errors"
	"log"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/repository/asset"
)

// CreateAlert create a new alert instance
func (i impl) CreateAlert(ctx context.Context, alertInstance model.AlertMessage) (int64, error) {
	cctvId, err := i.repo.Asset().GetCCTVKeyByName(ctx, alertInstance.CCTVName)
	if err != nil {
		log.Printf("[CreateAlert] error retrieve cctv by given cctv name %v\n", alertInstance.CCTVName)
		if errors.Is(err, asset.ErrCctvNotFound) {
			return 0, ErrCCTVNotFound
		}
		return 0, err
	}

	alertId, err := i.repo.Alert().CreateAlert(ctx, alertInstance, cctvId)
	if err != nil {
		log.Printf("[CreateAlert] create alert instance encounter error : %v\n", err)
		return 0, ErrCreateAlert
	}

	return alertId, err
}
