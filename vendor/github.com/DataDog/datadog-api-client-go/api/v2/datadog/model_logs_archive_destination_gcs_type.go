/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// LogsArchiveDestinationGCSType Type of the GCS archive destination.
type LogsArchiveDestinationGCSType string

// List of LogsArchiveDestinationGCSType
const (
	LOGSARCHIVEDESTINATIONGCSTYPE_GCS LogsArchiveDestinationGCSType = "gcs"
)

func (v *LogsArchiveDestinationGCSType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LogsArchiveDestinationGCSType(value)
	for _, existing := range []LogsArchiveDestinationGCSType{"gcs"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LogsArchiveDestinationGCSType", *v)
}

// Ptr returns reference to LogsArchiveDestinationGCSType value
func (v LogsArchiveDestinationGCSType) Ptr() *LogsArchiveDestinationGCSType {
	return &v
}

type NullableLogsArchiveDestinationGCSType struct {
	value *LogsArchiveDestinationGCSType
	isSet bool
}

func (v NullableLogsArchiveDestinationGCSType) Get() *LogsArchiveDestinationGCSType {
	return v.value
}

func (v *NullableLogsArchiveDestinationGCSType) Set(val *LogsArchiveDestinationGCSType) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsArchiveDestinationGCSType) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsArchiveDestinationGCSType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsArchiveDestinationGCSType(val *LogsArchiveDestinationGCSType) *NullableLogsArchiveDestinationGCSType {
	return &NullableLogsArchiveDestinationGCSType{value: val, isSet: true}
}

func (v NullableLogsArchiveDestinationGCSType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsArchiveDestinationGCSType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}