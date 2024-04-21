package main

import (
	"github.com/rs/zerolog/log"

	"github.com/barklan/postalert/internal"
)

func main() {
	if err := internal.Run(); err != nil {
		log.Err(err).Msg("error running server")
	}
}
