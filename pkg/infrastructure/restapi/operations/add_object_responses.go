// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// AddObjectOKCode is the HTTP code returned for type AddObjectOK
const AddObjectOKCode int = 200

/*
AddObjectOK OK

swagger:response addObjectOK
*/
type AddObjectOK struct {
}

// NewAddObjectOK creates AddObjectOK with default headers values
func NewAddObjectOK() *AddObjectOK {

	return &AddObjectOK{}
}

// WriteResponse to the client
func (o *AddObjectOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*
AddObjectDefault Any Error

swagger:response addObjectDefault
*/
type AddObjectDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewAddObjectDefault creates AddObjectDefault with default headers values
func NewAddObjectDefault(code int) *AddObjectDefault {
	if code <= 0 {
		code = 500
	}

	return &AddObjectDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the add object default response
func (o *AddObjectDefault) WithStatusCode(code int) *AddObjectDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the add object default response
func (o *AddObjectDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the add object default response
func (o *AddObjectDefault) WithPayload(payload string) *AddObjectDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add object default response
func (o *AddObjectDefault) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddObjectDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
