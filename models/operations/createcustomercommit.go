// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/Metronome-Industries/metronome-go-sdk/internal/utils"
	"github.com/Metronome-Industries/metronome-go-sdk/models/components"
	"time"
)

type CreateCustomerCommitType string

const (
	CreateCustomerCommitTypePrepaidUpper  CreateCustomerCommitType = "PREPAID"
	CreateCustomerCommitTypePrepaidLower  CreateCustomerCommitType = "prepaid"
	CreateCustomerCommitTypePostpaidUpper CreateCustomerCommitType = "POSTPAID"
	CreateCustomerCommitTypePostpaidLower CreateCustomerCommitType = "postpaid"
)

func (e CreateCustomerCommitType) ToPointer() *CreateCustomerCommitType {
	return &e
}
func (e *CreateCustomerCommitType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PREPAID":
		fallthrough
	case "prepaid":
		fallthrough
	case "POSTPAID":
		fallthrough
	case "postpaid":
		*e = CreateCustomerCommitType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateCustomerCommitType: %v", v)
	}
}

type CreateCustomerCommitScheduleItems struct {
	Amount float64 `json:"amount"`
	// RFC 3339 timestamp (inclusive)
	StartingAt time.Time `json:"starting_at"`
	// RFC 3339 timestamp (exclusive)
	EndingBefore time.Time `json:"ending_before"`
}

func (c CreateCustomerCommitScheduleItems) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateCustomerCommitScheduleItems) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateCustomerCommitScheduleItems) GetAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.Amount
}

func (o *CreateCustomerCommitScheduleItems) GetStartingAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartingAt
}

func (o *CreateCustomerCommitScheduleItems) GetEndingBefore() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.EndingBefore
}

// CreateCustomerCommitAccessSchedule - Schedule for distributing the commit to the customer. For "POSTPAID" commits only one schedule item is allowed and amount must match invoice_schedule total.
type CreateCustomerCommitAccessSchedule struct {
	CreditTypeID  *string                             `json:"credit_type_id,omitempty"`
	ScheduleItems []CreateCustomerCommitScheduleItems `json:"schedule_items"`
}

func (o *CreateCustomerCommitAccessSchedule) GetCreditTypeID() *string {
	if o == nil {
		return nil
	}
	return o.CreditTypeID
}

func (o *CreateCustomerCommitAccessSchedule) GetScheduleItems() []CreateCustomerCommitScheduleItems {
	if o == nil {
		return []CreateCustomerCommitScheduleItems{}
	}
	return o.ScheduleItems
}

type CreateCustomerCommitCustomerCommitsScheduleItems struct {
	// Unit price for the charge. Will be multiplied by quantity to determine the amount and must be specified with quantity. If specified amount cannot be provided.
	UnitPrice *float64 `json:"unit_price,omitempty"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the amount and must be specified with unit_price. If specified amount cannot be provided.
	Quantity *float64 `json:"quantity,omitempty"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If amount is sent, the unit_price is assumed to be the amount and quantity is inferred to be 1.
	Amount *float64 `json:"amount,omitempty"`
	// timestamp of the scheduled event
	Timestamp time.Time `json:"timestamp"`
}

func (c CreateCustomerCommitCustomerCommitsScheduleItems) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateCustomerCommitCustomerCommitsScheduleItems) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateCustomerCommitCustomerCommitsScheduleItems) GetUnitPrice() *float64 {
	if o == nil {
		return nil
	}
	return o.UnitPrice
}

func (o *CreateCustomerCommitCustomerCommitsScheduleItems) GetQuantity() *float64 {
	if o == nil {
		return nil
	}
	return o.Quantity
}

func (o *CreateCustomerCommitCustomerCommitsScheduleItems) GetAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *CreateCustomerCommitCustomerCommitsScheduleItems) GetTimestamp() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Timestamp
}

type CreateCustomerCommitFrequency string

