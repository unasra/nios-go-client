# RecordAAAA

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AwsRte53RecordInfo** | Pointer to **string** | Aws Route 53 record information. | [optional] 
**CloudInfo** | Pointer to **string** | Structure containing all cloud API related information for this object. | [optional] 
**Comment** | Pointer to **string** | Comment for the record; maximum 256 characters. | [optional] 
**CreationTime** | Pointer to **string** | The time of the record creation in Epoch seconds format. | [optional] 
**Creator** | Pointer to **string** | The record creator. Note that changing creator from or to &#39;SYSTEM&#39; value is not allowed. | [optional] 
**DdnsPrincipal** | Pointer to **string** | The GSS-TSIG principal that owns this record. | [optional] 
**DdnsProtected** | Pointer to **bool** | Determines if the DDNS updates for this record are allowed or not. | [optional] 
**Disable** | Pointer to **bool** | Determines if the record is disabled or not. False means that the record is enabled. | [optional] 
**DiscoveredData** | Pointer to **string** | The discovered data for this AAAA record. | [optional] 
**DnsName** | Pointer to **string** | The name for an AAAA record in punycode format. | [optional] 
**Extattrs** | Pointer to **string** | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ForbidReclamation** | Pointer to **bool** | Determines if the reclamation is allowed for the record or not. | [optional] 
**Ipv6addr** | Pointer to **string** | The IPv6 Address of the record. | [optional] 
**LastQueried** | Pointer to **string** | The time of the last DNS query in Epoch seconds format. | [optional] 
**MsAdUserData** | Pointer to **string** | The Microsoft Active Directory user related information. | [optional] 
**Name** | Pointer to **string** | Name for the AAAA record in FQDN format. This value can be in unicode format. | [optional] 
**Reclaimable** | Pointer to **bool** | Determines if the record is reclaimable or not. | [optional] 
**RemoveAssociatedPtr** | Pointer to **bool** | Delete option that indicates whether the associated PTR records should be removed while deleting the specified A record. | [optional] 
**SharedRecordGroup** | Pointer to **string** | The name of the shared record group in which the record resides. This field exists only on db_objects if this record is a shared record. | [optional] 
**Ttl** | Pointer to **int32** | The Time To Live (TTL) value for the record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached. | [optional] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] 
**View** | Pointer to **string** | The name of the DNS view in which the record resides. Example: external. | [optional] 
**Zone** | Pointer to **string** | The name of the zone in which the record resides. Example: zone.com. If a view is not specified when searching by zone, the default view is used. | [optional] 

## Methods

### NewRecordAAAA

`func NewRecordAAAA() *RecordAAAA`

NewRecordAAAA instantiates a new RecordAAAA object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordAAAAWithDefaults

`func NewRecordAAAAWithDefaults() *RecordAAAA`

NewRecordAAAAWithDefaults instantiates a new RecordAAAA object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *RecordAAAA) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *RecordAAAA) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *RecordAAAA) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *RecordAAAA) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAwsRte53RecordInfo

`func (o *RecordAAAA) GetAwsRte53RecordInfo() string`

GetAwsRte53RecordInfo returns the AwsRte53RecordInfo field if non-nil, zero value otherwise.

### GetAwsRte53RecordInfoOk

`func (o *RecordAAAA) GetAwsRte53RecordInfoOk() (*string, bool)`

GetAwsRte53RecordInfoOk returns a tuple with the AwsRte53RecordInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRte53RecordInfo

`func (o *RecordAAAA) SetAwsRte53RecordInfo(v string)`

SetAwsRte53RecordInfo sets AwsRte53RecordInfo field to given value.

### HasAwsRte53RecordInfo

`func (o *RecordAAAA) HasAwsRte53RecordInfo() bool`

HasAwsRte53RecordInfo returns a boolean if a field has been set.

### GetCloudInfo

`func (o *RecordAAAA) GetCloudInfo() string`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *RecordAAAA) GetCloudInfoOk() (*string, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *RecordAAAA) SetCloudInfo(v string)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *RecordAAAA) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *RecordAAAA) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RecordAAAA) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RecordAAAA) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RecordAAAA) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreationTime

`func (o *RecordAAAA) GetCreationTime() string`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *RecordAAAA) GetCreationTimeOk() (*string, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *RecordAAAA) SetCreationTime(v string)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *RecordAAAA) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetCreator

