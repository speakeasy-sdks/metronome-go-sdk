// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/metronome-go-sdk/internal/utils"
	"github.com/speakeasy-sdks/metronome-go-sdk/models/components"
	"time"
)

// GetRateCardRequestBody - The ID of the rate card to get
type GetRateCardRequestBody struct {
	ID string `json:"id"`
}

func (o *GetRateCardRequestBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetRateCardRateType string

const (
	GetRateCardRateTypeFlat         GetRateCardRateType = "FLAT"
	GetRateCardRateTypePercentage   GetRateCardRateType = "PERCENTAGE"
	GetRateCardRateTypeSubscription GetRateCardRateType = "SUBSCRIPTION"
	GetRateCardRateTypeCustom       GetRateCardRateType = "CUSTOM"
	GetRateCardRateTypeTiered       GetRateCardRateType = "TIERED"
)

func (e GetRateCardRateType) ToPointer() *GetRateCardRateType {
	return &e
}
func (e *GetRateCardRateType) UnmarshalJSON(data []byte) error {
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
		*e = GetRateCardRateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRateCardRateType: %v", v)
	}
}

type GetRateCardCreditType struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func (o *GetRateCardCreditType) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *GetRateCardCreditType) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetRateCardTiers struct {
	Size  *float64 `json:"size,omitempty"`
	Price float64  `json:"price"`
}

func (o *GetRateCardTiers) GetSize() *float64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *GetRateCardTiers) GetPrice() float64 {
	if o == nil {
		return 0.0
	}
	return o.Price
}

type GetRateCardCurrent struct {
	ID           *string                `json:"id,omitempty"`
	ProductID    *string                `json:"product_id,omitempty"`
	RateType     *GetRateCardRateType   `json:"rate_type,omitempty"`
	Price        *float64               `json:"price,omitempty"`
	CreditType   *GetRateCardCreditType `json:"credit_type,omitempty"`
	CustomRate   map[string]any         `json:"custom_rate,omitempty"`
	StartingAt   *time.Time             `json:"starting_at,omitempty"`
	Entitled     *bool                  `json:"entitled,omitempty"`
	CreatedAt    *time.Time             `json:"created_at,omitempty"`
	CreatedBy    *string                `json:"created_by,omitempty"`
	EndingBefore *time.Time             `json:"ending_before,omitempty"`
	Tiers        []GetRateCardTiers     `json:"tiers,omitempty"`
}

func (g GetRateCardCurrent) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetRateCardCurrent) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetRateCardCurrent) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetRateCardCurrent) GetProductID() *string {
	if o == nil {
		return nil
	}
	return o.ProductID
}

func (o *GetRateCardCurrent) GetRateType() *GetRateCardRateType {
	if o == nil {
		return nil
	}
	return o.RateType
}

func (o *GetRateCardCurrent) GetPrice() *float64 {
	if o == nil {
		return nil
	}
	return o.Price
}

func (o *GetRateCardCurrent) GetCreditType() *GetRateCardCreditType {
	if o == nil {
		return nil
	}
	return o.CreditType
}

func (o *GetRateCardCurrent) GetCustomRate() map[string]any {
	if o == nil {
		return nil
	}
	return o.CustomRate
}

func (o *GetRateCardCurrent) GetStartingAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartingAt
}

func (o *GetRateCardCurrent) GetEntitled() *bool {
	if o == nil {
		return nil
	}
	return o.Entitled
}

func (o *GetRateCardCurrent) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *GetRateCardCurrent) GetCreatedBy() *string {
	if o == nil {
		return nil
	}
	return o.CreatedBy
}

func (o *GetRateCardCurrent) GetEndingBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndingBefore
}

func (o *GetRateCardCurrent) GetTiers() []GetRateCardTiers {
	if o == nil {
		return nil
	}
	return o.Tiers
}

type GetRateCardRateCardsRateType string

const (
	GetRateCardRateCardsRateTypeFlat         GetRateCardRateCardsRateType = "FLAT"
	GetRateCardRateCardsRateTypePercentage   GetRateCardRateCardsRateType = "PERCENTAGE"
	GetRateCardRateCardsRateTypeSubscription GetRateCardRateCardsRateType = "SUBSCRIPTION"
	GetRateCardRateCardsRateTypeCustom       GetRateCardRateCardsRateType = "CUSTOM"
	GetRateCardRateCardsRateTypeTiered       GetRateCardRateCardsRateType = "TIERED"
)

