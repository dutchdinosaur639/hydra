// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RejectRequest The request payload used to accept a login or consent request.
//
// swagger:model rejectRequest
type RejectRequest struct {

	// The error should follow the OAuth2 error format (e.g. `invalid_request`, `login_required`).
	//
	// Defaults to `request_denied`.
	Error string `json:"error,omitempty"`

	// Debug contains information to help resolve the problem as a developer. Usually not exposed
	// to the public but only in the server logs.
	ErrorDebug string `json:"error_debug,omitempty"`

	// Description of the error in a human readable format.
	ErrorDescription string `json:"error_description,omitempty"`

	// Hint to help resolve the error.
	ErrorHint string `json:"error_hint,omitempty"`

	// Represents the HTTP status code of the error (e.g. 401 or 403)
	//
	// Defaults to 400
	StatusCode int64 `json:"status_code,omitempty"`
}

// Validate validates this reject request
func (m *RejectRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this reject request based on context it is used
func (m *RejectRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RejectRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RejectRequest) UnmarshalBinary(b []byte) error {
	var res RejectRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
