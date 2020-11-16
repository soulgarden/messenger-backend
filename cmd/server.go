package cmd

import (
	"github.com/rs/zerolog"
	"github.com/soulgarden/messenger-backend/conf"
	"github.com/soulgarden/messenger-backend/server"
	"github.com/soulgarden/messenger-backend/ws"
	"github.com/spf13/cobra"
)

//nolint: gochecknoglobals, exhaustivestruct
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run server",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := conf.New()

		if cfg.Debug {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		}

		s := server.NewServer(ws.NewStorage())
		router := s.NewRouter()
		s.Serve(cfg.ListenAddr, router)
	},
}
