/*
 * Account Management Service API
 *
 * Manage user subscriptions and clusters
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package amsclient

import (
	"encoding/json"
)

// SkuRules struct for SkuRules
type SkuRules struct {
	Href    *string `json:"href,omitempty"`
	Id      *string `json:"id,omitempty"`
	Kind    *string `json:"kind,omitempty"`
	Allowed *int32  `json:"allowed,omitempty"`
	QuotaId *string `json:"quota_id,omitempty"`
	Sku     *string `json:"sku,omitempty"`
}

// NewSkuRules instantiates a new SkuRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkuRules() *SkuRules {
	this := SkuRules{}
	return &this
}

// NewSkuRulesWithDefaults instantiates a new SkuRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkuRulesWithDefaults() *SkuRules {
	this := SkuRules{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *SkuRules) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuRules) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *SkuRules) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *SkuRules) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SkuRules) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuRules) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SkuRules) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SkuRules) SetId(v string) {
	o.Id = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *SkuRules) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuRules) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *SkuRules) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *SkuRules) SetKind(v string) {
	o.Kind = &v
}

// GetAllowed returns the Allowed field value if set, zero value otherwise.
func (o *SkuRules) GetAllowed() int32 {
	if o == nil || o.Allowed == nil {
		var ret int32
		return ret
	}
	return *o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuRules) GetAllowedOk() (*int32, bool) {
	if o == nil || o.Allowed == nil {
		return nil, false
	}
	return o.Allowed, true
}

// HasAllowed returns a boolean if a field has been set.
func (o *SkuRules) HasAllowed() bool {
	if o != nil && o.Allowed != nil {
		return true
	}

	return false
}

// SetAllowed gets a reference to the given int32 and assigns it to the Allowed field.
func (o *SkuRules) SetAllowed(v int32) {
	o.Allowed = &v
}

// GetQuotaId returns the QuotaId field value if set, zero value otherwise.
func (o *SkuRules) GetQuotaId() string {
	if o == nil || o.QuotaId == nil {
		var ret string
		return ret
	}
	return *o.QuotaId
}

// GetQuotaIdOk returns a tuple with the QuotaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuRules) GetQuotaIdOk() (*string, bool) {
	if o == nil || o.QuotaId == nil {
		return nil, false
	}
	return o.QuotaId, true
}

// HasQuotaId returns a boolean if a field has been set.
func (o *SkuRules) HasQuotaId() bool {
	if o != nil && o.QuotaId != nil {
		return true
	}

	return false
}

// SetQuotaId gets a reference to the given string and assigns it to the QuotaId field.
func (o *SkuRules) SetQuotaId(v string) {
	o.QuotaId = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *SkuRules) GetSku() string {
	if o == nil || o.Sku == nil {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkuRules) GetSkuOk() (*string, bool) {
	if o == nil || o.Sku == nil {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *SkuRules) HasSku() bool {
	if o != nil && o.Sku != nil {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *SkuRules) SetSku(v string) {
	o.Sku = &v
}

func (o SkuRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Allowed != nil {
		toSerialize["allowed"] = o.Allowed
	}
	if o.QuotaId != nil {
		toSerialize["quota_id"] = o.QuotaId
	}
	if o.Sku != nil {
		toSerialize["sku"] = o.Sku
	}
	return json.Marshal(toSerialize)
}

type NullableSkuRules struct {
	value *SkuRules
	isSet bool
}

func (v NullableSkuRules) Get() *SkuRules {
	return v.value
}

func (v *NullableSkuRules) Set(val *SkuRules) {
	v.value = val
	v.isSet = true
}

func (v NullableSkuRules) IsSet() bool {
	return v.isSet
}

func (v *NullableSkuRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkuRules(val *SkuRules) *NullableSkuRules {
	return &NullableSkuRules{value: val, isSet: true}
}

func (v NullableSkuRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkuRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}