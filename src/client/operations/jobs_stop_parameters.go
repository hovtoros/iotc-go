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
)

// NewJobsStopParams creates a new JobsStopParams object
// with the default values initialized.
func NewJobsStopParams() *JobsStopParams {
	var ()
	return &JobsStopParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJobsStopParamsWithTimeout creates a new JobsStopParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJobsStopParamsWithTimeout(timeout time.Duration) *JobsStopParams {
	var ()
	return &JobsStopParams{

		timeout: timeout,
	}
}

// NewJobsStopParamsWithContext creates a new JobsStopParams object
// with the default values initialized, and the ability to set a context for a request
func NewJobsStopParamsWithContext(ctx context.Context) *JobsStopParams {
	var ()
	return &JobsStopParams{

		Context: ctx,
	}
}

// NewJobsStopParamsWithHTTPClient creates a new JobsStopParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewJobsStopParamsWithHTTPClient(client *http.Client) *JobsStopParams {
	var ()
	return &JobsStopParams{
		HTTPClient: client,
	}
}

/*JobsStopParams contains all the parameters to send to the API endpoint
for the jobs stop operation typically these are written to a http.Request
*/
type JobsStopParams struct {

	/*JobID
	  Unique ID of the job.

	*/
	JobID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the jobs stop params
func (o *JobsStopParams) WithTimeout(timeout time.Duration) *JobsStopParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the jobs stop params
func (o *JobsStopParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the jobs stop params
func (o *JobsStopParams) WithContext(ctx context.Context) *JobsStopParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the jobs stop params
func (o *JobsStopParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the jobs stop params
func (o *JobsStopParams) WithHTTPClient(client *http.Client) *JobsStopParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the jobs stop params
func (o *JobsStopParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithJobID adds the jobID to the jobs stop params
func (o *JobsStopParams) WithJobID(jobID string) *JobsStopParams {
	o.SetJobID(jobID)
	return o
}

// SetJobID adds the jobId to the jobs stop params
func (o *JobsStopParams) SetJobID(jobID string) {
	o.JobID = jobID
}

// WriteToRequest writes these params to a swagger request
func (o *JobsStopParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param job_id
	if err := r.SetPathParam("job_id", o.JobID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}