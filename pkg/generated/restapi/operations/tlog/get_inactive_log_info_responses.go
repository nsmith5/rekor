// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright 2021 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package tlog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/sigstore/rekor/pkg/generated/models"
)

// GetInactiveLogInfoOKCode is the HTTP code returned for type GetInactiveLogInfoOK
const GetInactiveLogInfoOKCode int = 200

/*GetInactiveLogInfoOK A JSON object with the root hash and tree size as properties

swagger:response getInactiveLogInfoOK
*/
type GetInactiveLogInfoOK struct {

	/*
	  In: Body
	*/
	Payload []*models.LogInfo `json:"body,omitempty"`
}

// NewGetInactiveLogInfoOK creates GetInactiveLogInfoOK with default headers values
func NewGetInactiveLogInfoOK() *GetInactiveLogInfoOK {

	return &GetInactiveLogInfoOK{}
}

// WithPayload adds the payload to the get inactive log info o k response
func (o *GetInactiveLogInfoOK) WithPayload(payload []*models.LogInfo) *GetInactiveLogInfoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get inactive log info o k response
func (o *GetInactiveLogInfoOK) SetPayload(payload []*models.LogInfo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInactiveLogInfoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.LogInfo, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetInactiveLogInfoDefault There was an internal error in the server while processing the request

swagger:response getInactiveLogInfoDefault
*/
type GetInactiveLogInfoDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetInactiveLogInfoDefault creates GetInactiveLogInfoDefault with default headers values
func NewGetInactiveLogInfoDefault(code int) *GetInactiveLogInfoDefault {
	if code <= 0 {
		code = 500
	}

	return &GetInactiveLogInfoDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get inactive log info default response
func (o *GetInactiveLogInfoDefault) WithStatusCode(code int) *GetInactiveLogInfoDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get inactive log info default response
func (o *GetInactiveLogInfoDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get inactive log info default response
func (o *GetInactiveLogInfoDefault) WithPayload(payload *models.Error) *GetInactiveLogInfoDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get inactive log info default response
func (o *GetInactiveLogInfoDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetInactiveLogInfoDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
