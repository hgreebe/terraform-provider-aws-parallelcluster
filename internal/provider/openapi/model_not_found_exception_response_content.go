/*
ParallelCluster

ParallelCluster API

API version: 3.11.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the NotFoundExceptionResponseContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotFoundExceptionResponseContent{}

// NotFoundExceptionResponseContent This exception is thrown when the requested entity is not found.
type NotFoundExceptionResponseContent struct {
	Message *string `json:"message,omitempty"`
}

// NewNotFoundExceptionResponseContent instantiates a new NotFoundExceptionResponseContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotFoundExceptionResponseContent() *NotFoundExceptionResponseContent {
	this := NotFoundExceptionResponseContent{}
	return &this
}

// NewNotFoundExceptionResponseContentWithDefaults instantiates a new NotFoundExceptionResponseContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotFoundExceptionResponseContentWithDefaults() *NotFoundExceptionResponseContent {
	this := NotFoundExceptionResponseContent{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *NotFoundExceptionResponseContent) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotFoundExceptionResponseContent) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *NotFoundExceptionResponseContent) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *NotFoundExceptionResponseContent) SetMessage(v string) {
	o.Message = &v
}

func (o NotFoundExceptionResponseContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotFoundExceptionResponseContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableNotFoundExceptionResponseContent struct {
	value *NotFoundExceptionResponseContent
	isSet bool
}

func (v NullableNotFoundExceptionResponseContent) Get() *NotFoundExceptionResponseContent {
	return v.value
}

func (v *NullableNotFoundExceptionResponseContent) Set(val *NotFoundExceptionResponseContent) {
	v.value = val
	v.isSet = true
}

func (v NullableNotFoundExceptionResponseContent) IsSet() bool {
	return v.isSet
}

func (v *NullableNotFoundExceptionResponseContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotFoundExceptionResponseContent(val *NotFoundExceptionResponseContent) *NullableNotFoundExceptionResponseContent {
	return &NullableNotFoundExceptionResponseContent{value: val, isSet: true}
}

func (v NullableNotFoundExceptionResponseContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotFoundExceptionResponseContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


