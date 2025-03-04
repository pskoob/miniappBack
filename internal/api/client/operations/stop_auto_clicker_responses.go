// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	models "github.com/pskoob/miniappBack/internal/api/definition"
)

// StopAutoClickerReader is a Reader for the StopAutoClicker structure.
type StopAutoClickerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopAutoClickerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopAutoClickerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStopAutoClickerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStopAutoClickerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewStopAutoClickerUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStopAutoClickerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /stop_auto_clicker/{tg_id}] StopAutoClicker", response, response.Code())
	}
}

// NewStopAutoClickerOK creates a StopAutoClickerOK with default headers values
func NewStopAutoClickerOK() *StopAutoClickerOK {
	return &StopAutoClickerOK{}
}

/*
StopAutoClickerOK describes a response with status code 200, with default header values.

Save Progress Response
*/
type StopAutoClickerOK struct {
	Payload *models.Error
}

// IsSuccess returns true when this stop auto clicker o k response has a 2xx status code
func (o *StopAutoClickerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stop auto clicker o k response has a 3xx status code
func (o *StopAutoClickerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop auto clicker o k response has a 4xx status code
func (o *StopAutoClickerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stop auto clicker o k response has a 5xx status code
func (o *StopAutoClickerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stop auto clicker o k response a status code equal to that given
func (o *StopAutoClickerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the stop auto clicker o k response
func (o *StopAutoClickerOK) Code() int {
	return 200
}

func (o *StopAutoClickerOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /stop_auto_clicker/{tg_id}][%d] stopAutoClickerOK %s", 200, payload)
}

func (o *StopAutoClickerOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /stop_auto_clicker/{tg_id}][%d] stopAutoClickerOK %s", 200, payload)
}

func (o *StopAutoClickerOK) GetPayload() *models.Error {
	return o.Payload
}

func (o *StopAutoClickerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopAutoClickerBadRequest creates a StopAutoClickerBadRequest with default headers values
func NewStopAutoClickerBadRequest() *StopAutoClickerBadRequest {
	return &StopAutoClickerBadRequest{}
}

/*
StopAutoClickerBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type StopAutoClickerBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this stop auto clicker bad request response has a 2xx status code
func (o *StopAutoClickerBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop auto clicker bad request response has a 3xx status code
func (o *StopAutoClickerBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop auto clicker bad request response has a 4xx status code
func (o *StopAutoClickerBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop auto clicker bad request response has a 5xx status code
func (o *StopAutoClickerBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stop auto clicker bad request response a status code equal to that given
func (o *StopAutoClickerBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the stop auto clicker bad request response
func (o *StopAutoClickerBadRequest) Code() int {
	return 400
}

func (o *StopAutoClickerBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /stop_auto_clicker/{tg_id}][%d] stopAutoClickerBadRequest %s", 400, payload)
}

func (o *StopAutoClickerBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /stop_auto_clicker/{tg_id}][%d] stopAutoClickerBadRequest %s", 400, payload)
}

func (o *StopAutoClickerBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *StopAutoClickerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopAutoClickerForbidden creates a StopAutoClickerForbidden with default headers values
func NewStopAutoClickerForbidden() *StopAutoClickerForbidden {
	return &StopAutoClickerForbidden{}
}

/*
StopAutoClickerForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StopAutoClickerForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this stop auto clicker forbidden response has a 2xx status code
func (o *StopAutoClickerForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop auto clicker forbidden response has a 3xx status code
func (o *StopAutoClickerForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop auto clicker forbidden response has a 4xx status code
func (o *StopAutoClickerForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop auto clicker forbidden response has a 5xx status code
func (o *StopAutoClickerForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stop auto clicker forbidden response a status code equal to that given
func (o *StopAutoClickerForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the stop auto clicker forbidden response
func (o *StopAutoClickerForbidden) Code() int {
	return 403
}

func (o *StopAutoClickerForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /stop_auto_clicker/{tg_id}][%d] stopAutoClickerForbidden %s", 403, payload)
}

func (o *StopAutoClickerForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /stop_auto_clicker/{tg_id}][%d] stopAutoClickerForbidden %s", 403, payload)
}

func (o *StopAutoClickerForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *StopAutoClickerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopAutoClickerUnprocessableEntity creates a StopAutoClickerUnprocessableEntity with default headers values
func NewStopAutoClickerUnprocessableEntity() *StopAutoClickerUnprocessableEntity {
	return &StopAutoClickerUnprocessableEntity{}
}

/*
StopAutoClickerUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type StopAutoClickerUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this stop auto clicker unprocessable entity response has a 2xx status code
func (o *StopAutoClickerUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop auto clicker unprocessable entity response has a 3xx status code
func (o *StopAutoClickerUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop auto clicker unprocessable entity response has a 4xx status code
func (o *StopAutoClickerUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop auto clicker unprocessable entity response has a 5xx status code
func (o *StopAutoClickerUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this stop auto clicker unprocessable entity response a status code equal to that given
func (o *StopAutoClickerUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the stop auto clicker unprocessable entity response
func (o *StopAutoClickerUnprocessableEntity) Code() int {
	return 422
}

func (o *StopAutoClickerUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /stop_auto_clicker/{tg_id}][%d] stopAutoClickerUnprocessableEntity %s", 422, payload)
}

func (o *StopAutoClickerUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /stop_auto_clicker/{tg_id}][%d] stopAutoClickerUnprocessableEntity %s", 422, payload)
}

func (o *StopAutoClickerUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *StopAutoClickerUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopAutoClickerInternalServerError creates a StopAutoClickerInternalServerError with default headers values
func NewStopAutoClickerInternalServerError() *StopAutoClickerInternalServerError {
	return &StopAutoClickerInternalServerError{}
}

/*
StopAutoClickerInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type StopAutoClickerInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this stop auto clicker internal server error response has a 2xx status code
func (o *StopAutoClickerInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop auto clicker internal server error response has a 3xx status code
func (o *StopAutoClickerInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop auto clicker internal server error response has a 4xx status code
func (o *StopAutoClickerInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stop auto clicker internal server error response has a 5xx status code
func (o *StopAutoClickerInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stop auto clicker internal server error response a status code equal to that given
func (o *StopAutoClickerInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the stop auto clicker internal server error response
func (o *StopAutoClickerInternalServerError) Code() int {
	return 500
}

func (o *StopAutoClickerInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /stop_auto_clicker/{tg_id}][%d] stopAutoClickerInternalServerError %s", 500, payload)
}

func (o *StopAutoClickerInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /stop_auto_clicker/{tg_id}][%d] stopAutoClickerInternalServerError %s", 500, payload)
}

func (o *StopAutoClickerInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *StopAutoClickerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
