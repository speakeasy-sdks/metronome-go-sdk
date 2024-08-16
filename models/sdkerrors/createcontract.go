// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"github.com/Metronome-Industries/metronome-go-sdk/models/components"
)

// CreateContractContractsResponseBody - The specified resource was not found
type CreateContractContractsResponseBody struct {
	Message  string                  `json:"message"`
	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &CreateContractContractsResponseBody{}

func (e *CreateContractContractsResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// CreateContractResponseBody - Bad request
type CreateContractResponseBody struct {
	Message  string                  `json:"message"`
	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &CreateContractResponseBody{}

func (e *CreateContractResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
