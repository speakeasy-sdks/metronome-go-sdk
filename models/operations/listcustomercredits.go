// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy-sdks/metronome-go-sdk/internal/utils"
	"github.com/speakeasy-sdks/metronome-go-sdk/models/components"
	"time"
)

// ListCustomerCreditsRequestBody - List all credits for a customer
type ListCustomerCreditsRequestBody struct {
	CustomerID string  `json:"customer_id"`
	CreditID   *string `json:"credit_id,omitempty"`
	// Return only credits that have access schedules that "cover" the provided date
	CoveringDate *time.Time `json:"covering_date,omitempty"`
	// Include only credits that have any access on or after the provided date
	StartingAt *time.Time `json:"starting_at,omitempty"`
	// Include only credits that have any access before the provided date (exclusive)
	EffectiveBefore *time.Time `json:"effective_before,omitempty"`
	// Include credits on the contract level.
	IncludeContractCredits *bool `json:"include_contract_credits,omitempty"`
	// Include credits from archived contracts.
	IncludeArchived *bool `json:"include_archived,omitempty"`
	// Include credit ledgers in the response. Setting this flag may cause the query to be slower.
	IncludeLedgers *bool `json:"include_ledgers,omitempty"`
	// The next page token from a previous response.
	NextPage *string `json:"next_page,omitempty"`
}

func (l ListCustomerCreditsRequestBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListCustomerCreditsRequestBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListCustomerCreditsRequestBody) GetCustomerID() string {
	if o == nil {
		return ""
	}
	return o.CustomerID
}

func (o *ListCustomerCreditsRequestBody) GetCreditID() *string {
	if o == nil {
		return nil
	}
	return o.CreditID
}

func (o *ListCustomerCreditsRequestBody) GetCoveringDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.CoveringDate
}

func (o *ListCustomerCreditsRequestBody) GetStartingAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartingAt
}

func (o *ListCustomerCreditsRequestBody) GetEffectiveBefore() *time.Time {
	if o == nil {
		return nil
	}
	return o.EffectiveBefore
}

func (o *ListCustomerCreditsRequestBody) GetIncludeContractCredits() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeContractCredits
}

func (o *ListCustomerCreditsRequestBody) GetIncludeArchived() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeArchived
}

func (o *ListCustomerCreditsRequestBody) GetIncludeLedgers() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeLedgers
}

func (o *ListCustomerCreditsRequestBody) GetNextPage() *string {
	if o == nil {
		return nil
	}
	return o.NextPage
}

type ListCustomerCreditsContract struct {
	ID string `json:"id"`
}

func (o *ListCustomerCreditsContract) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type ListCustomerCreditsType string

const (
	ListCustomerCreditsTypeCredit ListCustomerCreditsType = "CREDIT"
)

func (e ListCustomerCreditsType) ToPointer() *ListCustomerCreditsType {
	return &e
}
func (e *ListCustomerCreditsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CREDIT":
		*e = ListCustomerCreditsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListCustomerCreditsType: %v", v)
	}
}

type ListCustomerCreditsProduct struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (o *ListCustomerCreditsProduct) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ListCustomerCreditsProduct) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type ListCustomerCreditsCreditType struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func (o *ListCustomerCreditsCreditType) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ListCustomerCreditsCreditType) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type ListCustomerCreditsScheduleItems struct {
	ID           string    `json:"id"`
	Amount       float64   `json:"amount"`
	StartingAt   time.Time `json:"starting_at"`
	EndingBefore time.Time `json:"ending_before"`
}

func (l ListCustomerCreditsScheduleItems) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListCustomerCreditsScheduleItems) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListCustomerCreditsScheduleItems) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ListCustomerCreditsScheduleItems) GetAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.Amount
}

func (o *ListCustomerCreditsScheduleItems) GetStartingAt() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.StartingAt
}

func (o *ListCustomerCreditsScheduleItems) GetEndingBefore() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.EndingBefore
}

// ListCustomerCreditsAccessSchedule - The schedule that the customer will gain access to the credits.
type ListCustomerCreditsAccessSchedule struct {
	CreditType    *ListCustomerCreditsCreditType     `json:"credit_type,omitempty"`
	ScheduleItems []ListCustomerCreditsScheduleItems `json:"schedule_items"`
}

