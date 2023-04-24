package config

import "errors"

// Transport знает конфигурацию транспортного уровня сервиса.
type Transport struct {
	Address   string `env:"TRANSPORT_ADDRESS" env-default:"0.0.0.0"`
	Port      int    `env:"TRANSPORT_PORT" env-default:"8040"`
	DebugPort int    `env:"TRANSPORT_DEBUG_PORT" env-default:"3110"`
}

// Validate проверяет валидность конфигурации транспортного уровня.
func (t Transport) Validate() error {
	if t.Port < 0 {
		return errors.New("invalid transport port (TRANSPORT_PORT)")
	}
	if t.Address == "" {
		return errors.New("invalid transport address (TRANSPORT_ADDRESS)")
	}
	if t.DebugPort < 0 {
		return errors.New("invalid debug transport port (TRANSPORT_DEBUG_PORT)")
	}

	return nil
}
