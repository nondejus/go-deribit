// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// OriginalOrderType Original order type. Optional field
// swagger:model original_order_type
type OriginalOrderType string

const (

	// OriginalOrderTypeMarket captures enum value "market"
	OriginalOrderTypeMarket OriginalOrderType = "market"
)

// for schema
var originalOrderTypeEnum []interface{}

func init() {
	var res []OriginalOrderType
	if err := json.Unmarshal([]byte(`["market"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		originalOrderTypeEnum = append(originalOrderTypeEnum, v)
	}
}

func (m OriginalOrderType) validateOriginalOrderTypeEnum(path, location string, value OriginalOrderType) error {
	if err := validate.Enum(path, location, value, originalOrderTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this original order type
func (m OriginalOrderType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOriginalOrderTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
