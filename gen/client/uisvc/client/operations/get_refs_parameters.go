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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetRefsParams creates a new GetRefsParams object
// with the default values initialized.
func NewGetRefsParams() *GetRefsParams {
	var ()
	return &GetRefsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRefsParamsWithTimeout creates a new GetRefsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRefsParamsWithTimeout(timeout time.Duration) *GetRefsParams {
	var ()
	return &GetRefsParams{

		timeout: timeout,
	}
}

// NewGetRefsParamsWithContext creates a new GetRefsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRefsParamsWithContext(ctx context.Context) *GetRefsParams {
	var ()
	return &GetRefsParams{

		Context: ctx,
	}
}

// NewGetRefsParamsWithHTTPClient creates a new GetRefsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRefsParamsWithHTTPClient(client *http.Client) *GetRefsParams {
	var ()
	return &GetRefsParams{
		HTTPClient: client,
	}
}

/*GetRefsParams contains all the parameters to send to the API endpoint
for the get refs operation typically these are written to a http.Request
*/
type GetRefsParams struct {

	/*Repository
	  Name of repository (eg. 'erikh/foo')

	*/
	Repository string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get refs params
func (o *GetRefsParams) WithTimeout(timeout time.Duration) *GetRefsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get refs params
func (o *GetRefsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get refs params
func (o *GetRefsParams) WithContext(ctx context.Context) *GetRefsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get refs params
func (o *GetRefsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get refs params
func (o *GetRefsParams) WithHTTPClient(client *http.Client) *GetRefsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get refs params
func (o *GetRefsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepository adds the repository to the get refs params
func (o *GetRefsParams) WithRepository(repository string) *GetRefsParams {
	o.SetRepository(repository)
	return o
}

// SetRepository adds the repository to the get refs params
func (o *GetRefsParams) SetRepository(repository string) {
	o.Repository = repository
}

// WriteToRequest writes these params to a swagger request
func (o *GetRefsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param repository
	qrRepository := o.Repository
	qRepository := qrRepository
	if qRepository != "" {
		if err := r.SetQueryParam("repository", qRepository); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
