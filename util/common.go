// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package ytypes implements YANG type validation logic.
package util

import (
	"reflect"

	"github.com/openconfig/goyang/pkg/yang"
)

var (
	// YangMaxNumber represents the maximum value for any integer type.
	YangMaxNumber = yang.Number{Kind: yang.MaxNumber}
	// YangMinNumber represents the minimum value for any integer type.
	YangMinNumber = yang.Number{Kind: yang.MinNumber}
)

// IsValueNil is a general purpose nil check for the kinds of value types expected in
// this package.
func IsValueNil(value interface{}) bool {
	if value == nil {
		return true
	}
	switch reflect.TypeOf(value).Kind() {
	case reflect.Slice, reflect.Ptr, reflect.Map:
		return reflect.ValueOf(value).IsNil()
	}
	return false
}

// stringMapKeys returns the keys for map m.
func stringMapKeys(m map[string]*yang.Entry) []string {
	var out []string
	for k := range m {
		out = append(out, k)
	}
	return out
}

// stringMapSetToSlice converts a string set expressed as a map m, into a slice
// of strings.
func stringMapSetToSlice(m map[string]interface{}) []string {
	var out []string
	for k := range m {
		out = append(out, k)
	}
	return out
}
