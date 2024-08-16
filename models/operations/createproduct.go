// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/metronome-go-sdk/models/components"
)

type Type string

const (
	TypeFixedUpper               Type = "FIXED"
	TypeFixedLower               Type = "fixed"
	TypeUsageUpper               Type = "USAGE"
	TypeUsageLower               Type = "usage"
	TypeCompositeUpper           Type = "COMPOSITE"
	TypeCompositeLower           Type = "composite"
	TypeSubscriptionUpper        Type = "SUBSCRIPTION"
	TypeSubscriptionLower        Type = "subscription"
	TypeProfessionalServiceUpper Type = "PROFESSIONAL_SERVICE"
	TypeProfessionalServiceLower Type = "professional_service"
	TypeProServiceUpper          Type = "PRO_SERVICE"
	TypeProServiceLower          Type = "pro_service"
)

func (e Type) ToPointer() *Type {
	return &e
}
func (e *Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "FIXED":
		fallthrough
	case "fixed":
		fallthrough
	case "USAGE":
		fallthrough
	case "usage":
		fallthrough
	case "COMPOSITE":
		fallthrough
	case "composite":
		fallthrough
	case "SUBSCRIPTION":
		fallthrough
	case "subscription":
		fallthrough
	case "PROFESSIONAL_SERVICE":
		fallthrough
	case "professional_service":
		fallthrough
	case "PRO_SERVICE":
		fallthrough
	case "pro_service":
		*e = Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Type: %v", v)
	}
}

// Operation - The operation to perform on the quantity
type Operation string

const (
	OperationMultiplyLower Operation = "multiply"
	OperationDivideLower   Operation = "divide"
	OperationMultiplyUpper Operation = "MULTIPLY"
	OperationDivideUpper   Operation = "DIVIDE"
)

func (e Operation) ToPointer() *Operation {
	return &e
}
func (e *Operation) UnmarshalJSON(data []byte) error {
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
		*e = Operation(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Operation: %v", v)
	}
}

// QuantityConversion - Optional. Only valid for USAGE products. If provided, the quantity will be converted using the provided conversion factor and operation. For example, if the operation is "multiply" and the conversion factor is 100, then the quantity will be multiplied by 100. This can be used in cases where data is sent in one unit and priced in another.  For example, data could be sent in MB and priced in GB. In this case, the conversion factor would be 1024 and the operation would be "divide".
type QuantityConversion struct {
	// Optional name for this conversion.
	Name *string `json:"name,omitempty"`
	// The factor to multiply or divide the quantity by.
	ConversionFactor float64 `json:"conversion_factor"`
	// The operation to perform on the quantity
	Operation Operation `json:"operation"`
}

func (o *QuantityConversion) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *QuantityConversion) GetConversionFactor() float64 {
	if o == nil {
		return 0.0
	}
	return o.ConversionFactor
}

func (o *QuantityConversion) GetOperation() Operation {
	if o == nil {
		return Operation("")
	}
	return o.Operation
}

type RoundingMethod string

const (
	RoundingMethodRoundUpLower     RoundingMethod = "round_up"
	RoundingMethodRoundDownLower   RoundingMethod = "round_down"
	RoundingMethodRoundHalfUpLower RoundingMethod = "round_half_up"
	RoundingMethodRoundUpUpper     RoundingMethod = "ROUND_UP"
	RoundingMethodRoundDownUpper   RoundingMethod = "ROUND_DOWN"
	RoundingMethodRoundHalfUpUpper RoundingMethod = "ROUND_HALF_UP"
)

func (e RoundingMethod) ToPointer() *RoundingMethod {
	return &e
}
func (e *RoundingMethod) UnmarshalJSON(data []byte) error {
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
		*e = RoundingMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RoundingMethod: %v", v)
	}
}

// QuantityRounding - Optional. Only valid for USAGE products. If provided, the quantity will be rounded using the provided rounding method and decimal places. For example, if the method is "round up" and the decimal places is 0, then the quantity will be rounded up to the nearest integer.
type QuantityRounding struct {
	RoundingMethod RoundingMethod `json:"rounding_method"`
	DecimalPlaces  float64        `json:"decimal_places"`
}

func (o *QuantityRounding) GetRoundingMethod() RoundingMethod {
	if o == nil {
		return RoundingMethod("")
	}
	return o.RoundingMethod
}

func (o *QuantityRounding) GetDecimalPlaces() float64 {
	if o == nil {
		return 0.0
	}
	return o.DecimalPlaces
}

