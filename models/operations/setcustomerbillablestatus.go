// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/metronome-go-sdk/internal/utils"
	"github.com/speakeasy-sdks/metronome-go-sdk/models/components"
	"time"
)

type BillableStatus string

const (
	BillableStatusBillable   BillableStatus = "billable"
	BillableStatusUnbillable BillableStatus = "unbillable"
)

func (e BillableStatus) ToPointer() *BillableStatus {
	return &e
}
func (e *BillableStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "billable":
		fallthrough
	case "unbillable":
		*e = BillableStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for BillableStatus: %v", v)
	}
}

type SetCustomerBillableStatusRequestBody struct {
	CustomerID     string         `json:"customer_id"`
	BillableStatus BillableStatus `json:"billable_status"`
	EffectiveAt    time.Time      `json:"effective_at"`
}

func (s SetCustomerBillableStatusRequestBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SetCustomerBillableStatusRequestBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SetCustomerBillableStatusRequestBody) GetCustomerID() string {
	if o == nil {
		return ""
	}
	return o.CustomerID
}

func (o *SetCustomerBillableStatusRequestBody) GetBillableStatus() BillableStatus {
	if o == nil {
		return BillableStatus("")
	}
	return o.BillableStatus
}

func (o *SetCustomerBillableStatusRequestBody) GetEffectiveAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.EffectiveAt
}

type CurrentBillableStatus string

const (
	CurrentBillableStatusBillable   CurrentBillableStatus = "billable"
	CurrentBillableStatusUnbillable CurrentBillableStatus = "unbillable"
)

func (e CurrentBillableStatus) ToPointer() *CurrentBillableStatus {
	return &e
}
func (e *CurrentBillableStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "billable":
		fallthrough
	case "unbillable":
		*e = CurrentBillableStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CurrentBillableStatus: %v", v)
	}
}

type Data struct {
	ID                    string                `json:"id"`
	CurrentBillableStatus CurrentBillableStatus `json:"current_billable_status"`
}

func (o *Data) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *Data) GetCurrentBillableStatus() CurrentBillableStatus {
	if o == nil {
		return CurrentBillableStatus("")
	}
	return o.CurrentBillableStatus
}

// SetCustomerBillableStatusResponseBody - Success
type SetCustomerBillableStatusResponseBody struct {
	Data Data `json:"data"`
}

func (o *SetCustomerBillableStatusResponseBody) GetData() Data {
	if o == nil {
		return Data{}
	}
	return o.Data
}

type SetCustomerBillableStatusResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	Object *SetCustomerBillableStatusResponseBody
}

func (o *SetCustomerBillableStatusResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *SetCustomerBillableStatusResponse) GetObject() *SetCustomerBillableStatusResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
