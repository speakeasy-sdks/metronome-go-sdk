// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/metronome-go-sdk/internal/utils"
	"github.com/speakeasy-sdks/metronome-go-sdk/models/components"
	"time"
)

// ArchiveFilter - Filter options for the product list
type ArchiveFilter string

const (
	ArchiveFilterArchived    ArchiveFilter = "ARCHIVED"
	ArchiveFilterNotArchived ArchiveFilter = "NOT_ARCHIVED"
	ArchiveFilterAll         ArchiveFilter = "ALL"
)

func (e ArchiveFilter) ToPointer() *ArchiveFilter {
	return &e
}
func (e *ArchiveFilter) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ARCHIVED":
		fallthrough
	case "NOT_ARCHIVED":
		fallthrough
	case "ALL":
		*e = ArchiveFilter(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ArchiveFilter: %v", v)
	}
}

// ListProductsRequestBody - Get list of products
type ListProductsRequestBody struct {
	// Filter options for the product list
	ArchiveFilter *ArchiveFilter `json:"archive_filter,omitempty"`
}

func (o *ListProductsRequestBody) GetArchiveFilter() *ArchiveFilter {
	if o == nil {
		return nil
	}
	return o.ArchiveFilter
}

type ListProductsRequest struct {
	// Max number of results that should be returned
	Limit *int64 `queryParam:"style=form,explode=true,name=limit"`
	// Cursor that indicates where the next page of results should start.
	NextPage *string `queryParam:"style=form,explode=true,name=next_page"`
	// Get list of products
	RequestBody *ListProductsRequestBody `request:"mediaType=application/json"`
}

func (o *ListProductsRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *ListProductsRequest) GetNextPage() *string {
	if o == nil {
		return nil
	}
	return o.NextPage
}

func (o *ListProductsRequest) GetRequestBody() *ListProductsRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

type ListProductsType string

const (
	ListProductsTypeUsage        ListProductsType = "USAGE"
	ListProductsTypeSubscription ListProductsType = "SUBSCRIPTION"
	ListProductsTypeComposite    ListProductsType = "COMPOSITE"
	ListProductsTypeFixed        ListProductsType = "FIXED"
	ListProductsTypeProService   ListProductsType = "PRO_SERVICE"
)

func (e ListProductsType) ToPointer() *ListProductsType {
	return &e
}
func (e *ListProductsType) UnmarshalJSON(data []byte) error {
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
		*e = ListProductsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListProductsType: %v", v)
	}
}

// ListProductsOperation - The operation to perform on the quantity
type ListProductsOperation string

const (
	ListProductsOperationMultiplyLower ListProductsOperation = "multiply"
	ListProductsOperationDivideLower   ListProductsOperation = "divide"
	ListProductsOperationMultiplyUpper ListProductsOperation = "MULTIPLY"
	ListProductsOperationDivideUpper   ListProductsOperation = "DIVIDE"
)

func (e ListProductsOperation) ToPointer() *ListProductsOperation {
	return &e
}
func (e *ListProductsOperation) UnmarshalJSON(data []byte) error {
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
		*e = ListProductsOperation(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListProductsOperation: %v", v)
	}
}

// ListProductsQuantityConversion - Optional. Only valid for USAGE products. If provided, the quantity will be converted using the provided conversion factor and operation. For example, if the operation is "multiply" and the conversion factor is 100, then the quantity will be multiplied by 100. This can be used in cases where data is sent in one unit and priced in another.  For example, data could be sent in MB and priced in GB. In this case, the conversion factor would be 1024 and the operation would be "divide".
type ListProductsQuantityConversion struct {
	// Optional name for this conversion.
	Name *string `json:"name,omitempty"`
	// The factor to multiply or divide the quantity by.
	ConversionFactor float64 `json:"conversion_factor"`
	// The operation to perform on the quantity
	Operation ListProductsOperation `json:"operation"`
}

func (o *ListProductsQuantityConversion) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ListProductsQuantityConversion) GetConversionFactor() float64 {
	if o == nil {
		return 0.0
	}
	return o.ConversionFactor
}

func (o *ListProductsQuantityConversion) GetOperation() ListProductsOperation {
	if o == nil {
		return ListProductsOperation("")
	}
	return o.Operation
}

type ListProductsRoundingMethod string

