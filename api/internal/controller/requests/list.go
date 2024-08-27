package requests

import (
	"context"
	"fmt"
	"log"

	"github.com/nhan1603/CloneScc/api/internal/model"
)

// List handles retrieve list alerts
func (i impl) List(ctx context.Context, inp model.GetRequestsInput) ([]model.RequestSummary, int64, error) {
	// get request
	requests, totalCount, err := i.repo.Request().GetRequests(ctx, inp)
	if err != nil {
		log.Println(fmt.Printf("[GetRequests] Error calling GetRequests from repository: %v", err.Error()))
		return nil, 0, err
	}

	return requests, totalCount, nil
}
