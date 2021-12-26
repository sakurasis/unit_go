package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Print("hello go")
	log.Debug().
		Int32("msgCode", 9999).
		Msg("Can not find any one.")
	log.Debug().
		Str("context", "Error Handler").Send()
}
