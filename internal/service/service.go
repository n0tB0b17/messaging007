package service

import (
	"fmt"
	"sync"

	"github.com/bob17/msg/internal/db"
	"github.com/bob17/msg/internal/models"
	"github.com/bob17/msg/internal/nats"
)

var (
	inMemDBInstance *db.DB
	natsInstance    *nats.NatsClient
	mu              sync.Mutex
	clients         = make(map[string]*models.Client)
)

func InitializeService(db *db.DB, nat *nats.NatsClient) {
	inMemDBInstance = db
	natsInstance = nat
}

func GetNatClient() *nats.NatsClient {
	return natsInstance
}

func AddClient(c *models.Client) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Printf("client id: %s", c.ID)
	clients[c.ID] = c
	natsInstance.Subscriber(c.ID, func(msg *models.Message) {
		c.Send <- []byte(msg.Content) // fix here
	})
}

func RemoveClient(c *models.Client) {
	mu.Lock()
	defer mu.Unlock()
	delete(clients, c.ID)
}
