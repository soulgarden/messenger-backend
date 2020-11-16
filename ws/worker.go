package ws

import (
	"errors"
	"sync"
	"time"

	"github.com/dgrr/fastws"
	"github.com/rs/zerolog/log"
	"github.com/soulgarden/messenger-backend/server/controller/handlers"
)

const deadline = 60

type Worker struct {
	handlers []handlers.Handler
}

func NewWsWorker(handlers []handlers.Handler) *Worker {
	return &Worker{handlers}
}

func (w *Worker) Reader(wg *sync.WaitGroup, toStopCH chan<- int8, conn *fastws.Conn, sendCh chan []byte) {
	defer wg.Done()

	conn.ReadTimeout = deadline * time.Second

	var msg []byte

	var err error

	for {
		_, msg, err = conn.ReadMessage(msg[:0])
		if err != nil {
			if errors.Is(err, fastws.EOF) {
				log.Err(err).Msg("reading message")
			}

			toStopCH <- 1

			return
		}

		log.Debug().Bytes("message", msg).Msg("message received")

		conn.ReadTimeout = deadline * time.Second

		for _, handler := range w.handlers {
			if handler.Supports(msg) {
				handler.Handler(sendCh)
			}
		}
	}
}

func (w *Worker) Writer(
	wg *sync.WaitGroup,
	toStopCH chan int8,
	stopCh chan int8,
	conn *fastws.Conn,
	sendCh <-chan []byte,
) {
	defer wg.Done()

	conn.WriteTimeout = deadline * time.Second

	for {
		select {
		case <-stopCh:
			return
		case m := <-sendCh:
			_, err := conn.Write(m)
			if err != nil {
				log.Err(err).Msg("writing message")

				toStopCH <- 1

				return
			}

			log.Debug().Bytes("message", m).Msg("message sent")
		}
	}
}
