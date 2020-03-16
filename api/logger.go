package api

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Log is the global logger
var Log zerolog.Logger

func init() {
	logFile, err := os.OpenFile("apip.log", os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	Log = zerolog.New(logFile)
}
