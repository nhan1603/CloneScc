package alerts

import (
	"context"
	"log"

	"github.com/gorilla/websocket"
)

// BroadCast broadcasts alert to client
func (i impl) BroadCast(_ context.Context, ws *websocket.Conn) {
	i.clients[ws] = true

	go func(i impl) {
		for {
			val := <-i.broadcast
			// send to every client that is currently connected
			for client := range i.clients {
				err := client.WriteJSON(val)
				if err != nil {
					log.Printf("Websocket error: %s", err)
					client.Close()
					delete(i.clients, client)
				}
			}
		}
	}(i)
}