func (o *ListCustomerCreditsAccessSchedule) GetCreditType() *ListCustomerCreditsCreditType {
	if o == nil {
		return nil
	}
	return o.CreditType
}

func (o *ListCustomerCreditsAccessSchedule) GetScheduleItems() []ListCustomerCreditsScheduleItems {
	if o == nil {
		return []ListCustomerCreditsScheduleItems{}
	}
	return o.ScheduleItems
}

type ListCustomerCreditsLedgerCustomerCreditsResponse200ApplicationJSONResponseBodyType string

const (
	ListCustomerCreditsLedgerCustomerCreditsResponse200ApplicationJSONResponseBodyTypeCreditManual ListCustomerCreditsLedgerCustomerCreditsResponse200ApplicationJSONResponseBodyType = "CREDIT_MANUAL"
)

func (e ListCustomerCreditsLedgerCustomerCreditsResponse200ApplicationJSONResponseBodyType) ToPointer() *ListCustomerCreditsLedgerCustomerCreditsResponse200ApplicationJSONResponseBodyType {
	return &e
}
func (e *ListCustomerCreditsLedgerCustomerCreditsResponse200ApplicationJSONResponseBodyType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CREDIT_MANUAL":
		*e = ListCustomerCreditsLedgerCustomerCreditsResponse200ApplicationJSONResponseBodyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListCustomerCreditsLedgerCustomerCreditsResponse200ApplicationJSONResponseBodyType: %v", v)
	}
}

type Ledger6 struct {
	Type      ListCustomerCreditsLedgerCustomerCreditsResponse200ApplicationJSONResponseBodyType `json:"type"`
	Timestamp time.Time                                                                          `json:"timestamp"`
	Amount    float64                                                                            `json:"amount"`
	Reason    string                                                                             `json:"reason"`
}

func (l Ledger6) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *Ledger6) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *Ledger6) GetType() ListCustomerCreditsLedgerCustomerCreditsResponse200ApplicationJSONResponseBodyType {
	if o == nil {
		return ListCustomerCreditsLedgerCustomerCreditsResponse200ApplicationJSONResponseBodyType("")
	}
	return o.Type
}

func (o *Ledger6) GetTimestamp() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Timestamp
}

func (o *Ledger6) GetAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.Amount
}

func (o *Ledger6) GetReason() string {
	if o == nil {
		return ""
	}
	return o.Reason
}

type ListCustomerCreditsLedgerCustomerCreditsResponse200ApplicationJSONType string

const (
	ListCustomerCreditsLedgerCustomerCreditsResponse200ApplicationJSONTypeCreditCredited ListCustomerCreditsLedgerCustomerCreditsResponse200ApplicationJSONType = "CREDIT_CREDITED"
)

func (e ListCustomerCreditsLedgerCustomerCreditsResponse200ApplicationJSONType) ToPointer() *ListCustomerCreditsLedgerCustomerCreditsResponse200ApplicationJSONType {
	return &e
}
func (e *ListCustomerCreditsLedgerCustomerCreditsResponse200ApplicationJSONType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CREDIT_CREDITED":
		*e = ListCustomerCreditsLedgerCustomerCreditsResponse200ApplicationJSONType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListCustomerCreditsLedgerCustomerCreditsResponse200ApplicationJSONType: %v", v)
	}
}

type Ledger5 struct {
	Type      ListCustomerCreditsLedgerCustomerCreditsResponse200ApplicationJSONType `json:"type"`
	Timestamp time.Time                                                              `json:"timestamp"`
	Amount    float64                                                                `json:"amount"`
	SegmentID string                                                                 `json:"segment_id"`
	InvoiceID string                                                                 `json:"invoice_id"`
}

func (l Ledger5) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *Ledger5) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *Ledger5) GetType() ListCustomerCreditsLedgerCustomerCreditsResponse200ApplicationJSONType {
	if o == nil {
		return ListCustomerCreditsLedgerCustomerCreditsResponse200ApplicationJSONType("")
	}
	return o.Type
}

func (o *Ledger5) GetTimestamp() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Timestamp
}

