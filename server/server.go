package server

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/dgrr/fastws"
	"github.com/rs/zerolog/log"
	"github.com/soulgarden/messenger-backend/server/controller"
	wsHandlers "github.com/soulgarden/messenger-backend/server/controller/handlers"
	"github.com/soulgarden/messenger-backend/ws"
	"github.com/valyala/fasthttp"
)

type Server struct {
	storage *ws.Storage
}

func NewServer(storage *ws.Storage) *Server {
	return &Server{storage: storage}
}

func (s *Server) NewRouter() *fasthttprouter.Router {
	router := fasthttprouter.New()
	handlers := make([]wsHandlers.Handler, 0)
	handlers = append(handlers, wsHandlers.NewPingHandler())

	router.GET("/ws", fastws.Upgrade(controller.NewWsController(s.storage, handlers).Handle))

	return router
}

func (s *Server) Serve(listenAddr string, router *fasthttprouter.Router) {
	log.Info().Msg("Starting the service on " + listenAddr)

	log.Fatal().Msg(fasthttp.ListenAndServe(listenAddr, router.Handler).Error())
}
