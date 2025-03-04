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

// NewStartAutoClickerParams creates a new StartAutoClickerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStartAutoClickerParams() *StartAutoClickerParams {
	return &StartAutoClickerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStartAutoClickerParamsWithTimeout creates a new StartAutoClickerParams object
// with the ability to set a timeout on a request.
func NewStartAutoClickerParamsWithTimeout(timeout time.Duration) *StartAutoClickerParams {
	return &StartAutoClickerParams{
		timeout: timeout,
	}
}

// NewStartAutoClickerParamsWithContext creates a new StartAutoClickerParams object
// with the ability to set a context for a request.
func NewStartAutoClickerParamsWithContext(ctx context.Context) *StartAutoClickerParams {
	return &StartAutoClickerParams{
		Context: ctx,
	}
}

// NewStartAutoClickerParamsWithHTTPClient creates a new StartAutoClickerParams object
// with the ability to set a custom HTTPClient for a request.
func NewStartAutoClickerParamsWithHTTPClient(client *http.Client) *StartAutoClickerParams {
	return &StartAutoClickerParams{
		HTTPClient: client,
	}
}

/*
StartAutoClickerParams contains all the parameters to send to the API endpoint

	for the start auto clicker operation.

	Typically these are written to a http.Request.
*/
type StartAutoClickerParams struct {

	/* TgID.

	   The tg ID of user
	*/
	TgID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the start auto clicker params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartAutoClickerParams) WithDefaults() *StartAutoClickerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the start auto clicker params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StartAutoClickerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the start auto clicker params
func (o *StartAutoClickerParams) WithTimeout(timeout time.Duration) *StartAutoClickerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start auto clicker params
func (o *StartAutoClickerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start auto clicker params
func (o *StartAutoClickerParams) WithContext(ctx context.Context) *StartAutoClickerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start auto clicker params
func (o *StartAutoClickerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start auto clicker params
func (o *StartAutoClickerParams) WithHTTPClient(client *http.Client) *StartAutoClickerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start auto clicker params
func (o *StartAutoClickerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTgID adds the tgID to the start auto clicker params
func (o *StartAutoClickerParams) WithTgID(tgID int64) *StartAutoClickerParams {
	o.SetTgID(tgID)
	return o
}

// SetTgID adds the tgId to the start auto clicker params
func (o *StartAutoClickerParams) SetTgID(tgID int64) {
	o.TgID = tgID
}

// WriteToRequest writes these params to a swagger request
func (o *StartAutoClickerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
