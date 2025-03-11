# RecordAAAARequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** |  | [optional] 
**Creator** | Pointer to **string** |  | [optional] 
**DdnsPrincipal** | Pointer to **string** |  | [optional] 
**DdnsProtected** | Pointer to **string** |  | [optional] 
**Disable** | Pointer to **string** |  | [optional] 
**Extattrs** | Pointer to **string** |  | [optional] 
**ForbidReclamation** | Pointer to **string** |  | [optional] 
**Ipv6addr** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Ttl** | Pointer to **string** |  | [optional] 
**UseTtl** | Pointer to **string** |  | [optional] 
**View** | Pointer to **string** |  | [optional] 

## Methods

### NewRecordAAAARequest

`func NewRecordAAAARequest() *RecordAAAARequest`

NewRecordAAAARequest instantiates a new RecordAAAARequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordAAAARequestWithDefaults

`func NewRecordAAAARequestWithDefaults() *RecordAAAARequest`

NewRecordAAAARequestWithDefaults instantiates a new RecordAAAARequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *RecordAAAARequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RecordAAAARequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RecordAAAARequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RecordAAAARequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreator

`func (o *RecordAAAARequest) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *RecordAAAARequest) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *RecordAAAARequest) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *RecordAAAARequest) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDdnsPrincipal

`func (o *RecordAAAARequest) GetDdnsPrincipal() string`

GetDdnsPrincipal returns the DdnsPrincipal field if non-nil, zero value otherwise.

### GetDdnsPrincipalOk

`func (o *RecordAAAARequest) GetDdnsPrincipalOk() (*string, bool)`

GetDdnsPrincipalOk returns a tuple with the DdnsPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsPrincipal

`func (o *RecordAAAARequest) SetDdnsPrincipal(v string)`

SetDdnsPrincipal sets DdnsPrincipal field to given value.

### HasDdnsPrincipal

`func (o *RecordAAAARequest) HasDdnsPrincipal() bool`

HasDdnsPrincipal returns a boolean if a field has been set.

### GetDdnsProtected

`func (o *RecordAAAARequest) GetDdnsProtected() string`

GetDdnsProtected returns the DdnsProtected field if non-nil, zero value otherwise.

### GetDdnsProtectedOk

`func (o *RecordAAAARequest) GetDdnsProtectedOk() (*string, bool)`

GetDdnsProtectedOk returns a tuple with the DdnsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsProtected

`func (o *RecordAAAARequest) SetDdnsProtected(v string)`

SetDdnsProtected sets DdnsProtected field to given value.

### HasDdnsProtected

`func (o *RecordAAAARequest) HasDdnsProtected() bool`

HasDdnsProtected returns a boolean if a field has been set.

### GetDisable

`func (o *RecordAAAARequest) GetDisable() string`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *RecordAAAARequest) GetDisableOk() (*string, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *RecordAAAARequest) SetDisable(v string)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *RecordAAAARequest) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetExtattrs

`func (o *RecordAAAARequest) GetExtattrs() string`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *RecordAAAARequest) GetExtattrsOk() (*string, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *RecordAAAARequest) SetExtattrs(v string)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *RecordAAAARequest) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetForbidReclamation

`func (o *RecordAAAARequest) GetForbidReclamation() string`

GetForbidReclamation returns the ForbidReclamation field if non-nil, zero value otherwise.

### GetForbidReclamationOk

`func (o *RecordAAAARequest) GetForbidReclamationOk() (*string, bool)`

GetForbidReclamationOk returns a tuple with the ForbidReclamation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbidReclamation

`func (o *RecordAAAARequest) SetForbidReclamation(v string)`

SetForbidReclamation sets ForbidReclamation field to given value.

### HasForbidReclamation

`func (o *RecordAAAARequest) HasForbidReclamation() bool`

HasForbidReclamation returns a boolean if a field has been set.

### GetIpv6addr

`func (o *RecordAAAARequest) GetIpv6addr() string`

GetIpv6addr returns the Ipv6addr field if non-nil, zero value otherwise.

### GetIpv6addrOk

`func (o *RecordAAAARequest) GetIpv6addrOk() (*string, bool)`

GetIpv6addrOk returns a tuple with the Ipv6addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6addr

`func (o *RecordAAAARequest) SetIpv6addr(v string)`

SetIpv6addr sets Ipv6addr field to given value.

### HasIpv6addr

`func (o *RecordAAAARequest) HasIpv6addr() bool`

HasIpv6addr returns a boolean if a field has been set.

### GetName

`func (o *RecordAAAARequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecordAAAARequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecordAAAARequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecordAAAARequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTtl

`func (o *RecordAAAARequest) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *RecordAAAARequest) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *RecordAAAARequest) SetTtl(v string)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *RecordAAAARequest) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *RecordAAAARequest) GetUseTtl() string`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *RecordAAAARequest) GetUseTtlOk() (*string, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *RecordAAAARequest) SetUseTtl(v string)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *RecordAAAARequest) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetView

`func (o *RecordAAAARequest) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *RecordAAAARequest) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *RecordAAAARequest) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *RecordAAAARequest) HasView() bool`

HasView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


