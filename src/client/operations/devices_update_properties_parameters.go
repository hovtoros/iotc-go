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

	"com.azure.iot/iotcentral/iotcgo/models"
)

// NewDevicesUpdatePropertiesParams creates a new DevicesUpdatePropertiesParams object
// with the default values initialized.
func NewDevicesUpdatePropertiesParams() *DevicesUpdatePropertiesParams {
	var ()
	return &DevicesUpdatePropertiesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDevicesUpdatePropertiesParamsWithTimeout creates a new DevicesUpdatePropertiesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDevicesUpdatePropertiesParamsWithTimeout(timeout time.Duration) *DevicesUpdatePropertiesParams {
	var ()
	return &DevicesUpdatePropertiesParams{

		timeout: timeout,
	}
}

// NewDevicesUpdatePropertiesParamsWithContext creates a new DevicesUpdatePropertiesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDevicesUpdatePropertiesParamsWithContext(ctx context.Context) *DevicesUpdatePropertiesParams {
	var ()
	return &DevicesUpdatePropertiesParams{

		Context: ctx,
	}
}

// NewDevicesUpdatePropertiesParamsWithHTTPClient creates a new DevicesUpdatePropertiesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDevicesUpdatePropertiesParamsWithHTTPClient(client *http.Client) *DevicesUpdatePropertiesParams {
	var ()
	return &DevicesUpdatePropertiesParams{
		HTTPClient: client,
	}
}

/*DevicesUpdatePropertiesParams contains all the parameters to send to the API endpoint
for the devices update properties operation typically these are written to a http.Request
*/
type DevicesUpdatePropertiesParams struct {

	/*Body
	  Device properties.

	*/
	Body models.DeviceProperties
	/*DeviceID
	  Unique ID of the device.

	*/
	DeviceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the devices update properties params
func (o *DevicesUpdatePropertiesParams) WithTimeout(timeout time.Duration) *DevicesUpdatePropertiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the devices update properties params
func (o *DevicesUpdatePropertiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the devices update properties params
func (o *DevicesUpdatePropertiesParams) WithContext(ctx context.Context) *DevicesUpdatePropertiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the devices update properties params
func (o *DevicesUpdatePropertiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the devices update properties params
func (o *DevicesUpdatePropertiesParams) WithHTTPClient(client *http.Client) *DevicesUpdatePropertiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the devices update properties params
func (o *DevicesUpdatePropertiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the devices update properties params
func (o *DevicesUpdatePropertiesParams) WithBody(body models.DeviceProperties) *DevicesUpdatePropertiesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the devices update properties params
func (o *DevicesUpdatePropertiesParams) SetBody(body models.DeviceProperties) {
	o.Body = body
}

// WithDeviceID adds the deviceID to the devices update properties params
func (o *DevicesUpdatePropertiesParams) WithDeviceID(deviceID string) *DevicesUpdatePropertiesParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the devices update properties params
func (o *DevicesUpdatePropertiesParams) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WriteToRequest writes these params to a swagger request
func (o *DevicesUpdatePropertiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param device_id
	if err := r.SetPathParam("device_id", o.DeviceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}