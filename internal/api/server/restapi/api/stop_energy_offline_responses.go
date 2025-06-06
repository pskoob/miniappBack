// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/pskoob/miniappBack/internal/api/definition"
)

// StopEnergyOfflineOKCode is the HTTP code returned for type StopEnergyOfflineOK
const StopEnergyOfflineOKCode int = 200

/*
StopEnergyOfflineOK Stop Collect Energy Response

swagger:response stopEnergyOfflineOK
*/
type StopEnergyOfflineOK struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewStopEnergyOfflineOK creates StopEnergyOfflineOK with default headers values
func NewStopEnergyOfflineOK() *StopEnergyOfflineOK {

	return &StopEnergyOfflineOK{}
}

// WithPayload adds the payload to the stop energy offline o k response
func (o *StopEnergyOfflineOK) WithPayload(payload *models.Error) *StopEnergyOfflineOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the stop energy offline o k response
func (o *StopEnergyOfflineOK) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StopEnergyOfflineOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// StopEnergyOfflineBadRequestCode is the HTTP code returned for type StopEnergyOfflineBadRequest
const StopEnergyOfflineBadRequestCode int = 400

/*
StopEnergyOfflineBadRequest Bad request

swagger:response stopEnergyOfflineBadRequest
*/
type StopEnergyOfflineBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewStopEnergyOfflineBadRequest creates StopEnergyOfflineBadRequest with default headers values
func NewStopEnergyOfflineBadRequest() *StopEnergyOfflineBadRequest {

	return &StopEnergyOfflineBadRequest{}
}

// WithPayload adds the payload to the stop energy offline bad request response
func (o *StopEnergyOfflineBadRequest) WithPayload(payload *models.Error) *StopEnergyOfflineBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the stop energy offline bad request response
func (o *StopEnergyOfflineBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StopEnergyOfflineBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// StopEnergyOfflineForbiddenCode is the HTTP code returned for type StopEnergyOfflineForbidden
const StopEnergyOfflineForbiddenCode int = 403

/*
StopEnergyOfflineForbidden Forbidden

swagger:response stopEnergyOfflineForbidden
*/
type StopEnergyOfflineForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewStopEnergyOfflineForbidden creates StopEnergyOfflineForbidden with default headers values
func NewStopEnergyOfflineForbidden() *StopEnergyOfflineForbidden {

	return &StopEnergyOfflineForbidden{}
}

// WithPayload adds the payload to the stop energy offline forbidden response
func (o *StopEnergyOfflineForbidden) WithPayload(payload *models.Error) *StopEnergyOfflineForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the stop energy offline forbidden response
func (o *StopEnergyOfflineForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StopEnergyOfflineForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// StopEnergyOfflineUnprocessableEntityCode is the HTTP code returned for type StopEnergyOfflineUnprocessableEntity
const StopEnergyOfflineUnprocessableEntityCode int = 422

/*
StopEnergyOfflineUnprocessableEntity Unprocessable Entity

swagger:response stopEnergyOfflineUnprocessableEntity
*/
type StopEnergyOfflineUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewStopEnergyOfflineUnprocessableEntity creates StopEnergyOfflineUnprocessableEntity with default headers values
func NewStopEnergyOfflineUnprocessableEntity() *StopEnergyOfflineUnprocessableEntity {

	return &StopEnergyOfflineUnprocessableEntity{}
}

// WithPayload adds the payload to the stop energy offline unprocessable entity response
func (o *StopEnergyOfflineUnprocessableEntity) WithPayload(payload *models.Error) *StopEnergyOfflineUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the stop energy offline unprocessable entity response
func (o *StopEnergyOfflineUnprocessableEntity) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StopEnergyOfflineUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// StopEnergyOfflineInternalServerErrorCode is the HTTP code returned for type StopEnergyOfflineInternalServerError
const StopEnergyOfflineInternalServerErrorCode int = 500

/*
StopEnergyOfflineInternalServerError Internal server error

swagger:response stopEnergyOfflineInternalServerError
*/
type StopEnergyOfflineInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewStopEnergyOfflineInternalServerError creates StopEnergyOfflineInternalServerError with default headers values
func NewStopEnergyOfflineInternalServerError() *StopEnergyOfflineInternalServerError {

	return &StopEnergyOfflineInternalServerError{}
}

// WithPayload adds the payload to the stop energy offline internal server error response
func (o *StopEnergyOfflineInternalServerError) WithPayload(payload *models.Error) *StopEnergyOfflineInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the stop energy offline internal server error response
func (o *StopEnergyOfflineInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StopEnergyOfflineInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
