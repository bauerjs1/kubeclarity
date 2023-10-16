// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Progress progress
//
// swagger:model Progress
type Progress struct {

	// scan type
	ScanType ScanType `json:"scanType,omitempty"`

	// Percentage of scanned images from total images (0-100)
	// Minimum: 0
	Scanned *int64 `json:"scanned,omitempty"`

	// scanned namespaces
	ScannedNamespaces []string `json:"scannedNamespaces"`

	// start time
	// Format: date-time
	StartTime strfmt.DateTime `json:"startTime,omitempty"`

	// status
	Status RuntimeScanStatus `json:"status,omitempty"`
}

// Validate validates this progress
func (m *Progress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScanType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScanned(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Progress) validateScanType(formats strfmt.Registry) error {
	if swag.IsZero(m.ScanType) { // not required
		return nil
	}

	if err := m.ScanType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("scanType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("scanType")
		}
		return err
	}

	return nil
}

func (m *Progress) validateScanned(formats strfmt.Registry) error {
	if swag.IsZero(m.Scanned) { // not required
		return nil
	}

	if err := validate.MinimumInt("scanned", "body", *m.Scanned, 0, false); err != nil {
		return err
	}

	return nil
}

func (m *Progress) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startTime", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Progress) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

// ContextValidate validate this progress based on the context it is used
func (m *Progress) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateScanType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Progress) contextValidateScanType(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.ScanType) { // not required
		return nil
	}

	if err := m.ScanType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("scanType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("scanType")
		}
		return err
	}

	return nil
}

func (m *Progress) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Progress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Progress) UnmarshalBinary(b []byte) error {
	var res Progress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