func (e GetRateCardRateCardsRateType) ToPointer() *GetRateCardRateCardsRateType {
	return &e
}
func (e *GetRateCardRateCardsRateType) UnmarshalJSON(data []byte) error {
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
		*e = GetRateCardRateCardsRateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRateCardRateCardsRateType: %v", v)
	}
}

type GetRateCardRateCardsCreditType struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func (o *GetRateCardRateCardsCreditType) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *GetRateCardRateCardsCreditType) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetRateCardRateCardsTiers struct {
	Size  *float64 `json:"size,omitempty"`
	Price float64  `json:"price"`
}

func (o *GetRateCardRateCardsTiers) GetSize() *float64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *GetRateCardRateCardsTiers) GetPrice() float64 {
	if o == nil {
		return 0.0
	}
	return o.Price
}

type GetRateCardRateCardsResponseRateType string

const (
	GetRateCardRateCardsResponseRateTypeFlatUpper         GetRateCardRateCardsResponseRateType = "FLAT"
	GetRateCardRateCardsResponseRateTypeFlatLower         GetRateCardRateCardsResponseRateType = "flat"
	GetRateCardRateCardsResponseRateTypePercentageUpper   GetRateCardRateCardsResponseRateType = "PERCENTAGE"
	GetRateCardRateCardsResponseRateTypePercentageLower   GetRateCardRateCardsResponseRateType = "percentage"
	GetRateCardRateCardsResponseRateTypeSubscriptionUpper GetRateCardRateCardsResponseRateType = "SUBSCRIPTION"
	GetRateCardRateCardsResponseRateTypeSubscriptionLower GetRateCardRateCardsResponseRateType = "subscription"
	GetRateCardRateCardsResponseRateTypeTieredUpper       GetRateCardRateCardsResponseRateType = "TIERED"
	GetRateCardRateCardsResponseRateTypeTieredLower       GetRateCardRateCardsResponseRateType = "tiered"
	GetRateCardRateCardsResponseRateTypeCustomUpper       GetRateCardRateCardsResponseRateType = "CUSTOM"
	GetRateCardRateCardsResponseRateTypeCustomLower       GetRateCardRateCardsResponseRateType = "custom"
)

func (e GetRateCardRateCardsResponseRateType) ToPointer() *GetRateCardRateCardsResponseRateType {
	return &e
}
func (e *GetRateCardRateCardsResponseRateType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "FLAT":
		fallthrough
	case "flat":
		fallthrough
	case "PERCENTAGE":
		fallthrough
	case "percentage":
		fallthrough
	case "SUBSCRIPTION":
		fallthrough
	case "subscription":
		fallthrough
	case "TIERED":
		fallthrough
	case "tiered":
		fallthrough
	case "CUSTOM":
		fallthrough
	case "custom":
		*e = GetRateCardRateCardsResponseRateType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetRateCardRateCardsResponseRateType: %v", v)
	}
}

type GetRateCardRateCardsResponseTiers struct {
	Size  *float64 `json:"size,omitempty"`
	Price float64  `json:"price"`
}

func (o *GetRateCardRateCardsResponseTiers) GetSize() *float64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *GetRateCardRateCardsResponseTiers) GetPrice() float64 {
	if o == nil {
		return 0.0
	}
	return o.Price
}

type GetRateCardRateCardsResponseCreditType struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func (o *GetRateCardRateCardsResponseCreditType) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *GetRateCardRateCardsResponseCreditType) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// GetRateCardCommitRate - The rate that will be used to rate a product when it is paid for by a commit. This feature requires opt-in before it can be used. Please contact Metronome support to enable this feature.
type GetRateCardCommitRate struct {
	RateType GetRateCardRateCardsResponseRateType `json:"rate_type"`
	// Commit rate price. For FLAT rate_type, this must be >=0. For PERCENTAGE rate_type, this is a decimal fraction, e.g. use 0.1 for 10%; this must be >=0 and <=1.
	Price *float64 `json:"price,omitempty"`
	// Commit rate quantity. For SUBSCRIPTION rate_type, this must be >=0.
	Quantity *float64 `json:"quantity,omitempty"`
	// Commit rate proration configuration. Only valid for SUBSCRIPTION rate_type.
	IsProrated *bool `json:"is_prorated,omitempty"`
	// Only set for PERCENTAGE rate_type. Defaults to false. If true, rate is computed using list prices rather than the standard rates for this product on the contract.
	UseListPrices *bool `json:"use_list_prices,omitempty"`
	// Only set for TIERED rate_type.
	Tiers      []GetRateCardRateCardsResponseTiers     `json:"tiers,omitempty"`
	CreditType *GetRateCardRateCardsResponseCreditType `json:"credit_type,omitempty"`
}

