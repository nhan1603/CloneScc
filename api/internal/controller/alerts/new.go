package alerts

import (
	"context"

	"github.com/gorilla/websocket"
	"github.com/nhan1603/CloneScc/api/internal/model"
	kafkaPkg "github.com/nhan1603/CloneScc/api/internal/pkg/kafka"
	"github.com/nhan1603/CloneScc/api/internal/repository"
)

// Controller represents the specification of this pkg
type Controller interface {
	// List handles retrieve list alerts
	List(context.Context, model.GetAlertsInput) ([]model.Alert, int64, error)
	// Detail handles retrieve alert detail
	Detail(context.Context, int64) (model.Alert, error)
	// BroadCast broadcasts alert to client
	BroadCast(ctx context.Context, conn *websocket.Conn)
	// Push pushes alert from kafka to ws
	Push(ctx context.Context, m model.AlertMessage)
	// HandleMessage receive the message from kafka
	HandleMessage(ctx context.Context, msg kafkaPkg.ConsumerMessage) error
	// CreateAlert add a new alert instance in the database
	CreateAlert(ctx context.Context, alertInstance model.AlertMessage) (int64, error)
}

// New initializes a new Controller instance and returns it
func New(repo repository.Registry,
	clients map[*websocket.Conn]bool,
	broadcast chan model.AlertMessage,
) Controller {
	return impl{
		repo:      repo,
		clients:   clients,
		broadcast: broadcast,
	}
}

type impl struct {
	repo      repository.Registry
	clients   map[*websocket.Conn]bool
	broadcast chan model.AlertMessage
}
