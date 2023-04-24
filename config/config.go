package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

// Config структура конфигурации сервиса.
type Config struct {
	AppName   string `env:"APP_NAME" env-default:"vault-service"`
	Transport Transport
	Logger    Logger
}

// Validate проверяет валидность конфигурации сервиса.
func (c Config) Validate() error {
	var err error

	err = c.Transport.Validate()
	if err != nil {
		return err
	}

	return nil
}

// LoadConfig загружает конфигурацию из переменных среды ENV.
func LoadConfig() (*Config, error) {
	cfg := Config{}

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
