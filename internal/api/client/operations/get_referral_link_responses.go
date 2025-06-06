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

// GetReferralLinkReader is a Reader for the GetReferralLink structure.
type GetReferralLinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReferralLinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReferralLinkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetReferralLinkBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetReferralLinkForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetReferralLinkUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetReferralLinkInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /get_referral_link/{tg_id}] GetReferralLink", response, response.Code())
	}
}

// NewGetReferralLinkOK creates a GetReferralLinkOK with default headers values
func NewGetReferralLinkOK() *GetReferralLinkOK {
	return &GetReferralLinkOK{}
}

/*
GetReferralLinkOK describes a response with status code 200, with default header values.

Referral Link Response
*/
type GetReferralLinkOK struct {
	Payload *models.ReferralLink
}

// IsSuccess returns true when this get referral link o k response has a 2xx status code
func (o *GetReferralLinkOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get referral link o k response has a 3xx status code
func (o *GetReferralLinkOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get referral link o k response has a 4xx status code
func (o *GetReferralLinkOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get referral link o k response has a 5xx status code
func (o *GetReferralLinkOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get referral link o k response a status code equal to that given
func (o *GetReferralLinkOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get referral link o k response
func (o *GetReferralLinkOK) Code() int {
	return 200
}

func (o *GetReferralLinkOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /get_referral_link/{tg_id}][%d] getReferralLinkOK %s", 200, payload)
}

func (o *GetReferralLinkOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /get_referral_link/{tg_id}][%d] getReferralLinkOK %s", 200, payload)
}

func (o *GetReferralLinkOK) GetPayload() *models.ReferralLink {
	return o.Payload
}

func (o *GetReferralLinkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReferralLink)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReferralLinkBadRequest creates a GetReferralLinkBadRequest with default headers values
func NewGetReferralLinkBadRequest() *GetReferralLinkBadRequest {
	return &GetReferralLinkBadRequest{}
}

/*
GetReferralLinkBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetReferralLinkBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this get referral link bad request response has a 2xx status code
func (o *GetReferralLinkBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get referral link bad request response has a 3xx status code
func (o *GetReferralLinkBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get referral link bad request response has a 4xx status code
func (o *GetReferralLinkBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get referral link bad request response has a 5xx status code
func (o *GetReferralLinkBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get referral link bad request response a status code equal to that given
func (o *GetReferralLinkBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get referral link bad request response
func (o *GetReferralLinkBadRequest) Code() int {
	return 400
}

func (o *GetReferralLinkBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /get_referral_link/{tg_id}][%d] getReferralLinkBadRequest %s", 400, payload)
}

func (o *GetReferralLinkBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /get_referral_link/{tg_id}][%d] getReferralLinkBadRequest %s", 400, payload)
}

func (o *GetReferralLinkBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetReferralLinkBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReferralLinkForbidden creates a GetReferralLinkForbidden with default headers values
func NewGetReferralLinkForbidden() *GetReferralLinkForbidden {
	return &GetReferralLinkForbidden{}
}

/*
GetReferralLinkForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetReferralLinkForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this get referral link forbidden response has a 2xx status code
func (o *GetReferralLinkForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get referral link forbidden response has a 3xx status code
func (o *GetReferralLinkForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get referral link forbidden response has a 4xx status code
func (o *GetReferralLinkForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get referral link forbidden response has a 5xx status code
func (o *GetReferralLinkForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get referral link forbidden response a status code equal to that given
func (o *GetReferralLinkForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get referral link forbidden response
func (o *GetReferralLinkForbidden) Code() int {
	return 403
}

func (o *GetReferralLinkForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /get_referral_link/{tg_id}][%d] getReferralLinkForbidden %s", 403, payload)
}

func (o *GetReferralLinkForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /get_referral_link/{tg_id}][%d] getReferralLinkForbidden %s", 403, payload)
}

func (o *GetReferralLinkForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetReferralLinkForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReferralLinkUnprocessableEntity creates a GetReferralLinkUnprocessableEntity with default headers values
func NewGetReferralLinkUnprocessableEntity() *GetReferralLinkUnprocessableEntity {
	return &GetReferralLinkUnprocessableEntity{}
}

/*
GetReferralLinkUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type GetReferralLinkUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this get referral link unprocessable entity response has a 2xx status code
func (o *GetReferralLinkUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get referral link unprocessable entity response has a 3xx status code
func (o *GetReferralLinkUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get referral link unprocessable entity response has a 4xx status code
func (o *GetReferralLinkUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this get referral link unprocessable entity response has a 5xx status code
func (o *GetReferralLinkUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this get referral link unprocessable entity response a status code equal to that given
func (o *GetReferralLinkUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the get referral link unprocessable entity response
func (o *GetReferralLinkUnprocessableEntity) Code() int {
	return 422
}

func (o *GetReferralLinkUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /get_referral_link/{tg_id}][%d] getReferralLinkUnprocessableEntity %s", 422, payload)
}

func (o *GetReferralLinkUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /get_referral_link/{tg_id}][%d] getReferralLinkUnprocessableEntity %s", 422, payload)
}

func (o *GetReferralLinkUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetReferralLinkUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReferralLinkInternalServerError creates a GetReferralLinkInternalServerError with default headers values
func NewGetReferralLinkInternalServerError() *GetReferralLinkInternalServerError {
	return &GetReferralLinkInternalServerError{}
}

/*
GetReferralLinkInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetReferralLinkInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this get referral link internal server error response has a 2xx status code
func (o *GetReferralLinkInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get referral link internal server error response has a 3xx status code
func (o *GetReferralLinkInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get referral link internal server error response has a 4xx status code
func (o *GetReferralLinkInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get referral link internal server error response has a 5xx status code
func (o *GetReferralLinkInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get referral link internal server error response a status code equal to that given
func (o *GetReferralLinkInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get referral link internal server error response
func (o *GetReferralLinkInternalServerError) Code() int {
	return 500
}

func (o *GetReferralLinkInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /get_referral_link/{tg_id}][%d] getReferralLinkInternalServerError %s", 500, payload)
}

func (o *GetReferralLinkInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /get_referral_link/{tg_id}][%d] getReferralLinkInternalServerError %s", 500, payload)
}

func (o *GetReferralLinkInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetReferralLinkInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
