// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package daemonset

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

// NewPostIpamIpsParams creates a new PostIpamIpsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostIpamIpsParams() *PostIpamIpsParams {
	return &PostIpamIpsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostIpamIpsParamsWithTimeout creates a new PostIpamIpsParams object
// with the ability to set a timeout on a request.
func NewPostIpamIpsParamsWithTimeout(timeout time.Duration) *PostIpamIpsParams {
	return &PostIpamIpsParams{
		timeout: timeout,
	}
}

// NewPostIpamIpsParamsWithContext creates a new PostIpamIpsParams object
// with the ability to set a context for a request.
func NewPostIpamIpsParamsWithContext(ctx context.Context) *PostIpamIpsParams {
	return &PostIpamIpsParams{
		Context: ctx,
	}
}

// NewPostIpamIpsParamsWithHTTPClient creates a new PostIpamIpsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostIpamIpsParamsWithHTTPClient(client *http.Client) *PostIpamIpsParams {
	return &PostIpamIpsParams{
		HTTPClient: client,
	}
}

/*
PostIpamIpsParams contains all the parameters to send to the API endpoint

	for the post ipam ips operation.

	Typically these are written to a http.Request.
*/
type PostIpamIpsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post ipam ips params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostIpamIpsParams) WithDefaults() *PostIpamIpsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post ipam ips params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostIpamIpsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post ipam ips params
func (o *PostIpamIpsParams) WithTimeout(timeout time.Duration) *PostIpamIpsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post ipam ips params
func (o *PostIpamIpsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post ipam ips params
func (o *PostIpamIpsParams) WithContext(ctx context.Context) *PostIpamIpsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post ipam ips params
func (o *PostIpamIpsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post ipam ips params
func (o *PostIpamIpsParams) WithHTTPClient(client *http.Client) *PostIpamIpsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post ipam ips params
func (o *PostIpamIpsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PostIpamIpsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}