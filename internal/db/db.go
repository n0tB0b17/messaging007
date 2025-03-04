package db

import (
	"sync"

	"github.com/bob17/msg/internal/models"
)

type DB struct {
	mu       sync.Mutex
	messages map[string]*models.Message
	clients  map[string]*models.Client
	users    map[string]*models.User
}

func NewDB() *DB {
	return &DB{
		messages: make(map[string]*models.Message),
		clients:  make(map[string]*models.Client),
		users:    make(map[string]*models.User),
	}
}

// message CRUD
func (db *DB) AddMessages(msg *models.Message) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.messages[msg.ID] = msg
}

func (db *DB) GetMessage(id string) (*models.Message, bool) {
	db.mu.Lock()
	defer db.mu.Unlock()
	msg, ok := db.messages[id]
	return msg, ok
}

func (db *DB) DeleteMessage(id string) bool {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, ok := db.messages[id]; ok {
		delete(db.messages, id)
		return true
	}

	return false
}

func (db *DB) UpdateMessage(id string, msg *models.Message) (*models.Message, bool) {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, ok := db.messages[id]; ok {
		db.messages[id] = msg
		return db.messages[id], true
	}
	return db.messages[id], false
}

// clients CRUD
func (db *DB) AddClients(client *models.Client) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.clients[client.ID] = client
}

func (db *DB) GetClient(id string) (*models.Client, bool) {
	db.mu.Lock()
	defer db.mu.Unlock()
	client, ok := db.clients[id]
	return client, ok
}

// users CRUD
func (db *DB) AddUsers(user *models.User) {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.users[user.ID] = user
}

func (db *DB) GetUser(id string) (*models.User, bool) {
	db.mu.Lock()
	defer db.mu.Unlock()

	client, ok := db.users[id]
	return client, ok
}
