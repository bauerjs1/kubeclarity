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

// GetRuntimeScanProgressReader is a Reader for the GetRuntimeScanProgress structure.
type GetRuntimeScanProgressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRuntimeScanProgressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRuntimeScanProgressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetRuntimeScanProgressDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRuntimeScanProgressOK creates a GetRuntimeScanProgressOK with default headers values
func NewGetRuntimeScanProgressOK() *GetRuntimeScanProgressOK {
	return &GetRuntimeScanProgressOK{}
}

/*
GetRuntimeScanProgressOK describes a response with status code 200, with default header values.

Success
*/
type GetRuntimeScanProgressOK struct {
	Payload *models.Progress
}

// IsSuccess returns true when this get runtime scan progress o k response has a 2xx status code
func (o *GetRuntimeScanProgressOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get runtime scan progress o k response has a 3xx status code
func (o *GetRuntimeScanProgressOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get runtime scan progress o k response has a 4xx status code
func (o *GetRuntimeScanProgressOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get runtime scan progress o k response has a 5xx status code
func (o *GetRuntimeScanProgressOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get runtime scan progress o k response a status code equal to that given
func (o *GetRuntimeScanProgressOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get runtime scan progress o k response
func (o *GetRuntimeScanProgressOK) Code() int {
	return 200
}

func (o *GetRuntimeScanProgressOK) Error() string {
	return fmt.Sprintf("[GET /runtime/scan/progress][%d] getRuntimeScanProgressOK  %+v", 200, o.Payload)
}

func (o *GetRuntimeScanProgressOK) String() string {
	return fmt.Sprintf("[GET /runtime/scan/progress][%d] getRuntimeScanProgressOK  %+v", 200, o.Payload)
}

func (o *GetRuntimeScanProgressOK) GetPayload() *models.Progress {
	return o.Payload
}

func (o *GetRuntimeScanProgressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Progress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRuntimeScanProgressDefault creates a GetRuntimeScanProgressDefault with default headers values
func NewGetRuntimeScanProgressDefault(code int) *GetRuntimeScanProgressDefault {
	return &GetRuntimeScanProgressDefault{
		_statusCode: code,
	}
}

/*
GetRuntimeScanProgressDefault describes a response with status code -1, with default header values.

unknown error
*/
type GetRuntimeScanProgressDefault struct {
	_statusCode int

	Payload *models.APIResponse
}

// IsSuccess returns true when this get runtime scan progress default response has a 2xx status code
func (o *GetRuntimeScanProgressDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get runtime scan progress default response has a 3xx status code
func (o *GetRuntimeScanProgressDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get runtime scan progress default response has a 4xx status code
func (o *GetRuntimeScanProgressDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get runtime scan progress default response has a 5xx status code
func (o *GetRuntimeScanProgressDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get runtime scan progress default response a status code equal to that given
func (o *GetRuntimeScanProgressDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get runtime scan progress default response
func (o *GetRuntimeScanProgressDefault) Code() int {
	return o._statusCode
}

func (o *GetRuntimeScanProgressDefault) Error() string {
	return fmt.Sprintf("[GET /runtime/scan/progress][%d] GetRuntimeScanProgress default  %+v", o._statusCode, o.Payload)
}

func (o *GetRuntimeScanProgressDefault) String() string {
	return fmt.Sprintf("[GET /runtime/scan/progress][%d] GetRuntimeScanProgress default  %+v", o._statusCode, o.Payload)
}

func (o *GetRuntimeScanProgressDefault) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *GetRuntimeScanProgressDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
