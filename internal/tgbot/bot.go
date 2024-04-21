package tgbot

import (
	"time"

	"github.com/barklan/postalert/internal/config"
	tele "gopkg.in/telebot.v3"
)

type Bot struct {
	b   *tele.Bot
	cfg config.Config
}

func New(cfg config.Config) (*Bot, error) {
	pref := tele.Settings{
		Token:  cfg.BotKey,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		return nil, err
	}

	// b.Handle("/start", func(c tele.Context) error {
	// 	if c.Chat() != nil && c.Chat().ID == cfg.ChatID {
	// 		return c.Send("Hi!")
	// 	}
	//
	// 	return nil
	// })

	return &Bot{
		b:   b,
		cfg: cfg,
	}, nil
}

func (b *Bot) Send(what interface{}, opts ...interface{}) error {
	_, err := b.b.Send(&tele.Chat{ID: b.cfg.ChatID}, what, opts...)
	return err
}
