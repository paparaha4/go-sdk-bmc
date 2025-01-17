/*
Rancher Solution API

Simplify enterprise-grade Kubernetes cluster operations and management with Rancher on Bare Metal Cloud. Deploy Kubernetes clusters using a few API calls.<br> <br> <span class='pnap-api-knowledge-base-link'> Knowledge base articles to help you can be found <a href='https://phoenixnap.com/kb/rancher-bmc-integration-kubernetes' target='_blank'>here</a> </span><br> <br> <b>All URLs are relative to (https://api.phoenixnap.com/solutions/rancher/v1beta)</b>

API version: 0.1
Contact: support@phoenixnap.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ranchersolutionapi

import (
	"encoding/json"
)

// RancherClusterConfigCertificates struct for RancherClusterConfigCertificates
type RancherClusterConfigCertificates struct {
	// The SSL CA certificate to be used for rancher admin.
	CaCertificate *string `json:"caCertificate,omitempty"`
	// The SSL certificate to be used for rancher admin.
	Certificate *string `json:"certificate,omitempty"`
	// The SSL certificate key to be used for rancher admin.
	CertificateKey *string `json:"certificateKey,omitempty"`
}

// NewRancherClusterConfigCertificates instantiates a new RancherClusterConfigCertificates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRancherClusterConfigCertificates() *RancherClusterConfigCertificates {
	this := RancherClusterConfigCertificates{}
	return &this
}

// NewRancherClusterConfigCertificatesWithDefaults instantiates a new RancherClusterConfigCertificates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRancherClusterConfigCertificatesWithDefaults() *RancherClusterConfigCertificates {
	this := RancherClusterConfigCertificates{}
	return &this
}

// GetCaCertificate returns the CaCertificate field value if set, zero value otherwise.
func (o *RancherClusterConfigCertificates) GetCaCertificate() string {
	if o == nil || o.CaCertificate == nil {
		var ret string
		return ret
	}
	return *o.CaCertificate
}

// GetCaCertificateOk returns a tuple with the CaCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RancherClusterConfigCertificates) GetCaCertificateOk() (*string, bool) {
	if o == nil || o.CaCertificate == nil {
		return nil, false
	}
	return o.CaCertificate, true
}

// HasCaCertificate returns a boolean if a field has been set.
func (o *RancherClusterConfigCertificates) HasCaCertificate() bool {
	if o != nil && o.CaCertificate != nil {
		return true
	}

	return false
}

// SetCaCertificate gets a reference to the given string and assigns it to the CaCertificate field.
func (o *RancherClusterConfigCertificates) SetCaCertificate(v string) {
	o.CaCertificate = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *RancherClusterConfigCertificates) GetCertificate() string {
	if o == nil || o.Certificate == nil {
		var ret string
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RancherClusterConfigCertificates) GetCertificateOk() (*string, bool) {
	if o == nil || o.Certificate == nil {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *RancherClusterConfigCertificates) HasCertificate() bool {
	if o != nil && o.Certificate != nil {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given string and assigns it to the Certificate field.
func (o *RancherClusterConfigCertificates) SetCertificate(v string) {
	o.Certificate = &v
}

// GetCertificateKey returns the CertificateKey field value if set, zero value otherwise.
func (o *RancherClusterConfigCertificates) GetCertificateKey() string {
	if o == nil || o.CertificateKey == nil {
		var ret string
		return ret
	}
	return *o.CertificateKey
}

// GetCertificateKeyOk returns a tuple with the CertificateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RancherClusterConfigCertificates) GetCertificateKeyOk() (*string, bool) {
	if o == nil || o.CertificateKey == nil {
		return nil, false
	}
	return o.CertificateKey, true
}

// HasCertificateKey returns a boolean if a field has been set.
func (o *RancherClusterConfigCertificates) HasCertificateKey() bool {
	if o != nil && o.CertificateKey != nil {
		return true
	}

	return false
}

// SetCertificateKey gets a reference to the given string and assigns it to the CertificateKey field.
func (o *RancherClusterConfigCertificates) SetCertificateKey(v string) {
	o.CertificateKey = &v
}

func (o RancherClusterConfigCertificates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CaCertificate != nil {
		toSerialize["caCertificate"] = o.CaCertificate
	}
	if o.Certificate != nil {
		toSerialize["certificate"] = o.Certificate
	}
	if o.CertificateKey != nil {
		toSerialize["certificateKey"] = o.CertificateKey
	}
	return json.Marshal(toSerialize)
}

type NullableRancherClusterConfigCertificates struct {
	value *RancherClusterConfigCertificates
	isSet bool
}

func (v NullableRancherClusterConfigCertificates) Get() *RancherClusterConfigCertificates {
	return v.value
}

func (v *NullableRancherClusterConfigCertificates) Set(val *RancherClusterConfigCertificates) {
	v.value = val
	v.isSet = true
}

func (v NullableRancherClusterConfigCertificates) IsSet() bool {
	return v.isSet
}

func (v *NullableRancherClusterConfigCertificates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRancherClusterConfigCertificates(val *RancherClusterConfigCertificates) *NullableRancherClusterConfigCertificates {
	return &NullableRancherClusterConfigCertificates{value: val, isSet: true}
}

func (v NullableRancherClusterConfigCertificates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRancherClusterConfigCertificates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
