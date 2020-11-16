package handlers

type Handler interface {
	Handler(sendCh chan<- []byte)
	Supports(m []byte) bool
}