const (
	ListProductsRoundingMethodRoundUpLower     ListProductsRoundingMethod = "round_up"
	ListProductsRoundingMethodRoundDownLower   ListProductsRoundingMethod = "round_down"
	ListProductsRoundingMethodRoundHalfUpLower ListProductsRoundingMethod = "round_half_up"
	ListProductsRoundingMethodRoundUpUpper     ListProductsRoundingMethod = "ROUND_UP"
	ListProductsRoundingMethodRoundDownUpper   ListProductsRoundingMethod = "ROUND_DOWN"
	ListProductsRoundingMethodRoundHalfUpUpper ListProductsRoundingMethod = "ROUND_HALF_UP"
)

func (e ListProductsRoundingMethod) ToPointer() *ListProductsRoundingMethod {
	return &e
}
func (e *ListProductsRoundingMethod) UnmarshalJSON(data []byte) error {
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
		*e = ListProductsRoundingMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListProductsRoundingMethod: %v", v)
	}
}

// ListProductsQuantityRounding - Optional. Only valid for USAGE products. If provided, the quantity will be rounded using the provided rounding method and decimal places. For example, if the method is "round up" and the decimal places is 0, then the quantity will be rounded up to the nearest integer.
type ListProductsQuantityRounding struct {
	RoundingMethod ListProductsRoundingMethod `json:"rounding_method"`
	DecimalPlaces  float64                    `json:"decimal_places"`
}

func (o *ListProductsQuantityRounding) GetRoundingMethod() ListProductsRoundingMethod {
	if o == nil {
		return ListProductsRoundingMethod("")
	}
	return o.RoundingMethod
}

func (o *ListProductsQuantityRounding) GetDecimalPlaces() float64 {
	if o == nil {
		return 0.0
	}
	return o.DecimalPlaces
}

