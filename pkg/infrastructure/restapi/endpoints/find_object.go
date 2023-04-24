package endpoints

import (
	"github.com/go-openapi/runtime/middleware"
	"go.uber.org/zap"
	"net/http"

	apiOperations "github.com/alexjurev/vault-service/pkg/infrastructure/restapi/operations"
)

// FindObjectEndpoint endpoint для ручки /objects/{key}.
// Получить объект из хранилища по ключу.
func (e *Endpoints) FindObjectEndpoint(
	params apiOperations.FindObjectParams,
) middleware.Responder {
	logger := e.logger.With(zap.String("endpoint", "FindObjectEndpoint"))
	if params.Key == "" {
		logger.With(zap.String("key", "empty key")).Error("internal error")
		return apiOperations.NewFindObjectDefault(http.StatusInternalServerError)
	}
	object, ok := e.cache.Get(params.Key)
	if !ok {
		logger.With(zap.String("object", "object not found")).Error("internal error")

		return apiOperations.NewFindObjectDefault(http.StatusNotFound)
	}

	return apiOperations.NewFindObjectOK().WithPayload(object.(string))
}
