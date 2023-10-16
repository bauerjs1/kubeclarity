// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openclarity/kubeclarity/api/client/models"
)

// PutRuntimeScanStartReader is a Reader for the PutRuntimeScanStart structure.
type PutRuntimeScanStartReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutRuntimeScanStartReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPutRuntimeScanStartCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutRuntimeScanStartBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPutRuntimeScanStartDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutRuntimeScanStartCreated creates a PutRuntimeScanStartCreated with default headers values
func NewPutRuntimeScanStartCreated() *PutRuntimeScanStartCreated {
	return &PutRuntimeScanStartCreated{}
}

/*
PutRuntimeScanStartCreated describes a response with status code 201, with default header values.

Success
*/
type PutRuntimeScanStartCreated struct {
	Payload interface{}
}

// IsSuccess returns true when this put runtime scan start created response has a 2xx status code
func (o *PutRuntimeScanStartCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put runtime scan start created response has a 3xx status code
func (o *PutRuntimeScanStartCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put runtime scan start created response has a 4xx status code
func (o *PutRuntimeScanStartCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this put runtime scan start created response has a 5xx status code
func (o *PutRuntimeScanStartCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this put runtime scan start created response a status code equal to that given
func (o *PutRuntimeScanStartCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the put runtime scan start created response
func (o *PutRuntimeScanStartCreated) Code() int {
	return 201
}

func (o *PutRuntimeScanStartCreated) Error() string {
	return fmt.Sprintf("[PUT /runtime/scan/start][%d] putRuntimeScanStartCreated  %+v", 201, o.Payload)
}

func (o *PutRuntimeScanStartCreated) String() string {
	return fmt.Sprintf("[PUT /runtime/scan/start][%d] putRuntimeScanStartCreated  %+v", 201, o.Payload)
}

func (o *PutRuntimeScanStartCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *PutRuntimeScanStartCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRuntimeScanStartBadRequest creates a PutRuntimeScanStartBadRequest with default headers values
func NewPutRuntimeScanStartBadRequest() *PutRuntimeScanStartBadRequest {
	return &PutRuntimeScanStartBadRequest{}
}

/*
PutRuntimeScanStartBadRequest describes a response with status code 400, with default header values.

Scan failed to start
*/
type PutRuntimeScanStartBadRequest struct {
	Payload string
}

// IsSuccess returns true when this put runtime scan start bad request response has a 2xx status code
func (o *PutRuntimeScanStartBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put runtime scan start bad request response has a 3xx status code
func (o *PutRuntimeScanStartBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put runtime scan start bad request response has a 4xx status code
func (o *PutRuntimeScanStartBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put runtime scan start bad request response has a 5xx status code
func (o *PutRuntimeScanStartBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put runtime scan start bad request response a status code equal to that given
func (o *PutRuntimeScanStartBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put runtime scan start bad request response
func (o *PutRuntimeScanStartBadRequest) Code() int {
	return 400
}

func (o *PutRuntimeScanStartBadRequest) Error() string {
	return fmt.Sprintf("[PUT /runtime/scan/start][%d] putRuntimeScanStartBadRequest  %+v", 400, o.Payload)
}

func (o *PutRuntimeScanStartBadRequest) String() string {
	return fmt.Sprintf("[PUT /runtime/scan/start][%d] putRuntimeScanStartBadRequest  %+v", 400, o.Payload)
}

func (o *PutRuntimeScanStartBadRequest) GetPayload() string {
	return o.Payload
}

func (o *PutRuntimeScanStartBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRuntimeScanStartDefault creates a PutRuntimeScanStartDefault with default headers values
func NewPutRuntimeScanStartDefault(code int) *PutRuntimeScanStartDefault {
	return &PutRuntimeScanStartDefault{
		_statusCode: code,
	}
}

/*
PutRuntimeScanStartDefault describes a response with status code -1, with default header values.

unknown error
*/
type PutRuntimeScanStartDefault struct {
	_statusCode int

	Payload *models.APIResponse
}

// IsSuccess returns true when this put runtime scan start default response has a 2xx status code
func (o *PutRuntimeScanStartDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this put runtime scan start default response has a 3xx status code
func (o *PutRuntimeScanStartDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this put runtime scan start default response has a 4xx status code
func (o *PutRuntimeScanStartDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this put runtime scan start default response has a 5xx status code
func (o *PutRuntimeScanStartDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this put runtime scan start default response a status code equal to that given
func (o *PutRuntimeScanStartDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the put runtime scan start default response
func (o *PutRuntimeScanStartDefault) Code() int {
	return o._statusCode
}

func (o *PutRuntimeScanStartDefault) Error() string {
	return fmt.Sprintf("[PUT /runtime/scan/start][%d] PutRuntimeScanStart default  %+v", o._statusCode, o.Payload)
}

func (o *PutRuntimeScanStartDefault) String() string {
	return fmt.Sprintf("[PUT /runtime/scan/start][%d] PutRuntimeScanStart default  %+v", o._statusCode, o.Payload)
}

func (o *PutRuntimeScanStartDefault) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *PutRuntimeScanStartDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
