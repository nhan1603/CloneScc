package request

import (
	"context"
	"time"

	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/nhan1603/CloneScc/api/internal/model"
)

// Repository provides the specification of the functionality provided by this pkg
type Repository interface {
	// Insert a new request to database
	Insert(ctx context.Context, input model.VerificationRequest) (model.VerificationRequest, error)
	// GetRequests returns the requests list
	GetRequests(context.Context, model.GetRequestsInput) ([]model.RequestSummary, int64, error)
	// GetRequest returns the request detail
	GetRequest(context.Context, int64) (model.RequestDetail, error)
	// InsertRequestResponses inserts a new request response to database
	InsertRequestResponses(ctx context.Context, input model.VerificationRequestResponses) (model.VerificationRequestResponses, error)
	// UpdateStatusAndEndTime updates the status and end time of a request
	UpdateStatusAndEndTime(ctx context.Context, reqID int64, status model.VerificationRequestStatus, endTime time.Time) (model.VerificationRequest, error)
	// GetByID returns the request with the given ID
	GetByID(ctx context.Context, id int64) (model.VerificationRequest, error)
}

// New returns an implementation instance satisfying Repository
func New(dbConn boil.ContextExecutor) Repository {
	return impl{dbConn: dbConn}
}

type impl struct {
	dbConn boil.ContextExecutor
}
