// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Metronome-Industries/metronome-go-sdk/internal/utils"
	"github.com/Metronome-Industries/metronome-go-sdk/models/components"
	"time"
)

// UpdateCommitEndDateRequestBody - Update the access or invoice end date of a commit
type UpdateCommitEndDateRequestBody struct {
	// ID of the customer whose commit is to be updated
	CustomerID string `json:"customer_id"`
	// ID of the commit to update. Only supports "PREPAID" commits.
	CommitID string `json:"commit_id"`
	// RFC 3339 timestamp indicating when access to the commit will end and it will no longer be possible to draw it down (exclusive). If not provided, the access will not be updated.
	AccessEndingBefore *time.Time `json:"access_ending_before,omitempty"`
	// RFC 3339 timestamp indicating when the commit will stop being invoiced (exclusive). If not provided, the invoice schedule will not be updated.
	InvoicesEndingBefore *time.Time `json:"invoices_ending_before,omitempty"`
}

func (u UpdateCommitEndDateRequestBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpdateCommitEndDateRequestBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpdateCommitEndDateRequestBody) GetCustomerID() string {
	if o == nil {
		return ""
	}
	return o.CustomerID
}

func (o *UpdateCommitEndDateRequestBody) GetCommitID() string {
	if o == nil {
		return ""
	}
	return o.CommitID
}

func (o *UpdateCommitEndDateRequestBody) GetAccessEndingBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.AccessEndingBefore
}

func (o *UpdateCommitEndDateRequestBody) GetInvoicesEndingBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.InvoicesEndingBefore
}

type UpdateCommitEndDateData struct {
	ID string `json:"id"`
}

func (o *UpdateCommitEndDateData) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// UpdateCommitEndDateResponseBody - Success
type UpdateCommitEndDateResponseBody struct {
	Data UpdateCommitEndDateData `json:"data"`
}

func (o *UpdateCommitEndDateResponseBody) GetData() UpdateCommitEndDateData {
	if o == nil {
		return UpdateCommitEndDateData{}
	}
	return o.Data
}

type UpdateCommitEndDateResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	Object *UpdateCommitEndDateResponseBody
}

func (o *UpdateCommitEndDateResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateCommitEndDateResponse) GetObject() *UpdateCommitEndDateResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
