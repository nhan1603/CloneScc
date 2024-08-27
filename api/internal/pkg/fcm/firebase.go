package fcm

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	pkgerrors "github.com/pkg/errors"
	"google.golang.org/api/option"
)

// NewFCM creates a Firebase Admin SDK and Client connection
func NewFCM(ctx context.Context) (*messaging.Client, error) {
	log.Println("Initializing Firebase connection")

	if _, err := os.ReadFile(os.Getenv("FB_SA")); err != nil {
		return nil, pkgerrors.Wrap(err, "error loading Firebase Key Path file")
	}

	// Init Firebase Admin SDK
	opt := option.WithCredentialsFile(os.Getenv("FB_SA"))
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, pkgerrors.Wrap(err, "error initializing Firebase app")
	}

	// Init client to send notification
	client, err := app.Messaging(ctx)
	if err != nil {
		return nil, pkgerrors.Wrap(err, "error initializing Firebase messaging client")
	}

	return client, nil
}
