package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"

	"github.com/barklan/postalert/internal/config"
	"github.com/barklan/postalert/internal/tgbot"
)

type AlertInput struct {
	Namespace string `path:"namespace" maxLength:"30" example:"myservice" doc:"Namespace for alerts" required:"true"`
	Body      struct {
		Level   string `path:"level"`
		Message string `path:"message"`
	}
}

type AlertOutput struct {
	Body struct {
		Ok bool `json:"ok"`
	}
}

type Server struct {
	router *chi.Mux
	bot    *tgbot.Bot
	cfg    config.Config
}

const template = `Namespace: %s
Level: %s
Message: %s`

func New(cfg config.Config, bot *tgbot.Bot) *Server {
	router := chi.NewMux()

	api := humachi.New(router, huma.DefaultConfig("PostAlert API", "1.0.0"))

	huma.Post(api, "/{namespace}", func(ctx context.Context, input *AlertInput) (*AlertOutput, error) {
		log.Info().
			Str("event", "alert").
			Str("alert.namespace", input.Namespace).
			Str("alert.level", input.Body.Level).
			Str("alert.message", input.Body.Message).
			Msg("received alert")

		go func() {
			if err := bot.Send(fmt.Sprintf(template, input.Namespace, input.Body.Level, input.Body.Message)); err != nil {
				log.Err(err).Msg("failed to send notification")
			}
		}()

		resp := &AlertOutput{}
		resp.Body.Ok = true
		return resp, nil
	})

	return &Server{
		router: router,
		cfg:    cfg,
		bot:    bot,
	}
}

func (s *Server) Serve() error {
	_ = s.bot.Send("alerter started")
	return http.ListenAndServe("0.0.0.0:8000", s.router)
}
