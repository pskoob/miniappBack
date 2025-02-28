// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/pskoob/miniappBack/internal/api/definition"
)

// GetUserProgressOKCode is the HTTP code returned for type GetUserProgressOK
const GetUserProgressOKCode int = 200

/*
GetUserProgressOK Get User Progress Response

swagger:response getUserProgressOK
*/
type GetUserProgressOK struct {

	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewGetUserProgressOK creates GetUserProgressOK with default headers values
func NewGetUserProgressOK() *GetUserProgressOK {

	return &GetUserProgressOK{}
}

// WithPayload adds the payload to the get user progress o k response
func (o *GetUserProgressOK) WithPayload(payload *models.User) *GetUserProgressOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user progress o k response
func (o *GetUserProgressOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserProgressOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserProgressBadRequestCode is the HTTP code returned for type GetUserProgressBadRequest
const GetUserProgressBadRequestCode int = 400

/*
GetUserProgressBadRequest Bad request

swagger:response getUserProgressBadRequest
*/
type GetUserProgressBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUserProgressBadRequest creates GetUserProgressBadRequest with default headers values
func NewGetUserProgressBadRequest() *GetUserProgressBadRequest {

	return &GetUserProgressBadRequest{}
}

// WithPayload adds the payload to the get user progress bad request response
func (o *GetUserProgressBadRequest) WithPayload(payload *models.Error) *GetUserProgressBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user progress bad request response
func (o *GetUserProgressBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserProgressBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserProgressForbiddenCode is the HTTP code returned for type GetUserProgressForbidden
const GetUserProgressForbiddenCode int = 403

/*
GetUserProgressForbidden Forbidden

swagger:response getUserProgressForbidden
*/
type GetUserProgressForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUserProgressForbidden creates GetUserProgressForbidden with default headers values
func NewGetUserProgressForbidden() *GetUserProgressForbidden {

	return &GetUserProgressForbidden{}
}

// WithPayload adds the payload to the get user progress forbidden response
func (o *GetUserProgressForbidden) WithPayload(payload *models.Error) *GetUserProgressForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user progress forbidden response
func (o *GetUserProgressForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserProgressForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserProgressUnprocessableEntityCode is the HTTP code returned for type GetUserProgressUnprocessableEntity
const GetUserProgressUnprocessableEntityCode int = 422

/*
GetUserProgressUnprocessableEntity Unprocessable Entity

swagger:response getUserProgressUnprocessableEntity
*/
type GetUserProgressUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUserProgressUnprocessableEntity creates GetUserProgressUnprocessableEntity with default headers values
func NewGetUserProgressUnprocessableEntity() *GetUserProgressUnprocessableEntity {

	return &GetUserProgressUnprocessableEntity{}
}

// WithPayload adds the payload to the get user progress unprocessable entity response
func (o *GetUserProgressUnprocessableEntity) WithPayload(payload *models.Error) *GetUserProgressUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user progress unprocessable entity response
func (o *GetUserProgressUnprocessableEntity) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserProgressUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserProgressInternalServerErrorCode is the HTTP code returned for type GetUserProgressInternalServerError
const GetUserProgressInternalServerErrorCode int = 500

/*
GetUserProgressInternalServerError Internal server error

swagger:response getUserProgressInternalServerError
*/
type GetUserProgressInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetUserProgressInternalServerError creates GetUserProgressInternalServerError with default headers values
func NewGetUserProgressInternalServerError() *GetUserProgressInternalServerError {

	return &GetUserProgressInternalServerError{}
}

// WithPayload adds the payload to the get user progress internal server error response
func (o *GetUserProgressInternalServerError) WithPayload(payload *models.Error) *GetUserProgressInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user progress internal server error response
func (o *GetUserProgressInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserProgressInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