func (o *Ledger5) GetAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.Amount
}

func (o *Ledger5) GetSegmentID() string {
	if o == nil {
		return ""
	}
	return o.SegmentID
}

func (o *Ledger5) GetInvoiceID() string {
	if o == nil {
		return ""
	}
	return o.InvoiceID
}

type ListCustomerCreditsLedgerCustomerCreditsResponse200Type string

const (
	ListCustomerCreditsLedgerCustomerCreditsResponse200TypeCreditCanceled ListCustomerCreditsLedgerCustomerCreditsResponse200Type = "CREDIT_CANCELED"
)

func (e ListCustomerCreditsLedgerCustomerCreditsResponse200Type) ToPointer() *ListCustomerCreditsLedgerCustomerCreditsResponse200Type {
	return &e
}
func (e *ListCustomerCreditsLedgerCustomerCreditsResponse200Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CREDIT_CANCELED":
		*e = ListCustomerCreditsLedgerCustomerCreditsResponse200Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListCustomerCreditsLedgerCustomerCreditsResponse200Type: %v", v)
	}
}

type Ledger4 struct {
	Type      ListCustomerCreditsLedgerCustomerCreditsResponse200Type `json:"type"`
	Timestamp time.Time                                               `json:"timestamp"`
	Amount    float64                                                 `json:"amount"`
	SegmentID string                                                  `json:"segment_id"`
	InvoiceID string                                                  `json:"invoice_id"`
}

func (l Ledger4) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *Ledger4) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *Ledger4) GetType() ListCustomerCreditsLedgerCustomerCreditsResponse200Type {
	if o == nil {
		return ListCustomerCreditsLedgerCustomerCreditsResponse200Type("")
	}
	return o.Type
}

func (o *Ledger4) GetTimestamp() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Timestamp
}

func (o *Ledger4) GetAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.Amount
}

func (o *Ledger4) GetSegmentID() string {
	if o == nil {
		return ""
	}
	return o.SegmentID
}

func (o *Ledger4) GetInvoiceID() string {
	if o == nil {
		return ""
	}
	return o.InvoiceID
}

type ListCustomerCreditsLedgerCustomerCreditsResponseType string

const (
	ListCustomerCreditsLedgerCustomerCreditsResponseTypeCreditExpiration ListCustomerCreditsLedgerCustomerCreditsResponseType = "CREDIT_EXPIRATION"
)

func (e ListCustomerCreditsLedgerCustomerCreditsResponseType) ToPointer() *ListCustomerCreditsLedgerCustomerCreditsResponseType {
	return &e
}
func (e *ListCustomerCreditsLedgerCustomerCreditsResponseType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CREDIT_EXPIRATION":
		*e = ListCustomerCreditsLedgerCustomerCreditsResponseType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListCustomerCreditsLedgerCustomerCreditsResponseType: %v", v)
	}
}

type Ledger3 struct {
	Type      ListCustomerCreditsLedgerCustomerCreditsResponseType `json:"type"`
	Timestamp time.Time                                            `json:"timestamp"`
	Amount    float64                                              `json:"amount"`
	SegmentID string                                               `json:"segment_id"`
}

func (l Ledger3) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *Ledger3) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *Ledger3) GetType() ListCustomerCreditsLedgerCustomerCreditsResponseType {
	if o == nil {
		return ListCustomerCreditsLedgerCustomerCreditsResponseType("")
	}
	return o.Type
}

func (o *Ledger3) GetTimestamp() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Timestamp
}

func (o *Ledger3) GetAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.Amount
}

func (o *Ledger3) GetSegmentID() string {
	if o == nil {
		return ""
	}
	return o.SegmentID
}

type ListCustomerCreditsLedgerCustomerCreditsType string

const (
	ListCustomerCreditsLedgerCustomerCreditsTypeCreditAutomatedInvoiceDeduction ListCustomerCreditsLedgerCustomerCreditsType = "CREDIT_AUTOMATED_INVOICE_DEDUCTION"
)

func (e ListCustomerCreditsLedgerCustomerCreditsType) ToPointer() *ListCustomerCreditsLedgerCustomerCreditsType {
	return &e
}
func (e *ListCustomerCreditsLedgerCustomerCreditsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CREDIT_AUTOMATED_INVOICE_DEDUCTION":
		*e = ListCustomerCreditsLedgerCustomerCreditsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListCustomerCreditsLedgerCustomerCreditsType: %v", v)
	}
}

