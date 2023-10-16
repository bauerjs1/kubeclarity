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

// GetDashboardTrendsVulnerabilitiesReader is a Reader for the GetDashboardTrendsVulnerabilities structure.
type GetDashboardTrendsVulnerabilitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDashboardTrendsVulnerabilitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDashboardTrendsVulnerabilitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetDashboardTrendsVulnerabilitiesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDashboardTrendsVulnerabilitiesOK creates a GetDashboardTrendsVulnerabilitiesOK with default headers values
func NewGetDashboardTrendsVulnerabilitiesOK() *GetDashboardTrendsVulnerabilitiesOK {
	return &GetDashboardTrendsVulnerabilitiesOK{}
}

/*
GetDashboardTrendsVulnerabilitiesOK describes a response with status code 200, with default header values.

Success
*/
type GetDashboardTrendsVulnerabilitiesOK struct {
	Payload []*models.NewVulnerabilitiesTrend
}

// IsSuccess returns true when this get dashboard trends vulnerabilities o k response has a 2xx status code
func (o *GetDashboardTrendsVulnerabilitiesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get dashboard trends vulnerabilities o k response has a 3xx status code
func (o *GetDashboardTrendsVulnerabilitiesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dashboard trends vulnerabilities o k response has a 4xx status code
func (o *GetDashboardTrendsVulnerabilitiesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get dashboard trends vulnerabilities o k response has a 5xx status code
func (o *GetDashboardTrendsVulnerabilitiesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get dashboard trends vulnerabilities o k response a status code equal to that given
func (o *GetDashboardTrendsVulnerabilitiesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get dashboard trends vulnerabilities o k response
func (o *GetDashboardTrendsVulnerabilitiesOK) Code() int {
	return 200
}

func (o *GetDashboardTrendsVulnerabilitiesOK) Error() string {
	return fmt.Sprintf("[GET /dashboard/trends/vulnerabilities][%d] getDashboardTrendsVulnerabilitiesOK  %+v", 200, o.Payload)
}

func (o *GetDashboardTrendsVulnerabilitiesOK) String() string {
	return fmt.Sprintf("[GET /dashboard/trends/vulnerabilities][%d] getDashboardTrendsVulnerabilitiesOK  %+v", 200, o.Payload)
}

func (o *GetDashboardTrendsVulnerabilitiesOK) GetPayload() []*models.NewVulnerabilitiesTrend {
	return o.Payload
}

func (o *GetDashboardTrendsVulnerabilitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDashboardTrendsVulnerabilitiesDefault creates a GetDashboardTrendsVulnerabilitiesDefault with default headers values
func NewGetDashboardTrendsVulnerabilitiesDefault(code int) *GetDashboardTrendsVulnerabilitiesDefault {
	return &GetDashboardTrendsVulnerabilitiesDefault{
		_statusCode: code,
	}
}

/*
GetDashboardTrendsVulnerabilitiesDefault describes a response with status code -1, with default header values.

unknown error
*/
type GetDashboardTrendsVulnerabilitiesDefault struct {
	_statusCode int

	Payload *models.APIResponse
}

// IsSuccess returns true when this get dashboard trends vulnerabilities default response has a 2xx status code
func (o *GetDashboardTrendsVulnerabilitiesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get dashboard trends vulnerabilities default response has a 3xx status code
func (o *GetDashboardTrendsVulnerabilitiesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get dashboard trends vulnerabilities default response has a 4xx status code
func (o *GetDashboardTrendsVulnerabilitiesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get dashboard trends vulnerabilities default response has a 5xx status code
func (o *GetDashboardTrendsVulnerabilitiesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get dashboard trends vulnerabilities default response a status code equal to that given
func (o *GetDashboardTrendsVulnerabilitiesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get dashboard trends vulnerabilities default response
func (o *GetDashboardTrendsVulnerabilitiesDefault) Code() int {
	return o._statusCode
}

func (o *GetDashboardTrendsVulnerabilitiesDefault) Error() string {
	return fmt.Sprintf("[GET /dashboard/trends/vulnerabilities][%d] GetDashboardTrendsVulnerabilities default  %+v", o._statusCode, o.Payload)
}

func (o *GetDashboardTrendsVulnerabilitiesDefault) String() string {
	return fmt.Sprintf("[GET /dashboard/trends/vulnerabilities][%d] GetDashboardTrendsVulnerabilities default  %+v", o._statusCode, o.Payload)
}

func (o *GetDashboardTrendsVulnerabilitiesDefault) GetPayload() *models.APIResponse {
	return o.Payload
}

func (o *GetDashboardTrendsVulnerabilitiesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