`func (o *RecordAAAA) GetCreator() string`

GetCreator returns the Creator field if non-nil, zero value otherwise.

### GetCreatorOk

`func (o *RecordAAAA) GetCreatorOk() (*string, bool)`

GetCreatorOk returns a tuple with the Creator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreator

`func (o *RecordAAAA) SetCreator(v string)`

SetCreator sets Creator field to given value.

### HasCreator

`func (o *RecordAAAA) HasCreator() bool`

HasCreator returns a boolean if a field has been set.

### GetDdnsPrincipal

`func (o *RecordAAAA) GetDdnsPrincipal() string`

GetDdnsPrincipal returns the DdnsPrincipal field if non-nil, zero value otherwise.

### GetDdnsPrincipalOk

`func (o *RecordAAAA) GetDdnsPrincipalOk() (*string, bool)`

GetDdnsPrincipalOk returns a tuple with the DdnsPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsPrincipal

`func (o *RecordAAAA) SetDdnsPrincipal(v string)`

SetDdnsPrincipal sets DdnsPrincipal field to given value.

### HasDdnsPrincipal

`func (o *RecordAAAA) HasDdnsPrincipal() bool`

HasDdnsPrincipal returns a boolean if a field has been set.

### GetDdnsProtected

`func (o *RecordAAAA) GetDdnsProtected() bool`

GetDdnsProtected returns the DdnsProtected field if non-nil, zero value otherwise.

### GetDdnsProtectedOk

`func (o *RecordAAAA) GetDdnsProtectedOk() (*bool, bool)`

GetDdnsProtectedOk returns a tuple with the DdnsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsProtected

`func (o *RecordAAAA) SetDdnsProtected(v bool)`

SetDdnsProtected sets DdnsProtected field to given value.

### HasDdnsProtected

`func (o *RecordAAAA) HasDdnsProtected() bool`

HasDdnsProtected returns a boolean if a field has been set.

### GetDisable

`func (o *RecordAAAA) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *RecordAAAA) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *RecordAAAA) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *RecordAAAA) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDiscoveredData

`func (o *RecordAAAA) GetDiscoveredData() string`

GetDiscoveredData returns the DiscoveredData field if non-nil, zero value otherwise.

### GetDiscoveredDataOk

`func (o *RecordAAAA) GetDiscoveredDataOk() (*string, bool)`

GetDiscoveredDataOk returns a tuple with the DiscoveredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredData

`func (o *RecordAAAA) SetDiscoveredData(v string)`

SetDiscoveredData sets DiscoveredData field to given value.

### HasDiscoveredData

`func (o *RecordAAAA) HasDiscoveredData() bool`

HasDiscoveredData returns a boolean if a field has been set.

### GetDnsName

`func (o *RecordAAAA) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *RecordAAAA) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *RecordAAAA) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *RecordAAAA) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetExtattrs

`func (o *RecordAAAA) GetExtattrs() string`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *RecordAAAA) GetExtattrsOk() (*string, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *RecordAAAA) SetExtattrs(v string)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *RecordAAAA) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetForbidReclamation

`func (o *RecordAAAA) GetForbidReclamation() bool`

GetForbidReclamation returns the ForbidReclamation field if non-nil, zero value otherwise.

### GetForbidReclamationOk

`func (o *RecordAAAA) GetForbidReclamationOk() (*bool, bool)`

GetForbidReclamationOk returns a tuple with the ForbidReclamation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbidReclamation

`func (o *RecordAAAA) SetForbidReclamation(v bool)`

SetForbidReclamation sets ForbidReclamation field to given value.

### HasForbidReclamation

`func (o *RecordAAAA) HasForbidReclamation() bool`

HasForbidReclamation returns a boolean if a field has been set.

### GetIpv6addr

`func (o *RecordAAAA) GetIpv6addr() string`

GetIpv6addr returns the Ipv6addr field if non-nil, zero value otherwise.

### GetIpv6addrOk

`func (o *RecordAAAA) GetIpv6addrOk() (*string, bool)`

GetIpv6addrOk returns a tuple with the Ipv6addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6addr

`func (o *RecordAAAA) SetIpv6addr(v string)`

