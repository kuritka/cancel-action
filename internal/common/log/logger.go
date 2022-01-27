// Package log wraps zerolog logger and provides standard log functionality
package log

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

// Log is the global logger.
var Log *zerolog.Logger

// init initializes the logger
func init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	l := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05", NoColor: false}).
		With().
		Timestamp().
		Logger()
	Log = &l
}