type Ledger2 struct {
	Type      ListCustomerCreditsLedgerCustomerCreditsType `json:"type"`
	Timestamp time.Time                                    `json:"timestamp"`
	Amount    float64                                      `json:"amount"`
	SegmentID string                                       `json:"segment_id"`
	InvoiceID string                                       `json:"invoice_id"`
}

func (l Ledger2) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *Ledger2) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *Ledger2) GetType() ListCustomerCreditsLedgerCustomerCreditsType {
	if o == nil {
		return ListCustomerCreditsLedgerCustomerCreditsType("")
	}
	return o.Type
}

func (o *Ledger2) GetTimestamp() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Timestamp
}

func (o *Ledger2) GetAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.Amount
}

func (o *Ledger2) GetSegmentID() string {
	if o == nil {
		return ""
	}
	return o.SegmentID
}

func (o *Ledger2) GetInvoiceID() string {
	if o == nil {
		return ""
	}
	return o.InvoiceID
}

type ListCustomerCreditsLedgerType string

const (
	ListCustomerCreditsLedgerTypeCreditSegmentStart ListCustomerCreditsLedgerType = "CREDIT_SEGMENT_START"
)

func (e ListCustomerCreditsLedgerType) ToPointer() *ListCustomerCreditsLedgerType {
	return &e
}
func (e *ListCustomerCreditsLedgerType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CREDIT_SEGMENT_START":
		*e = ListCustomerCreditsLedgerType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListCustomerCreditsLedgerType: %v", v)
	}
}

type Ledger1 struct {
	Type      ListCustomerCreditsLedgerType `json:"type"`
	Timestamp time.Time                     `json:"timestamp"`
	Amount    float64                       `json:"amount"`
	SegmentID string                        `json:"segment_id"`
}

func (l Ledger1) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *Ledger1) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *Ledger1) GetType() ListCustomerCreditsLedgerType {
	if o == nil {
		return ListCustomerCreditsLedgerType("")
	}
	return o.Type
}

func (o *Ledger1) GetTimestamp() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.Timestamp
}

func (o *Ledger1) GetAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.Amount
}

func (o *Ledger1) GetSegmentID() string {
	if o == nil {
		return ""
	}
	return o.SegmentID
}

type ListCustomerCreditsLedgerUnionType string

const (
	ListCustomerCreditsLedgerUnionTypeLedger1 ListCustomerCreditsLedgerUnionType = "ledger_1"
	ListCustomerCreditsLedgerUnionTypeLedger2 ListCustomerCreditsLedgerUnionType = "ledger_2"
	ListCustomerCreditsLedgerUnionTypeLedger3 ListCustomerCreditsLedgerUnionType = "ledger_3"
	ListCustomerCreditsLedgerUnionTypeLedger4 ListCustomerCreditsLedgerUnionType = "ledger_4"
	ListCustomerCreditsLedgerUnionTypeLedger5 ListCustomerCreditsLedgerUnionType = "ledger_5"
	ListCustomerCreditsLedgerUnionTypeLedger6 ListCustomerCreditsLedgerUnionType = "ledger_6"
)

type ListCustomerCreditsLedger struct {
	Ledger1 *Ledger1
	Ledger2 *Ledger2
	Ledger3 *Ledger3
	Ledger4 *Ledger4
	Ledger5 *Ledger5
	Ledger6 *Ledger6

	Type ListCustomerCreditsLedgerUnionType
}

func CreateListCustomerCreditsLedgerLedger1(ledger1 Ledger1) ListCustomerCreditsLedger {
	typ := ListCustomerCreditsLedgerUnionTypeLedger1

	return ListCustomerCreditsLedger{
		Ledger1: &ledger1,
		Type:    typ,
	}
}

func CreateListCustomerCreditsLedgerLedger2(ledger2 Ledger2) ListCustomerCreditsLedger {
	typ := ListCustomerCreditsLedgerUnionTypeLedger2

	return ListCustomerCreditsLedger{
		Ledger2: &ledger2,
		Type:    typ,
	}
}

