// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/metronome-go-sdk/internal/utils"
	"github.com/speakeasy-sdks/metronome-go-sdk/models/components"
	"time"
)

type ListRateCardsRequestBody struct {
}

type ListRateCardsRequest struct {
	// Max number of results that should be returned
	Limit *int64 `queryParam:"style=form,explode=true,name=limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage    *string                   `queryParam:"style=form,explode=true,name=next_page"`
	RequestBody *ListRateCardsRequestBody `request:"mediaType=application/json"`
}

func (o *ListRateCardsRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListRateCardsRequest) GetNextPage() *string {
	if o == nil {
		return nil
	}
	return o.NextPage
}

func (o *ListRateCardsRequest) GetRequestBody() *ListRateCardsRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type ListRateCardsRateType string

const (
	ListRateCardsRateTypeFlat         ListRateCardsRateType = "FLAT"
	ListRateCardsRateTypePercentage   ListRateCardsRateType = "PERCENTAGE"
	ListRateCardsRateTypeSubscription ListRateCardsRateType = "SUBSCRIPTION"
	ListRateCardsRateTypeCustom       ListRateCardsRateType = "CUSTOM"
	ListRateCardsRateTypeTiered       ListRateCardsRateType = "TIERED"
)

func (e ListRateCardsRateType) ToPointer() *ListRateCardsRateType {
	return &e
}
func (e *ListRateCardsRateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "FLAT":
		fallthrough
	case "PERCENTAGE":
		fallthrough
	case "SUBSCRIPTION":
		fallthrough
	case "CUSTOM":
		fallthrough
	case "TIERED":
		*e = ListRateCardsRateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListRateCardsRateType: %v", v)
	}
}

type ListRateCardsCreditType struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func (o *ListRateCardsCreditType) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ListRateCardsCreditType) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type ListRateCardsTiers struct {
	Size  *float64 `json:"size,omitempty"`
	Price float64  `json:"price"`
}

func (o *ListRateCardsTiers) GetSize() *float64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *ListRateCardsTiers) GetPrice() float64 {
	if o == nil {
		return 0.0
	}
	return o.Price
}

type ListRateCardsCurrent struct {
	ID           *string                  `json:"id,omitempty"`
	ProductID    *string                  `json:"product_id,omitempty"`
	RateType     *ListRateCardsRateType   `json:"rate_type,omitempty"`
	Price        *float64                 `json:"price,omitempty"`
	CreditType   *ListRateCardsCreditType `json:"credit_type,omitempty"`
	CustomRate   map[string]any           `json:"custom_rate,omitempty"`
	StartingAt   *time.Time               `json:"starting_at,omitempty"`
	Entitled     *bool                    `json:"entitled,omitempty"`
	CreatedAt    *time.Time               `json:"created_at,omitempty"`
	CreatedBy    *string                  `json:"created_by,omitempty"`
	EndingBefore *time.Time               `json:"ending_before,omitempty"`
	Tiers        []ListRateCardsTiers     `json:"tiers,omitempty"`
}

func (l ListRateCardsCurrent) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListRateCardsCurrent) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListRateCardsCurrent) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ListRateCardsCurrent) GetProductID() *string {
	if o == nil {
		return nil
	}
	return o.ProductID
}

func (o *ListRateCardsCurrent) GetRateType() *ListRateCardsRateType {
	if o == nil {
		return nil
	}
	return o.RateType
}

func (o *ListRateCardsCurrent) GetPrice() *float64 {
	if o == nil {
		return nil
	}
	return o.Price
}

func (o *ListRateCardsCurrent) GetCreditType() *ListRateCardsCreditType {
	if o == nil {
		return nil
	}
	return o.CreditType
}

func (o *ListRateCardsCurrent) GetCustomRate() map[string]any {
	if o == nil {
		return nil
	}
	return o.CustomRate
}

func (o *ListRateCardsCurrent) GetStartingAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartingAt
}

func (o *ListRateCardsCurrent) GetEntitled() *bool {
	if o == nil {
		return nil
	}
	return o.Entitled
}

func (o *ListRateCardsCurrent) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *ListRateCardsCurrent) GetCreatedBy() *string {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}

func (o *ListRateCardsCurrent) GetEndingBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndingBefore
}

func (o *ListRateCardsCurrent) GetTiers() []ListRateCardsTiers {
	if o == nil {
		return nil
	}
	return o.Tiers
}

type ListRateCardsRateCardsRateType string

const (
	ListRateCardsRateCardsRateTypeFlat         ListRateCardsRateCardsRateType = "FLAT"
	ListRateCardsRateCardsRateTypePercentage   ListRateCardsRateCardsRateType = "PERCENTAGE"
	ListRateCardsRateCardsRateTypeSubscription ListRateCardsRateCardsRateType = "SUBSCRIPTION"
	ListRateCardsRateCardsRateTypeCustom       ListRateCardsRateCardsRateType = "CUSTOM"
	ListRateCardsRateCardsRateTypeTiered       ListRateCardsRateCardsRateType = "TIERED"
)

func (e ListRateCardsRateCardsRateType) ToPointer() *ListRateCardsRateCardsRateType {
	return &e
}
func (e *ListRateCardsRateCardsRateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "FLAT":
		fallthrough
	case "PERCENTAGE":
		fallthrough
	case "SUBSCRIPTION":
		fallthrough
	case "CUSTOM":
		fallthrough
	case "TIERED":
		*e = ListRateCardsRateCardsRateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListRateCardsRateCardsRateType: %v", v)
	}
}

type ListRateCardsRateCardsCreditType struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func (o *ListRateCardsRateCardsCreditType) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ListRateCardsRateCardsCreditType) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type ListRateCardsRateCardsTiers struct {
	Size  *float64 `json:"size,omitempty"`
	Price float64  `json:"price"`
}

func (o *ListRateCardsRateCardsTiers) GetSize() *float64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *ListRateCardsRateCardsTiers) GetPrice() float64 {
	if o == nil {
		return 0.0
	}
	return o.Price
}

type ListRateCardsUpdates struct {
	ID           string                            `json:"id"`
	ProductID    string                            `json:"product_id"`
	RateType     ListRateCardsRateCardsRateType    `json:"rate_type"`
	Price        *float64                          `json:"price,omitempty"`
	CreditType   *ListRateCardsRateCardsCreditType `json:"credit_type,omitempty"`
	CustomRate   map[string]any                    `json:"custom_rate,omitempty"`
	Quantity     *float64                          `json:"quantity,omitempty"`
	IsProrated   *bool                             `json:"is_prorated,omitempty"`
	StartingAt   time.Time                         `json:"starting_at"`
	Entitled     bool                              `json:"entitled"`
	CreatedAt    time.Time                         `json:"created_at"`
	CreatedBy    string                            `json:"created_by"`
	EndingBefore *time.Time                        `json:"ending_before,omitempty"`
	Tiers        []ListRateCardsRateCardsTiers     `json:"tiers,omitempty"`
}

func (l ListRateCardsUpdates) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListRateCardsUpdates) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListRateCardsUpdates) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ListRateCardsUpdates) GetProductID() string {
	if o == nil {
		return ""
	}
	return o.ProductID
}

func (o *ListRateCardsUpdates) GetRateType() ListRateCardsRateCardsRateType {
	if o == nil {
		return ListRateCardsRateCardsRateType("")
	}
	return o.RateType
}

func (o *ListRateCardsUpdates) GetPrice() *float64 {
	if o == nil {
		return nil
	}
	return o.Price
}

func (o *ListRateCardsUpdates) GetCreditType() *ListRateCardsRateCardsCreditType {
	if o == nil {
		return nil
	}
	return o.CreditType
}

func (o *ListRateCardsUpdates) GetCustomRate() map[string]any {
	if o == nil {
		return nil
	}
	return o.CustomRate
}

func (o *ListRateCardsUpdates) GetQuantity() *float64 {
	if o == nil {
		return nil
	}
	return o.Quantity
}

func (o *ListRateCardsUpdates) GetIsProrated() *bool {
	if o == nil {
		return nil
	}
	return o.IsProrated
}

func (o *ListRateCardsUpdates) GetStartingAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartingAt
}

func (o *ListRateCardsUpdates) GetEntitled() bool {
	if o == nil {
		return false
	}
	return o.Entitled
}

func (o *ListRateCardsUpdates) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *ListRateCardsUpdates) GetCreatedBy() string {
	if o == nil {
		return ""
	}
	return o.CreatedBy
}

func (o *ListRateCardsUpdates) GetEndingBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndingBefore
}

func (o *ListRateCardsUpdates) GetTiers() []ListRateCardsRateCardsTiers {
	if o == nil {
		return nil
	}
	return o.Tiers
}

type ListRateCardsRateCardEntries struct {
	Current *ListRateCardsCurrent  `json:"current,omitempty"`
	Updates []ListRateCardsUpdates `json:"updates,omitempty"`
}

func (o *ListRateCardsRateCardEntries) GetCurrent() *ListRateCardsCurrent {
	if o == nil {
		return nil
	}
	return o.Current
}

func (o *ListRateCardsRateCardEntries) GetUpdates() []ListRateCardsUpdates {
	if o == nil {
		return nil
	}
	return o.Updates
}

type ListRateCardsFiatCreditType struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func (o *ListRateCardsFiatCreditType) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ListRateCardsFiatCreditType) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type ListRateCardsCustomCreditType struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func (o *ListRateCardsCustomCreditType) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ListRateCardsCustomCreditType) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type ListRateCardsCreditTypeConversions struct {
	FiatPerCustomCredit string                        `json:"fiat_per_custom_credit"`
	CustomCreditType    ListRateCardsCustomCreditType `json:"custom_credit_type"`
}

func (o *ListRateCardsCreditTypeConversions) GetFiatPerCustomCredit() string {
	if o == nil {
		return ""
	}
	return o.FiatPerCustomCredit
}

func (o *ListRateCardsCreditTypeConversions) GetCustomCreditType() ListRateCardsCustomCreditType {
	if o == nil {
		return ListRateCardsCustomCreditType{}
	}
	return o.CustomCreditType
}

type ListRateCardsAliases struct {
	Name         string     `json:"name"`
	StartingAt   *time.Time `json:"starting_at,omitempty"`
	EndingBefore *time.Time `json:"ending_before,omitempty"`
}

func (l ListRateCardsAliases) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListRateCardsAliases) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListRateCardsAliases) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ListRateCardsAliases) GetStartingAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartingAt
}

func (o *ListRateCardsAliases) GetEndingBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndingBefore
}

type ListRateCardsData struct {
	ID                    string                                  `json:"id"`
	Name                  string                                  `json:"name"`
	RateCardEntries       map[string]ListRateCardsRateCardEntries `json:"rate_card_entries"`
	CreatedAt             time.Time                               `json:"created_at"`
	CreatedBy             string                                  `json:"created_by"`
	Description           *string                                 `json:"description,omitempty"`
	FiatCreditType        *ListRateCardsFiatCreditType            `json:"fiat_credit_type,omitempty"`
	CreditTypeConversions []ListRateCardsCreditTypeConversions    `json:"credit_type_conversions,omitempty"`
	Aliases               []ListRateCardsAliases                  `json:"aliases,omitempty"`
	CustomFields          map[string]string                       `json:"custom_fields,omitempty"`
}

func (l ListRateCardsData) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListRateCardsData) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListRateCardsData) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ListRateCardsData) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ListRateCardsData) GetRateCardEntries() map[string]ListRateCardsRateCardEntries {
	if o == nil {
		return map[string]ListRateCardsRateCardEntries{}
	}
	return o.RateCardEntries
}

func (o *ListRateCardsData) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *ListRateCardsData) GetCreatedBy() string {
	if o == nil {
		return ""
	}
	return o.CreatedBy
}

func (o *ListRateCardsData) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *ListRateCardsData) GetFiatCreditType() *ListRateCardsFiatCreditType {
	if o == nil {
		return nil
	}
	return o.FiatCreditType
}

func (o *ListRateCardsData) GetCreditTypeConversions() []ListRateCardsCreditTypeConversions {
	if o == nil {
		return nil
	}
	return o.CreditTypeConversions
}

func (o *ListRateCardsData) GetAliases() []ListRateCardsAliases {
	if o == nil {
		return nil
	}
	return o.Aliases
}

func (o *ListRateCardsData) GetCustomFields() map[string]string {
	if o == nil {
		return nil
	}
	return o.CustomFields
}

// ListRateCardsResponseBody - Success
type ListRateCardsResponseBody struct {
	Data     []ListRateCardsData `json:"data"`
	NextPage *string             `json:"next_page"`
}

func (o *ListRateCardsResponseBody) GetData() []ListRateCardsData {
	if o == nil {
		return []ListRateCardsData{}
	}
	return o.Data
}

func (o *ListRateCardsResponseBody) GetNextPage() *string {
	if o == nil {
		return nil
	}
	return o.NextPage
}

type ListRateCardsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	Object *ListRateCardsResponseBody

	Next func() (*ListRateCardsResponse, error)
}

func (o *ListRateCardsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListRateCardsResponse) GetObject() *ListRateCardsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
