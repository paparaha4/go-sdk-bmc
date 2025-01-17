/*
Bare Metal Cloud API

Create, power off, power on, reset, reboot, or shut down your server with the Bare Metal Cloud API.  Deprovision servers, get or edit SSH key details, assign public IPs, assign servers to networks and a lot more.  Manage your infrastructure more efficiently using just a few simple API calls.<br> <br> <span class='pnap-api-knowledge-base-link'> Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/how-to-deploy-bare-metal-cloud-server' target='_blank'>here</a> </span><br> <br> <b>All URLs are relative to (https://api.phoenixnap.com/bmc/v1/)</b>

API version: 0.1
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bmcapi

import (
	"encoding/json"
)

// DeleteResult Result of a successful delete action.
type DeleteResult struct {
	// Server has been deleted.
	Result string `json:"result"`
	// The unique identifier of the server.
	ServerId string `json:"serverId"`
}

// NewDeleteResult instantiates a new DeleteResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteResult(result string, serverId string) *DeleteResult {
	this := DeleteResult{}
	this.Result = result
	this.ServerId = serverId
	return &this
}

// NewDeleteResultWithDefaults instantiates a new DeleteResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteResultWithDefaults() *DeleteResult {
	this := DeleteResult{}
	return &this
}

// GetResult returns the Result field value
func (o *DeleteResult) GetResult() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *DeleteResult) GetResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *DeleteResult) SetResult(v string) {
	o.Result = v
}

// GetServerId returns the ServerId field value
func (o *DeleteResult) GetServerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value
// and a boolean to check if the value has been set.
func (o *DeleteResult) GetServerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerId, true
}

// SetServerId sets field value
func (o *DeleteResult) SetServerId(v string) {
	o.ServerId = v
}

func (o DeleteResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["result"] = o.Result
	}
	if true {
		toSerialize["serverId"] = o.ServerId
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteResult struct {
	value *DeleteResult
	isSet bool
}

func (v NullableDeleteResult) Get() *DeleteResult {
	return v.value
}

func (v *NullableDeleteResult) Set(val *DeleteResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteResult(val *DeleteResult) *NullableDeleteResult {
	return &NullableDeleteResult{value: val, isSet: true}
}

func (v NullableDeleteResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