func (o *GetRateCardCommitRate) GetRateType() GetRateCardRateCardsResponseRateType {
	if o == nil {
		return GetRateCardRateCardsResponseRateType("")
	}
	return o.RateType
}

func (o *GetRateCardCommitRate) GetPrice() *float64 {
	if o == nil {
		return nil
	}
	return o.Price
}

func (o *GetRateCardCommitRate) GetQuantity() *float64 {
	if o == nil {
		return nil
	}
	return o.Quantity
}

func (o *GetRateCardCommitRate) GetIsProrated() *bool {
	if o == nil {
		return nil
	}
	return o.IsProrated
}

func (o *GetRateCardCommitRate) GetUseListPrices() *bool {
	if o == nil {
		return nil
	}
	return o.UseListPrices
}

func (o *GetRateCardCommitRate) GetTiers() []GetRateCardRateCardsResponseTiers {
	if o == nil {
		return nil
	}
	return o.Tiers
}

func (o *GetRateCardCommitRate) GetCreditType() *GetRateCardRateCardsResponseCreditType {
	if o == nil {
		return nil
	}
	return o.CreditType
}

type GetRateCardUpdates struct {
	ID           string                          `json:"id"`
	ProductID    string                          `json:"product_id"`
	RateType     GetRateCardRateCardsRateType    `json:"rate_type"`
	Price        *float64                        `json:"price,omitempty"`
	CreditType   *GetRateCardRateCardsCreditType `json:"credit_type,omitempty"`
	CustomRate   map[string]any                  `json:"custom_rate,omitempty"`
	Quantity     *float64                        `json:"quantity,omitempty"`
	IsProrated   *bool                           `json:"is_prorated,omitempty"`
	StartingAt   time.Time                       `json:"starting_at"`
	Entitled     bool                            `json:"entitled"`
	CreatedAt    time.Time                       `json:"created_at"`
	CreatedBy    string                          `json:"created_by"`
	EndingBefore *time.Time                      `json:"ending_before,omitempty"`
	Tiers        []GetRateCardRateCardsTiers     `json:"tiers,omitempty"`
	// The rate that will be used to rate a product when it is paid for by a commit. This feature requires opt-in before it can be used. Please contact Metronome support to enable this feature.
	CommitRate *GetRateCardCommitRate `json:"commit_rate,omitempty"`
}

func (g GetRateCardUpdates) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetRateCardUpdates) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetRateCardUpdates) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetRateCardUpdates) GetProductID() string {
	if o == nil {
		return ""
	}
	return o.ProductID
}

func (o *GetRateCardUpdates) GetRateType() GetRateCardRateCardsRateType {
	if o == nil {
		return GetRateCardRateCardsRateType("")
	}
	return o.RateType
}

func (o *GetRateCardUpdates) GetPrice() *float64 {
	if o == nil {
		return nil
	}
	return o.Price
}

func (o *GetRateCardUpdates) GetCreditType() *GetRateCardRateCardsCreditType {
	if o == nil {
		return nil
	}
	return o.CreditType
}

func (o *GetRateCardUpdates) GetCustomRate() map[string]any {
	if o == nil {
		return nil
	}
	return o.CustomRate
}

func (o *GetRateCardUpdates) GetQuantity() *float64 {
	if o == nil {
		return nil
	}
	return o.Quantity
}

func (o *GetRateCardUpdates) GetIsProrated() *bool {
	if o == nil {
		return nil
	}
	return o.IsProrated
}

func (o *GetRateCardUpdates) GetStartingAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartingAt
}

func (o *GetRateCardUpdates) GetEntitled() bool {
	if o == nil {
		return false
	}
	return o.Entitled
}

func (o *GetRateCardUpdates) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *GetRateCardUpdates) GetCreatedBy() string {
	if o == nil {
		return ""
	}
	return o.CreatedBy
}

