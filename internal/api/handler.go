package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bob17/msg/internal/models"
	"github.com/bob17/msg/internal/nats"
	"github.com/bob17/msg/internal/service"
	"github.com/gorilla/websocket"
)

type APIClient struct {
	*models.Client
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func SocketRoute(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "unable to upgrade connection", http.StatusBadGateway)
		return
	}

	defer ws.Close()

	cli := &APIClient{
		Client: &models.Client{
			ID:     r.Header.Get("X-Client-ID"),
			Send:   make(chan []byte),
			Topics: make(map[string]struct{}),
		},
	}

	service.AddClient(cli.Client)
	go cli.Read(ws, service.GetNatClient())
	go cli.Write(ws)
}

func (a *APIClient) Read(conn *websocket.Conn, natsClient *nats.NatsClient) {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("error while reading message: %v", err)
			break
		}

		var newMsg models.Message
		err = json.Unmarshal(msg, &newMsg)
		if err != nil {
			fmt.Printf("error while unmarshing bytes data: %v", err)
			continue
		}

		natsClient.Publish(newMsg.RecipientID, &newMsg)
	}
}

func (a *APIClient) Write(conn *websocket.Conn) {
	for msgBytes := range a.Send {
		err := conn.WriteMessage(websocket.TextMessage, msgBytes)
		if err != nil {
			fmt.Printf("error while writing message: %v", err)
			break
		}
	}

}
