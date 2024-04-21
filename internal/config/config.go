package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

const Path = "./internal/config/config.json"

type Config struct {
	BotKey string `env:"POSTALERT_TG_BOT_KEY,notEmpty"`
	ChatID int64  `env:"POSTALERT_TG_CHAT_ID,notEmpty"`
}

func Load() (Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return Config{}, fmt.Errorf("failed to parse config from env: %w", err)
	}

	return cfg, nil
}