const (
	CreateCustomerCommitFrequencyMonthlyUpper    CreateCustomerCommitFrequency = "MONTHLY"
	CreateCustomerCommitFrequencyMonthlyLower    CreateCustomerCommitFrequency = "monthly"
	CreateCustomerCommitFrequencyQuarterlyUpper  CreateCustomerCommitFrequency = "QUARTERLY"
	CreateCustomerCommitFrequencyQuarterlyLower  CreateCustomerCommitFrequency = "quarterly"
	CreateCustomerCommitFrequencySemiAnnualUpper CreateCustomerCommitFrequency = "SEMI_ANNUAL"
	CreateCustomerCommitFrequencySemiAnnualLower CreateCustomerCommitFrequency = "semi_annual"
	CreateCustomerCommitFrequencyAnnualUpper     CreateCustomerCommitFrequency = "ANNUAL"
	CreateCustomerCommitFrequencyAnnualLower     CreateCustomerCommitFrequency = "annual"
)

func (e CreateCustomerCommitFrequency) ToPointer() *CreateCustomerCommitFrequency {
	return &e
}
func (e *CreateCustomerCommitFrequency) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MONTHLY":
		fallthrough
	case "monthly":
		fallthrough
	case "QUARTERLY":
		fallthrough
	case "quarterly":
		fallthrough
	case "SEMI_ANNUAL":
		fallthrough
	case "semi_annual":
		fallthrough
	case "ANNUAL":
		fallthrough
	case "annual":
		*e = CreateCustomerCommitFrequency(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateCustomerCommitFrequency: %v", v)
	}
}

type CreateCustomerCommitAmountDistribution string

const (
	CreateCustomerCommitAmountDistributionDividedUpper        CreateCustomerCommitAmountDistribution = "DIVIDED"
	CreateCustomerCommitAmountDistributionDividedLower        CreateCustomerCommitAmountDistribution = "divided"
	CreateCustomerCommitAmountDistributionDividedRoundedUpper CreateCustomerCommitAmountDistribution = "DIVIDED_ROUNDED"
	CreateCustomerCommitAmountDistributionDividedRoundedLower CreateCustomerCommitAmountDistribution = "divided_rounded"
	CreateCustomerCommitAmountDistributionEachUpper           CreateCustomerCommitAmountDistribution = "EACH"
	CreateCustomerCommitAmountDistributionEachLower           CreateCustomerCommitAmountDistribution = "each"
)

func (e CreateCustomerCommitAmountDistribution) ToPointer() *CreateCustomerCommitAmountDistribution {
	return &e
}
func (e *CreateCustomerCommitAmountDistribution) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DIVIDED":
		fallthrough
	case "divided":
		fallthrough
	case "DIVIDED_ROUNDED":
		fallthrough
	case "divided_rounded":
		fallthrough
	case "EACH":
		fallthrough
	case "each":
		*e = CreateCustomerCommitAmountDistribution(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateCustomerCommitAmountDistribution: %v", v)
	}
}

// CreateCustomerCommitRecurringSchedule - Enter the unit price and quantity for the charge or instead only send the amount. If amount is sent, the unit price is assumed to be the amount and quantity is inferred to be 1.
type CreateCustomerCommitRecurringSchedule struct {
	// RFC 3339 timestamp (inclusive).
	StartingAt time.Time `json:"starting_at"`
	// RFC 3339 timestamp (exclusive).
	EndingBefore time.Time                     `json:"ending_before"`
	Frequency    CreateCustomerCommitFrequency `json:"frequency"`
	// Unit price for the charge. Will be multiplied by quantity to determine the amount and must be specified with quantity. If specified amount cannot be provided.
	UnitPrice *float64 `json:"unit_price,omitempty"`
	// Quantity for the charge. Will be multiplied by unit_price to determine the amount and must be specified with unit_price. If specified amount cannot be provided.
	Quantity *float64 `json:"quantity,omitempty"`
	// Amount for the charge. Can be provided instead of unit_price and quantity. If amount is sent, the unit_price is assumed to be the amount and quantity is inferred to be 1.
	Amount             *float64                               `json:"amount,omitempty"`
	AmountDistribution CreateCustomerCommitAmountDistribution `json:"amount_distribution"`
}

