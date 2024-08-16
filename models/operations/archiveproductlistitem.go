// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/Metronome-Industries/metronome-go-sdk/models/components"
)

// ArchiveProductListItemRequestBody - Archive a product
type ArchiveProductListItemRequestBody struct {
	// ID of the product to be archived
	ProductID string `json:"product_id"`
}

func (o *ArchiveProductListItemRequestBody) GetProductID() string {
	if o == nil {
		return ""
	}
	return o.ProductID
}

type ArchiveProductListItemData struct {
	ID string `json:"id"`
}

func (o *ArchiveProductListItemData) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

// ArchiveProductListItemResponseBody - Success
type ArchiveProductListItemResponseBody struct {
	Data ArchiveProductListItemData `json:"data"`
}

func (o *ArchiveProductListItemResponseBody) GetData() ArchiveProductListItemData {
	if o == nil {
		return ArchiveProductListItemData{}
	}
	return o.Data
}

type ArchiveProductListItemResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Success
	Object *ArchiveProductListItemResponseBody
}

func (o *ArchiveProductListItemResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ArchiveProductListItemResponse) GetObject() *ArchiveProductListItemResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
