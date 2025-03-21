/*
IbClient

OpenAPI 3.x.x specification for the IbClient API

API version: 3.0.0
Contact: jkhatri@infoblox.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns

import (
	"encoding/json"
)

// checks if the RecordAAAARequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecordAAAARequest{}

// RecordAAAARequest struct for RecordAAAARequest
type RecordAAAARequest struct {
	Comment              *string `json:"comment,omitempty"`
	Creator              *string `json:"creator,omitempty"`
	DdnsPrincipal        *string `json:"ddns_principal,omitempty"`
	DdnsProtected        *string `json:"ddns_protected,omitempty"`
	Disable              *string `json:"disable,omitempty"`
	Extattrs             *string `json:"extattrs,omitempty"`
	ForbidReclamation    *string `json:"forbid_reclamation,omitempty"`
	Ipv6addr             *string `json:"ipv6addr,omitempty"`
	Name                 *string `json:"name,omitempty"`
	Ttl                  *string `json:"ttl,omitempty"`
	UseTtl               *string `json:"use_ttl,omitempty"`
	View                 *string `json:"view,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecordAAAARequest RecordAAAARequest

// NewRecordAAAARequest instantiates a new RecordAAAARequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordAAAARequest() *RecordAAAARequest {
	this := RecordAAAARequest{}
	return &this
}

// NewRecordAAAARequestWithDefaults instantiates a new RecordAAAARequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordAAAARequestWithDefaults() *RecordAAAARequest {
	this := RecordAAAARequest{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *RecordAAAARequest) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordAAAARequest) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *RecordAAAARequest) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *RecordAAAARequest) SetComment(v string) {
	o.Comment = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *RecordAAAARequest) GetCreator() string {
	if o == nil || IsNil(o.Creator) {
		var ret string
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordAAAARequest) GetCreatorOk() (*string, bool) {
	if o == nil || IsNil(o.Creator) {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *RecordAAAARequest) HasCreator() bool {
	if o != nil && !IsNil(o.Creator) {
		return true
	}

	return false
}

// SetCreator gets a reference to the given string and assigns it to the Creator field.
func (o *RecordAAAARequest) SetCreator(v string) {
	o.Creator = &v
}

// GetDdnsPrincipal returns the DdnsPrincipal field value if set, zero value otherwise.
func (o *RecordAAAARequest) GetDdnsPrincipal() string {
	if o == nil || IsNil(o.DdnsPrincipal) {
		var ret string
		return ret
	}
	return *o.DdnsPrincipal
}

// GetDdnsPrincipalOk returns a tuple with the DdnsPrincipal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordAAAARequest) GetDdnsPrincipalOk() (*string, bool) {
	if o == nil || IsNil(o.DdnsPrincipal) {
		return nil, false
	}
	return o.DdnsPrincipal, true
}

// HasDdnsPrincipal returns a boolean if a field has been set.
func (o *RecordAAAARequest) HasDdnsPrincipal() bool {
	if o != nil && !IsNil(o.DdnsPrincipal) {
		return true
	}

	return false
}

// SetDdnsPrincipal gets a reference to the given string and assigns it to the DdnsPrincipal field.
func (o *RecordAAAARequest) SetDdnsPrincipal(v string) {
	o.DdnsPrincipal = &v
}

// GetDdnsProtected returns the DdnsProtected field value if set, zero value otherwise.
func (o *RecordAAAARequest) GetDdnsProtected() string {
	if o == nil || IsNil(o.DdnsProtected) {
		var ret string
		return ret
	}
	return *o.DdnsProtected
}

// GetDdnsProtectedOk returns a tuple with the DdnsProtected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordAAAARequest) GetDdnsProtectedOk() (*string, bool) {
	if o == nil || IsNil(o.DdnsProtected) {
		return nil, false
	}
	return o.DdnsProtected, true
}

// HasDdnsProtected returns a boolean if a field has been set.
func (o *RecordAAAARequest) HasDdnsProtected() bool {
	if o != nil && !IsNil(o.DdnsProtected) {
		return true
	}

	return false
}

// SetDdnsProtected gets a reference to the given string and assigns it to the DdnsProtected field.
func (o *RecordAAAARequest) SetDdnsProtected(v string) {
	o.DdnsProtected = &v
}

// GetDisable returns the Disable field value if set, zero value otherwise.
func (o *RecordAAAARequest) GetDisable() string {
	if o == nil || IsNil(o.Disable) {
		var ret string
		return ret
	}
	return *o.Disable
}

// GetDisableOk returns a tuple with the Disable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordAAAARequest) GetDisableOk() (*string, bool) {
	if o == nil || IsNil(o.Disable) {
		return nil, false
	}
	return o.Disable, true
}

// HasDisable returns a boolean if a field has been set.
func (o *RecordAAAARequest) HasDisable() bool {
	if o != nil && !IsNil(o.Disable) {
		return true
	}

	return false
}

// SetDisable gets a reference to the given string and assigns it to the Disable field.
func (o *RecordAAAARequest) SetDisable(v string) {
	o.Disable = &v
}

// GetExtattrs returns the Extattrs field value if set, zero value otherwise.
func (o *RecordAAAARequest) GetExtattrs() string {
	if o == nil || IsNil(o.Extattrs) {
		var ret string
		return ret
	}
	return *o.Extattrs
}

// GetExtattrsOk returns a tuple with the Extattrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordAAAARequest) GetExtattrsOk() (*string, bool) {
	if o == nil || IsNil(o.Extattrs) {
		return nil, false
	}
	return o.Extattrs, true
}

// HasExtattrs returns a boolean if a field has been set.
func (o *RecordAAAARequest) HasExtattrs() bool {
	if o != nil && !IsNil(o.Extattrs) {
		return true
	}

	return false
}

// SetExtattrs gets a reference to the given string and assigns it to the Extattrs field.
func (o *RecordAAAARequest) SetExtattrs(v string) {
	o.Extattrs = &v
}

// GetForbidReclamation returns the ForbidReclamation field value if set, zero value otherwise.
func (o *RecordAAAARequest) GetForbidReclamation() string {
	if o == nil || IsNil(o.ForbidReclamation) {
		var ret string
		return ret
	}
	return *o.ForbidReclamation
}

// GetForbidReclamationOk returns a tuple with the ForbidReclamation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordAAAARequest) GetForbidReclamationOk() (*string, bool) {
	if o == nil || IsNil(o.ForbidReclamation) {
		return nil, false
	}
	return o.ForbidReclamation, true
}

// HasForbidReclamation returns a boolean if a field has been set.
func (o *RecordAAAARequest) HasForbidReclamation() bool {
	if o != nil && !IsNil(o.ForbidReclamation) {
		return true
	}

	return false
}

// SetForbidReclamation gets a reference to the given string and assigns it to the ForbidReclamation field.
func (o *RecordAAAARequest) SetForbidReclamation(v string) {
	o.ForbidReclamation = &v
}

// GetIpv6addr returns the Ipv6addr field value if set, zero value otherwise.
func (o *RecordAAAARequest) GetIpv6addr() string {
	if o == nil || IsNil(o.Ipv6addr) {
		var ret string
		return ret
	}
	return *o.Ipv6addr
}

// GetIpv6addrOk returns a tuple with the Ipv6addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordAAAARequest) GetIpv6addrOk() (*string, bool) {
	if o == nil || IsNil(o.Ipv6addr) {
		return nil, false
	}
	return o.Ipv6addr, true
}

// HasIpv6addr returns a boolean if a field has been set.
func (o *RecordAAAARequest) HasIpv6addr() bool {
	if o != nil && !IsNil(o.Ipv6addr) {
		return true
	}

	return false
}

// SetIpv6addr gets a reference to the given string and assigns it to the Ipv6addr field.
func (o *RecordAAAARequest) SetIpv6addr(v string) {
	o.Ipv6addr = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RecordAAAARequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordAAAARequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RecordAAAARequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RecordAAAARequest) SetName(v string) {
	o.Name = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *RecordAAAARequest) GetTtl() string {
	if o == nil || IsNil(o.Ttl) {
		var ret string
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordAAAARequest) GetTtlOk() (*string, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *RecordAAAARequest) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given string and assigns it to the Ttl field.
func (o *RecordAAAARequest) SetTtl(v string) {
	o.Ttl = &v
}

// GetUseTtl returns the UseTtl field value if set, zero value otherwise.
func (o *RecordAAAARequest) GetUseTtl() string {
	if o == nil || IsNil(o.UseTtl) {
		var ret string
		return ret
	}
	return *o.UseTtl
}

// GetUseTtlOk returns a tuple with the UseTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordAAAARequest) GetUseTtlOk() (*string, bool) {
	if o == nil || IsNil(o.UseTtl) {
		return nil, false
	}
	return o.UseTtl, true
}

// HasUseTtl returns a boolean if a field has been set.
func (o *RecordAAAARequest) HasUseTtl() bool {
	if o != nil && !IsNil(o.UseTtl) {
		return true
	}

	return false
}

// SetUseTtl gets a reference to the given string and assigns it to the UseTtl field.
func (o *RecordAAAARequest) SetUseTtl(v string) {
	o.UseTtl = &v
}

// GetView returns the View field value if set, zero value otherwise.
func (o *RecordAAAARequest) GetView() string {
	if o == nil || IsNil(o.View) {
		var ret string
		return ret
	}
	return *o.View
}

// GetViewOk returns a tuple with the View field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecordAAAARequest) GetViewOk() (*string, bool) {
	if o == nil || IsNil(o.View) {
		return nil, false
	}
	return o.View, true
}

// HasView returns a boolean if a field has been set.
func (o *RecordAAAARequest) HasView() bool {
	if o != nil && !IsNil(o.View) {
		return true
	}

	return false
}

// SetView gets a reference to the given string and assigns it to the View field.
func (o *RecordAAAARequest) SetView(v string) {
	o.View = &v
}

func (o RecordAAAARequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecordAAAARequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Creator) {
		toSerialize["creator"] = o.Creator
	}
	if !IsNil(o.DdnsPrincipal) {
		toSerialize["ddns_principal"] = o.DdnsPrincipal
	}
	if !IsNil(o.DdnsProtected) {
		toSerialize["ddns_protected"] = o.DdnsProtected
	}
	if !IsNil(o.Disable) {
		toSerialize["disable"] = o.Disable
	}
	if !IsNil(o.Extattrs) {
		toSerialize["extattrs"] = o.Extattrs
	}
	if !IsNil(o.ForbidReclamation) {
		toSerialize["forbid_reclamation"] = o.ForbidReclamation
	}
	if !IsNil(o.Ipv6addr) {
		toSerialize["ipv6addr"] = o.Ipv6addr
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	if !IsNil(o.UseTtl) {
		toSerialize["use_ttl"] = o.UseTtl
	}
	if !IsNil(o.View) {
		toSerialize["view"] = o.View
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RecordAAAARequest) UnmarshalJSON(data []byte) (err error) {
	varRecordAAAARequest := _RecordAAAARequest{}

	err = json.Unmarshal(data, &varRecordAAAARequest)

	if err != nil {
		return err
	}

	*o = RecordAAAARequest(varRecordAAAARequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "comment")
		delete(additionalProperties, "creator")
		delete(additionalProperties, "ddns_principal")
		delete(additionalProperties, "ddns_protected")
		delete(additionalProperties, "disable")
		delete(additionalProperties, "extattrs")
		delete(additionalProperties, "forbid_reclamation")
		delete(additionalProperties, "ipv6addr")
		delete(additionalProperties, "name")
		delete(additionalProperties, "ttl")
		delete(additionalProperties, "use_ttl")
		delete(additionalProperties, "view")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecordAAAARequest struct {
	value *RecordAAAARequest
	isSet bool
}

func (v NullableRecordAAAARequest) Get() *RecordAAAARequest {
	return v.value
}

func (v *NullableRecordAAAARequest) Set(val *RecordAAAARequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordAAAARequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordAAAARequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordAAAARequest(val *RecordAAAARequest) *NullableRecordAAAARequest {
	return &NullableRecordAAAARequest{value: val, isSet: true}
}

func (v NullableRecordAAAARequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordAAAARequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
