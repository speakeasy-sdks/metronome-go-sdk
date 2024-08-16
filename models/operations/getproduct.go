// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/Metronome-Industries/metronome-go-sdk/internal/utils"
	"github.com/Metronome-Industries/metronome-go-sdk/models/components"
	"time"
)

// GetProductRequestBody - The ID of the product to get
type GetProductRequestBody struct {
	ID string `json:"id"`
}

func (o *GetProductRequestBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetProductType string

const (
	GetProductTypeUsage        GetProductType = "USAGE"
	GetProductTypeSubscription GetProductType = "SUBSCRIPTION"
	GetProductTypeComposite    GetProductType = "COMPOSITE"
	GetProductTypeFixed        GetProductType = "FIXED"
	GetProductTypeProService   GetProductType = "PRO_SERVICE"
)

func (e GetProductType) ToPointer() *GetProductType {
	return &e
}
func (e *GetProductType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "USAGE":
		fallthrough
	case "SUBSCRIPTION":
		fallthrough
	case "COMPOSITE":
		fallthrough
	case "FIXED":
		fallthrough
	case "PRO_SERVICE":
		*e = GetProductType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetProductType: %v", v)
	}
}

// GetProductProductsResponseOperation - The operation to perform on the quantity
type GetProductProductsResponseOperation string

const (
	GetProductProductsResponseOperationMultiplyLower GetProductProductsResponseOperation = "multiply"
	GetProductProductsResponseOperationDivideLower   GetProductProductsResponseOperation = "divide"
	GetProductProductsResponseOperationMultiplyUpper GetProductProductsResponseOperation = "MULTIPLY"
	GetProductProductsResponseOperationDivideUpper   GetProductProductsResponseOperation = "DIVIDE"
)

func (e GetProductProductsResponseOperation) ToPointer() *GetProductProductsResponseOperation {
	return &e
}
func (e *GetProductProductsResponseOperation) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "multiply":
		fallthrough
	case "divide":
		fallthrough
	case "MULTIPLY":
		fallthrough
	case "DIVIDE":
		*e = GetProductProductsResponseOperation(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetProductProductsResponseOperation: %v", v)
	}
}

// GetProductProductsResponseQuantityConversion - Optional. Only valid for USAGE products. If provided, the quantity will be converted using the provided conversion factor and operation. For example, if the operation is "multiply" and the conversion factor is 100, then the quantity will be multiplied by 100. This can be used in cases where data is sent in one unit and priced in another.  For example, data could be sent in MB and priced in GB. In this case, the conversion factor would be 1024 and the operation would be "divide".
type GetProductProductsResponseQuantityConversion struct {
	// Optional name for this conversion.
	Name *string `json:"name,omitempty"`
	// The factor to multiply or divide the quantity by.
	ConversionFactor float64 `json:"conversion_factor"`
	// The operation to perform on the quantity
	Operation GetProductProductsResponseOperation `json:"operation"`
}

func (o *GetProductProductsResponseQuantityConversion) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *GetProductProductsResponseQuantityConversion) GetConversionFactor() float64 {
	if o == nil {
		return 0.0
	}
	return o.ConversionFactor
}

func (o *GetProductProductsResponseQuantityConversion) GetOperation() GetProductProductsResponseOperation {
	if o == nil {
		return GetProductProductsResponseOperation("")
	}
	return o.Operation
}

type GetProductProductsResponseRoundingMethod string

const (
	GetProductProductsResponseRoundingMethodRoundUpLower     GetProductProductsResponseRoundingMethod = "round_up"
	GetProductProductsResponseRoundingMethodRoundDownLower   GetProductProductsResponseRoundingMethod = "round_down"
	GetProductProductsResponseRoundingMethodRoundHalfUpLower GetProductProductsResponseRoundingMethod = "round_half_up"
	GetProductProductsResponseRoundingMethodRoundUpUpper     GetProductProductsResponseRoundingMethod = "ROUND_UP"
	GetProductProductsResponseRoundingMethodRoundDownUpper   GetProductProductsResponseRoundingMethod = "ROUND_DOWN"
	GetProductProductsResponseRoundingMethodRoundHalfUpUpper GetProductProductsResponseRoundingMethod = "ROUND_HALF_UP"
)

