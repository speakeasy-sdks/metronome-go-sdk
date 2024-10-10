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

func TestRateCards_GetRateSchedule_(t *testing.T) {
	assert.Fail(t, "incomplete test found please make sure to address the following errors: [`missing response values`]")
}

func TestRateCards_GetRates_(t *testing.T) {
	assert.Fail(t, "incomplete test found please make sure to address the following errors: [`missing response values`]")
}

func TestRateCards_GetRateCard_(t *testing.T) {
	s := metronomegosdk.New(
		metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)
	var request *operations.GetRateCardRequestBody = &operations.GetRateCardRequestBody{
		ID: "f3d51ae8-f283-44e1-9933-a3cf9ad7a6fe",
	}
	ctx := context.Background()
	res, err := s.RateCards.Get(ctx, request)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.Equal(t, operations.GetRateCardResponseBody{
		Data: operations.GetRateCardData{
			ID:              "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
			Name:            "Test rate card",
			RateCardEntries: map[string]operations.RateCardEntries{},
			CreatedAt:       types.MustTimeFromString("2019-12-30T04:24:55.123Z"),
			CreatedBy:       "Bob",
			Description:     metronomegosdk.String("Test rate card description"),
			FiatCreditType: &operations.FiatCreditType{
				Name: "USD (cents)",
				ID:   "2714e483-4ff1-48e4-9e25-ac732e8f24f2",
			},
			Aliases: []operations.GetRateCardAliases{
				operations.GetRateCardAliases{
					Name: "test-rate-card",
				},
			},
			CustomFields: map[string]string{
				"x_account_id": "KyVnHhSBWl7eY2bl",
			},
		},
	}, *res.Object)
}

func TestRateCards_ListRateCards_(t *testing.T) {
	s := metronomegosdk.New(
		metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)
	var requestBody *operations.ListRateCardsRequestBody = &operations.ListRateCardsRequestBody{}
	ctx := context.Background()
	res, err := s.RateCards.List(ctx, nil, nil, requestBody)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.Equal(t, operations.ListRateCardsResponseBody{
		Data: []operations.ListRateCardsData{
			operations.ListRateCardsData{
				ID:              "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
				Name:            "Test rate card",
				RateCardEntries: map[string]operations.ListRateCardsRateCardEntries{},
				CreatedAt:       types.MustTimeFromString("2019-12-30T04:24:55.123Z"),
				CreatedBy:       "Bob",
				Description:     metronomegosdk.String("Test rate card description"),
				FiatCreditType: &operations.ListRateCardsFiatCreditType{
					Name: "USD (cents)",
					ID:   "2714e483-4ff1-48e4-9e25-ac732e8f24f2",
				},
				Aliases: []operations.ListRateCardsAliases{
					operations.ListRateCardsAliases{
						Name: "<value>",
					},
				},
				CustomFields: map[string]string{
					"x_account_id": "KyVnHhSBWl7eY2bl",
				},
			},
		},
		NextPage: nil,
	}, *res.Object)
}

func TestRateCards_CreateRateCard_(t *testing.T) {
	s := metronomegosdk.New(
		metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)
	var request *operations.CreateRateCardRequestBody = &operations.CreateRateCardRequestBody{
		Name:             "My Rate Card",
		Description:      metronomegosdk.String("My Rate Card Description"),
		FiatCreditTypeID: metronomegosdk.String("2714e483-4ff1-48e4-9e25-ac732e8f24f2"),
		CreditTypeConversions: []operations.CreditTypeConversions{
			operations.CreditTypeConversions{
				CustomCreditTypeID:  "2714e483-4ff1-48e4-9e25-ac732e8f24f2",
				FiatPerCustomCredit: 2,
			},
		},
		Aliases: []operations.Aliases{
			operations.Aliases{
				Name: "my-rate-card",
			},
		},
	}
	ctx := context.Background()
	res, err := s.RateCards.Create(ctx, request)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.Equal(t, operations.CreateRateCardResponseBody{
		Data: operations.CreateRateCardData{
			ID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		},
	}, *res.Object)
}

func TestRateCards_UpdateRateCard_(t *testing.T) {
	s := metronomegosdk.New(
		metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)
	var request *operations.UpdateRateCardRequestBody = &operations.UpdateRateCardRequestBody{
		RateCardID:  "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		Name:        metronomegosdk.String("My Updated Rate Card"),
		Description: metronomegosdk.String("My Updated Rate Card Description"),
	}
	ctx := context.Background()
	res, err := s.RateCards.Update(ctx, request)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.Equal(t, operations.UpdateRateCardResponseBody{
		Data: operations.UpdateRateCardData{
			ID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		},
	}, *res.Object)
}

