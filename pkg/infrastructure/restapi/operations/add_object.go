// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AddObjectHandlerFunc turns a function with the right signature into a add object handler
type AddObjectHandlerFunc func(AddObjectParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddObjectHandlerFunc) Handle(params AddObjectParams) middleware.Responder {
	return fn(params)
}

// AddObjectHandler interface for that can handle valid add object params
type AddObjectHandler interface {
	Handle(AddObjectParams) middleware.Responder
}

// NewAddObject creates a new http.Handler for the add object operation
func NewAddObject(ctx *middleware.Context, handler AddObjectHandler) *AddObject {
	return &AddObject{Context: ctx, Handler: handler}
}

/*
	AddObject swagger:route PUT /objects/{key} addObject

Добавить объект в хранилище по ключу
*/
type AddObject struct {
	Context *middleware.Context
	Handler AddObjectHandler
}

func (o *AddObject) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewAddObjectParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}