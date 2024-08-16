// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"github.com/Metronome-Industries/metronome-go-sdk/models/components"
)

// UpdateCommitEndDateContractsResponseBody - The specified resource was not found
type UpdateCommitEndDateContractsResponseBody struct {
	Message  string                  `json:"message"`
	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &UpdateCommitEndDateContractsResponseBody{}

func (e *UpdateCommitEndDateContractsResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

// UpdateCommitEndDateResponseBody - Bad request
type UpdateCommitEndDateResponseBody struct {
	Message  string                  `json:"message"`
	HTTPMeta components.HTTPMetadata `json:"-"`
}

var _ error = &UpdateCommitEndDateResponseBody{}

func (e *UpdateCommitEndDateResponseBody) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