func (c CreateCustomerCommitRecurringSchedule) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateCustomerCommitRecurringSchedule) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateCustomerCommitRecurringSchedule) GetStartingAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartingAt
}

func (o *CreateCustomerCommitRecurringSchedule) GetEndingBefore() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.EndingBefore
}

func (o *CreateCustomerCommitRecurringSchedule) GetFrequency() CreateCustomerCommitFrequency {
	if o == nil {
		return CreateCustomerCommitFrequency("")
	}
	return o.Frequency
}

func (o *CreateCustomerCommitRecurringSchedule) GetUnitPrice() *float64 {
	if o == nil {
		return nil
	}
	return o.UnitPrice
}

func (o *CreateCustomerCommitRecurringSchedule) GetQuantity() *float64 {
	if o == nil {
		return nil
	}
	return o.Quantity
}

func (o *CreateCustomerCommitRecurringSchedule) GetAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *CreateCustomerCommitRecurringSchedule) GetAmountDistribution() CreateCustomerCommitAmountDistribution {
	if o == nil {
		return CreateCustomerCommitAmountDistribution("")
	}
	return o.AmountDistribution
}

// CreateCustomerCommitInvoiceSchedule - Required for "POSTPAID" commits: the true up invoice will be generated at this time and only one schedule item is allowed; the total must match accesss_schedule amount. Optional for "PREPAID" commits: if not provided, this will be a "complimentary" commit with no invoice.
type CreateCustomerCommitInvoiceSchedule struct {
	// Defaults to USD if not passed. Only USD is supported at this time.
	CreditTypeID *string `json:"credit_type_id,omitempty"`
	// Either provide amount or provide both unit_price and quantity.
	ScheduleItems []CreateCustomerCommitCustomerCommitsScheduleItems `json:"schedule_items,omitempty"`
	// Enter the unit price and quantity for the charge or instead only send the amount. If amount is sent, the unit price is assumed to be the amount and quantity is inferred to be 1.
	RecurringSchedule *CreateCustomerCommitRecurringSchedule `json:"recurring_schedule,omitempty"`
}

func (o *CreateCustomerCommitInvoiceSchedule) GetCreditTypeID() *string {
	if o == nil {
		return nil
	}
	return o.CreditTypeID
}

func (o *CreateCustomerCommitInvoiceSchedule) GetScheduleItems() []CreateCustomerCommitCustomerCommitsScheduleItems {
	if o == nil {
		return nil
	}
	return o.ScheduleItems
}

func (o *CreateCustomerCommitInvoiceSchedule) GetRecurringSchedule() *CreateCustomerCommitRecurringSchedule {
	if o == nil {
		return nil
	}
	return o.RecurringSchedule
}

// CreateCustomerCommitRequestBody - Create a commit
type CreateCustomerCommitRequestBody struct {
	CustomerID string                   `json:"customer_id"`
	Type       CreateCustomerCommitType `json:"type"`
	// displayed on invoices
	Name *string `json:"name,omitempty"`
	// Used only in UI/API. It is not exposed to end customers.
	Description *string `json:"description,omitempty"`
	// If multiple credits or commits are applicable, the one with the lower priority will apply first.
	Priority  float64 `json:"priority"`
	ProductID string  `json:"product_id"`
	// Schedule for distributing the commit to the customer. For "POSTPAID" commits only one schedule item is allowed and amount must match invoice_schedule total.
	AccessSchedule CreateCustomerCommitAccessSchedule `json:"access_schedule"`
	// Required for "POSTPAID" commits: the true up invoice will be generated at this time and only one schedule item is allowed; the total must match accesss_schedule amount. Optional for "PREPAID" commits: if not provided, this will be a "complimentary" commit with no invoice.
	InvoiceSchedule *CreateCustomerCommitInvoiceSchedule `json:"invoice_schedule,omitempty"`
	// The contract that this commit will be billed on. This is required for "POSTPAID" commits and for "PREPAID" commits unless there is no invoice schedule above (i.e., the commit is 'free').
	InvoiceContractID *string `json:"invoice_contract_id,omitempty"`
	// Which products the commit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the commit applies to all products.
	ApplicableProductIds []string `json:"applicable_product_ids,omitempty"`
	// Which tags the commit applies to. If both applicable_product_ids and applicable_product_tags are not provided, the commit applies to all products.
	ApplicableProductTags []string `json:"applicable_product_tags,omitempty"`
	// Which contract the commit applies to. If not provided, the commit applies to all contracts.
	ApplicableContractIds []string `json:"applicable_contract_ids,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID *string `json:"netsuite_sales_order_id,omitempty"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID *string           `json:"salesforce_opportunity_id,omitempty"`
	CustomFields            map[string]string `json:"custom_fields,omitempty"`
}

