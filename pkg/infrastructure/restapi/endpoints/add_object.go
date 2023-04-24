package endpoints

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/patrickmn/go-cache"
	"go.uber.org/zap"
	"net/http"
	"time"

	apiOperations "github.com/alexjurev/vault-service/pkg/infrastructure/restapi/operations"
)

// AddObjectEndpoint endpoint для ручки /objects/{key}.
// Добавить объект в хранилище по ключу.
func (e *Endpoints) AddObjectEndpoint(
	params apiOperations.AddObjectParams,
) middleware.Responder {
	logger := e.logger.With(zap.String("endpoint", "AddObjectEndpoint"))
	if params.Body == "" {
		logger.With(zap.String("body", "empty body")).Error("internal error")
		return apiOperations.NewAddObjectDefault(http.StatusInternalServerError)
	}
	if params.Expires != nil {
		e.cache.Set(params.Key, params.Body, time.Duration(*params.Expires)*time.Minute)
	}
	e.cache.Set(params.Key, params.Body, cache.DefaultExpiration)

	return apiOperations.NewAddObjectOK()
}
