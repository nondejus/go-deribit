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

// NewGetPrivateCancelParams creates a new GetPrivateCancelParams object
// with the default values initialized.
func NewGetPrivateCancelParams() *GetPrivateCancelParams {
	var ()
	return &GetPrivateCancelParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPrivateCancelParamsWithTimeout creates a new GetPrivateCancelParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPrivateCancelParamsWithTimeout(timeout time.Duration) *GetPrivateCancelParams {
	var ()
	return &GetPrivateCancelParams{

		timeout: timeout,
	}
}

// NewGetPrivateCancelParamsWithContext creates a new GetPrivateCancelParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPrivateCancelParamsWithContext(ctx context.Context) *GetPrivateCancelParams {
	var ()
	return &GetPrivateCancelParams{

		Context: ctx,
	}
}

// NewGetPrivateCancelParamsWithHTTPClient creates a new GetPrivateCancelParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPrivateCancelParamsWithHTTPClient(client *http.Client) *GetPrivateCancelParams {
	var ()
	return &GetPrivateCancelParams{
		HTTPClient: client,
	}
}

/*GetPrivateCancelParams contains all the parameters to send to the API endpoint
for the get private cancel operation typically these are written to a http.Request
*/
type GetPrivateCancelParams struct {

	/*OrderID
	  The order id

	*/
	OrderID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get private cancel params
func (o *GetPrivateCancelParams) WithTimeout(timeout time.Duration) *GetPrivateCancelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get private cancel params
func (o *GetPrivateCancelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get private cancel params
func (o *GetPrivateCancelParams) WithContext(ctx context.Context) *GetPrivateCancelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get private cancel params
func (o *GetPrivateCancelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get private cancel params
func (o *GetPrivateCancelParams) WithHTTPClient(client *http.Client) *GetPrivateCancelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get private cancel params
func (o *GetPrivateCancelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderID adds the orderID to the get private cancel params
func (o *GetPrivateCancelParams) WithOrderID(orderID string) *GetPrivateCancelParams {
	o.SetOrderID(orderID)
	return o
}

// SetOrderID adds the orderId to the get private cancel params
func (o *GetPrivateCancelParams) SetOrderID(orderID string) {
	o.OrderID = orderID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivateCancelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param order_id
	qrOrderID := o.OrderID
	qOrderID := qrOrderID
	if qOrderID != "" {
		if err := r.SetQueryParam("order_id", qOrderID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
