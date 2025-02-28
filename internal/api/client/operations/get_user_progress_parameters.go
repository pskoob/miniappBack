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

// NewGetUserProgressParams creates a new GetUserProgressParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetUserProgressParams() *GetUserProgressParams {
	return &GetUserProgressParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserProgressParamsWithTimeout creates a new GetUserProgressParams object
// with the ability to set a timeout on a request.
func NewGetUserProgressParamsWithTimeout(timeout time.Duration) *GetUserProgressParams {
	return &GetUserProgressParams{
		timeout: timeout,
	}
}

// NewGetUserProgressParamsWithContext creates a new GetUserProgressParams object
// with the ability to set a context for a request.
func NewGetUserProgressParamsWithContext(ctx context.Context) *GetUserProgressParams {
	return &GetUserProgressParams{
		Context: ctx,
	}
}

// NewGetUserProgressParamsWithHTTPClient creates a new GetUserProgressParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetUserProgressParamsWithHTTPClient(client *http.Client) *GetUserProgressParams {
	return &GetUserProgressParams{
		HTTPClient: client,
	}
}

/*
GetUserProgressParams contains all the parameters to send to the API endpoint

	for the get user progress operation.

	Typically these are written to a http.Request.
*/
type GetUserProgressParams struct {

	/* TgID.

	   The tg ID of user
	*/
	TgID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get user progress params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserProgressParams) WithDefaults() *GetUserProgressParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get user progress params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetUserProgressParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get user progress params
func (o *GetUserProgressParams) WithTimeout(timeout time.Duration) *GetUserProgressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user progress params
func (o *GetUserProgressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user progress params
func (o *GetUserProgressParams) WithContext(ctx context.Context) *GetUserProgressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user progress params
func (o *GetUserProgressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user progress params
func (o *GetUserProgressParams) WithHTTPClient(client *http.Client) *GetUserProgressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user progress params
func (o *GetUserProgressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTgID adds the tgID to the get user progress params
func (o *GetUserProgressParams) WithTgID(tgID int64) *GetUserProgressParams {
	o.SetTgID(tgID)
	return o
}

// SetTgID adds the tgId to the get user progress params
func (o *GetUserProgressParams) SetTgID(tgID int64) {
	o.TgID = tgID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserProgressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
