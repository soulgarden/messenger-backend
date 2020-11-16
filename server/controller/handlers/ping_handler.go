package handlers

import "bytes"

const (
	pingRequest  = "2"
	pingResponse = "3"
)

type PingHandler struct {
}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (h *PingHandler) Handler(sendCh chan<- []byte) {
	sendCh <- []byte(pingResponse)
}

func (h *PingHandler) Supports(m []byte) bool {
	return bytes.Equal(m, []byte(pingRequest))
}
