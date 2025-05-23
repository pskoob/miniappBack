// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/pskoob/miniappBack/internal/api/definition"
)

// TransitNearOKCode is the HTTP code returned for type TransitNearOK
const TransitNearOKCode int = 200

/*
TransitNearOK Successful Transit Near Response

swagger:response transitNearOK
*/
type TransitNearOK struct {

	/*
	  In: Body
	*/
	Payload *models.NearTransit `json:"body,omitempty"`
}

// NewTransitNearOK creates TransitNearOK with default headers values
func NewTransitNearOK() *TransitNearOK {

	return &TransitNearOK{}
}

// WithPayload adds the payload to the transit near o k response
func (o *TransitNearOK) WithPayload(payload *models.NearTransit) *TransitNearOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the transit near o k response
func (o *TransitNearOK) SetPayload(payload *models.NearTransit) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TransitNearOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// TransitNearBadRequestCode is the HTTP code returned for type TransitNearBadRequest
const TransitNearBadRequestCode int = 400

/*
TransitNearBadRequest Bad request

swagger:response transitNearBadRequest
*/
type TransitNearBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewTransitNearBadRequest creates TransitNearBadRequest with default headers values
func NewTransitNearBadRequest() *TransitNearBadRequest {

	return &TransitNearBadRequest{}
}

// WithPayload adds the payload to the transit near bad request response
func (o *TransitNearBadRequest) WithPayload(payload *models.Error) *TransitNearBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the transit near bad request response
func (o *TransitNearBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TransitNearBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// TransitNearForbiddenCode is the HTTP code returned for type TransitNearForbidden
const TransitNearForbiddenCode int = 403

/*
TransitNearForbidden Forbidden

swagger:response transitNearForbidden
*/
type TransitNearForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewTransitNearForbidden creates TransitNearForbidden with default headers values
func NewTransitNearForbidden() *TransitNearForbidden {

	return &TransitNearForbidden{}
}

// WithPayload adds the payload to the transit near forbidden response
func (o *TransitNearForbidden) WithPayload(payload *models.Error) *TransitNearForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the transit near forbidden response
func (o *TransitNearForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TransitNearForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// TransitNearUnprocessableEntityCode is the HTTP code returned for type TransitNearUnprocessableEntity
const TransitNearUnprocessableEntityCode int = 422

/*
TransitNearUnprocessableEntity Unprocessable Entity

swagger:response transitNearUnprocessableEntity
*/
type TransitNearUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewTransitNearUnprocessableEntity creates TransitNearUnprocessableEntity with default headers values
func NewTransitNearUnprocessableEntity() *TransitNearUnprocessableEntity {

	return &TransitNearUnprocessableEntity{}
}

// WithPayload adds the payload to the transit near unprocessable entity response
func (o *TransitNearUnprocessableEntity) WithPayload(payload *models.Error) *TransitNearUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the transit near unprocessable entity response
func (o *TransitNearUnprocessableEntity) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TransitNearUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// TransitNearInternalServerErrorCode is the HTTP code returned for type TransitNearInternalServerError
const TransitNearInternalServerErrorCode int = 500

/*
TransitNearInternalServerError Internal server error

swagger:response transitNearInternalServerError
*/
type TransitNearInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewTransitNearInternalServerError creates TransitNearInternalServerError with default headers values
func NewTransitNearInternalServerError() *TransitNearInternalServerError {

	return &TransitNearInternalServerError{}
}

// WithPayload adds the payload to the transit near internal server error response
func (o *TransitNearInternalServerError) WithPayload(payload *models.Error) *TransitNearInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the transit near internal server error response
func (o *TransitNearInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TransitNearInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