func TestRateCards_AddRate_(t *testing.T) {
	s := metronomegosdk.New(
		metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)
	var request *operations.AddRateRequestBody = &operations.AddRateRequestBody{
		RateCardID:   "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		ProductID:    "13117714-3f05-48e5-a6e9-a66093f13b4d",
		StartingAt:   types.MustTimeFromString("2020-01-01T00:00:00.000Z"),
		Entitled:     true,
		RateType:     operations.RateTypeFlatUpper,
		Price:        metronomegosdk.Float64(100),
		CreditTypeID: metronomegosdk.String("2714e483-4ff1-48e4-9e25-ac732e8f24f2"),
	}
	ctx := context.Background()
	res, err := s.RateCards.AddRate(ctx, request)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.Equal(t, operations.AddRateResponseBody{
		Data: operations.AddRateData{
			RateType: operations.AddRateRateCardsRateTypeFlatUpper,
			Price:    metronomegosdk.Float64(100),
		},
	}, *res.Object)
}

func TestRateCards_AddRates_(t *testing.T) {
	s := metronomegosdk.New(
		metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)
	var request *operations.AddRatesRequestBody = &operations.AddRatesRequestBody{
		RateCardID: metronomegosdk.String("d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc"),
		Rates: []operations.Rates{
			operations.Rates{
				ProductID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
				PricingGroupValues: map[string]string{
					"region": "us-west-2",
					"cloud":  "aws",
				},
				StartingAt:   types.MustTimeFromString("2020-01-01T00:00:00.000Z"),
				Entitled:     true,
				RateType:     operations.AddRatesRateTypeFlatUpper,
				Price:        metronomegosdk.Float64(100),
				CreditTypeID: metronomegosdk.String("2714e483-4ff1-48e4-9e25-ac732e8f24f2"),
			},
			operations.Rates{
				ProductID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
				PricingGroupValues: map[string]string{
					"region": "us-east-2",
					"cloud":  "aws",
				},
				StartingAt:   types.MustTimeFromString("2020-01-01T00:00:00.000Z"),
				Entitled:     true,
				RateType:     operations.AddRatesRateTypeFlatUpper,
				Price:        metronomegosdk.Float64(120),
				CreditTypeID: metronomegosdk.String("2714e483-4ff1-48e4-9e25-ac732e8f24f2"),
			},
		},
	}
	ctx := context.Background()
	res, err := s.RateCards.AddRates(ctx, request)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.Equal(t, operations.AddRatesResponseBody{
		Data: operations.AddRatesData{
			ID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		},
	}, *res.Object)
}

func TestRateCards_SetRateCardProductsOrder_(t *testing.T) {
	s := metronomegosdk.New(
		metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)
	var request *operations.SetRateCardProductsOrderRequestBody = &operations.SetRateCardProductsOrderRequestBody{
		RateCardID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		ProductOrder: []string{
			"13117714-3f05-48e5-a6e9-a66093f13b4d",
			"b086f2f4-9851-4466-9ca0-30d53e6a42ac",
		},
	}
	ctx := context.Background()
	res, err := s.RateCards.SetProductsOrder(ctx, request)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.Equal(t, operations.SetRateCardProductsOrderResponseBody{
		Data: operations.SetRateCardProductsOrderData{
			ID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		},
	}, *res.Object)
}

func TestRateCards_MoveRateCardProducts_(t *testing.T) {
	s := metronomegosdk.New(
		metronomegosdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)
	var request *operations.MoveRateCardProductsRequestBody = &operations.MoveRateCardProductsRequestBody{
		RateCardID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		ProductMoves: []operations.ProductMoves{
			operations.ProductMoves{
				ProductID: "13117714-3f05-48e5-a6e9-a66093f13b4d",
				Position:  0,
			},
			operations.ProductMoves{
				ProductID: "b086f2f4-9851-4466-9ca0-30d53e6a42ac",
				Position:  1,
			},
		},
	}
	ctx := context.Background()
	res, err := s.RateCards.MoveProducts(ctx, request)
	require.NoError(t, err)
	assert.Equal(t, 200, res.HTTPMeta.Response.StatusCode)
	assert.Equal(t, operations.MoveRateCardProductsResponseBody{
		Data: operations.MoveRateCardProductsData{
			ID: "d7abd0cd-4ae9-4db7-8676-e986a4ebd8dc",
		},
	}, *res.Object)
}
