// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// FindObjectHandlerFunc turns a function with the right signature into a find object handler
type FindObjectHandlerFunc func(FindObjectParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FindObjectHandlerFunc) Handle(params FindObjectParams) middleware.Responder {
	return fn(params)
}

// FindObjectHandler interface for that can handle valid find object params
type FindObjectHandler interface {
	Handle(FindObjectParams) middleware.Responder
}

// NewFindObject creates a new http.Handler for the find object operation
func NewFindObject(ctx *middleware.Context, handler FindObjectHandler) *FindObject {
	return &FindObject{Context: ctx, Handler: handler}
}

/*
	FindObject swagger:route GET /objects/{key} findObject

Получить объект из хранилища по ключу
*/
type FindObject struct {
	Context *middleware.Context
	Handler FindObjectHandler
}

func (o *FindObject) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewFindObjectParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
