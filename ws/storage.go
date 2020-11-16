package ws

import "sync"

type Storage struct {
	// connection ID => channel
	connectionsChannels map[string]chan []byte

	// user ID => connection ID
	usersConnections map[int64]string

	mx sync.RWMutex
}

func (s *Storage) Add(connID string, sendChan chan []byte) {
	s.mx.Lock()
	s.connectionsChannels[connID] = sendChan
	s.mx.Unlock()
}

func (s *Storage) Remove(connID string) {
	s.mx.Lock()
	delete(s.connectionsChannels, connID)
	s.mx.Unlock()
}

func (s *Storage) SignIn(userID int64, connID string) {
	s.mx.Lock()
	s.usersConnections[userID] = connID
	s.mx.Unlock()
}

func NewStorage() *Storage {
	return &Storage{
		connectionsChannels: make(map[string]chan []byte),
		usersConnections:    make(map[int64]string),
	}
}
