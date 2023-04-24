// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"github.com/go-openapi/loads"
	"go.uber.org/zap"
	"net/http"

	"github.com/alexjurev/vault-service/config"
	"github.com/alexjurev/vault-service/pkg/infrastructure/restapi/endpoints"
	"github.com/alexjurev/vault-service/pkg/infrastructure/restapi/operations"
)

//go:generate swagger generate server --target ../../infrastructure --name VaultService --spec ../../../docs/swagger.yaml --template-dir ./swagger-gen/templates --principal interface{}

// API swagger api container.
type API struct {
	*operations.VaultServiceAPI
	endpoints *endpoints.Endpoints
	logger    *zap.Logger

	server *Server
}

// NewAPI create API instance.
func NewAPI(
	e *endpoints.Endpoints,
	logger *zap.Logger,

	cfg config.Transport,
) (*API, error) {
	spec, err := loads.Embedded(SwaggerJSON, FlatSwaggerJSON)
	if err != nil {
		return nil, err
	}

	api := operations.NewVaultServiceAPI(spec)
	setupSwaggerEndpoints(api, e)

	server := NewServer(api)
	server.Host, server.Port = cfg.Address, cfg.Port

	return &API{
		VaultServiceAPI: api,
		server:          server,
		logger:          logger,
	}, nil
}

// Serve run api server.
func (api *API) Serve() error {
	api.server.SetHandler(setupGlobalMiddleware(
		api.VaultServiceAPI.Serve(setupMiddlewares),
	))

	return api.server.Serve()
}

// Shutdown api server.
func (api *API) Shutdown() error {
	return api.server.Shutdown()
}

func configureFlags(_ *operations.VaultServiceAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

// метод используется в сгенерированном сервере, не удалять
func configureAPI(_ *operations.VaultServiceAPI) http.Handler {
	return nil
}

// The TLS configuration before HTTPS server starts.
func configureTLS(_ *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(_ *http.Server, _, _ string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