func (e GetProductProductsResponseRoundingMethod) ToPointer() *GetProductProductsResponseRoundingMethod {
	return &e
}
func (e *GetProductProductsResponseRoundingMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "round_up":
		fallthrough
	case "round_down":
		fallthrough
	case "round_half_up":
		fallthrough
	case "ROUND_UP":
		fallthrough
	case "ROUND_DOWN":
		fallthrough
	case "ROUND_HALF_UP":
		*e = GetProductProductsResponseRoundingMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetProductProductsResponseRoundingMethod: %v", v)
	}
}

// GetProductProductsResponseQuantityRounding - Optional. Only valid for USAGE products. If provided, the quantity will be rounded using the provided rounding method and decimal places. For example, if the method is "round up" and the decimal places is 0, then the quantity will be rounded up to the nearest integer.
type GetProductProductsResponseQuantityRounding struct {
	RoundingMethod GetProductProductsResponseRoundingMethod `json:"rounding_method"`
	DecimalPlaces  float64                                  `json:"decimal_places"`
}

func (o *GetProductProductsResponseQuantityRounding) GetRoundingMethod() GetProductProductsResponseRoundingMethod {
	if o == nil {
		return GetProductProductsResponseRoundingMethod("")
	}
	return o.RoundingMethod
}

func (o *GetProductProductsResponseQuantityRounding) GetDecimalPlaces() float64 {
	if o == nil {
		return 0.0
	}
	return o.DecimalPlaces
}

type Initial struct {
	Name       string     `json:"name"`
	StartingAt *time.Time `json:"starting_at,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteInternalItemID *string   `json:"netsuite_internal_item_id,omitempty"`
	CreatedAt              time.Time `json:"created_at"`
	CreatedBy              string    `json:"created_by"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteOverageItemID *string  `json:"netsuite_overage_item_id,omitempty"`
	BillableMetricID      *string  `json:"billable_metric_id,omitempty"`
	CompositeProductIds   []string `json:"composite_product_ids,omitempty"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be converted using the provided conversion factor and operation. For example, if the operation is "multiply" and the conversion factor is 100, then the quantity will be multiplied by 100. This can be used in cases where data is sent in one unit and priced in another.  For example, data could be sent in MB and priced in GB. In this case, the conversion factor would be 1024 and the operation would be "divide".
	QuantityConversion *GetProductProductsResponseQuantityConversion `json:"quantity_conversion,omitempty"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be rounded using the provided rounding method and decimal places. For example, if the method is "round up" and the decimal places is 0, then the quantity will be rounded up to the nearest integer.
	QuantityRounding *GetProductProductsResponseQuantityRounding `json:"quantity_rounding,omitempty"`
	CompositeTags    []string                                    `json:"composite_tags,omitempty"`
	// This field's availability is dependent on your client's configuration.
	IsRefundable     *bool    `json:"is_refundable,omitempty"`
	Tags             []string `json:"tags,omitempty"`
	ExcludeFreeUsage *bool    `json:"exclude_free_usage,omitempty"`
	// For USAGE products only. If set, pricing for this product will be determined for each pricing_group_key value, as opposed to the product as a whole.
	PricingGroupKey []string `json:"pricing_group_key,omitempty"`
	// For USAGE products only. Groups usage line items on invoices.
	PresentationGroupKey []string `json:"presentation_group_key,omitempty"`
}

func (i Initial) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *Initial) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Initial) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Initial) GetStartingAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartingAt
}

func (o *Initial) GetNetsuiteInternalItemID() *string {
	if o == nil {
		return nil
	}
	return o.NetsuiteInternalItemID
}

func (o *Initial) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *Initial) GetCreatedBy() string {
	if o == nil {
		return ""
	}
	return o.CreatedBy
}

func (o *Initial) GetNetsuiteOverageItemID() *string {
	if o == nil {
		return nil
	}
	return o.NetsuiteOverageItemID
}

func (o *Initial) GetBillableMetricID() *string {
	if o == nil {
		return nil
	}
	return o.BillableMetricID
}

func (o *Initial) GetCompositeProductIds() []string {
	if o == nil {
		return nil
	}
	return o.CompositeProductIds
}

func (o *Initial) GetQuantityConversion() *GetProductProductsResponseQuantityConversion {
	if o == nil {
		return nil
	}
	return o.QuantityConversion
}

func (o *Initial) GetQuantityRounding() *GetProductProductsResponseQuantityRounding {
	if o == nil {
		return nil
	}
	return o.QuantityRounding
}

