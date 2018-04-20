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

package testutil

import (
	"testing"

	gnmipb "github.com/openconfig/gnmi/proto/gnmi"
)

func TestPathLess(t *testing.T) {
	tests := []struct {
		name string
		inA  *gnmipb.Path
		inB  *gnmipb.Path
		want bool
	}{{
		name: "equal - a < b",
		inA: &gnmipb.Path{
			Elem: []*gnmipb.PathElem{{
				Name: "one",
			}},
		},
		inB: &gnmipb.Path{
			Elem: []*gnmipb.PathElem{{
				Name: "one",
			}},
		},
		want: true,
	}, {
		name: "a < b due to path element name",
		inA: &gnmipb.Path{
			Elem: []*gnmipb.PathElem{{
				Name: "a",
			}},
		},
		inB: &gnmipb.Path{
			Elem: []*gnmipb.PathElem{{
				Name: "b",
			}},
		},
		want: true,
	}, {
		name: "b < a due to path element name",
		inA: &gnmipb.Path{
			Elem: []*gnmipb.PathElem{{
				Name: "b",
			}},
		},
		inB: &gnmipb.Path{
			Elem: []*gnmipb.PathElem{{
				Name: "a",
			}},
		},
		want: false,
	}, {
		name: "equal: a < b with path elem keys",
		inA: &gnmipb.Path{
			Elem: []*gnmipb.PathElem{{
				Name: "a",
				Key:  map[string]string{"a": "a"},
			}},
		},
		inB: &gnmipb.Path{
			Elem: []*gnmipb.PathElem{{
				Name: "a",
				Key:  map[string]string{"a": "a"},
			}},
		},
		want: true,
	}, {
		name: "a < b due to path elem key name",
		inA: &gnmipb.Path{
			Elem: []*gnmipb.PathElem{{
				Name: "a",
				Key:  map[string]string{"a": "a"},
			}},
		},
		inB: &gnmipb.Path{
			Elem: []*gnmipb.PathElem{{
				Name: "a",
				Key:  map[string]string{"b": "a"},
			}},
		},
		want: true,
	}, {
		name: "b < a due to path elem key name",
		inA: &gnmipb.Path{
			Elem: []*gnmipb.PathElem{{
				Name: "a",
				Key:  map[string]string{"b": "a"},
			}},
		},
		inB: &gnmipb.Path{
			Elem: []*gnmipb.PathElem{{
				Name: "a",
				Key:  map[string]string{"a": "a"},
			}},
		},
		want: false,
	}, {
		name: "a < b due to path elem key value",
		inA: &gnmipb.Path{
			Elem: []*gnmipb.PathElem{{
				Name: "a",
				Key:  map[string]string{"a": "a"},
			}},
		},
		inB: &gnmipb.Path{
			Elem: []*gnmipb.PathElem{{
				Name: "a",
				Key:  map[string]string{"a": "z"},
			}},
		},
		want: true,
	}, {
		name: "b < a due to path elem key value",
		inA: &gnmipb.Path{
			Elem: []*gnmipb.PathElem{{
				Name: "a",
				Key:  map[string]string{"a": "z"},
			}},
		},
		inB: &gnmipb.Path{
			Elem: []*gnmipb.PathElem{{
				Name: "a",
				Key:  map[string]string{"a": "a"},
			}},
		},
		want: false,
	}, {
		name: "a < b due to more specific path",
		inA: &gnmipb.Path{
			Elem: []*gnmipb.PathElem{{
				Name: "a",
			}},
		},
		inB: &gnmipb.Path{
			Elem: []*gnmipb.PathElem{{
				Name: "a",
			}, {
				Name: "b",
			}},
		},
		want: true,
	}, {
		name: "b < a due to more specific path",
		inA: &gnmipb.Path{
			Elem: []*gnmipb.PathElem{{
				Name: "a",
			}, {
				Name: "b",
			}},
		},
		inB: &gnmipb.Path{
			Elem: []*gnmipb.PathElem{{
				Name: "a",
			}},
		},
		want: false,
	}, {
		name: "a < b due to number of keys",
		inA: &gnmipb.Path{
			Elem: []*gnmipb.PathElem{{
				Name: "a",
				Key:  map[string]string{"one": "1"},
			}},
		},
		inB: &gnmipb.Path{
			Elem: []*gnmipb.PathElem{{
				Name: "a",
				Key:  map[string]string{"one": "1", "two": "2"},
			}},
		},
		want: true,
	}, {
		name: "b < a due to number of keys",
		inA: &gnmipb.Path{
			Elem: []*gnmipb.PathElem{{
				Name: "a",
				Key:  map[string]string{"one": "1", "two": "2"},
			}},
		},
		inB: &gnmipb.Path{
			Elem: []*gnmipb.PathElem{{
				Name: "a",
				Key:  map[string]string{"one": "1"},
			}},
		},
		want: false,
	}, {
		name: "equal - a < b with origin",
		inA: &gnmipb.Path{
			Elem: []*gnmipb.PathElem{{
				Name: "a",
			}, {
				Name: "b",
			}},
			Origin: "a",
		},
		inB: &gnmipb.Path{
			Elem: []*gnmipb.PathElem{{
				Name: "a",
			}, {
				Name: "b",
			}},
			Origin: "a",
		},
		want: true,
	}, {
		name: "a < b due to origin",
		inA: &gnmipb.Path{
			Elem: []*gnmipb.PathElem{{
				Name: "a",
			}, {
				Name: "b",
			}},
			Origin: "a",
		},
		inB: &gnmipb.Path{
			Elem: []*gnmipb.PathElem{{
				Name: "a",
			}, {
				Name: "b",
			}},
			Origin: "z",
		},
		want: true,
	}, {
		name: "b < a due to origin",
		inA: &gnmipb.Path{
			Elem: []*gnmipb.PathElem{{
				Name: "a",
			}, {
				Name: "b",
			}},
			Origin: "z",
		},
		inB: &gnmipb.Path{
			Elem: []*gnmipb.PathElem{{
				Name: "a",
			}, {
				Name: "b",
			}},
			Origin: "a",
		},
		want: false,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathLess(tt.inA, tt.inB); got != tt.want {
				t.Fatalf("PathLess(%#v, %#v): did not get expected result, got: %v, want: %v", tt.inA, tt.inB, got, tt.want)
			}
		})
	}
}

