package requests

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"firebase.google.com/go/v4/messaging"
	"github.com/nhan1603/CloneScc/api/internal/model"
)

// CreateRequestInput is the input type for the CreateRequest controller
type CreateRequestInput struct {
	AlertID        int64
	RequestBy      int64
	AssignedUserID int64
	Content        string
}

func (inp CreateRequestInput) toNewVerificationRequest() model.VerificationRequest {
	return model.VerificationRequest{
		AlertID:        inp.AlertID,
		RequestBy:      inp.RequestBy,
		AssignedUserID: inp.AssignedUserID,
		Message:        inp.Content,
		Status:         model.VerificationRequestStatusNew,
		StartTime:      time.Now(),
	}
}

// CreateRequest Create a new request
func (i impl) CreateRequest(ctx context.Context, input CreateRequestInput) error {
	log.Println("[CreateRequest] START creating new request in ctrl")
	requestInp := input.toNewVerificationRequest()
	result, err := i.repo.Request().Insert(ctx, requestInp)
	if err != nil {
		log.Println(fmt.Printf("[CreateRequest] Error calling Request.Insert %#v from repository: %v", requestInp, err))
		return err
	}

	// Push notification, intentionally skipping errors from Firebase and write logs for later tracking
	if err := i.pushNotification(ctx, requestInp.AssignedUserID, result.ID); err != nil {
		log.Println(fmt.Printf("[CreateRequest] pushing notification failed, %v", err))
	}

	return nil
}

func (i impl) pushNotification(ctx context.Context, userID int64, reqID int64) error {
	log.Printf("[pushNotification] START pushing notification to (userID: %d, reqID: %d)\n", userID, reqID)
	deviceToken, err := i.repo.Asset().GetDeviceToken(ctx, userID)
	if err != nil {
		return fmt.Errorf("get device token failed, %v", err)
	}
	log.Println("Device token: ", deviceToken)

	reqURL := fmt.Sprintf("https://scc.demo.com/request/%d", reqID)
	log.Println("Request URL: ", reqURL)

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "SCC App Request",
			Body:  "Please checking the issue",
		},
		Data: map[string]string{
			"url": reqURL,
		},
		Token: deviceToken,
	}

	// Create a wait group to ensure all goroutines are finished before returning
	var wg sync.WaitGroup
	wg.Add(1)

	// Start the goroutine for sending the message
	go func() {
		defer wg.Done()
		response, err := i.clientFCM.Send(ctx, message)
		if err != nil {
			log.Printf("[pushNotification] Failed to send message: %v", err)
			return
		}
		log.Printf("[pushNotification] Sent message successfully, %v\n", response)
	}()

	// Wait for the goroutine to finish
	wg.Wait()

	log.Println("[pushNotification] END")
	return nil
}