// CreateProductRequestBody - Create a new product
type CreateProductRequestBody struct {
	// displayed on invoices
	Name string `json:"name"`
	Type Type   `json:"type"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteInternalItemID *string `json:"netsuite_internal_item_id,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteOverageItemID *string `json:"netsuite_overage_item_id,omitempty"`
	// Required for USAGE products
	BillableMetricID *string `json:"billable_metric_id,omitempty"`
	// Required for COMPOSITE products
	CompositeProductIds []string `json:"composite_product_ids,omitempty"`
	// Required for COMPOSITE products
	CompositeTags []string `json:"composite_tags,omitempty"`
	// This field's availability is dependent on your client's configuration. Defaults to true
	IsRefundable *bool `json:"is_refundable,omitempty"`
	// Beta feature only available for composite products. If true, products with $0 will not be included when computing composite usage. Defaults to false
	ExcludeFreeUsage *bool    `json:"exclude_free_usage,omitempty"`
	Tags             []string `json:"tags,omitempty"`
	// For USAGE products only. If set, pricing for this product will be determined for each pricing_group_key value, as opposed to the product as a whole.
	PricingGroupKey []string `json:"pricing_group_key,omitempty"`
	// For USAGE products only. Groups usage line items on invoices.
	PresentationGroupKey []string `json:"presentation_group_key,omitempty"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be converted using the provided conversion factor and operation. For example, if the operation is "multiply" and the conversion factor is 100, then the quantity will be multiplied by 100. This can be used in cases where data is sent in one unit and priced in another.  For example, data could be sent in MB and priced in GB. In this case, the conversion factor would be 1024 and the operation would be "divide".
	QuantityConversion *QuantityConversion `json:"quantity_conversion,omitempty"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be rounded using the provided rounding method and decimal places. For example, if the method is "round up" and the decimal places is 0, then the quantity will be rounded up to the nearest integer.
	QuantityRounding *QuantityRounding `json:"quantity_rounding,omitempty"`
}

func (o *CreateProductRequestBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateProductRequestBody) GetType() Type {
	if o == nil {
		return Type("")
	}
	return o.Type
}

func (o *CreateProductRequestBody) GetNetsuiteInternalItemID() *string {
	if o == nil {
		return nil
	}
	return o.NetsuiteInternalItemID
}

func (o *CreateProductRequestBody) GetNetsuiteOverageItemID() *string {
	if o == nil {
		return nil
	}
	return o.NetsuiteOverageItemID
}

func (o *CreateProductRequestBody) GetBillableMetricID() *string {
	if o == nil {
		return nil
	}
	return o.BillableMetricID
}

func (o *CreateProductRequestBody) GetCompositeProductIds() []string {
	if o == nil {
		return nil
	}
	return o.CompositeProductIds
}

func (o *CreateProductRequestBody) GetCompositeTags() []string {
	if o == nil {
		return nil
	}
	return o.CompositeTags
}

func (o *CreateProductRequestBody) GetIsRefundable() *bool {
	if o == nil {
		return nil
	}
	return o.IsRefundable
}

func (o *CreateProductRequestBody) GetExcludeFreeUsage() *bool {
	if o == nil {
		return nil
	}
	return o.ExcludeFreeUsage
}

func (o *CreateProductRequestBody) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateProductRequestBody) GetPricingGroupKey() []string {
	if o == nil {
		return nil
	}
	return o.PricingGroupKey
}

func (o *CreateProductRequestBody) GetPresentationGroupKey() []string {
	if o == nil {
		return nil
	}
	return o.PresentationGroupKey
}

func (o *CreateProductRequestBody) GetQuantityConversion() *QuantityConversion {
	if o == nil {
		return nil
	}
	return o.QuantityConversion
}

func (o *CreateProductRequestBody) GetQuantityRounding() *QuantityRounding {
	if o == nil {
		return nil
	}
	return o.QuantityRounding
}

type CreateProductData struct {
	ID string `json:"id"`
}

func (o *CreateProductData) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// CreateProductResponseBody - Success
type CreateProductResponseBody struct {
	Data CreateProductData `json:"data"`
}

func (o *CreateProductResponseBody) GetData() CreateProductData {
	if o == nil {
		return CreateProductData{}
	}
	return o.Data
}

type CreateProductResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	Object *CreateProductResponseBody
}

func (o *CreateProductResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateProductResponse) GetObject() *CreateProductResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
