package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/nhan1603/CloneScc/api/internal/appconfig/httpserver"
	"github.com/nhan1603/CloneScc/api/internal/controller/alerts"
	"github.com/nhan1603/CloneScc/api/internal/controller/asset"
	"github.com/nhan1603/CloneScc/api/internal/controller/auth"
	"github.com/nhan1603/CloneScc/api/internal/controller/requests"
	"github.com/nhan1603/CloneScc/api/internal/controller/users"
	alertHandler "github.com/nhan1603/CloneScc/api/internal/handler/rest/authenticated/v1/alert"
	assetHandler "github.com/nhan1603/CloneScc/api/internal/handler/rest/authenticated/v1/asset"
	requestHandler "github.com/nhan1603/CloneScc/api/internal/handler/rest/authenticated/v1/request"
	userHandler "github.com/nhan1603/CloneScc/api/internal/handler/rest/authenticated/v1/user"
	authHandler "github.com/nhan1603/CloneScc/api/internal/handler/rest/public/v1/auth"
	cctvHandler "github.com/nhan1603/CloneScc/api/internal/handler/rest/public/v1/cctv"
	"github.com/nhan1603/CloneScc/api/internal/pkg/kafka"
	"github.com/nhan1603/CloneScc/api/internal/pkg/runner"
)

type router struct {
	ctx         context.Context
	authCtrl    auth.Controller
	assetCtrl   asset.Controller
	alertCtrl   alerts.Controller
	requestCtrl requests.Controller
	userCtrl    users.Controller
}

func (rtr router) initKafkaConsumer() {
	// Inital consumer kafka
	consumer, err := kafka.NewConsumer(
		rtr.ctx,
		os.Getenv("SCC_TOPIC"),
		os.Getenv("KAFKA_BROKER"),
		rtr.alertCtrl.HandleMessage,
	)
	if err != nil {
		log.Printf("Error when init consumer, %v", err)
		return
	}

	log.Printf("Kafka consumer start successfully\n")
	go runner.ExecParallel(rtr.ctx, consumer.Consume)
}

func (rtr router) routes(r chi.Router) {
	rtr.initKafkaConsumer()
	r.Group(rtr.authenticated)
	r.Group(rtr.public)
}

func (rtr router) authenticated(r chi.Router) {
	prefix := "/api/authenticated"

	r.Group(func(r chi.Router) {
		r.Use(httpserver.AuthenticateUserMiddleware())
		prefix = prefix + "/v1"

		userH := userHandler.New(rtr.userCtrl)
		r.Get(prefix+"/users", userH.GetUsers())

		assetH := assetHandler.New(rtr.assetCtrl)
		r.Get(prefix+"/premises", assetH.GetPremises())
		r.Get(prefix+"/devices", assetH.GetDevices())
		r.Post(prefix+"/device-token", assetH.UpdateDeviceToken())

		alertH := alertHandler.New(rtr.alertCtrl)
		r.Get(prefix+"/alerts", alertHandler.New(rtr.alertCtrl).GetAlerts())
		r.Get(prefix+"/alerts/{alertID}", alertH.GetAlertDetail())

		requestH := requestHandler.New(rtr.requestCtrl)
		r.Get(prefix+"/requests", requestH.GetRequests())
		r.Get(prefix+"/requests/{requestID}", requestH.GetRequestDetail())
		r.Post(prefix+"/requests/response", requestH.CreateResponse())
		r.Post(prefix+"/requests", requestH.CreateNewRequest())
	})
}

func (rtr router) public(r chi.Router) {
	prefix := "/api/public"

	r.Use(middleware.Logger)
	r.Group(func(r chi.Router) {
		r.Get(prefix+"/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		})
	})

	// Websocket routing
	r.Group(func(r chi.Router) {
		alertH := alertHandler.New(rtr.alertCtrl)
		r.Get("/broadcast-alert", alertH.BroadCastAlert())

		requestH := requestHandler.New(rtr.requestCtrl)
		r.Get("/broadcast-sg-result", requestH.BroadCastResponse())
	})

	// v1
	r.Group(func(r chi.Router) {
		prefix = prefix + "/v1"

		r.Group(func(r chi.Router) {
			authH := authHandler.New(rtr.authCtrl)
			r.Post(prefix+"/auth/ou", authH.AuthenticateOperationUser())
			r.Post(prefix+"/auth/sg", authH.AuthenticateSecurityGuard())

			// CCTV streaming video
			cctvH := cctvHandler.New(rtr.assetCtrl)
			r.Get(prefix+"/camera/{deviceCode}", cctvH.StreamingVideo())
		})
	})
}
