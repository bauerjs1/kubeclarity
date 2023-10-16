// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewVulnerabilitiesTrend new vulnerabilities trend
//
// swagger:model NewVulnerabilitiesTrend
type NewVulnerabilitiesTrend struct {

	// num of vuls
	NumOfVuls []*VulnerabilityCount `json:"numOfVuls"`

	// time
	// Format: date-time
	Time strfmt.DateTime `json:"time,omitempty"`
}

// Validate validates this new vulnerabilities trend
func (m *NewVulnerabilitiesTrend) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNumOfVuls(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewVulnerabilitiesTrend) validateNumOfVuls(formats strfmt.Registry) error {
	if swag.IsZero(m.NumOfVuls) { // not required
		return nil
	}

	for i := 0; i < len(m.NumOfVuls); i++ {
		if swag.IsZero(m.NumOfVuls[i]) { // not required
			continue
		}

		if m.NumOfVuls[i] != nil {
			if err := m.NumOfVuls[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("numOfVuls" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("numOfVuls" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NewVulnerabilitiesTrend) validateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := validate.FormatOf("time", "body", "date-time", m.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this new vulnerabilities trend based on the context it is used
func (m *NewVulnerabilitiesTrend) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNumOfVuls(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NewVulnerabilitiesTrend) contextValidateNumOfVuls(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NumOfVuls); i++ {

		if m.NumOfVuls[i] != nil {

			if swag.IsZero(m.NumOfVuls[i]) { // not required
				return nil
			}

			if err := m.NumOfVuls[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("numOfVuls" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("numOfVuls" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NewVulnerabilitiesTrend) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewVulnerabilitiesTrend) UnmarshalBinary(b []byte) error {
	var res NewVulnerabilitiesTrend
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