type ListProductsInitial struct {
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
	QuantityConversion *ListProductsQuantityConversion `json:"quantity_conversion,omitempty"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be rounded using the provided rounding method and decimal places. For example, if the method is "round up" and the decimal places is 0, then the quantity will be rounded up to the nearest integer.
	QuantityRounding *ListProductsQuantityRounding `json:"quantity_rounding,omitempty"`
	CompositeTags    []string                      `json:"composite_tags,omitempty"`
	// This field's availability is dependent on your client's configuration.
	IsRefundable     *bool    `json:"is_refundable,omitempty"`
	Tags             []string `json:"tags,omitempty"`
	ExcludeFreeUsage *bool    `json:"exclude_free_usage,omitempty"`
	// For USAGE products only. If set, pricing for this product will be determined for each pricing_group_key value, as opposed to the product as a whole.
	PricingGroupKey []string `json:"pricing_group_key,omitempty"`
	// For USAGE products only. Groups usage line items on invoices.
	PresentationGroupKey []string `json:"presentation_group_key,omitempty"`
}

func (l ListProductsInitial) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListProductsInitial) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListProductsInitial) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ListProductsInitial) GetStartingAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartingAt
}

func (o *ListProductsInitial) GetNetsuiteInternalItemID() *string {
	if o == nil {
		return nil
	}
	return o.NetsuiteInternalItemID
}

func (o *ListProductsInitial) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *ListProductsInitial) GetCreatedBy() string {
	if o == nil {
		return ""
	}
	return o.CreatedBy
}

func (o *ListProductsInitial) GetNetsuiteOverageItemID() *string {
	if o == nil {
		return nil
	}
	return o.NetsuiteOverageItemID
}

func (o *ListProductsInitial) GetBillableMetricID() *string {
	if o == nil {
		return nil
	}
	return o.BillableMetricID
}

func (o *ListProductsInitial) GetCompositeProductIds() []string {
	if o == nil {
		return nil
	}
	return o.CompositeProductIds
}

func (o *ListProductsInitial) GetQuantityConversion() *ListProductsQuantityConversion {
	if o == nil {
		return nil
	}
	return o.QuantityConversion
}

func (o *ListProductsInitial) GetQuantityRounding() *ListProductsQuantityRounding {
	if o == nil {
		return nil
	}
	return o.QuantityRounding
}

func (o *ListProductsInitial) GetCompositeTags() []string {
	if o == nil {
		return nil
	}
	return o.CompositeTags
}

func (o *ListProductsInitial) GetIsRefundable() *bool {
	if o == nil {
		return nil
	}
	return o.IsRefundable
}

func (o *ListProductsInitial) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *ListProductsInitial) GetExcludeFreeUsage() *bool {
	if o == nil {
		return nil
	}
	return o.ExcludeFreeUsage
}

func (o *ListProductsInitial) GetPricingGroupKey() []string {
	if o == nil {
		return nil
	}
	return o.PricingGroupKey
}

func (o *ListProductsInitial) GetPresentationGroupKey() []string {
	if o == nil {
		return nil
	}
	return o.PresentationGroupKey
}

// ListProductsProductsOperation - The operation to perform on the quantity
type ListProductsProductsOperation string

const (
	ListProductsProductsOperationMultiplyLower ListProductsProductsOperation = "multiply"
	ListProductsProductsOperationDivideLower   ListProductsProductsOperation = "divide"
	ListProductsProductsOperationMultiplyUpper ListProductsProductsOperation = "MULTIPLY"
	ListProductsProductsOperationDivideUpper   ListProductsProductsOperation = "DIVIDE"
)

func (e ListProductsProductsOperation) ToPointer() *ListProductsProductsOperation {
	return &e
}
func (e *ListProductsProductsOperation) UnmarshalJSON(data []byte) error {
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
		*e = ListProductsProductsOperation(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListProductsProductsOperation: %v", v)
	}
}

// ListProductsProductsQuantityConversion - Optional. Only valid for USAGE products. If provided, the quantity will be converted using the provided conversion factor and operation. For example, if the operation is "multiply" and the conversion factor is 100, then the quantity will be multiplied by 100. This can be used in cases where data is sent in one unit and priced in another.  For example, data could be sent in MB and priced in GB. In this case, the conversion factor would be 1024 and the operation would be "divide".
type ListProductsProductsQuantityConversion struct {
	// Optional name for this conversion.
	Name *string `json:"name,omitempty"`
	// The factor to multiply or divide the quantity by.
	ConversionFactor float64 `json:"conversion_factor"`
	// The operation to perform on the quantity
	Operation ListProductsProductsOperation `json:"operation"`
}

func (o *ListProductsProductsQuantityConversion) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ListProductsProductsQuantityConversion) GetConversionFactor() float64 {
	if o == nil {
		return 0.0
	}
	return o.ConversionFactor
}

func (o *ListProductsProductsQuantityConversion) GetOperation() ListProductsProductsOperation {
	if o == nil {
		return ListProductsProductsOperation("")
	}
	return o.Operation
}

type ListProductsProductsRoundingMethod string

const (
	ListProductsProductsRoundingMethodRoundUpLower     ListProductsProductsRoundingMethod = "round_up"
	ListProductsProductsRoundingMethodRoundDownLower   ListProductsProductsRoundingMethod = "round_down"
	ListProductsProductsRoundingMethodRoundHalfUpLower ListProductsProductsRoundingMethod = "round_half_up"
	ListProductsProductsRoundingMethodRoundUpUpper     ListProductsProductsRoundingMethod = "ROUND_UP"
	ListProductsProductsRoundingMethodRoundDownUpper   ListProductsProductsRoundingMethod = "ROUND_DOWN"
	ListProductsProductsRoundingMethodRoundHalfUpUpper ListProductsProductsRoundingMethod = "ROUND_HALF_UP"
)

func (e ListProductsProductsRoundingMethod) ToPointer() *ListProductsProductsRoundingMethod {
	return &e
}
func (e *ListProductsProductsRoundingMethod) UnmarshalJSON(data []byte) error {
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
		*e = ListProductsProductsRoundingMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListProductsProductsRoundingMethod: %v", v)
	}
}

// ListProductsProductsQuantityRounding - Optional. Only valid for USAGE products. If provided, the quantity will be rounded using the provided rounding method and decimal places. For example, if the method is "round up" and the decimal places is 0, then the quantity will be rounded up to the nearest integer.
type ListProductsProductsQuantityRounding struct {
	RoundingMethod ListProductsProductsRoundingMethod `json:"rounding_method"`
	DecimalPlaces  float64                            `json:"decimal_places"`
}

func (o *ListProductsProductsQuantityRounding) GetRoundingMethod() ListProductsProductsRoundingMethod {
	if o == nil {
		return ListProductsProductsRoundingMethod("")
	}
	return o.RoundingMethod
}

func (o *ListProductsProductsQuantityRounding) GetDecimalPlaces() float64 {
	if o == nil {
		return 0.0
	}
	return o.DecimalPlaces
}

type ListProductsCurrent struct {
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
	QuantityConversion *ListProductsProductsQuantityConversion `json:"quantity_conversion,omitempty"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be rounded using the provided rounding method and decimal places. For example, if the method is "round up" and the decimal places is 0, then the quantity will be rounded up to the nearest integer.
	QuantityRounding *ListProductsProductsQuantityRounding `json:"quantity_rounding,omitempty"`
	CompositeTags    []string                              `json:"composite_tags,omitempty"`
	// This field's availability is dependent on your client's configuration.
	IsRefundable     *bool    `json:"is_refundable,omitempty"`
	Tags             []string `json:"tags,omitempty"`
	ExcludeFreeUsage *bool    `json:"exclude_free_usage,omitempty"`
	// For USAGE products only. If set, pricing for this product will be determined for each pricing_group_key value, as opposed to the product as a whole.
	PricingGroupKey []string `json:"pricing_group_key,omitempty"`
	// For USAGE products only. Groups usage line items on invoices.
	PresentationGroupKey []string `json:"presentation_group_key,omitempty"`
}