func CreateListCustomerCreditsLedgerLedger3(ledger3 Ledger3) ListCustomerCreditsLedger {
	typ := ListCustomerCreditsLedgerUnionTypeLedger3

	return ListCustomerCreditsLedger{
		Ledger3: &ledger3,
		Type:    typ,
	}
}

func CreateListCustomerCreditsLedgerLedger4(ledger4 Ledger4) ListCustomerCreditsLedger {
	typ := ListCustomerCreditsLedgerUnionTypeLedger4

	return ListCustomerCreditsLedger{
		Ledger4: &ledger4,
		Type:    typ,
	}
}

func CreateListCustomerCreditsLedgerLedger5(ledger5 Ledger5) ListCustomerCreditsLedger {
	typ := ListCustomerCreditsLedgerUnionTypeLedger5

	return ListCustomerCreditsLedger{
		Ledger5: &ledger5,
		Type:    typ,
	}
}

func CreateListCustomerCreditsLedgerLedger6(ledger6 Ledger6) ListCustomerCreditsLedger {
	typ := ListCustomerCreditsLedgerUnionTypeLedger6

	return ListCustomerCreditsLedger{
		Ledger6: &ledger6,
		Type:    typ,
	}
}

func (u *ListCustomerCreditsLedger) UnmarshalJSON(data []byte) error {

	var ledger1 Ledger1 = Ledger1{}
	if err := utils.UnmarshalJSON(data, &ledger1, "", true, true); err == nil {
		u.Ledger1 = &ledger1
		u.Type = ListCustomerCreditsLedgerUnionTypeLedger1
		return nil
	}

	var ledger3 Ledger3 = Ledger3{}
	if err := utils.UnmarshalJSON(data, &ledger3, "", true, true); err == nil {
		u.Ledger3 = &ledger3
		u.Type = ListCustomerCreditsLedgerUnionTypeLedger3
		return nil
	}

	var ledger6 Ledger6 = Ledger6{}
	if err := utils.UnmarshalJSON(data, &ledger6, "", true, true); err == nil {
		u.Ledger6 = &ledger6
		u.Type = ListCustomerCreditsLedgerUnionTypeLedger6
		return nil
	}

	var ledger2 Ledger2 = Ledger2{}
	if err := utils.UnmarshalJSON(data, &ledger2, "", true, true); err == nil {
		u.Ledger2 = &ledger2
		u.Type = ListCustomerCreditsLedgerUnionTypeLedger2
		return nil
	}

	var ledger4 Ledger4 = Ledger4{}
	if err := utils.UnmarshalJSON(data, &ledger4, "", true, true); err == nil {
		u.Ledger4 = &ledger4
		u.Type = ListCustomerCreditsLedgerUnionTypeLedger4
		return nil
	}

	var ledger5 Ledger5 = Ledger5{}
	if err := utils.UnmarshalJSON(data, &ledger5, "", true, true); err == nil {
		u.Ledger5 = &ledger5
		u.Type = ListCustomerCreditsLedgerUnionTypeLedger5
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for ListCustomerCreditsLedger", string(data))
}

func (u ListCustomerCreditsLedger) MarshalJSON() ([]byte, error) {
	if u.Ledger1 != nil {
		return utils.MarshalJSON(u.Ledger1, "", true)
	}

	if u.Ledger2 != nil {
		return utils.MarshalJSON(u.Ledger2, "", true)
	}

	if u.Ledger3 != nil {
		return utils.MarshalJSON(u.Ledger3, "", true)
	}

	if u.Ledger4 != nil {
		return utils.MarshalJSON(u.Ledger4, "", true)
	}

	if u.Ledger5 != nil {
		return utils.MarshalJSON(u.Ledger5, "", true)
	}

	if u.Ledger6 != nil {
		return utils.MarshalJSON(u.Ledger6, "", true)
	}

	return nil, errors.New("could not marshal union type ListCustomerCreditsLedger: all fields are null")
}

type ListCustomerCreditsData struct {
	ID       string                       `json:"id"`
	Contract *ListCustomerCreditsContract `json:"contract,omitempty"`
	Type     ListCustomerCreditsType      `json:"type"`
	Name     *string                      `json:"name,omitempty"`
	// If multiple credits or commits are applicable, the one with the lower priority will apply first.
	Priority *float64                   `json:"priority,omitempty"`
	Product  ListCustomerCreditsProduct `json:"product"`
	// The schedule that the customer will gain access to the credits.
	AccessSchedule        *ListCustomerCreditsAccessSchedule `json:"access_schedule,omitempty"`
	Description           *string                            `json:"description,omitempty"`
	ApplicableProductIds  []string                           `json:"applicable_product_ids,omitempty"`
	ApplicableProductTags []string                           `json:"applicable_product_tags,omitempty"`
	ApplicableContractIds []string                           `json:"applicable_contract_ids,omitempty"`
	// This field's availability is dependent on your client's configuration.
	NetsuiteSalesOrderID *string `json:"netsuite_sales_order_id,omitempty"`
	// This field's availability is dependent on your client's configuration.
	SalesforceOpportunityID *string `json:"salesforce_opportunity_id,omitempty"`
	// A list of ordered events that impact the balance of a credit. For example, an invoice deduction or an expiration.
	Ledger       []ListCustomerCreditsLedger `json:"ledger,omitempty"`
	CustomFields map[string]string           `json:"custom_fields,omitempty"`
}

func (o *ListCustomerCreditsData) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ListCustomerCreditsData) GetContract() *ListCustomerCreditsContract {
	if o == nil {
		return nil
	}
	return o.Contract
}

