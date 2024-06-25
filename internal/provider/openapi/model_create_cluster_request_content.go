/*
ParallelCluster

ParallelCluster API

API version: 3.10.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the CreateClusterRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateClusterRequestContent{}

// CreateClusterRequestContent struct for CreateClusterRequestContent
type CreateClusterRequestContent struct {
	// Name of the cluster that will be created.
	ClusterName string `json:"clusterName"`
	// Cluster configuration as a YAML document.
	ClusterConfiguration string `json:"clusterConfiguration"`
}

type _CreateClusterRequestContent CreateClusterRequestContent

// NewCreateClusterRequestContent instantiates a new CreateClusterRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateClusterRequestContent(clusterName string, clusterConfiguration string) *CreateClusterRequestContent {
	this := CreateClusterRequestContent{}
	this.ClusterName = clusterName
	this.ClusterConfiguration = clusterConfiguration
	return &this
}

// NewCreateClusterRequestContentWithDefaults instantiates a new CreateClusterRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateClusterRequestContentWithDefaults() *CreateClusterRequestContent {
	this := CreateClusterRequestContent{}
	return &this
}

// GetClusterName returns the ClusterName field value
func (o *CreateClusterRequestContent) GetClusterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value
// and a boolean to check if the value has been set.
func (o *CreateClusterRequestContent) GetClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterName, true
}

// SetClusterName sets field value
func (o *CreateClusterRequestContent) SetClusterName(v string) {
	o.ClusterName = v
}

// GetClusterConfiguration returns the ClusterConfiguration field value
func (o *CreateClusterRequestContent) GetClusterConfiguration() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterConfiguration
}

// GetClusterConfigurationOk returns a tuple with the ClusterConfiguration field value
// and a boolean to check if the value has been set.
func (o *CreateClusterRequestContent) GetClusterConfigurationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterConfiguration, true
}

// SetClusterConfiguration sets field value
func (o *CreateClusterRequestContent) SetClusterConfiguration(v string) {
	o.ClusterConfiguration = v
}

func (o CreateClusterRequestContent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateClusterRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clusterName"] = o.ClusterName
	toSerialize["clusterConfiguration"] = o.ClusterConfiguration
	return toSerialize, nil
}

func (o *CreateClusterRequestContent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"clusterName",
		"clusterConfiguration",
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

	varCreateClusterRequestContent := _CreateClusterRequestContent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateClusterRequestContent)

	if err != nil {
		return err
	}

	*o = CreateClusterRequestContent(varCreateClusterRequestContent)

	return err
}

type NullableCreateClusterRequestContent struct {
	value *CreateClusterRequestContent
	isSet bool
}

func (v NullableCreateClusterRequestContent) Get() *CreateClusterRequestContent {
	return v.value
}

func (v *NullableCreateClusterRequestContent) Set(val *CreateClusterRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateClusterRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateClusterRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateClusterRequestContent(val *CreateClusterRequestContent) *NullableCreateClusterRequestContent {
	return &NullableCreateClusterRequestContent{value: val, isSet: true}
}

func (v NullableCreateClusterRequestContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateClusterRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
