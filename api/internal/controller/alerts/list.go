package alerts

import (
	"context"
	"fmt"
	"log"

	"github.com/nhan1603/CloneScc/api/internal/model"
)

// List handles retrieve list alerts
func (i impl) List(ctx context.Context, inp model.GetAlertsInput) ([]model.Alert, int64, error) {
	// get alerts
	alerts, totalCount, err := i.repo.Alert().GetAlerts(ctx, inp)
	if err != nil {
		log.Println(fmt.Printf("[GetAlerts] Error calling GetAlerts from repository: %v", err.Error()))
		return nil, 0, err
	}

	return alerts, totalCount, nil
}
