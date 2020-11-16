package main

import (
	"github.com/rs/zerolog"
	"github.com/soulgarden/messenger-backend/cmd"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	cmd.Execute()
}
