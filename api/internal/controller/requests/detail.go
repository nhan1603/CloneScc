package requests

import (
	"context"
	"errors"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/repository/request"
)

// Detail handles retrieve detail of request
func (i impl) Detail(ctx context.Context, requestID int64) (model.RequestDetail, error) {
	r, err := i.repo.Request().GetRequest(ctx, requestID)
	if err != nil {
		if errors.Is(err, request.ErrNotFound) {
			return model.RequestDetail{}, ErrNotFound
		}

		return model.RequestDetail{}, err
	}

	return r, nil
}
