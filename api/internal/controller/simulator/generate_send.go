package simulator

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/pkg/kafka"
)

// SendMessage send the message to kafka
func sendMessage(ctx context.Context, message model.AlertMessage, topic string, producer *kafka.SyncProducer) error {
	b, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("marshal input failed: %w", err)
	}

	log.Printf("Sending message to kafka: (topic: %s, payload: %s)\n", topic, string(b))
	_, _, err = producer.SendMessage(ctx, topic, b, kafka.ProducerMessageOption{})
	if err != nil {
		return err
	}

	return nil
}

// GenerateMessage generate a message instance
func generateMessage(CCTVname, alerType, floorNumber, description string) model.AlertMessage {
	return model.AlertMessage{
		CCTVName:    CCTVname,
		FloorNumber: floorNumber,
		Type:        alerType,
		Description: description,
		IncidentAt:  time.Now(),
	}
}
