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

	models "github.com/pskoob/miniappBack/internal/api/definition"
)

// NewSaveProgressParams creates a new SaveProgressParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSaveProgressParams() *SaveProgressParams {
	return &SaveProgressParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSaveProgressParamsWithTimeout creates a new SaveProgressParams object
// with the ability to set a timeout on a request.
func NewSaveProgressParamsWithTimeout(timeout time.Duration) *SaveProgressParams {
	return &SaveProgressParams{
		timeout: timeout,
	}
}

// NewSaveProgressParamsWithContext creates a new SaveProgressParams object
// with the ability to set a context for a request.
func NewSaveProgressParamsWithContext(ctx context.Context) *SaveProgressParams {
	return &SaveProgressParams{
		Context: ctx,
	}
}

// NewSaveProgressParamsWithHTTPClient creates a new SaveProgressParams object
// with the ability to set a custom HTTPClient for a request.
func NewSaveProgressParamsWithHTTPClient(client *http.Client) *SaveProgressParams {
	return &SaveProgressParams{
		HTTPClient: client,
	}
}

/*
SaveProgressParams contains all the parameters to send to the API endpoint

	for the save progress operation.

	Typically these are written to a http.Request.
*/
type SaveProgressParams struct {

	/* Progress.

	   Save progress body
	*/
	Progress *models.Progress

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the save progress params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SaveProgressParams) WithDefaults() *SaveProgressParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the save progress params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SaveProgressParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the save progress params
func (o *SaveProgressParams) WithTimeout(timeout time.Duration) *SaveProgressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the save progress params
func (o *SaveProgressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the save progress params
func (o *SaveProgressParams) WithContext(ctx context.Context) *SaveProgressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the save progress params
func (o *SaveProgressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the save progress params
func (o *SaveProgressParams) WithHTTPClient(client *http.Client) *SaveProgressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the save progress params
func (o *SaveProgressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProgress adds the progress to the save progress params
func (o *SaveProgressParams) WithProgress(progress *models.Progress) *SaveProgressParams {
	o.SetProgress(progress)
	return o
}

// SetProgress adds the progress to the save progress params
func (o *SaveProgressParams) SetProgress(progress *models.Progress) {
	o.Progress = progress
}

// WriteToRequest writes these params to a swagger request
func (o *SaveProgressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Progress != nil {
		if err := r.SetBodyParam(o.Progress); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
