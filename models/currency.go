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

// Currency Currency, i.e `"BTC"`, `"ETH"`
// swagger:model currency
type Currency string

const (

	// CurrencyBTC captures enum value "BTC"
	CurrencyBTC Currency = "BTC"

	// CurrencyETH captures enum value "ETH"
	CurrencyETH Currency = "ETH"
)

// for schema
var currencyEnum []interface{}

func init() {
	var res []Currency
	if err := json.Unmarshal([]byte(`["BTC","ETH"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		currencyEnum = append(currencyEnum, v)
	}
}

func (m Currency) validateCurrencyEnum(path, location string, value Currency) error {
	if err := validate.Enum(path, location, value, currencyEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this currency
func (m Currency) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCurrencyEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
