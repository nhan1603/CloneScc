package requests

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/nhan1603/CloneScc/api/internal/model"
)

// Push pushes alert from kafka to ws
func (i impl) Push(ctx context.Context, requestID int64) {
	contextNew := context.Background()
	go func() {
		request, err := i.repo.Request().GetByID(contextNew, requestID)
		if err != nil {
			log.Printf("[Push] getting request with id %v from resolve response encounter error:%v", requestID, err)
			return
		}
		alert, err := i.repo.Alert().GetAlert(contextNew, request.AlertID)
		if err != nil {
			log.Printf("[Push] getting alert with id %v from resolve response encounter error:%v", request.AlertID, err)
			return
		}
		message := model.ResponseMessage{
			RequestID:   strconv.FormatInt(requestID, 10),
			PremiseName: alert.PremiseName,
			CctvName:    alert.CCTVDevice,
			AlertID:     strconv.FormatInt(alert.ID, 10),
			IncidentAt:  time.Now(),
		}
		i.broadcast <- message
	}()
}