SetIpv6addr sets Ipv6addr field to given value.

### HasIpv6addr

`func (o *RecordAAAA) HasIpv6addr() bool`

HasIpv6addr returns a boolean if a field has been set.

### GetLastQueried

`func (o *RecordAAAA) GetLastQueried() string`

GetLastQueried returns the LastQueried field if non-nil, zero value otherwise.

### GetLastQueriedOk

`func (o *RecordAAAA) GetLastQueriedOk() (*string, bool)`

GetLastQueriedOk returns a tuple with the LastQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueried

`func (o *RecordAAAA) SetLastQueried(v string)`

SetLastQueried sets LastQueried field to given value.

### HasLastQueried

`func (o *RecordAAAA) HasLastQueried() bool`

HasLastQueried returns a boolean if a field has been set.

### GetMsAdUserData

`func (o *RecordAAAA) GetMsAdUserData() string`

GetMsAdUserData returns the MsAdUserData field if non-nil, zero value otherwise.

### GetMsAdUserDataOk

`func (o *RecordAAAA) GetMsAdUserDataOk() (*string, bool)`

GetMsAdUserDataOk returns a tuple with the MsAdUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdUserData

`func (o *RecordAAAA) SetMsAdUserData(v string)`

SetMsAdUserData sets MsAdUserData field to given value.

### HasMsAdUserData

`func (o *RecordAAAA) HasMsAdUserData() bool`

HasMsAdUserData returns a boolean if a field has been set.

### GetName

`func (o *RecordAAAA) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RecordAAAA) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RecordAAAA) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RecordAAAA) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReclaimable

`func (o *RecordAAAA) GetReclaimable() bool`

GetReclaimable returns the Reclaimable field if non-nil, zero value otherwise.

### GetReclaimableOk

`func (o *RecordAAAA) GetReclaimableOk() (*bool, bool)`

GetReclaimableOk returns a tuple with the Reclaimable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReclaimable

`func (o *RecordAAAA) SetReclaimable(v bool)`

SetReclaimable sets Reclaimable field to given value.

### HasReclaimable

`func (o *RecordAAAA) HasReclaimable() bool`

HasReclaimable returns a boolean if a field has been set.

### GetRemoveAssociatedPtr

`func (o *RecordAAAA) GetRemoveAssociatedPtr() bool`

GetRemoveAssociatedPtr returns the RemoveAssociatedPtr field if non-nil, zero value otherwise.

### GetRemoveAssociatedPtrOk

`func (o *RecordAAAA) GetRemoveAssociatedPtrOk() (*bool, bool)`

GetRemoveAssociatedPtrOk returns a tuple with the RemoveAssociatedPtr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveAssociatedPtr

`func (o *RecordAAAA) SetRemoveAssociatedPtr(v bool)`

SetRemoveAssociatedPtr sets RemoveAssociatedPtr field to given value.

### HasRemoveAssociatedPtr

`func (o *RecordAAAA) HasRemoveAssociatedPtr() bool`

HasRemoveAssociatedPtr returns a boolean if a field has been set.

### GetSharedRecordGroup

`func (o *RecordAAAA) GetSharedRecordGroup() string`

GetSharedRecordGroup returns the SharedRecordGroup field if non-nil, zero value otherwise.

### GetSharedRecordGroupOk

`func (o *RecordAAAA) GetSharedRecordGroupOk() (*string, bool)`

GetSharedRecordGroupOk returns a tuple with the SharedRecordGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedRecordGroup

`func (o *RecordAAAA) SetSharedRecordGroup(v string)`

SetSharedRecordGroup sets SharedRecordGroup field to given value.

### HasSharedRecordGroup

`func (o *RecordAAAA) HasSharedRecordGroup() bool`

HasSharedRecordGroup returns a boolean if a field has been set.

### GetTtl

`func (o *RecordAAAA) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *RecordAAAA) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *RecordAAAA) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *RecordAAAA) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseTtl

`func (o *RecordAAAA) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *RecordAAAA) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *RecordAAAA) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *RecordAAAA) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetView

`func (o *RecordAAAA) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *RecordAAAA) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *RecordAAAA) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *RecordAAAA) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *RecordAAAA) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *RecordAAAA) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *RecordAAAA) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *RecordAAAA) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