func TestTypedValueLess(t *testing.T) {
	tests := []struct {
		name string
		inA  *gnmipb.TypedValue
		inB  *gnmipb.TypedValue
		want bool
	}{{
		name: "different types: a < b",
		inA: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_UintVal{42},
		},
		inB: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_StringVal{"ab"},
		},
		want: true,
	}, {
		name: "different types: b < a",
		inA: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_StringVal{"zzxx"},
		},
		inB: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_IntVal{42},
		},
		want: false,
	}, {
		name: "different types: a < b",
		inA: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_DecimalVal{&gnmipb.Decimal64{
				Digits:    1234,
				Precision: 4,
			}},
		},
		inB: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_StringVal{"forty-two"},
		},
		want: true,
	}, {
		name: "different types: b < a",
		inA: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_StringVal{"forty-two"},
		},
		inB: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_DecimalVal{&gnmipb.Decimal64{
				Digits:    1234,
				Precision: 4,
			}},
		},
		want: false,
	}, {
		name: "a and b nil: a < b",
		want: true,
	}, {
		name: "a nil, b non-nil: b < a",
		inB:  &gnmipb.TypedValue{},
		want: false,
	}, {
		name: "a non-nil, b nil: a < b",
		inA:  &gnmipb.TypedValue{},
		want: true,
	}, {
		name: "non-scalar: a < b",
		inA: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_JsonVal{[]byte("json")},
		},
		inB: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_JsonVal{[]byte("zzz")},
		},
		want: true,
	}, {
		name: "non-scalar: b < a",
		inA: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_JsonIetfVal{[]byte("aa")},
		},
		inB: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_JsonIetfVal{[]byte("zz")},
		},
		want: false,
	}, {
		name: "scalar string: a < b",
		inA: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_StringVal{"a"},
		},
		inB: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_StringVal{"z"},
		},
		want: true,
	}, {
		name: "scalar string: a < b",
		inA: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_StringVal{"z"},
		},
		inB: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_StringVal{"a"},
		},
		want: false,
	}, {
		name: "scalar float32: a < b",
		inA: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_DecimalVal{&gnmipb.Decimal64{
				Digits:    1234,
				Precision: 4,
			}},
		},
		inB: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_DecimalVal{&gnmipb.Decimal64{
				Digits:    1234,
				Precision: 2,
			}},
		},
		want: true,
	}, {
		name: "scalar float32: b < a",
		inA: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_DecimalVal{&gnmipb.Decimal64{
				Digits:    1234,
				Precision: 0,
			}},
		},
		inB: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_DecimalVal{&gnmipb.Decimal64{
				Digits:    1234,
				Precision: 10,
			}},
		},
		want: false,
	}, {
		name: "scalar float64: a < b",
		inA: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_FloatVal{42.42},
		},
		inB: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_FloatVal{84.84},
		},
		want: true,
	}, {
		name: "scalar float64: b < a",
		inA: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_FloatVal{84.84},
		},
		inB: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_FloatVal{42.42},
		},
	}, {
		name: "scalar int64: a < b",
		inA: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_IntVal{-42},
		},
		inB: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_IntVal{42},
		},
		want: true,
	}, {
		name: "scalar int64: b < a",
		inA: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_IntVal{42},
		},
		inB: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_IntVal{-42},
		},
	}, {
		name: "scalar int64: a < b",
		inA: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_UintVal{0},
		},
		inB: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_UintVal{42},
		},
		want: true,
	}, {
		name: "scalar int64: b < a",
		inA: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_UintVal{42},
		},
		inB: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_UintVal{0},
		},
		want: false,
	}, {
		name: "scalar bool: a < b",
		inA: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_BoolVal{false},
		},
		inB: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_BoolVal{true},
		},
		want: true,
	}, {
		name: "scalar bool: a < b but equal",
		inA: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_BoolVal{true},
		},
		inB: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_BoolVal{true},
		},
		want: true,
	}, {
		name: "scalar bool: b < a",
		inA: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_BoolVal{true},
		},
		inB: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_BoolVal{false},
		},
		want: false,
	}, {
		name: "non-scalar: a < b",
		inA: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_LeaflistVal{&gnmipb.ScalarArray{
				Element: []*gnmipb.TypedValue{{
					Value: &gnmipb.TypedValue_StringVal{"a"},
				}},
			}},
		},
		inB: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_LeaflistVal{&gnmipb.ScalarArray{
				Element: []*gnmipb.TypedValue{{
					Value: &gnmipb.TypedValue_StringVal{"z"},
				}},
			}},
		},
		want: true,
	}, {
		name: "non-scalar: b < a",
		inA: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_LeaflistVal{&gnmipb.ScalarArray{
				Element: []*gnmipb.TypedValue{{
					Value: &gnmipb.TypedValue_StringVal{"z"},
				}},
			}},
		},
		inB: &gnmipb.TypedValue{
			Value: &gnmipb.TypedValue_LeaflistVal{&gnmipb.ScalarArray{
				Element: []*gnmipb.TypedValue{{
					Value: &gnmipb.TypedValue_StringVal{"a"},
				}},
			}},
		},
		want: false,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := typedValueLess(tt.inA, tt.inB); got != tt.want {
				t.Fatalf("typedValueLess(%#v, %#v): did not get expected value, got: %v, want: %v", tt.inA, tt.inB, got, tt.want)
			}
		})
	}
}
