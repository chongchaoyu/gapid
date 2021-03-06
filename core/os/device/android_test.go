// Copyright (C) 2017 Google Inc.
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

package device_test

import (
	"testing"

	"github.com/google/gapid/core/assert"
	"github.com/google/gapid/core/os/device"
)

func TestAndroidOS(t *testing.T) {
	ctx := assert.Context(t)
	const point int32 = 2
	for _, test := range []struct {
		major int32
		minor int32
		name  string
	}{
		{major: 0, minor: 7, name: "Android 0.7.2"},
		{major: 6, minor: 0, name: "Marshmallow"},
		{major: 5, minor: 0, name: "Lollipop"},
		{major: 4, minor: 9, name: "KitKat"},
		{major: 4, minor: 4, name: "KitKat"},
		{major: 4, minor: 3, name: "Jelly Bean"},
		{major: 4, minor: 1, name: "Jelly Bean"},
		{major: 4, minor: 0, name: "Ice Cream Sandwich"},
		{major: 3, minor: 0, name: "Honeycomb"},
		{major: 2, minor: 9, name: "Gingerbread"},
		{major: 2, minor: 3, name: "Gingerbread"},
		{major: 2, minor: 2, name: "Froyo"},
		{major: 2, minor: 1, name: "Eclair"},
		{major: 2, minor: 0, name: "Eclair"},
		{major: 1, minor: 6, name: "Donut"},
		{major: 1, minor: 5, name: "Cupcake"},
		{major: 1, minor: 0, name: "Cupcake"},
	} {
		os := device.AndroidOS(test.major, test.minor, point)
		assert.For(ctx, "OS Kind").That(os.Kind).Equals(device.Android)
		assert.For(ctx, "OS Name").That(os.Name).Equals(test.name)
		assert.For(ctx, "OS Build").That(os.Build).Equals("")
		assert.For(ctx, "OS Major").That(os.Major).Equals(test.major)
		assert.For(ctx, "OS Minor").That(os.Minor).Equals(test.minor)
		assert.For(ctx, "OS Point").That(os.Point).Equals(point)
	}
}
