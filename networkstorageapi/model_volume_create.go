/*
Network Storage API

Create, list, edit, and delete storage networks with the Network Storage API. Use storage networks to expand storage capacity on a private network. <br> <span class='pnap-api-knowledge-base-link'> Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/bare-metal-cloud-storage' target='_blank'>here</a> </span> <br> <b>All URLs are relative to (https://api-dev.phoenixnap.com/network-storage/v1/)</b>

API version: 1.0
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package networkstorageapi

import (
	"encoding/json"
)

// VolumeCreate Create Volume.
type VolumeCreate struct {
	// Volume friendly name.
	Name string `json:"name"`
	// Volume description.
	Description *string `json:"description,omitempty"`
	// Last part of volume's path.
	PathSuffix *string `json:"pathSuffix,omitempty"`
	// Capacity of Volume in GB. Currently only whole numbers and multiples of 1000GB are supported.
	CapacityInGb int32 `json:"capacityInGb"`
}

// NewVolumeCreate instantiates a new VolumeCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeCreate(name string, capacityInGb int32) *VolumeCreate {
	this := VolumeCreate{}
	this.Name = name
	this.CapacityInGb = capacityInGb
	return &this
}

// NewVolumeCreateWithDefaults instantiates a new VolumeCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeCreateWithDefaults() *VolumeCreate {
	this := VolumeCreate{}
	return &this
}

// GetName returns the Name field value
func (o *VolumeCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VolumeCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VolumeCreate) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VolumeCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VolumeCreate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VolumeCreate) SetDescription(v string) {
	o.Description = &v
}

// GetPathSuffix returns the PathSuffix field value if set, zero value otherwise.
func (o *VolumeCreate) GetPathSuffix() string {
	if o == nil || o.PathSuffix == nil {
		var ret string
		return ret
	}
	return *o.PathSuffix
}

// GetPathSuffixOk returns a tuple with the PathSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeCreate) GetPathSuffixOk() (*string, bool) {
	if o == nil || o.PathSuffix == nil {
		return nil, false
	}
	return o.PathSuffix, true
}

// HasPathSuffix returns a boolean if a field has been set.
func (o *VolumeCreate) HasPathSuffix() bool {
	if o != nil && o.PathSuffix != nil {
		return true
	}

	return false
}

// SetPathSuffix gets a reference to the given string and assigns it to the PathSuffix field.
func (o *VolumeCreate) SetPathSuffix(v string) {
	o.PathSuffix = &v
}

// GetCapacityInGb returns the CapacityInGb field value
func (o *VolumeCreate) GetCapacityInGb() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CapacityInGb
}

// GetCapacityInGbOk returns a tuple with the CapacityInGb field value
// and a boolean to check if the value has been set.
func (o *VolumeCreate) GetCapacityInGbOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CapacityInGb, true
}

// SetCapacityInGb sets field value
func (o *VolumeCreate) SetCapacityInGb(v int32) {
	o.CapacityInGb = v
}

func (o VolumeCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.PathSuffix != nil {
		toSerialize["pathSuffix"] = o.PathSuffix
	}
	if true {
		toSerialize["capacityInGb"] = o.CapacityInGb
	}
	return json.Marshal(toSerialize)
}

type NullableVolumeCreate struct {
	value *VolumeCreate
	isSet bool
}

func (v NullableVolumeCreate) Get() *VolumeCreate {
	return v.value
}

func (v *NullableVolumeCreate) Set(val *VolumeCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeCreate(val *VolumeCreate) *NullableVolumeCreate {
	return &NullableVolumeCreate{value: val, isSet: true}
}

func (v NullableVolumeCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
