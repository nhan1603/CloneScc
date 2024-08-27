package simulator

import (
	"context"
	"math/rand"
	"time"

	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/pkg/kafka"
)

const alertInterval = 1

// Simulate simulates alerts
func (i impl) Simulate(ctx context.Context) {
	mapData := make(map[int]string, 3)
	mapData[0] = model.AlertTypeUnauthorizedAccess.ToString()
	mapData[1] = model.AlertTypePropertyDamage.ToString()
	mapData[2] = model.AlertTypeSuspiciousActivities.ToString()

	listCctv, err := i.repo.Asset().GetAllCCTV(ctx)
	if err != nil {
		return
	}

	executeAtInterval(ctx, listCctv, mapData, alertInterval*time.Minute, i.topic, i.producer)

	select {}
}

// ExecuteAtInterval execute the alert simulation
func executeAtInterval(ctx context.Context, listCctv []model.CctvData, mapData map[int]string, interval time.Duration, topic string, producer *kafka.SyncProducer) {
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			duration := rand.Intn(300) + 1
			time.Sleep(time.Duration(duration) * time.Second)
			randomCctv := rand.Intn(len(listCctv))
			cctv := listCctv[randomCctv]
			alertType := mapData[rand.Intn(3)]
			desc := "There is " + alertType + " at floor " + cctv.FloorNumber
			message := generateMessage(cctv.CctvName, alertType, cctv.FloorNumber, desc)
			sendMessage(ctx, message, topic, producer)
		}
	}()
}
