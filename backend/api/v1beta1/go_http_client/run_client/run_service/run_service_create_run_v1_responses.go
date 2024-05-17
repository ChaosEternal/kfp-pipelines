// Code generated by go-swagger; DO NOT EDIT.

package run_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	run_model "github.com/kubeflow/pipelines/backend/api/v1beta1/go_http_client/run_model"
)

// RunServiceCreateRunV1Reader is a Reader for the RunServiceCreateRunV1 structure.
type RunServiceCreateRunV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RunServiceCreateRunV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRunServiceCreateRunV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewRunServiceCreateRunV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRunServiceCreateRunV1OK creates a RunServiceCreateRunV1OK with default headers values
func NewRunServiceCreateRunV1OK() *RunServiceCreateRunV1OK {
	return &RunServiceCreateRunV1OK{}
}

/*RunServiceCreateRunV1OK handles this case with default header values.

A successful response.
*/
type RunServiceCreateRunV1OK struct {
	Payload *run_model.APIRunDetail
}

func (o *RunServiceCreateRunV1OK) Error() string {
	return fmt.Sprintf("[POST /apis/v1beta1/runs][%d] runServiceCreateRunV1OK  %+v", 200, o.Payload)
}

func (o *RunServiceCreateRunV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(run_model.APIRunDetail)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRunServiceCreateRunV1Default creates a RunServiceCreateRunV1Default with default headers values
func NewRunServiceCreateRunV1Default(code int) *RunServiceCreateRunV1Default {
	return &RunServiceCreateRunV1Default{
		_statusCode: code,
	}
}

/*RunServiceCreateRunV1Default handles this case with default header values.

An unexpected error response.
*/
type RunServiceCreateRunV1Default struct {
	_statusCode int

	Payload *run_model.GatewayruntimeError
}

// Code gets the status code for the run service create run v1 default response
func (o *RunServiceCreateRunV1Default) Code() int {
	return o._statusCode
}

func (o *RunServiceCreateRunV1Default) Error() string {
	return fmt.Sprintf("[POST /apis/v1beta1/runs][%d] RunService_CreateRunV1 default  %+v", o._statusCode, o.Payload)
}

func (o *RunServiceCreateRunV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(run_model.GatewayruntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}