func (o *CreateCustomerCommitRequestBody) GetCustomerID() string {
	if o == nil {
		return ""
	}
	return o.CustomerID
}

func (o *CreateCustomerCommitRequestBody) GetType() CreateCustomerCommitType {
	if o == nil {
		return CreateCustomerCommitType("")
	}
	return o.Type
}

func (o *CreateCustomerCommitRequestBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CreateCustomerCommitRequestBody) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateCustomerCommitRequestBody) GetPriority() float64 {
	if o == nil {
		return 0.0
	}
	return o.Priority
}

func (o *CreateCustomerCommitRequestBody) GetProductID() string {
	if o == nil {
		return ""
	}
	return o.ProductID
}

func (o *CreateCustomerCommitRequestBody) GetAccessSchedule() CreateCustomerCommitAccessSchedule {
	if o == nil {
		return CreateCustomerCommitAccessSchedule{}
	}
	return o.AccessSchedule
}

func (o *CreateCustomerCommitRequestBody) GetInvoiceSchedule() *CreateCustomerCommitInvoiceSchedule {
	if o == nil {
		return nil
	}
	return o.InvoiceSchedule
}

func (o *CreateCustomerCommitRequestBody) GetInvoiceContractID() *string {
	if o == nil {
		return nil
	}
	return o.InvoiceContractID
}

func (o *CreateCustomerCommitRequestBody) GetApplicableProductIds() []string {
	if o == nil {
		return nil
	}
	return o.ApplicableProductIds
}

func (o *CreateCustomerCommitRequestBody) GetApplicableProductTags() []string {
	if o == nil {
		return nil
	}
	return o.ApplicableProductTags
}

func (o *CreateCustomerCommitRequestBody) GetApplicableContractIds() []string {
	if o == nil {
		return nil
	}
	return o.ApplicableContractIds
}

func (o *CreateCustomerCommitRequestBody) GetNetsuiteSalesOrderID() *string {
	if o == nil {
		return nil
	}
	return o.NetsuiteSalesOrderID
}

func (o *CreateCustomerCommitRequestBody) GetSalesforceOpportunityID() *string {
	if o == nil {
		return nil
	}
	return o.SalesforceOpportunityID
}

func (o *CreateCustomerCommitRequestBody) GetCustomFields() map[string]string {
	if o == nil {
		return nil
	}
	return o.CustomFields
}

type CreateCustomerCommitData struct {
	ID string `json:"id"`
}

func (o *CreateCustomerCommitData) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// CreateCustomerCommitResponseBody - Success
type CreateCustomerCommitResponseBody struct {
	Data CreateCustomerCommitData `json:"data"`
}

func (o *CreateCustomerCommitResponseBody) GetData() CreateCustomerCommitData {
	if o == nil {
		return CreateCustomerCommitData{}
	}
	return o.Data
}

type CreateCustomerCommitResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	Object *CreateCustomerCommitResponseBody
}

func (o *CreateCustomerCommitResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateCustomerCommitResponse) GetObject() *CreateCustomerCommitResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