func (o *ListCustomerCreditsData) GetType() ListCustomerCreditsType {
	if o == nil {
		return ListCustomerCreditsType("")
	}
	return o.Type
}

func (o *ListCustomerCreditsData) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ListCustomerCreditsData) GetPriority() *float64 {
	if o == nil {
		return nil
	}
	return o.Priority
}

func (o *ListCustomerCreditsData) GetProduct() ListCustomerCreditsProduct {
	if o == nil {
		return ListCustomerCreditsProduct{}
	}
	return o.Product
}

func (o *ListCustomerCreditsData) GetAccessSchedule() *ListCustomerCreditsAccessSchedule {
	if o == nil {
		return nil
	}
	return o.AccessSchedule
}

func (o *ListCustomerCreditsData) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *ListCustomerCreditsData) GetApplicableProductIds() []string {
	if o == nil {
		return nil
	}
	return o.ApplicableProductIds
}

func (o *ListCustomerCreditsData) GetApplicableProductTags() []string {
	if o == nil {
		return nil
	}
	return o.ApplicableProductTags
}

func (o *ListCustomerCreditsData) GetApplicableContractIds() []string {
	if o == nil {
		return nil
	}
	return o.ApplicableContractIds
}

func (o *ListCustomerCreditsData) GetNetsuiteSalesOrderID() *string {
	if o == nil {
		return nil
	}
	return o.NetsuiteSalesOrderID
}

func (o *ListCustomerCreditsData) GetSalesforceOpportunityID() *string {
	if o == nil {
		return nil
	}
	return o.SalesforceOpportunityID
}

func (o *ListCustomerCreditsData) GetLedger() []ListCustomerCreditsLedger {
	if o == nil {
		return nil
	}
	return o.Ledger
}

func (o *ListCustomerCreditsData) GetCustomFields() map[string]string {
	if o == nil {
		return nil
	}
	return o.CustomFields
}

// ListCustomerCreditsResponseBody - Success
type ListCustomerCreditsResponseBody struct {
	Data     []ListCustomerCreditsData `json:"data"`
	NextPage *string                   `json:"next_page"`
}

func (o *ListCustomerCreditsResponseBody) GetData() []ListCustomerCreditsData {
	if o == nil {
		return []ListCustomerCreditsData{}
	}
	return o.Data
}

func (o *ListCustomerCreditsResponseBody) GetNextPage() *string {
	if o == nil {
		return nil
	}
	return o.NextPage
}

type ListCustomerCreditsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	Object *ListCustomerCreditsResponseBody

	Next func() (*ListCustomerCreditsResponse, error)
}

func (o *ListCustomerCreditsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ListCustomerCreditsResponse) GetObject() *ListCustomerCreditsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
