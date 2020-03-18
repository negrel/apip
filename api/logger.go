package api

import (
	"github.com/rs/zerolog"
)

// Log is the global logger
var Log zerolog.Logger

func init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}