func (o *Initial) GetCompositeTags() []string {
	if o == nil {
		return nil
	}
	return o.CompositeTags
}

func (o *Initial) GetIsRefundable() *bool {
	if o == nil {
		return nil
	}
	return o.IsRefundable
}

func (o *Initial) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *Initial) GetExcludeFreeUsage() *bool {
	if o == nil {
		return nil
	}
	return o.ExcludeFreeUsage
}

func (o *Initial) GetPricingGroupKey() []string {
	if o == nil {
		return nil
	}
	return o.PricingGroupKey
}

func (o *Initial) GetPresentationGroupKey() []string {
	if o == nil {
		return nil
	}
	return o.PresentationGroupKey
}

// GetProductOperation - The operation to perform on the quantity
type GetProductOperation string

const (
	GetProductOperationMultiplyLower GetProductOperation = "multiply"
	GetProductOperationDivideLower   GetProductOperation = "divide"
	GetProductOperationMultiplyUpper GetProductOperation = "MULTIPLY"
	GetProductOperationDivideUpper   GetProductOperation = "DIVIDE"
)

func (e GetProductOperation) ToPointer() *GetProductOperation {
	return &e
}
func (e *GetProductOperation) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "multiply":
		fallthrough
	case "divide":
		fallthrough
	case "MULTIPLY":
		fallthrough
	case "DIVIDE":
		*e = GetProductOperation(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetProductOperation: %v", v)
	}
}

// GetProductQuantityConversion - Optional. Only valid for USAGE products. If provided, the quantity will be converted using the provided conversion factor and operation. For example, if the operation is "multiply" and the conversion factor is 100, then the quantity will be multiplied by 100. This can be used in cases where data is sent in one unit and priced in another.  For example, data could be sent in MB and priced in GB. In this case, the conversion factor would be 1024 and the operation would be "divide".
type GetProductQuantityConversion struct {
	// Optional name for this conversion.
	Name *string `json:"name,omitempty"`
	// The factor to multiply or divide the quantity by.
	ConversionFactor float64 `json:"conversion_factor"`
	// The operation to perform on the quantity
	Operation GetProductOperation `json:"operation"`
}

func (o *GetProductQuantityConversion) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *GetProductQuantityConversion) GetConversionFactor() float64 {
	if o == nil {
		return 0.0
	}
	return o.ConversionFactor
}

func (o *GetProductQuantityConversion) GetOperation() GetProductOperation {
	if o == nil {
		return GetProductOperation("")
	}
	return o.Operation
}

type GetProductRoundingMethod string

const (
	GetProductRoundingMethodRoundUpLower     GetProductRoundingMethod = "round_up"
	GetProductRoundingMethodRoundDownLower   GetProductRoundingMethod = "round_down"
	GetProductRoundingMethodRoundHalfUpLower GetProductRoundingMethod = "round_half_up"
	GetProductRoundingMethodRoundUpUpper     GetProductRoundingMethod = "ROUND_UP"
	GetProductRoundingMethodRoundDownUpper   GetProductRoundingMethod = "ROUND_DOWN"
	GetProductRoundingMethodRoundHalfUpUpper GetProductRoundingMethod = "ROUND_HALF_UP"
)

func (e GetProductRoundingMethod) ToPointer() *GetProductRoundingMethod {
	return &e
}
func (e *GetProductRoundingMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "round_up":
		fallthrough
	case "round_down":
		fallthrough
	case "round_half_up":
		fallthrough
	case "ROUND_UP":
		fallthrough
	case "ROUND_DOWN":
		fallthrough
	case "ROUND_HALF_UP":
		*e = GetProductRoundingMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetProductRoundingMethod: %v", v)
	}
}

// GetProductQuantityRounding - Optional. Only valid for USAGE products. If provided, the quantity will be rounded using the provided rounding method and decimal places. For example, if the method is "round up" and the decimal places is 0, then the quantity will be rounded up to the nearest integer.
type GetProductQuantityRounding struct {
	RoundingMethod GetProductRoundingMethod `json:"rounding_method"`
	DecimalPlaces  float64                  `json:"decimal_places"`
}

func (o *GetProductQuantityRounding) GetRoundingMethod() GetProductRoundingMethod {
	if o == nil {
		return GetProductRoundingMethod("")
	}
	return o.RoundingMethod
}

