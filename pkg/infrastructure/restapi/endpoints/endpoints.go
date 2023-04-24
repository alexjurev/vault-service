package endpoints

import (
	"fmt"
	"github.com/patrickmn/go-cache"

	"go.uber.org/zap"
)

// Endpoints restapi endpoints.
type Endpoints struct {
	cache  *cache.Cache
	logger *zap.Logger
}

// NewRespAPIEndpoints create endpoints.
func NewRespAPIEndpoints(
	cache *cache.Cache,
	logger *zap.Logger,
) (*Endpoints, error) {
	if cache == nil {
		return nil, fmt.Errorf("cache should be passed")
	}
	if logger == nil {
		return nil, fmt.Errorf("logger should be passed")
	}

	return &Endpoints{
		cache:  cache,
		logger: logger,
	}, nil
}
