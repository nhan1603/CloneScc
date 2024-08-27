package alerts

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	"github.com/nhan1603/CloneScc/api/internal/model"
	kafkaPkg "github.com/nhan1603/CloneScc/api/internal/pkg/kafka"
)

// Push pushes alert from kafka to ws
func (i impl) Push(ctx context.Context, m model.AlertMessage) {
	go func() {
		i.broadcast <- m
	}()
}

// HandleMessage handles and processes a message
func (i impl) HandleMessage(ctx context.Context, msg kafkaPkg.ConsumerMessage) error {
	log.Printf("[HandleMessage] START payload: (%s) in ctrl\n", string(msg.Value))

	var input model.AlertMessage
	if err := json.Unmarshal(msg.Value, &input); err != nil {
		// err when unmarshal message
		log.Printf("[HandleMessage] cannot parse message as alert info")
		return err
	}

	// insert a new record into the database
	alertId, errCreate := i.CreateAlert(ctx, input)
	if errCreate != nil {
		log.Printf("[HandleMessage] cannot insert alert")
		return errCreate
	}

	input.ID = strconv.FormatInt(alertId, 10)
	// push alert instance to websocket
	i.Push(ctx, input)
	return nil
}
