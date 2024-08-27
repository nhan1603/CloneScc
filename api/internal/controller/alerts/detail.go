package alerts

import (
	"context"
	"errors"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/repository/alert"
)

// Detail handles retrieve alert detail
func (i impl) Detail(ctx context.Context, alertID int64) (model.Alert, error) {
	a, err := i.repo.Alert().GetAlert(ctx, alertID)
	if err != nil {
		if errors.Is(err, alert.ErrNotFound) {
			return model.Alert{}, ErrNotFound
		}

		return model.Alert{}, err
	}

	return a, nil
}
