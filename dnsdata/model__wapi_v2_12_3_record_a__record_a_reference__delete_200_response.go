/*
Infoblox WAPI Swagger for DNS ( wapi version - 2.12.3 )

OpenAPI 3.x.x specification for the IbClient API

API version: 2.12.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dnsdata

import (
	"encoding/json"
)

// checks if the WapiV2123RecordARecordAReferenceDelete200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WapiV2123RecordARecordAReferenceDelete200Response{}

// WapiV2123RecordARecordAReferenceDelete200Response struct for WapiV2123RecordARecordAReferenceDelete200Response
type WapiV2123RecordARecordAReferenceDelete200Response struct {
	// Reference of the deleted record:a
	Ref                  *string `json:"ref,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WapiV2123RecordARecordAReferenceDelete200Response WapiV2123RecordARecordAReferenceDelete200Response

// NewWapiV2123RecordARecordAReferenceDelete200Response instantiates a new WapiV2123RecordARecordAReferenceDelete200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWapiV2123RecordARecordAReferenceDelete200Response() *WapiV2123RecordARecordAReferenceDelete200Response {
	this := WapiV2123RecordARecordAReferenceDelete200Response{}
	return &this
}

// NewWapiV2123RecordARecordAReferenceDelete200ResponseWithDefaults instantiates a new WapiV2123RecordARecordAReferenceDelete200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWapiV2123RecordARecordAReferenceDelete200ResponseWithDefaults() *WapiV2123RecordARecordAReferenceDelete200Response {
	this := WapiV2123RecordARecordAReferenceDelete200Response{}
	return &this
}

// GetRef returns the Ref field value if set, zero value otherwise.
func (o *WapiV2123RecordARecordAReferenceDelete200Response) GetRef() string {
	if o == nil || IsNil(o.Ref) {
		var ret string
		return ret
	}
	return *o.Ref
}

// GetRefOk returns a tuple with the Ref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WapiV2123RecordARecordAReferenceDelete200Response) GetRefOk() (*string, bool) {
	if o == nil || IsNil(o.Ref) {
		return nil, false
	}
	return o.Ref, true
}

// HasRef returns a boolean if a field has been set.
func (o *WapiV2123RecordARecordAReferenceDelete200Response) HasRef() bool {
	if o != nil && !IsNil(o.Ref) {
		return true
	}

	return false
}

// SetRef gets a reference to the given string and assigns it to the Ref field.
func (o *WapiV2123RecordARecordAReferenceDelete200Response) SetRef(v string) {
	o.Ref = &v
}

func (o WapiV2123RecordARecordAReferenceDelete200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WapiV2123RecordARecordAReferenceDelete200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ref) {
		toSerialize["ref"] = o.Ref
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WapiV2123RecordARecordAReferenceDelete200Response) UnmarshalJSON(data []byte) (err error) {
	varWapiV2123RecordARecordAReferenceDelete200Response := _WapiV2123RecordARecordAReferenceDelete200Response{}

	err = json.Unmarshal(data, &varWapiV2123RecordARecordAReferenceDelete200Response)

	if err != nil {
		return err
	}

	*o = WapiV2123RecordARecordAReferenceDelete200Response(varWapiV2123RecordARecordAReferenceDelete200Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ref")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWapiV2123RecordARecordAReferenceDelete200Response struct {
	value *WapiV2123RecordARecordAReferenceDelete200Response
	isSet bool
}

func (v NullableWapiV2123RecordARecordAReferenceDelete200Response) Get() *WapiV2123RecordARecordAReferenceDelete200Response {
	return v.value
}

func (v *NullableWapiV2123RecordARecordAReferenceDelete200Response) Set(val *WapiV2123RecordARecordAReferenceDelete200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableWapiV2123RecordARecordAReferenceDelete200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableWapiV2123RecordARecordAReferenceDelete200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWapiV2123RecordARecordAReferenceDelete200Response(val *WapiV2123RecordARecordAReferenceDelete200Response) *NullableWapiV2123RecordARecordAReferenceDelete200Response {
	return &NullableWapiV2123RecordARecordAReferenceDelete200Response{value: val, isSet: true}
}

func (v NullableWapiV2123RecordARecordAReferenceDelete200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWapiV2123RecordARecordAReferenceDelete200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