func (o *GetProductQuantityRounding) GetDecimalPlaces() float64 {
	if o == nil {
		return 0.0
	}
	return o.DecimalPlaces
}

type Current struct {
	Name       string     `json:"name"`
	StartingAt *time.Time `json:"starting_at,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteInternalItemID *string   `json:"netsuite_internal_item_id,omitempty"`
	CreatedAt              time.Time `json:"created_at"`
	CreatedBy              string    `json:"created_by"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteOverageItemID *string  `json:"netsuite_overage_item_id,omitempty"`
	BillableMetricID      *string  `json:"billable_metric_id,omitempty"`
	CompositeProductIds   []string `json:"composite_product_ids,omitempty"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be converted using the provided conversion factor and operation. For example, if the operation is "multiply" and the conversion factor is 100, then the quantity will be multiplied by 100. This can be used in cases where data is sent in one unit and priced in another.  For example, data could be sent in MB and priced in GB. In this case, the conversion factor would be 1024 and the operation would be "divide".
	QuantityConversion *GetProductQuantityConversion `json:"quantity_conversion,omitempty"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be rounded using the provided rounding method and decimal places. For example, if the method is "round up" and the decimal places is 0, then the quantity will be rounded up to the nearest integer.
	QuantityRounding *GetProductQuantityRounding `json:"quantity_rounding,omitempty"`
	CompositeTags    []string                    `json:"composite_tags,omitempty"`
	// This field's availability is dependent on your client's configuration.
	IsRefundable     *bool    `json:"is_refundable,omitempty"`
	Tags             []string `json:"tags,omitempty"`
	ExcludeFreeUsage *bool    `json:"exclude_free_usage,omitempty"`
	// For USAGE products only. If set, pricing for this product will be determined for each pricing_group_key value, as opposed to the product as a whole.
	PricingGroupKey []string `json:"pricing_group_key,omitempty"`
	// For USAGE products only. Groups usage line items on invoices.
	PresentationGroupKey []string `json:"presentation_group_key,omitempty"`
}

func (c Current) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *Current) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Current) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Current) GetStartingAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartingAt
}

func (o *Current) GetNetsuiteInternalItemID() *string {
	if o == nil {
		return nil
	}
	return o.NetsuiteInternalItemID
}

func (o *Current) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *Current) GetCreatedBy() string {
	if o == nil {
		return ""
	}
	return o.CreatedBy
}

func (o *Current) GetNetsuiteOverageItemID() *string {
	if o == nil {
		return nil
	}
	return o.NetsuiteOverageItemID
}

func (o *Current) GetBillableMetricID() *string {
	if o == nil {
		return nil
	}
	return o.BillableMetricID
}

func (o *Current) GetCompositeProductIds() []string {
	if o == nil {
		return nil
	}
	return o.CompositeProductIds
}

func (o *Current) GetQuantityConversion() *GetProductQuantityConversion {
	if o == nil {
		return nil
	}
	return o.QuantityConversion
}

func (o *Current) GetQuantityRounding() *GetProductQuantityRounding {
	if o == nil {
		return nil
	}
	return o.QuantityRounding
}

func (o *Current) GetCompositeTags() []string {
	if o == nil {
		return nil
	}
	return o.CompositeTags
}

func (o *Current) GetIsRefundable() *bool {
	if o == nil {
		return nil
	}
	return o.IsRefundable
}

func (o *Current) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *Current) GetExcludeFreeUsage() *bool {
	if o == nil {
		return nil
	}
	return o.ExcludeFreeUsage
}

func (o *Current) GetPricingGroupKey() []string {
	if o == nil {
		return nil
	}
	return o.PricingGroupKey
}

func (o *Current) GetPresentationGroupKey() []string {
	if o == nil {
		return nil
	}
	return o.PresentationGroupKey
}

// GetProductProductsOperation - The operation to perform on the quantity
type GetProductProductsOperation string

const (
	GetProductProductsOperationMultiplyLower GetProductProductsOperation = "multiply"
	GetProductProductsOperationDivideLower   GetProductProductsOperation = "divide"
	GetProductProductsOperationMultiplyUpper GetProductProductsOperation = "MULTIPLY"
	GetProductProductsOperationDivideUpper   GetProductProductsOperation = "DIVIDE"
)

