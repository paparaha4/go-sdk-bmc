/*
Billing API

Automate your infrastructure billing with the Bare Metal Cloud Billing API. Reserve your server instances to ensure guaranteed resource availability for 12, 24, and 36 months. Retrieve your server’s rated usage for a given period and enable or disable auto-renewals.<br> <br> <span class='pnap-api-knowledge-base-link'> Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/phoenixnap-bare-metal-cloud-billing-models' target='_blank'>here</a> </span><br> <br> <b>All URLs are relative to (https://api.phoenixnap.com/billing/v1/)</b>

API version: 0.1
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package billingapi

import (
	"encoding/json"
)

// ServerProduct struct for ServerProduct
type ServerProduct struct {
	// The code identifying the product. This code has significant across all locations.
	ProductCode string `json:"productCode"`
	// The product category.
	ProductCategory string `json:"productCategory"`
	// The pricing plans available for this product.
	Plans    []PricingPlan         `json:"plans,omitempty"`
	Metadata ServerProductMetadata `json:"metadata"`
}

// NewServerProduct instantiates a new ServerProduct object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerProduct(productCode string, productCategory string, metadata ServerProductMetadata) *ServerProduct {
	this := ServerProduct{}
	this.ProductCode = productCode
	this.ProductCategory = productCategory
	this.Metadata = metadata
	return &this
}

// NewServerProductWithDefaults instantiates a new ServerProduct object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerProductWithDefaults() *ServerProduct {
	this := ServerProduct{}
	return &this
}

// GetProductCode returns the ProductCode field value
func (o *ServerProduct) GetProductCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductCode
}

// GetProductCodeOk returns a tuple with the ProductCode field value
// and a boolean to check if the value has been set.
func (o *ServerProduct) GetProductCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductCode, true
}

// SetProductCode sets field value
func (o *ServerProduct) SetProductCode(v string) {
	o.ProductCode = v
}

// GetProductCategory returns the ProductCategory field value
func (o *ServerProduct) GetProductCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductCategory
}

// GetProductCategoryOk returns a tuple with the ProductCategory field value
// and a boolean to check if the value has been set.
func (o *ServerProduct) GetProductCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductCategory, true
}

// SetProductCategory sets field value
func (o *ServerProduct) SetProductCategory(v string) {
	o.ProductCategory = v
}

// GetPlans returns the Plans field value if set, zero value otherwise.
func (o *ServerProduct) GetPlans() []PricingPlan {
	if o == nil || o.Plans == nil {
		var ret []PricingPlan
		return ret
	}
	return o.Plans
}

// GetPlansOk returns a tuple with the Plans field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerProduct) GetPlansOk() ([]PricingPlan, bool) {
	if o == nil || o.Plans == nil {
		return nil, false
	}
	return o.Plans, true
}

// HasPlans returns a boolean if a field has been set.
func (o *ServerProduct) HasPlans() bool {
	if o != nil && o.Plans != nil {
		return true
	}

	return false
}

// SetPlans gets a reference to the given []PricingPlan and assigns it to the Plans field.
func (o *ServerProduct) SetPlans(v []PricingPlan) {
	o.Plans = v
}

// GetMetadata returns the Metadata field value
func (o *ServerProduct) GetMetadata() ServerProductMetadata {
	if o == nil {
		var ret ServerProductMetadata
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ServerProduct) GetMetadataOk() (*ServerProductMetadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *ServerProduct) SetMetadata(v ServerProductMetadata) {
	o.Metadata = v
}

func (o ServerProduct) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["productCode"] = o.ProductCode
	}
	if true {
		toSerialize["productCategory"] = o.ProductCategory
	}
	if o.Plans != nil {
		toSerialize["plans"] = o.Plans
	}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableServerProduct struct {
	value *ServerProduct
	isSet bool
}

func (v NullableServerProduct) Get() *ServerProduct {
	return v.value
}

func (v *NullableServerProduct) Set(val *ServerProduct) {
	v.value = val
	v.isSet = true
}

func (v NullableServerProduct) IsSet() bool {
	return v.isSet
}

func (v *NullableServerProduct) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerProduct(val *ServerProduct) *NullableServerProduct {
	return &NullableServerProduct{value: val, isSet: true}
}

func (v NullableServerProduct) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerProduct) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
