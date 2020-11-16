package controller

import (
	"sync"

	"github.com/dgrr/fastws"
	"github.com/google/uuid"
	"github.com/soulgarden/messenger-backend/server/controller/handlers"
	"github.com/soulgarden/messenger-backend/ws"
)

const GoroutinesNum = 3

type WsController struct {
	storage  *ws.Storage
	handlers []handlers.Handler
}

func (c *WsController) Handle(conn *fastws.Conn) {
	sendCh := make(chan []byte)
	stopCh := make(chan int8)
	toStopCH := make(chan int8)

	connID := uuid.New().String()

	c.storage.Add(connID, sendCh)

	var wg sync.WaitGroup

	wg.Add(GoroutinesNum)

	worker := ws.NewWsWorker(c.handlers)

	go worker.Reader(&wg, toStopCH, conn, sendCh)
	go worker.Writer(&wg, toStopCH, stopCh, conn, sendCh)

	go func(wg *sync.WaitGroup, toStopCH chan int8, stopCh chan int8) {
		defer wg.Done()

		<-toStopCH
		close(stopCh)

		_ = conn.Close()
	}(&wg, toStopCH, stopCh)

	wg.Wait()

	c.storage.Remove(connID)

	close(sendCh)
	close(toStopCH)
}

func NewWsController(storage *ws.Storage, handlers []handlers.Handler) *WsController {
	return &WsController{storage, handlers}
}