func (l ListProductsCurrent) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListProductsCurrent) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListProductsCurrent) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ListProductsCurrent) GetStartingAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartingAt
}

func (o *ListProductsCurrent) GetNetsuiteInternalItemID() *string {
	if o == nil {
		return nil
	}
	return o.NetsuiteInternalItemID
}

func (o *ListProductsCurrent) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *ListProductsCurrent) GetCreatedBy() string {
	if o == nil {
		return ""
	}
	return o.CreatedBy
}

func (o *ListProductsCurrent) GetNetsuiteOverageItemID() *string {
	if o == nil {
		return nil
	}
	return o.NetsuiteOverageItemID
}

func (o *ListProductsCurrent) GetBillableMetricID() *string {
	if o == nil {
		return nil
	}
	return o.BillableMetricID
}

func (o *ListProductsCurrent) GetCompositeProductIds() []string {
	if o == nil {
		return nil
	}
	return o.CompositeProductIds
}

func (o *ListProductsCurrent) GetQuantityConversion() *ListProductsProductsQuantityConversion {
	if o == nil {
		return nil
	}
	return o.QuantityConversion
}

func (o *ListProductsCurrent) GetQuantityRounding() *ListProductsProductsQuantityRounding {
	if o == nil {
		return nil
	}
	return o.QuantityRounding
}

func (o *ListProductsCurrent) GetCompositeTags() []string {
	if o == nil {
		return nil
	}
	return o.CompositeTags
}

func (o *ListProductsCurrent) GetIsRefundable() *bool {
	if o == nil {
		return nil
	}
	return o.IsRefundable
}

func (o *ListProductsCurrent) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *ListProductsCurrent) GetExcludeFreeUsage() *bool {
	if o == nil {
		return nil
	}
	return o.ExcludeFreeUsage
}

func (o *ListProductsCurrent) GetPricingGroupKey() []string {
	if o == nil {
		return nil
	}
	return o.PricingGroupKey
}

func (o *ListProductsCurrent) GetPresentationGroupKey() []string {
	if o == nil {
		return nil
	}
	return o.PresentationGroupKey
}

// ListProductsProductsResponseOperation - The operation to perform on the quantity
type ListProductsProductsResponseOperation string

const (
	ListProductsProductsResponseOperationMultiplyLower ListProductsProductsResponseOperation = "multiply"
	ListProductsProductsResponseOperationDivideLower   ListProductsProductsResponseOperation = "divide"
	ListProductsProductsResponseOperationMultiplyUpper ListProductsProductsResponseOperation = "MULTIPLY"
	ListProductsProductsResponseOperationDivideUpper   ListProductsProductsResponseOperation = "DIVIDE"
)

func (e ListProductsProductsResponseOperation) ToPointer() *ListProductsProductsResponseOperation {
	return &e
}
func (e *ListProductsProductsResponseOperation) UnmarshalJSON(data []byte) error {
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
		*e = ListProductsProductsResponseOperation(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListProductsProductsResponseOperation: %v", v)
	}
}

