package nats

import (
	"encoding/json"
	"fmt"

	"github.com/bob17/msg/internal/models"
	"github.com/nats-io/nats.go"
)

type NatsClient struct {
	Conn *nats.Conn
}

func GetNatsClient(url string) (*NatsClient, error) {
	conn, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}

	return &NatsClient{Conn: conn}, nil
}

func (nc *NatsClient) Publish(subject string, msg *models.Message) error {
	docs, err := json.Marshal(msg)
	if err != nil {
		fmt.Printf("error while publishing message to subject: %s | error: %v \n", subject, err)
		return err
	}

	return nc.Conn.Publish(subject, docs)
}

func (nc *NatsClient) Subscriber(subject string, handler func(*models.Message)) error {
	_, err := nc.Conn.Subscribe(subject, func(m *nats.Msg) {
		var msg models.Message
		err := json.Unmarshal(m.Data, &msg)
		if err != nil {
			fmt.Printf("failed to unmarshal published data on subject: %s \n", subject)
			return
		}

		handler(&msg)
	})

	return err
}
