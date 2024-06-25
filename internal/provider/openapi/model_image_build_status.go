/*
ParallelCluster

ParallelCluster API

API version: 3.10.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// ImageBuildStatus the model 'ImageBuildStatus'
type ImageBuildStatus string

// List of ImageBuildStatus
const (
	IMAGEBUILDSTATUS_BUILD_IN_PROGRESS ImageBuildStatus = "BUILD_IN_PROGRESS"
	IMAGEBUILDSTATUS_BUILD_FAILED ImageBuildStatus = "BUILD_FAILED"
	IMAGEBUILDSTATUS_BUILD_COMPLETE ImageBuildStatus = "BUILD_COMPLETE"
	IMAGEBUILDSTATUS_DELETE_IN_PROGRESS ImageBuildStatus = "DELETE_IN_PROGRESS"
	IMAGEBUILDSTATUS_DELETE_FAILED ImageBuildStatus = "DELETE_FAILED"
	IMAGEBUILDSTATUS_DELETE_COMPLETE ImageBuildStatus = "DELETE_COMPLETE"
)

// All allowed values of ImageBuildStatus enum
var AllowedImageBuildStatusEnumValues = []ImageBuildStatus{
	"BUILD_IN_PROGRESS",
	"BUILD_FAILED",
	"BUILD_COMPLETE",
	"DELETE_IN_PROGRESS",
	"DELETE_FAILED",
	"DELETE_COMPLETE",
}

func (v *ImageBuildStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ImageBuildStatus(value)
	for _, existing := range AllowedImageBuildStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ImageBuildStatus", value)
}

// NewImageBuildStatusFromValue returns a pointer to a valid ImageBuildStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewImageBuildStatusFromValue(v string) (*ImageBuildStatus, error) {
	ev := ImageBuildStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ImageBuildStatus: valid values are %v", v, AllowedImageBuildStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ImageBuildStatus) IsValid() bool {
	for _, existing := range AllowedImageBuildStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ImageBuildStatus value
func (v ImageBuildStatus) Ptr() *ImageBuildStatus {
	return &v
}

type NullableImageBuildStatus struct {
	value *ImageBuildStatus
	isSet bool
}

func (v NullableImageBuildStatus) Get() *ImageBuildStatus {
	return v.value
}

func (v *NullableImageBuildStatus) Set(val *ImageBuildStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableImageBuildStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableImageBuildStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageBuildStatus(val *ImageBuildStatus) *NullableImageBuildStatus {
	return &NullableImageBuildStatus{value: val, isSet: true}
}

func (v NullableImageBuildStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageBuildStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