// ListProductsProductsResponseQuantityConversion - Optional. Only valid for USAGE products. If provided, the quantity will be converted using the provided conversion factor and operation. For example, if the operation is "multiply" and the conversion factor is 100, then the quantity will be multiplied by 100. This can be used in cases where data is sent in one unit and priced in another.  For example, data could be sent in MB and priced in GB. In this case, the conversion factor would be 1024 and the operation would be "divide".
type ListProductsProductsResponseQuantityConversion struct {
	// Optional name for this conversion.
	Name *string `json:"name,omitempty"`
	// The factor to multiply or divide the quantity by.
	ConversionFactor float64 `json:"conversion_factor"`
	// The operation to perform on the quantity
	Operation ListProductsProductsResponseOperation `json:"operation"`
}

func (o *ListProductsProductsResponseQuantityConversion) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ListProductsProductsResponseQuantityConversion) GetConversionFactor() float64 {
	if o == nil {
		return 0.0
	}
	return o.ConversionFactor
}

func (o *ListProductsProductsResponseQuantityConversion) GetOperation() ListProductsProductsResponseOperation {
	if o == nil {
		return ListProductsProductsResponseOperation("")
	}
	return o.Operation
}

type ListProductsProductsResponseRoundingMethod string

const (
	ListProductsProductsResponseRoundingMethodRoundUpLower     ListProductsProductsResponseRoundingMethod = "round_up"
	ListProductsProductsResponseRoundingMethodRoundDownLower   ListProductsProductsResponseRoundingMethod = "round_down"
	ListProductsProductsResponseRoundingMethodRoundHalfUpLower ListProductsProductsResponseRoundingMethod = "round_half_up"
	ListProductsProductsResponseRoundingMethodRoundUpUpper     ListProductsProductsResponseRoundingMethod = "ROUND_UP"
	ListProductsProductsResponseRoundingMethodRoundDownUpper   ListProductsProductsResponseRoundingMethod = "ROUND_DOWN"
	ListProductsProductsResponseRoundingMethodRoundHalfUpUpper ListProductsProductsResponseRoundingMethod = "ROUND_HALF_UP"
)

func (e ListProductsProductsResponseRoundingMethod) ToPointer() *ListProductsProductsResponseRoundingMethod {
	return &e
}
func (e *ListProductsProductsResponseRoundingMethod) UnmarshalJSON(data []byte) error {
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
		*e = ListProductsProductsResponseRoundingMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListProductsProductsResponseRoundingMethod: %v", v)
	}
}

// ListProductsProductsResponseQuantityRounding - Optional. Only valid for USAGE products. If provided, the quantity will be rounded using the provided rounding method and decimal places. For example, if the method is "round up" and the decimal places is 0, then the quantity will be rounded up to the nearest integer.
type ListProductsProductsResponseQuantityRounding struct {
	RoundingMethod ListProductsProductsResponseRoundingMethod `json:"rounding_method"`
	DecimalPlaces  float64                                    `json:"decimal_places"`
}

func (o *ListProductsProductsResponseQuantityRounding) GetRoundingMethod() ListProductsProductsResponseRoundingMethod {
	if o == nil {
		return ListProductsProductsResponseRoundingMethod("")
	}
	return o.RoundingMethod
}

func (o *ListProductsProductsResponseQuantityRounding) GetDecimalPlaces() float64 {
	if o == nil {
		return 0.0
	}
	return o.DecimalPlaces
}

