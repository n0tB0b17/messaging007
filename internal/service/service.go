package service

import (
	"github.com/bob17/msg/internal/db"
	"github.com/bob17/msg/internal/models"
	"github.com/bob17/msg/internal/nats"
)

var inMemDBInstance *db.DB
var natsInstance *nats.NatsClient

func InitializeService(db *db.DB, nat *nats.NatsClient) {
	inMemDBInstance = db
	natsInstance = nat
}

func GetNatClient() *nats.NatsClient {
	return natsInstance
}

func AddClient(c *models.Client) {
	inMemDBInstance.AddClients(c)
	natsInstance.Subscriber(c.ID, func(msg *models.Message) {
		c.Send <- []byte(msg.Content)
	})
}
