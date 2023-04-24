// Package logging is package that contains all that is needed for creating logger.
//
// Level is the level for creating a zap logger with different output formats.
//   - Production is the level passed which the NewProduction function is called
//   - Development is the level passed which the NewDevelopment function is called
//
// ErrNoLevel is thrown when an unknown level is passed to NewLogger.
package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Level is logger level
type Level int

// list of logger levels
const (
	Production Level = iota + 1
	Stage
	Development
)

// NewLogger creates logger by level
func NewLogger(level Level) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	switch level {
	case Production:
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case Stage:
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case Development:
		cfg = zap.NewDevelopmentConfig()
	}
	return cfg.Build(zap.AddStacktrace(zap.DPanicLevel))
}
