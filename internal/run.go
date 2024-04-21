package internal

import (
	"github.com/barklan/postalert/internal/config"
	"github.com/barklan/postalert/internal/server"
	"github.com/barklan/postalert/internal/tgbot"
)

func Run() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	bot, err := tgbot.New(cfg)
	if err != nil {
		return err
	}

	defer bot.Send("alerter exiting")

	s := server.New(cfg, bot)

	if err := s.Serve(); err != nil {
		return err
	}

	return nil
}