func (e GetProductProductsOperation) ToPointer() *GetProductProductsOperation {
	return &e
}
func (e *GetProductProductsOperation) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "multiply":
		fallthrough
	case "divide":
		fallthrough
	case "MULTIPLY":
		fallthrough
	case "DIVIDE":
		*e = GetProductProductsOperation(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetProductProductsOperation: %v", v)
	}
}

// GetProductProductsQuantityConversion - Optional. Only valid for USAGE products. If provided, the quantity will be converted using the provided conversion factor and operation. For example, if the operation is "multiply" and the conversion factor is 100, then the quantity will be multiplied by 100. This can be used in cases where data is sent in one unit and priced in another.  For example, data could be sent in MB and priced in GB. In this case, the conversion factor would be 1024 and the operation would be "divide".
type GetProductProductsQuantityConversion struct {
	// Optional name for this conversion.
	Name *string `json:"name,omitempty"`
	// The factor to multiply or divide the quantity by.
	ConversionFactor float64 `json:"conversion_factor"`
	// The operation to perform on the quantity
	Operation GetProductProductsOperation `json:"operation"`
}

func (o *GetProductProductsQuantityConversion) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *GetProductProductsQuantityConversion) GetConversionFactor() float64 {
	if o == nil {
		return 0.0
	}
	return o.ConversionFactor
}

func (o *GetProductProductsQuantityConversion) GetOperation() GetProductProductsOperation {
	if o == nil {
		return GetProductProductsOperation("")
	}
	return o.Operation
}

type GetProductProductsRoundingMethod string

const (
	GetProductProductsRoundingMethodRoundUpLower     GetProductProductsRoundingMethod = "round_up"
	GetProductProductsRoundingMethodRoundDownLower   GetProductProductsRoundingMethod = "round_down"
	GetProductProductsRoundingMethodRoundHalfUpLower GetProductProductsRoundingMethod = "round_half_up"
	GetProductProductsRoundingMethodRoundUpUpper     GetProductProductsRoundingMethod = "ROUND_UP"
	GetProductProductsRoundingMethodRoundDownUpper   GetProductProductsRoundingMethod = "ROUND_DOWN"
	GetProductProductsRoundingMethodRoundHalfUpUpper GetProductProductsRoundingMethod = "ROUND_HALF_UP"
)

func (e GetProductProductsRoundingMethod) ToPointer() *GetProductProductsRoundingMethod {
	return &e
}
func (e *GetProductProductsRoundingMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "round_up":
		fallthrough
	case "round_down":
		fallthrough
	case "round_half_up":
		fallthrough
	case "ROUND_UP":
		fallthrough
	case "ROUND_DOWN":
		fallthrough
	case "ROUND_HALF_UP":
		*e = GetProductProductsRoundingMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetProductProductsRoundingMethod: %v", v)
	}
}

// GetProductProductsQuantityRounding - Optional. Only valid for USAGE products. If provided, the quantity will be rounded using the provided rounding method and decimal places. For example, if the method is "round up" and the decimal places is 0, then the quantity will be rounded up to the nearest integer.
type GetProductProductsQuantityRounding struct {
	RoundingMethod GetProductProductsRoundingMethod `json:"rounding_method"`
	DecimalPlaces  float64                          `json:"decimal_places"`
}

func (o *GetProductProductsQuantityRounding) GetRoundingMethod() GetProductProductsRoundingMethod {
	if o == nil {
		return GetProductProductsRoundingMethod("")
	}
	return o.RoundingMethod
}

func (o *GetProductProductsQuantityRounding) GetDecimalPlaces() float64 {
	if o == nil {
		return 0.0
	}
	return o.DecimalPlaces
}

