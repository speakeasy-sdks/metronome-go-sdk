// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"github.com/Metronome-Industries/metronome-go-sdk/models/components"
)

// CreateCustomerCreditContractsResponseBody - The specified resource was not found
type CreateCustomerCreditContractsResponseBody struct {
	Message  string                  `json:"message"`
	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &CreateCustomerCreditContractsResponseBody{}

func (e *CreateCustomerCreditContractsResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// CreateCustomerCreditResponseBody - Bad request
type CreateCustomerCreditResponseBody struct {
	Message  string                  `json:"message"`
	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &CreateCustomerCreditResponseBody{}

func (e *CreateCustomerCreditResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
