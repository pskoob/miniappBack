// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetUserCardsParams creates a new GetUserCardsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUserCardsParams() *GetUserCardsParams {
	return &GetUserCardsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserCardsParamsWithTimeout creates a new GetUserCardsParams object
// with the ability to set a timeout on a request.
func NewGetUserCardsParamsWithTimeout(timeout time.Duration) *GetUserCardsParams {
	return &GetUserCardsParams{
		timeout: timeout,
	}
}

// NewGetUserCardsParamsWithContext creates a new GetUserCardsParams object
// with the ability to set a context for a request.
func NewGetUserCardsParamsWithContext(ctx context.Context) *GetUserCardsParams {
	return &GetUserCardsParams{
		Context: ctx,
	}
}

// NewGetUserCardsParamsWithHTTPClient creates a new GetUserCardsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUserCardsParamsWithHTTPClient(client *http.Client) *GetUserCardsParams {
	return &GetUserCardsParams{
		HTTPClient: client,
	}
}

/*
GetUserCardsParams contains all the parameters to send to the API endpoint

	for the get user cards operation.

	Typically these are written to a http.Request.
*/
type GetUserCardsParams struct {

	/* TgID.

	   The tg ID of user
	*/
	TgID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get user cards params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserCardsParams) WithDefaults() *GetUserCardsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get user cards params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserCardsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get user cards params
func (o *GetUserCardsParams) WithTimeout(timeout time.Duration) *GetUserCardsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user cards params
func (o *GetUserCardsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user cards params
func (o *GetUserCardsParams) WithContext(ctx context.Context) *GetUserCardsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user cards params
func (o *GetUserCardsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user cards params
func (o *GetUserCardsParams) WithHTTPClient(client *http.Client) *GetUserCardsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user cards params
func (o *GetUserCardsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTgID adds the tgID to the get user cards params
func (o *GetUserCardsParams) WithTgID(tgID int64) *GetUserCardsParams {
	o.SetTgID(tgID)
	return o
}

// SetTgID adds the tgId to the get user cards params
func (o *GetUserCardsParams) SetTgID(tgID int64) {
	o.TgID = tgID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserCardsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param tg_id
	if err := r.SetPathParam("tg_id", swag.FormatInt64(o.TgID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