type Updates struct {
	Name             *string    `json:"name,omitempty"`
	StartingAt       *time.Time `json:"starting_at,omitempty"`
	IsRefundable     *bool      `json:"is_refundable,omitempty"`
	CreatedAt        time.Time  `json:"created_at"`
	CreatedBy        string     `json:"created_by"`
	BillableMetricID *string    `json:"billable_metric_id,omitempty"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be converted using the provided conversion factor and operation. For example, if the operation is "multiply" and the conversion factor is 100, then the quantity will be multiplied by 100. This can be used in cases where data is sent in one unit and priced in another.  For example, data could be sent in MB and priced in GB. In this case, the conversion factor would be 1024 and the operation would be "divide".
	QuantityConversion *GetProductProductsQuantityConversion `json:"quantity_conversion,omitempty"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be rounded using the provided rounding method and decimal places. For example, if the method is "round up" and the decimal places is 0, then the quantity will be rounded up to the nearest integer.
	QuantityRounding *GetProductProductsQuantityRounding `json:"quantity_rounding,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteInternalItemID *string `json:"netsuite_internal_item_id,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteOverageItemID *string  `json:"netsuite_overage_item_id,omitempty"`
	CompositeProductIds   []string `json:"composite_product_ids,omitempty"`
	CompositeTags         []string `json:"composite_tags,omitempty"`
	Tags                  []string `json:"tags,omitempty"`
	ExcludeFreeUsage      *bool    `json:"exclude_free_usage,omitempty"`
	// For USAGE products only. If set, pricing for this product will be determined for each pricing_group_key value, as opposed to the product as a whole.
	PricingGroupKey []string `json:"pricing_group_key,omitempty"`
	// For USAGE products only. Groups usage line items on invoices.
	PresentationGroupKey []string `json:"presentation_group_key,omitempty"`
}

func (u Updates) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *Updates) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Updates) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Updates) GetStartingAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartingAt
}

func (o *Updates) GetIsRefundable() *bool {
	if o == nil {
		return nil
	}
	return o.IsRefundable
}

func (o *Updates) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *Updates) GetCreatedBy() string {
	if o == nil {
		return ""
	}
	return o.CreatedBy
}

func (o *Updates) GetBillableMetricID() *string {
	if o == nil {
		return nil
	}
	return o.BillableMetricID
}

func (o *Updates) GetQuantityConversion() *GetProductProductsQuantityConversion {
	if o == nil {
		return nil
	}
	return o.QuantityConversion
}

func (o *Updates) GetQuantityRounding() *GetProductProductsQuantityRounding {
	if o == nil {
		return nil
	}
	return o.QuantityRounding
}

func (o *Updates) GetNetsuiteInternalItemID() *string {
	if o == nil {
		return nil
	}
	return o.NetsuiteInternalItemID
}

func (o *Updates) GetNetsuiteOverageItemID() *string {
	if o == nil {
		return nil
	}
	return o.NetsuiteOverageItemID
}

func (o *Updates) GetCompositeProductIds() []string {
	if o == nil {
		return nil
	}
	return o.CompositeProductIds
}

func (o *Updates) GetCompositeTags() []string {
	if o == nil {
		return nil
	}
	return o.CompositeTags
}

func (o *Updates) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *Updates) GetExcludeFreeUsage() *bool {
	if o == nil {
		return nil
	}
	return o.ExcludeFreeUsage
}

func (o *Updates) GetPricingGroupKey() []string {
	if o == nil {
		return nil
	}
	return o.PricingGroupKey
}

func (o *Updates) GetPresentationGroupKey() []string {
	if o == nil {
		return nil
	}
	return o.PresentationGroupKey
}

type GetProductData struct {
	ID           string            `json:"id"`
	Type         GetProductType    `json:"type"`
	ArchivedAt   *time.Time        `json:"archived_at,omitempty"`
	Initial      Initial           `json:"initial"`
	Current      Current           `json:"current"`
	Updates      []Updates         `json:"updates"`
	CustomFields map[string]string `json:"custom_fields,omitempty"`
}

func (g GetProductData) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetProductData) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetProductData) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetProductData) GetType() GetProductType {
	if o == nil {
		return GetProductType("")
	}
	return o.Type
}

func (o *GetProductData) GetArchivedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ArchivedAt
}

func (o *GetProductData) GetInitial() Initial {
	if o == nil {
		return Initial{}
	}
	return o.Initial
}

func (o *GetProductData) GetCurrent() Current {
	if o == nil {
		return Current{}
	}
	return o.Current
}

func (o *GetProductData) GetUpdates() []Updates {
	if o == nil {
		return []Updates{}
	}
	return o.Updates
}

func (o *GetProductData) GetCustomFields() map[string]string {
	if o == nil {
		return nil
	}
	return o.CustomFields
}

// GetProductResponseBody - Success
type GetProductResponseBody struct {
	Data GetProductData `json:"data"`
}

func (o *GetProductResponseBody) GetData() GetProductData {
	if o == nil {
		return GetProductData{}
	}
	return o.Data
}

type GetProductResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	Object *GetProductResponseBody
}

func (o *GetProductResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetProductResponse) GetObject() *GetProductResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
