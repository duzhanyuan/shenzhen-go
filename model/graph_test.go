// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"reflect"
	"strings"
	"testing"

	"github.com/google/shenzhen-go/model/parts"
)

func TestLoadJSON(t *testing.T) {
	got, err := LoadJSON(strings.NewReader(`{
	"nodes": {
		"foo": {"part_type": "Code", "part": {}},
		"bar": {"part_type": "Code", "part": {}}
	},
	"channels": {
		"baz": {},
		"qux": {}
	}
}`), "filePath", "urlPath")
	if err != nil {
		t.Fatalf("LoadJSON() error = %v", err)
	}

	if got, want := got.FilePath, "filePath"; got != want {
		t.Errorf("LoadJSON().FilePath = %q, want %q", got, want)
	}
	if got, want := got.URLPath, "urlPath"; got != want {
		t.Errorf("LoadJSON().URLPath = %q, want %q", got, want)
	}

	wantNodes := map[string]*Node{
		"foo": {
			Name:         "foo",
			Multiplicity: 1,
			Part:         &parts.Code{},
		},
		"bar": {
			Name:         "bar",
			Multiplicity: 1,
			Part:         &parts.Code{},
		},
	}
	if got, want := got.Nodes, wantNodes; !reflect.DeepEqual(got, want) {
		t.Errorf("LoadJSON().Nodes = %#v, want %#v", got, want)
	}
	wantChans := map[string]*Channel{
		"baz": {Name: "baz"},
		"qux": {Name: "qux"},
	}
	if got, want := got.Channels, wantChans; !reflect.DeepEqual(got, want) {
		t.Errorf("LoadJSON().Channels = %#v, want %#v", got, want)
	}
}