func (o *GetRateCardUpdates) GetEndingBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndingBefore
}

func (o *GetRateCardUpdates) GetTiers() []GetRateCardRateCardsTiers {
	if o == nil {
		return nil
	}
	return o.Tiers
}

func (o *GetRateCardUpdates) GetCommitRate() *GetRateCardCommitRate {
	if o == nil {
		return nil
	}
	return o.CommitRate
}

type RateCardEntries struct {
	Current *GetRateCardCurrent  `json:"current,omitempty"`
	Updates []GetRateCardUpdates `json:"updates,omitempty"`
}

func (o *RateCardEntries) GetCurrent() *GetRateCardCurrent {
	if o == nil {
		return nil
	}
	return o.Current
}

func (o *RateCardEntries) GetUpdates() []GetRateCardUpdates {
	if o == nil {
		return nil
	}
	return o.Updates
}

type FiatCreditType struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func (o *FiatCreditType) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *FiatCreditType) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type CustomCreditType struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func (o *CustomCreditType) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CustomCreditType) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetRateCardCreditTypeConversions struct {
	FiatPerCustomCredit string           `json:"fiat_per_custom_credit"`
	CustomCreditType    CustomCreditType `json:"custom_credit_type"`
}

func (o *GetRateCardCreditTypeConversions) GetFiatPerCustomCredit() string {
	if o == nil {
		return ""
	}
	return o.FiatPerCustomCredit
}

func (o *GetRateCardCreditTypeConversions) GetCustomCreditType() CustomCreditType {
	if o == nil {
		return CustomCreditType{}
	}
	return o.CustomCreditType
}

type GetRateCardAliases struct {
	Name         string     `json:"name"`
	StartingAt   *time.Time `json:"starting_at,omitempty"`
	EndingBefore *time.Time `json:"ending_before,omitempty"`
}

func (g GetRateCardAliases) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetRateCardAliases) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetRateCardAliases) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *GetRateCardAliases) GetStartingAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartingAt
}

func (o *GetRateCardAliases) GetEndingBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndingBefore
}

type GetRateCardData struct {
	ID                    string                             `json:"id"`
	Name                  string                             `json:"name"`
	RateCardEntries       map[string]RateCardEntries         `json:"rate_card_entries"`
	CreatedAt             time.Time                          `json:"created_at"`
	CreatedBy             string                             `json:"created_by"`
	Description           *string                            `json:"description,omitempty"`
	FiatCreditType        *FiatCreditType                    `json:"fiat_credit_type,omitempty"`
	CreditTypeConversions []GetRateCardCreditTypeConversions `json:"credit_type_conversions,omitempty"`
	Aliases               []GetRateCardAliases               `json:"aliases,omitempty"`
	CustomFields          map[string]string                  `json:"custom_fields,omitempty"`
}

func (g GetRateCardData) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetRateCardData) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetRateCardData) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetRateCardData) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *GetRateCardData) GetRateCardEntries() map[string]RateCardEntries {
	if o == nil {
		return map[string]RateCardEntries{}
	}
	return o.RateCardEntries
}

func (o *GetRateCardData) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *GetRateCardData) GetCreatedBy() string {
	if o == nil {
		return ""
	}
	return o.CreatedBy
}

func (o *GetRateCardData) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *GetRateCardData) GetFiatCreditType() *FiatCreditType {
	if o == nil {
		return nil
	}
	return o.FiatCreditType
}

func (o *GetRateCardData) GetCreditTypeConversions() []GetRateCardCreditTypeConversions {
	if o == nil {
		return nil
	}
	return o.CreditTypeConversions
}

func (o *GetRateCardData) GetAliases() []GetRateCardAliases {
	if o == nil {
		return nil
	}
	return o.Aliases
}

func (o *GetRateCardData) GetCustomFields() map[string]string {
	if o == nil {
		return nil
	}
	return o.CustomFields
}

// GetRateCardResponseBody - Success
type GetRateCardResponseBody struct {
	Data GetRateCardData `json:"data"`
}

func (o *GetRateCardResponseBody) GetData() GetRateCardData {
	if o == nil {
		return GetRateCardData{}
	}
	return o.Data
}

type GetRateCardResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	Object *GetRateCardResponseBody
}

func (o *GetRateCardResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetRateCardResponse) GetObject() *GetRateCardResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