type ListProductsUpdates struct {
	Name             *string    `json:"name,omitempty"`
	StartingAt       *time.Time `json:"starting_at,omitempty"`
	IsRefundable     *bool      `json:"is_refundable,omitempty"`
	CreatedAt        time.Time  `json:"created_at"`
	CreatedBy        string     `json:"created_by"`
	BillableMetricID *string    `json:"billable_metric_id,omitempty"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be converted using the provided conversion factor and operation. For example, if the operation is "multiply" and the conversion factor is 100, then the quantity will be multiplied by 100. This can be used in cases where data is sent in one unit and priced in another.  For example, data could be sent in MB and priced in GB. In this case, the conversion factor would be 1024 and the operation would be "divide".
	QuantityConversion *ListProductsProductsResponseQuantityConversion `json:"quantity_conversion,omitempty"`
	// Optional. Only valid for USAGE products. If provided, the quantity will be rounded using the provided rounding method and decimal places. For example, if the method is "round up" and the decimal places is 0, then the quantity will be rounded up to the nearest integer.
	QuantityRounding *ListProductsProductsResponseQuantityRounding `json:"quantity_rounding,omitempty"`
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

func (l ListProductsUpdates) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListProductsUpdates) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListProductsUpdates) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ListProductsUpdates) GetStartingAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartingAt
}

func (o *ListProductsUpdates) GetIsRefundable() *bool {
	if o == nil {
		return nil
	}
	return o.IsRefundable
}

func (o *ListProductsUpdates) GetCreatedAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedAt
}

func (o *ListProductsUpdates) GetCreatedBy() string {
	if o == nil {
		return ""
	}
	return o.CreatedBy
}

func (o *ListProductsUpdates) GetBillableMetricID() *string {
	if o == nil {
		return nil
	}
	return o.BillableMetricID
}

func (o *ListProductsUpdates) GetQuantityConversion() *ListProductsProductsResponseQuantityConversion {
	if o == nil {
		return nil
	}
	return o.QuantityConversion
}

func (o *ListProductsUpdates) GetQuantityRounding() *ListProductsProductsResponseQuantityRounding {
	if o == nil {
		return nil
	}
	return o.QuantityRounding
}

func (o *ListProductsUpdates) GetNetsuiteInternalItemID() *string {
	if o == nil {
		return nil
	}
	return o.NetsuiteInternalItemID
}

func (o *ListProductsUpdates) GetNetsuiteOverageItemID() *string {
	if o == nil {
		return nil
	}
	return o.NetsuiteOverageItemID
}

func (o *ListProductsUpdates) GetCompositeProductIds() []string {
	if o == nil {
		return nil
	}
	return o.CompositeProductIds
}

func (o *ListProductsUpdates) GetCompositeTags() []string {
	if o == nil {
		return nil
	}
	return o.CompositeTags
}

func (o *ListProductsUpdates) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *ListProductsUpdates) GetExcludeFreeUsage() *bool {
	if o == nil {
		return nil
	}
	return o.ExcludeFreeUsage
}

func (o *ListProductsUpdates) GetPricingGroupKey() []string {
	if o == nil {
		return nil
	}
	return o.PricingGroupKey
}

func (o *ListProductsUpdates) GetPresentationGroupKey() []string {
	if o == nil {
		return nil
	}
	return o.PresentationGroupKey
}

type ListProductsData struct {
	ID           string                `json:"id"`
	Type         ListProductsType      `json:"type"`
	ArchivedAt   *time.Time            `json:"archived_at,omitempty"`
	Initial      ListProductsInitial   `json:"initial"`
	Current      ListProductsCurrent   `json:"current"`
	Updates      []ListProductsUpdates `json:"updates"`
	CustomFields map[string]string     `json:"custom_fields,omitempty"`
}

func (l ListProductsData) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListProductsData) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListProductsData) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ListProductsData) GetType() ListProductsType {
	if o == nil {
		return ListProductsType("")
	}
	return o.Type
}

func (o *ListProductsData) GetArchivedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ArchivedAt
}

func (o *ListProductsData) GetInitial() ListProductsInitial {
	if o == nil {
		return ListProductsInitial{}
	}
	return o.Initial
}

func (o *ListProductsData) GetCurrent() ListProductsCurrent {
	if o == nil {
		return ListProductsCurrent{}
	}
	return o.Current
}

func (o *ListProductsData) GetUpdates() []ListProductsUpdates {
	if o == nil {
		return []ListProductsUpdates{}
	}
	return o.Updates
}

func (o *ListProductsData) GetCustomFields() map[string]string {
	if o == nil {
		return nil
	}
	return o.CustomFields
}

// ListProductsResponseBody - Success
type ListProductsResponseBody struct {
	Data     []ListProductsData `json:"data"`
	NextPage *string            `json:"next_page"`
}

func (o *ListProductsResponseBody) GetData() []ListProductsData {
	if o == nil {
		return []ListProductsData{}
	}
	return o.Data
}

func (o *ListProductsResponseBody) GetNextPage() *string {
	if o == nil {
		return nil
	}
	return o.NextPage
}

type ListProductsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	Object *ListProductsResponseBody

	Next func() (*ListProductsResponse, error)
}

func (o *ListProductsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListProductsResponse) GetObject() *ListProductsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
