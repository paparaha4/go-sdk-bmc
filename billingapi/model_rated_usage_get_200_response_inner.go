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
	"fmt"
)

// RatedUsageGet200ResponseInner - struct for RatedUsageGet200ResponseInner
type RatedUsageGet200ResponseInner struct {
	BandwidthRecord       *BandwidthRecord
	OperatingSystemRecord *OperatingSystemRecord
	PublicSubnetRecord    *PublicSubnetRecord
	ServerRecord          *ServerRecord
	StorageRecord         *StorageRecord
}

// BandwidthRecordAsRatedUsageGet200ResponseInner is a convenience function that returns BandwidthRecord wrapped in RatedUsageGet200ResponseInner
func BandwidthRecordAsRatedUsageGet200ResponseInner(v *BandwidthRecord) RatedUsageGet200ResponseInner {
	return RatedUsageGet200ResponseInner{
		BandwidthRecord: v,
	}
}

// OperatingSystemRecordAsRatedUsageGet200ResponseInner is a convenience function that returns OperatingSystemRecord wrapped in RatedUsageGet200ResponseInner
func OperatingSystemRecordAsRatedUsageGet200ResponseInner(v *OperatingSystemRecord) RatedUsageGet200ResponseInner {
	return RatedUsageGet200ResponseInner{
		OperatingSystemRecord: v,
	}
}

// PublicSubnetRecordAsRatedUsageGet200ResponseInner is a convenience function that returns PublicSubnetRecord wrapped in RatedUsageGet200ResponseInner
func PublicSubnetRecordAsRatedUsageGet200ResponseInner(v *PublicSubnetRecord) RatedUsageGet200ResponseInner {
	return RatedUsageGet200ResponseInner{
		PublicSubnetRecord: v,
	}
}

// ServerRecordAsRatedUsageGet200ResponseInner is a convenience function that returns ServerRecord wrapped in RatedUsageGet200ResponseInner
func ServerRecordAsRatedUsageGet200ResponseInner(v *ServerRecord) RatedUsageGet200ResponseInner {
	return RatedUsageGet200ResponseInner{
		ServerRecord: v,
	}
}

