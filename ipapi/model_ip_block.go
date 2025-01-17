/*
IP Addresses API

Public IP blocks are a set of contiguous IPs that allow you to access your servers or networks from the internet. Use the IP Addresses API to request and delete IP blocks.<br> <br> <span class='pnap-api-knowledge-base-link'> Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/public-ip-management#bmc-public-ip-allocations-api' target='_blank'>here</a> </span><br> <br> <b>All URLs are relative to (https://api.phoenixnap.com/ips/v1/)</b>

API version: 1.0
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipapi

import (
	"encoding/json"
	"time"
)

// IpBlock IP Block Details.
type IpBlock struct {
	// IP Block identifier.
	Id string `json:"id"`
	// IP Block location ID. Currently this field should be set to `PHX`, `ASH`, `SGP`, `NLD`, `CHI`, `SEA` or `AUS`.
	Location string `json:"location"`
	// CIDR IP Block Size. Currently this field should be set to either `/31`, `/30`, `/29`, `/28`, `/27`, `/26`, `/25`, `/24`, `/23` or `/22`.
	CidrBlockSize string `json:"cidrBlockSize"`
	// The IP Block in CIDR notation.
	Cidr string `json:"cidr"`
	// The status of the IP Block.
	Status string `json:"status"`
	// ID of the resource assigned to the IP Block.
	AssignedResourceId *string `json:"assignedResourceId,omitempty"`
	// Type of the resource assigned to the IP Block.
	AssignedResourceType *string `json:"assignedResourceType,omitempty"`
	// The description of the IP Block.
	Description *string `json:"description,omitempty"`
	// The tags assigned if any.
	Tags []TagAssignment `json:"tags,omitempty"`
	// True if the IP block is a `bring your own` block.
	IsBringYourOwn bool `json:"isBringYourOwn"`
	// Date and time when the IP block was created.
	CreatedOn time.Time `json:"createdOn"`
}

// NewIpBlock instantiates a new IpBlock object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpBlock(id string, location string, cidrBlockSize string, cidr string, status string, isBringYourOwn bool, createdOn time.Time) *IpBlock {
	this := IpBlock{}
	this.Id = id
	this.Location = location
	this.CidrBlockSize = cidrBlockSize
	this.Cidr = cidr
	this.Status = status
	this.IsBringYourOwn = isBringYourOwn
	this.CreatedOn = createdOn
	return &this
}

// NewIpBlockWithDefaults instantiates a new IpBlock object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpBlockWithDefaults() *IpBlock {
	this := IpBlock{}
	return &this
}

// GetId returns the Id field value
func (o *IpBlock) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IpBlock) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IpBlock) SetId(v string) {
	o.Id = v
}

// GetLocation returns the Location field value
func (o *IpBlock) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *IpBlock) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *IpBlock) SetLocation(v string) {
	o.Location = v
}

// GetCidrBlockSize returns the CidrBlockSize field value
func (o *IpBlock) GetCidrBlockSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CidrBlockSize
}

// GetCidrBlockSizeOk returns a tuple with the CidrBlockSize field value
// and a boolean to check if the value has been set.
func (o *IpBlock) GetCidrBlockSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CidrBlockSize, true
}

// SetCidrBlockSize sets field value
func (o *IpBlock) SetCidrBlockSize(v string) {
	o.CidrBlockSize = v
}

// GetCidr returns the Cidr field value
func (o *IpBlock) GetCidr() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cidr
}

// GetCidrOk returns a tuple with the Cidr field value
// and a boolean to check if the value has been set.
func (o *IpBlock) GetCidrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cidr, true
}

// SetCidr sets field value
func (o *IpBlock) SetCidr(v string) {
	o.Cidr = v
}

// GetStatus returns the Status field value
func (o *IpBlock) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *IpBlock) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *IpBlock) SetStatus(v string) {
	o.Status = v
}

// GetAssignedResourceId returns the AssignedResourceId field value if set, zero value otherwise.
func (o *IpBlock) GetAssignedResourceId() string {
	if o == nil || o.AssignedResourceId == nil {
		var ret string
		return ret
	}
	return *o.AssignedResourceId
}

// GetAssignedResourceIdOk returns a tuple with the AssignedResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpBlock) GetAssignedResourceIdOk() (*string, bool) {
	if o == nil || o.AssignedResourceId == nil {
		return nil, false
	}
	return o.AssignedResourceId, true
}

// HasAssignedResourceId returns a boolean if a field has been set.
func (o *IpBlock) HasAssignedResourceId() bool {
	if o != nil && o.AssignedResourceId != nil {
		return true
	}

	return false
}

// SetAssignedResourceId gets a reference to the given string and assigns it to the AssignedResourceId field.
func (o *IpBlock) SetAssignedResourceId(v string) {
	o.AssignedResourceId = &v
}

// GetAssignedResourceType returns the AssignedResourceType field value if set, zero value otherwise.
func (o *IpBlock) GetAssignedResourceType() string {
	if o == nil || o.AssignedResourceType == nil {
		var ret string
		return ret
	}
	return *o.AssignedResourceType
}

// GetAssignedResourceTypeOk returns a tuple with the AssignedResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpBlock) GetAssignedResourceTypeOk() (*string, bool) {
	if o == nil || o.AssignedResourceType == nil {
		return nil, false
	}
	return o.AssignedResourceType, true
}

// HasAssignedResourceType returns a boolean if a field has been set.
func (o *IpBlock) HasAssignedResourceType() bool {
	if o != nil && o.AssignedResourceType != nil {
		return true
	}

	return false
}

// SetAssignedResourceType gets a reference to the given string and assigns it to the AssignedResourceType field.
func (o *IpBlock) SetAssignedResourceType(v string) {
	o.AssignedResourceType = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IpBlock) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpBlock) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IpBlock) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IpBlock) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *IpBlock) GetTags() []TagAssignment {
	if o == nil || o.Tags == nil {
		var ret []TagAssignment
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpBlock) GetTagsOk() ([]TagAssignment, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *IpBlock) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagAssignment and assigns it to the Tags field.
func (o *IpBlock) SetTags(v []TagAssignment) {
	o.Tags = v
}

// GetIsBringYourOwn returns the IsBringYourOwn field value
func (o *IpBlock) GetIsBringYourOwn() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsBringYourOwn
}

// GetIsBringYourOwnOk returns a tuple with the IsBringYourOwn field value
// and a boolean to check if the value has been set.
func (o *IpBlock) GetIsBringYourOwnOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsBringYourOwn, true
}

// SetIsBringYourOwn sets field value
func (o *IpBlock) SetIsBringYourOwn(v bool) {
	o.IsBringYourOwn = v
}

// GetCreatedOn returns the CreatedOn field value
func (o *IpBlock) GetCreatedOn() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedOn
}

// GetCreatedOnOk returns a tuple with the CreatedOn field value
// and a boolean to check if the value has been set.
func (o *IpBlock) GetCreatedOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedOn, true
}

// SetCreatedOn sets field value
func (o *IpBlock) SetCreatedOn(v time.Time) {
	o.CreatedOn = v
}

func (o IpBlock) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["location"] = o.Location
	}
	if true {
		toSerialize["cidrBlockSize"] = o.CidrBlockSize
	}
	if true {
		toSerialize["cidr"] = o.Cidr
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.AssignedResourceId != nil {
		toSerialize["assignedResourceId"] = o.AssignedResourceId
	}
	if o.AssignedResourceType != nil {
		toSerialize["assignedResourceType"] = o.AssignedResourceType
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["isBringYourOwn"] = o.IsBringYourOwn
	}
	if true {
		toSerialize["createdOn"] = o.CreatedOn
	}
	return json.Marshal(toSerialize)
}

type NullableIpBlock struct {
	value *IpBlock
	isSet bool
}

func (v NullableIpBlock) Get() *IpBlock {
	return v.value
}

func (v *NullableIpBlock) Set(val *IpBlock) {
	v.value = val
	v.isSet = true
}

func (v NullableIpBlock) IsSet() bool {
	return v.isSet
}

func (v *NullableIpBlock) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpBlock(val *IpBlock) *NullableIpBlock {
	return &NullableIpBlock{value: val, isSet: true}
}

func (v NullableIpBlock) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpBlock) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
