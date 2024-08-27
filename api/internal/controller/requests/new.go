package requests

import (
	"context"

	"firebase.google.com/go/v4/messaging"
	"github.com/gorilla/websocket"
	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/repository"
)

// Controller represents the specification of this pkg
type Controller interface {
	// List handles retrieve list requests
	List(context.Context, model.GetRequestsInput) ([]model.RequestSummary, int64, error)
	// Detail handles retrieve detail of request
	Detail(context.Context, int64) (model.RequestDetail, error)
	// CreateRequest Create a new request
	CreateRequest(ctx context.Context, input CreateRequestInput) error
	// CreateResponse Create a new response
	CreateResponse(ctx context.Context, input CreateResponseInput) error
	// Push send a response message to ws
	Push(ctx context.Context, requestId int64)
	// BroadCastResponse create a ws for broadcast response
	BroadCastResponse(ctx context.Context, ws *websocket.Conn)
}

// New initializes a new Controller instance and returns it
func New(
	repo repository.Registry,
	clientFCM *messaging.Client,
	clients map[*websocket.Conn]bool,
	broadcast chan model.ResponseMessage,
) Controller {
	return impl{
		repo:      repo,
		clients:   clients,
		broadcast: broadcast,
		clientFCM: clientFCM,
	}
}

type impl struct {
	repo      repository.Registry
	clientFCM *messaging.Client
	clients   map[*websocket.Conn]bool
	broadcast chan model.ResponseMessage
}
