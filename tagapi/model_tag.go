/*
Tags API

The Tags Manager API. </br></br>**All URLs are relative to (https://api.phoenixnap.com/tag-manager/v1/)**

API version: 1.0
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tagapi

import (
	"encoding/json"
)

// Tag Tag model.
type Tag struct {
	// The unique id of the tag.
	Id string `json:"id"`
	// The name of the tag.
	Name string `json:"name"`
	// The optional values of the tag.
	Values *[]string `json:"values,omitempty"`
	// The description of the tag.
	Description *string `json:"description,omitempty"`
	// Whether or not to show the tag as part of billing and invoices.
	IsBillingTag bool `json:"isBillingTag"`
	// The tag's assigned resources.
	ResourceAssignments *[]ResourceAssignment `json:"resourceAssignments,omitempty"`
	// The tag's creator.
	CreatedBy *string `json:"createdBy,omitempty"`
}

// NewTag instantiates a new Tag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTag(id string, name string, isBillingTag bool) *Tag {
	this := Tag{}
	this.Id = id
	this.Name = name
	this.IsBillingTag = isBillingTag
	var createdBy string = "USER"
	this.CreatedBy = &createdBy
	return &this
}

// NewTagWithDefaults instantiates a new Tag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagWithDefaults() *Tag {
	this := Tag{}
	var createdBy string = "USER"
	this.CreatedBy = &createdBy
	return &this
}

// GetId returns the Id field value
func (o *Tag) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Tag) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Tag) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Tag) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Tag) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Tag) SetName(v string) {
	o.Name = v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *Tag) GetValues() []string {
	if o == nil || o.Values == nil {
		var ret []string
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetValuesOk() (*[]string, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *Tag) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *Tag) SetValues(v []string) {
	o.Values = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Tag) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Tag) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Tag) SetDescription(v string) {
	o.Description = &v
}

// GetIsBillingTag returns the IsBillingTag field value
func (o *Tag) GetIsBillingTag() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsBillingTag
}

// GetIsBillingTagOk returns a tuple with the IsBillingTag field value
// and a boolean to check if the value has been set.
func (o *Tag) GetIsBillingTagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsBillingTag, true
}

// SetIsBillingTag sets field value
func (o *Tag) SetIsBillingTag(v bool) {
	o.IsBillingTag = v
}

// GetResourceAssignments returns the ResourceAssignments field value if set, zero value otherwise.
func (o *Tag) GetResourceAssignments() []ResourceAssignment {
	if o == nil || o.ResourceAssignments == nil {
		var ret []ResourceAssignment
		return ret
	}
	return *o.ResourceAssignments
}

// GetResourceAssignmentsOk returns a tuple with the ResourceAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetResourceAssignmentsOk() (*[]ResourceAssignment, bool) {
	if o == nil || o.ResourceAssignments == nil {
		return nil, false
	}
	return o.ResourceAssignments, true
}

// HasResourceAssignments returns a boolean if a field has been set.
func (o *Tag) HasResourceAssignments() bool {
	if o != nil && o.ResourceAssignments != nil {
		return true
	}

	return false
}

// SetResourceAssignments gets a reference to the given []ResourceAssignment and assigns it to the ResourceAssignments field.
func (o *Tag) SetResourceAssignments(v []ResourceAssignment) {
	o.ResourceAssignments = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Tag) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Tag) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *Tag) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

func (o Tag) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["isBillingTag"] = o.IsBillingTag
	}
	if o.ResourceAssignments != nil {
		toSerialize["resourceAssignments"] = o.ResourceAssignments
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	return json.Marshal(toSerialize)
}

type NullableTag struct {
	value *Tag
	isSet bool
}

func (v NullableTag) Get() *Tag {
	return v.value
}

func (v *NullableTag) Set(val *Tag) {
	v.value = val
	v.isSet = true
}

func (v NullableTag) IsSet() bool {
	return v.isSet
}

func (v *NullableTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTag(val *Tag) *NullableTag {
	return &NullableTag{value: val, isSet: true}
}

func (v NullableTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
