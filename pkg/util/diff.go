/*
Copyright 2014 Google Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package util

import (
	"encoding/json"
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

// StringDiff diffs a and b and returns a human readable diff.
func StringDiff(a, b string) string {
	ba := []byte(a)
	bb := []byte(b)
	out := []byte{}
	i := 0
	for ; i < len(ba) && i < len(bb); i++ {
		if ba[i] != bb[i] {
			break
		}
		out = append(out, ba[i])
	}
	out = append(out, []byte("\n\nA: ")...)
	out = append(out, ba[i:]...)
	out = append(out, []byte("\n\nB: ")...)
	out = append(out, bb[i:]...)
	out = append(out, []byte("\n\n")...)
	return string(out)
}

// ObjectDiff writes the two objects out as JSON and prints out the identical part of
// the objects followed by the remaining part of 'a' and finally the remaining part of 'b'.
// For debugging tests.
func ObjectDiff(a, b interface{}) string {
	ab, err := json.Marshal(a)
	if err != nil {
		panic(fmt.Sprintf("a: %v", err))
	}
	bb, err := json.Marshal(b)
	if err != nil {
		panic(fmt.Sprintf("b: %v", err))
	}
	return StringDiff(string(ab), string(bb))
}

// ObjectGoPrintDiff is like ObjectDiff, but uses go-spew to print the objects,
// which shows absolutely everything by recursing into every single pointer
// (go's %#v formatters OTOH stop at a certain point). This is needed when you
// can't figure out why reflect.DeepEqual is returning false and nothing is
// showing you differences. This will.
func ObjectGoPrintDiff(a, b interface{}) string {
	s := spew.ConfigState{
		Indent: " ",
		// Extra deep spew.
		DisableMethods: true,
	}
	return StringDiff(
		s.Sdump(a),
		s.Sdump(b),
	)
}
