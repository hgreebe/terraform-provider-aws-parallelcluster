/*
ParallelCluster

ParallelCluster API

API version: 3.10.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the LogEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogEvent{}

// LogEvent struct for LogEvent
type LogEvent struct {
	Timestamp time.Time `json:"timestamp"`
	Message string `json:"message"`
}

type _LogEvent LogEvent

// NewLogEvent instantiates a new LogEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogEvent(timestamp time.Time, message string) *LogEvent {
	this := LogEvent{}
	this.Timestamp = timestamp
	this.Message = message
	return &this
}

// NewLogEventWithDefaults instantiates a new LogEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogEventWithDefaults() *LogEvent {
	this := LogEvent{}
	return &this
}

// GetTimestamp returns the Timestamp field value
func (o *LogEvent) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *LogEvent) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *LogEvent) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetMessage returns the Message field value
func (o *LogEvent) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *LogEvent) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *LogEvent) SetMessage(v string) {
	o.Message = v
}

func (o LogEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *LogEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"timestamp",
		"message",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varLogEvent := _LogEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLogEvent)

	if err != nil {
		return err
	}

	*o = LogEvent(varLogEvent)

	return err
}

type NullableLogEvent struct {
	value *LogEvent
	isSet bool
}

func (v NullableLogEvent) Get() *LogEvent {
	return v.value
}

func (v *NullableLogEvent) Set(val *LogEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableLogEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableLogEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogEvent(val *LogEvent) *NullableLogEvent {
	return &NullableLogEvent{value: val, isSet: true}
}

func (v NullableLogEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
