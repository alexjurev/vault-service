package config

// Logger is logger config
type Logger struct {
	Level int `env:"LOGGER_LEVEL" env-default:"1"`
}
