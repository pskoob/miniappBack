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

// CollectReferralEarnReader is a Reader for the CollectReferralEarn structure.
type CollectReferralEarnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CollectReferralEarnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCollectReferralEarnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCollectReferralEarnBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCollectReferralEarnForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCollectReferralEarnUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCollectReferralEarnInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /collect_referral_earn/{tg_id}] CollectReferralEarn", response, response.Code())
	}
}

// NewCollectReferralEarnOK creates a CollectReferralEarnOK with default headers values
func NewCollectReferralEarnOK() *CollectReferralEarnOK {
	return &CollectReferralEarnOK{}
}

/*
CollectReferralEarnOK describes a response with status code 200, with default header values.

Referral Balance Response
*/
type CollectReferralEarnOK struct {
	Payload *models.ReferralBalance
}

// IsSuccess returns true when this collect referral earn o k response has a 2xx status code
func (o *CollectReferralEarnOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this collect referral earn o k response has a 3xx status code
func (o *CollectReferralEarnOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this collect referral earn o k response has a 4xx status code
func (o *CollectReferralEarnOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this collect referral earn o k response has a 5xx status code
func (o *CollectReferralEarnOK) IsServerError() bool {
	return false
}

// IsCode returns true when this collect referral earn o k response a status code equal to that given
func (o *CollectReferralEarnOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the collect referral earn o k response
func (o *CollectReferralEarnOK) Code() int {
	return 200
}

func (o *CollectReferralEarnOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /collect_referral_earn/{tg_id}][%d] collectReferralEarnOK %s", 200, payload)
}

func (o *CollectReferralEarnOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /collect_referral_earn/{tg_id}][%d] collectReferralEarnOK %s", 200, payload)
}

func (o *CollectReferralEarnOK) GetPayload() *models.ReferralBalance {
	return o.Payload
}

func (o *CollectReferralEarnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReferralBalance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCollectReferralEarnBadRequest creates a CollectReferralEarnBadRequest with default headers values
func NewCollectReferralEarnBadRequest() *CollectReferralEarnBadRequest {
	return &CollectReferralEarnBadRequest{}
}

/*
CollectReferralEarnBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CollectReferralEarnBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this collect referral earn bad request response has a 2xx status code
func (o *CollectReferralEarnBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this collect referral earn bad request response has a 3xx status code
func (o *CollectReferralEarnBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this collect referral earn bad request response has a 4xx status code
func (o *CollectReferralEarnBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this collect referral earn bad request response has a 5xx status code
func (o *CollectReferralEarnBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this collect referral earn bad request response a status code equal to that given
func (o *CollectReferralEarnBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the collect referral earn bad request response
func (o *CollectReferralEarnBadRequest) Code() int {
	return 400
}

func (o *CollectReferralEarnBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /collect_referral_earn/{tg_id}][%d] collectReferralEarnBadRequest %s", 400, payload)
}

func (o *CollectReferralEarnBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /collect_referral_earn/{tg_id}][%d] collectReferralEarnBadRequest %s", 400, payload)
}

func (o *CollectReferralEarnBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *CollectReferralEarnBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCollectReferralEarnForbidden creates a CollectReferralEarnForbidden with default headers values
func NewCollectReferralEarnForbidden() *CollectReferralEarnForbidden {
	return &CollectReferralEarnForbidden{}
}

/*
CollectReferralEarnForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CollectReferralEarnForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this collect referral earn forbidden response has a 2xx status code
func (o *CollectReferralEarnForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this collect referral earn forbidden response has a 3xx status code
func (o *CollectReferralEarnForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this collect referral earn forbidden response has a 4xx status code
func (o *CollectReferralEarnForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this collect referral earn forbidden response has a 5xx status code
func (o *CollectReferralEarnForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this collect referral earn forbidden response a status code equal to that given
func (o *CollectReferralEarnForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the collect referral earn forbidden response
func (o *CollectReferralEarnForbidden) Code() int {
	return 403
}

func (o *CollectReferralEarnForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /collect_referral_earn/{tg_id}][%d] collectReferralEarnForbidden %s", 403, payload)
}

func (o *CollectReferralEarnForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /collect_referral_earn/{tg_id}][%d] collectReferralEarnForbidden %s", 403, payload)
}

func (o *CollectReferralEarnForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *CollectReferralEarnForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCollectReferralEarnUnprocessableEntity creates a CollectReferralEarnUnprocessableEntity with default headers values
func NewCollectReferralEarnUnprocessableEntity() *CollectReferralEarnUnprocessableEntity {
	return &CollectReferralEarnUnprocessableEntity{}
}

/*
CollectReferralEarnUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type CollectReferralEarnUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this collect referral earn unprocessable entity response has a 2xx status code
func (o *CollectReferralEarnUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this collect referral earn unprocessable entity response has a 3xx status code
func (o *CollectReferralEarnUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this collect referral earn unprocessable entity response has a 4xx status code
func (o *CollectReferralEarnUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this collect referral earn unprocessable entity response has a 5xx status code
func (o *CollectReferralEarnUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this collect referral earn unprocessable entity response a status code equal to that given
func (o *CollectReferralEarnUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the collect referral earn unprocessable entity response
func (o *CollectReferralEarnUnprocessableEntity) Code() int {
	return 422
}

func (o *CollectReferralEarnUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /collect_referral_earn/{tg_id}][%d] collectReferralEarnUnprocessableEntity %s", 422, payload)
}

func (o *CollectReferralEarnUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /collect_referral_earn/{tg_id}][%d] collectReferralEarnUnprocessableEntity %s", 422, payload)
}

func (o *CollectReferralEarnUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *CollectReferralEarnUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCollectReferralEarnInternalServerError creates a CollectReferralEarnInternalServerError with default headers values
func NewCollectReferralEarnInternalServerError() *CollectReferralEarnInternalServerError {
	return &CollectReferralEarnInternalServerError{}
}

/*
CollectReferralEarnInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type CollectReferralEarnInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this collect referral earn internal server error response has a 2xx status code
func (o *CollectReferralEarnInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this collect referral earn internal server error response has a 3xx status code
func (o *CollectReferralEarnInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this collect referral earn internal server error response has a 4xx status code
func (o *CollectReferralEarnInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this collect referral earn internal server error response has a 5xx status code
func (o *CollectReferralEarnInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this collect referral earn internal server error response a status code equal to that given
func (o *CollectReferralEarnInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the collect referral earn internal server error response
func (o *CollectReferralEarnInternalServerError) Code() int {
	return 500
}

func (o *CollectReferralEarnInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /collect_referral_earn/{tg_id}][%d] collectReferralEarnInternalServerError %s", 500, payload)
}

func (o *CollectReferralEarnInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /collect_referral_earn/{tg_id}][%d] collectReferralEarnInternalServerError %s", 500, payload)
}

func (o *CollectReferralEarnInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *CollectReferralEarnInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
