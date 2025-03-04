package models

import "sync"

type Client struct {
	ID     string
	Send   chan []byte
	Mu     sync.Mutex
	Topics map[string]struct{}
}
