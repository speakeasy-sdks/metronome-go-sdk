// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/metronome-go-sdk/internal/utils"
	"github.com/speakeasy-sdks/metronome-go-sdk/models/components"
	"time"
)

// GetRateCardNamedScheduleRequestBody - Which rate card, schedule name, and date to retrieve.
type GetRateCardNamedScheduleRequestBody struct {
	// ID of the rate card whose named schedule is to be retrieved
	RateCardID string `json:"rate_card_id"`
	// The identifier for the schedule to be retrieved
	ScheduleName string `json:"schedule_name"`
	// If provided, at most one schedule segment will be returned (the one that covers this date). If not provided, all segments will be returned.
	CoveringDate *time.Time `json:"covering_date,omitempty"`
}

func (g GetRateCardNamedScheduleRequestBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetRateCardNamedScheduleRequestBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetRateCardNamedScheduleRequestBody) GetRateCardID() string {
	if o == nil {
		return ""
	}
	return o.RateCardID
}

func (o *GetRateCardNamedScheduleRequestBody) GetScheduleName() string {
	if o == nil {
		return ""
	}
	return o.ScheduleName
}

func (o *GetRateCardNamedScheduleRequestBody) GetCoveringDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.CoveringDate
}

type GetRateCardNamedScheduleData struct {
	Value        any        `json:"value"`
	StartingAt   time.Time  `json:"starting_at"`
	EndingBefore *time.Time `json:"ending_before,omitempty"`
}

func (g GetRateCardNamedScheduleData) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetRateCardNamedScheduleData) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetRateCardNamedScheduleData) GetValue() any {
	if o == nil {
		return nil
	}
	return o.Value
}

func (o *GetRateCardNamedScheduleData) GetStartingAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartingAt
}

func (o *GetRateCardNamedScheduleData) GetEndingBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndingBefore
}

// GetRateCardNamedScheduleResponseBody - Success
type GetRateCardNamedScheduleResponseBody struct {
	Data []GetRateCardNamedScheduleData `json:"data"`
}

func (o *GetRateCardNamedScheduleResponseBody) GetData() []GetRateCardNamedScheduleData {
	if o == nil {
		return []GetRateCardNamedScheduleData{}
	}
	return o.Data
}

type GetRateCardNamedScheduleResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	Object *GetRateCardNamedScheduleResponseBody
}

func (o *GetRateCardNamedScheduleResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetRateCardNamedScheduleResponse) GetObject() *GetRateCardNamedScheduleResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