// StorageRecordAsRatedUsageGet200ResponseInner is a convenience function that returns StorageRecord wrapped in RatedUsageGet200ResponseInner
func StorageRecordAsRatedUsageGet200ResponseInner(v *StorageRecord) RatedUsageGet200ResponseInner {
	return RatedUsageGet200ResponseInner{
		StorageRecord: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *RatedUsageGet200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BandwidthRecord'
	if jsonDict["productCategory"] == "BandwidthRecord" {
		// try to unmarshal JSON data into BandwidthRecord
		err = json.Unmarshal(data, &dst.BandwidthRecord)
		if err == nil {
			return nil // data stored in dst.BandwidthRecord, return on the first match
		} else {
			dst.BandwidthRecord = nil
			return fmt.Errorf("Failed to unmarshal RatedUsageGet200ResponseInner as BandwidthRecord: %s", err.Error())
		}
	}

	// check if the discriminator value is 'OperatingSystemRecord'
	if jsonDict["productCategory"] == "OperatingSystemRecord" {
		// try to unmarshal JSON data into OperatingSystemRecord
		err = json.Unmarshal(data, &dst.OperatingSystemRecord)
		if err == nil {
			return nil // data stored in dst.OperatingSystemRecord, return on the first match
		} else {
			dst.OperatingSystemRecord = nil
			return fmt.Errorf("Failed to unmarshal RatedUsageGet200ResponseInner as OperatingSystemRecord: %s", err.Error())
		}
	}

	// check if the discriminator value is 'PublicSubnetRecord'
	if jsonDict["productCategory"] == "PublicSubnetRecord" {
		// try to unmarshal JSON data into PublicSubnetRecord
		err = json.Unmarshal(data, &dst.PublicSubnetRecord)
		if err == nil {
			return nil // data stored in dst.PublicSubnetRecord, return on the first match
		} else {
			dst.PublicSubnetRecord = nil
			return fmt.Errorf("Failed to unmarshal RatedUsageGet200ResponseInner as PublicSubnetRecord: %s", err.Error())
		}
	}

	// check if the discriminator value is 'ServerRecord'
	if jsonDict["productCategory"] == "ServerRecord" {
		// try to unmarshal JSON data into ServerRecord
		err = json.Unmarshal(data, &dst.ServerRecord)
		if err == nil {
			return nil // data stored in dst.ServerRecord, return on the first match
		} else {
			dst.ServerRecord = nil
			return fmt.Errorf("Failed to unmarshal RatedUsageGet200ResponseInner as ServerRecord: %s", err.Error())
		}
	}

	// check if the discriminator value is 'StorageRecord'
	if jsonDict["productCategory"] == "StorageRecord" {
		// try to unmarshal JSON data into StorageRecord
		err = json.Unmarshal(data, &dst.StorageRecord)
		if err == nil {
			return nil // data stored in dst.StorageRecord, return on the first match
		} else {
			dst.StorageRecord = nil
			return fmt.Errorf("Failed to unmarshal RatedUsageGet200ResponseInner as StorageRecord: %s", err.Error())
		}
	}

	// check if the discriminator value is 'bandwidth'
	if jsonDict["productCategory"] == "bandwidth" {
		// try to unmarshal JSON data into BandwidthRecord
		err = json.Unmarshal(data, &dst.BandwidthRecord)
		if err == nil {
			return nil // data stored in dst.BandwidthRecord, return on the first match
		} else {
			dst.BandwidthRecord = nil
			return fmt.Errorf("Failed to unmarshal RatedUsageGet200ResponseInner as BandwidthRecord: %s", err.Error())
		}
	}

	// check if the discriminator value is 'bmc-server'
	if jsonDict["productCategory"] == "bmc-server" {
		// try to unmarshal JSON data into ServerRecord
		err = json.Unmarshal(data, &dst.ServerRecord)
		if err == nil {
			return nil // data stored in dst.ServerRecord, return on the first match
		} else {
			dst.ServerRecord = nil
			return fmt.Errorf("Failed to unmarshal RatedUsageGet200ResponseInner as ServerRecord: %s", err.Error())
		}
	}

	// check if the discriminator value is 'operating-system'
	if jsonDict["productCategory"] == "operating-system" {
		// try to unmarshal JSON data into OperatingSystemRecord
		err = json.Unmarshal(data, &dst.OperatingSystemRecord)
		if err == nil {
			return nil // data stored in dst.OperatingSystemRecord, return on the first match
		} else {
			dst.OperatingSystemRecord = nil
			return fmt.Errorf("Failed to unmarshal RatedUsageGet200ResponseInner as OperatingSystemRecord: %s", err.Error())
		}
	}

	// check if the discriminator value is 'public-ip'
	if jsonDict["productCategory"] == "public-ip" {
		// try to unmarshal JSON data into PublicSubnetRecord
		err = json.Unmarshal(data, &dst.PublicSubnetRecord)
		if err == nil {
			return nil // data stored in dst.PublicSubnetRecord, return on the first match
		} else {
			dst.PublicSubnetRecord = nil
			return fmt.Errorf("Failed to unmarshal RatedUsageGet200ResponseInner as PublicSubnetRecord: %s", err.Error())
		}
	}

	// check if the discriminator value is 'storage'
	if jsonDict["productCategory"] == "storage" {
		// try to unmarshal JSON data into StorageRecord
		err = json.Unmarshal(data, &dst.StorageRecord)
		if err == nil {
			return nil // data stored in dst.StorageRecord, return on the first match
		} else {
			dst.StorageRecord = nil
			return fmt.Errorf("Failed to unmarshal RatedUsageGet200ResponseInner as StorageRecord: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RatedUsageGet200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.BandwidthRecord != nil {
		return json.Marshal(&src.BandwidthRecord)
	}

	if src.OperatingSystemRecord != nil {
		return json.Marshal(&src.OperatingSystemRecord)
	}

	if src.PublicSubnetRecord != nil {
		return json.Marshal(&src.PublicSubnetRecord)
	}

	if src.ServerRecord != nil {
		return json.Marshal(&src.ServerRecord)
	}

	if src.StorageRecord != nil {
		return json.Marshal(&src.StorageRecord)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RatedUsageGet200ResponseInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.BandwidthRecord != nil {
		return obj.BandwidthRecord
	}

	if obj.OperatingSystemRecord != nil {
		return obj.OperatingSystemRecord
	}

	if obj.PublicSubnetRecord != nil {
		return obj.PublicSubnetRecord
	}

	if obj.ServerRecord != nil {
		return obj.ServerRecord
	}

	if obj.StorageRecord != nil {
		return obj.StorageRecord
	}

	// all schemas are nil
	return nil
}

type NullableRatedUsageGet200ResponseInner struct {
	value *RatedUsageGet200ResponseInner
	isSet bool
}

func (v NullableRatedUsageGet200ResponseInner) Get() *RatedUsageGet200ResponseInner {
	return v.value
}

func (v *NullableRatedUsageGet200ResponseInner) Set(val *RatedUsageGet200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableRatedUsageGet200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableRatedUsageGet200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRatedUsageGet200ResponseInner(val *RatedUsageGet200ResponseInner) *NullableRatedUsageGet200ResponseInner {
	return &NullableRatedUsageGet200ResponseInner{value: val, isSet: true}
}

func (v NullableRatedUsageGet200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRatedUsageGet200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
