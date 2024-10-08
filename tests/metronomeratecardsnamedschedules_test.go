// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package tests

import (
	"context"
	metronomegosdk "github.com/speakeasy-sdks/metronome-go-sdk"
	"github.com/speakeasy-sdks/metronome-go-sdk/models/operations"
	"github.com/speakeasy-sdks/metronome-go-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMetronomeRateCardsNamedSchedules_GetRateCardNamedSchedule_(t *testing.T) {
	s := metronomegosdk.New(
		metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)
	var request *operations.GetRateCardNamedScheduleRequestBody = &operations.GetRateCardNamedScheduleRequestBody{
		RateCardID:   "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		ScheduleName: "my-schedule",
		CoveringDate: types.MustNewTimeFromString("2022-02-15T00:00:00Z"),
	}
	ctx := context.Background()
	res, err := s.RateCards.NamedSchedules.Get(ctx, request)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.Equal(t, operations.GetRateCardNamedScheduleResponseBody{
		Data: []operations.GetRateCardNamedScheduleData{
			operations.GetRateCardNamedScheduleData{
				Value:        "my-value",
				StartingAt:   types.MustTimeFromString("2022-02-01T00:00:00Z"),
				EndingBefore: types.MustNewTimeFromString("2022-03-01T00:00:00Z"),
			},
		},
	}, *res.Object)
}

func TestMetronomeRateCardsNamedSchedules_UpdateRateCardNamedSchedule_(t *testing.T) {
	assert.Fail(t, "incomplete test found please make sure to address the following errors: [`missing response values`]")
}
