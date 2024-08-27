package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	"firebase.google.com/go/v4/messaging"
	"github.com/gorilla/websocket"
	"github.com/nhan1603/CloneScc/api/internal/appconfig/db/pg"
	"github.com/nhan1603/CloneScc/api/internal/appconfig/httpserver"
	"github.com/nhan1603/CloneScc/api/internal/controller/alerts"
	"github.com/nhan1603/CloneScc/api/internal/controller/asset"
	"github.com/nhan1603/CloneScc/api/internal/controller/auth"
	"github.com/nhan1603/CloneScc/api/internal/controller/requests"
	"github.com/nhan1603/CloneScc/api/internal/controller/users"
	"github.com/nhan1603/CloneScc/api/internal/model"
	"github.com/nhan1603/CloneScc/api/internal/pkg/fcm"
	"github.com/nhan1603/CloneScc/api/internal/pkg/kafka"
	"github.com/nhan1603/CloneScc/api/internal/repository"
	"github.com/nhan1603/CloneScc/api/internal/repository/generator"
)

func main() {
	log.Println("Security Command Center API")

	ctx := context.Background()

	// Initial DB connection
	conn, err := pg.Connect(os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Initial snowflake generator
	if err := generator.InitSnowflakeGenerators(); err != nil {
		log.Fatal(err)
	}

	// Initial producer kafka
	producer, err := kafka.NewSyncProducer(ctx, os.Getenv("KAFKA_BROKER"))
	if err != nil {
		log.Fatal(err)
	}

	// Initial Firebase Client
	clientFCM, err := fcm.NewFCM(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Initial router
	rtr, err := initRouter(ctx, conn, producer, clientFCM)
	if err != nil {
		log.Fatal(err)
	}

	// Start server
	httpserver.Start(httpserver.Handler(ctx, rtr.routes))
}

func initRouter(
	ctx context.Context,
	db *sql.DB,
	producer *kafka.SyncProducer,
	clientFCM *messaging.Client,
) (router, error) {
	repo := repository.New(db)

	return router{
		ctx:         ctx,
		authCtrl:    auth.New(repo),
		assetCtrl:   asset.New(repo),
		alertCtrl:   alerts.New(repo, make(map[*websocket.Conn]bool), make(chan model.AlertMessage)),
		requestCtrl: requests.New(repo, clientFCM, make(map[*websocket.Conn]bool), make(chan model.ResponseMessage)),
		userCtrl:    users.New(repo),
	}, nil
}
