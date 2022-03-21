package log

import (
	"os"

	"github.com/rs/zerolog"
)

var Logger zerolog.Logger

func Init() {
	Logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
}
