package alert

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/nhan1603/CloneScc/api/internal/appconfig/httpserver"
	"github.com/nhan1603/CloneScc/api/internal/model"
)

// BroadCastAlert broadcasts alert to client
func (h Handler) BroadCastAlert() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		log.Println("[BroadCastAlert] START processing requests")
		var upgrader = websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true },
		}

		// init websocket
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal(err)
		}

		h.alertCtrl.BroadCast(r.Context(), ws)
		return nil
	})
}

// PushAlert pushes message to socket so that populate to client
func (h Handler) PushAlert() http.HandlerFunc {
	return httpserver.HandlerErr(func(w http.ResponseWriter, r *http.Request) error {
		log.Println("[PushAlert] START processing requests")
		var coordinates model.AlertMessage
		if err := json.NewDecoder(r.Body).Decode(&coordinates); err != nil {
			log.Printf("ERROR: %s", err)
			return err
		}
		defer r.Body.Close()

		h.alertCtrl.Push(r.Context(), coordinates)

		return nil
	})
}
