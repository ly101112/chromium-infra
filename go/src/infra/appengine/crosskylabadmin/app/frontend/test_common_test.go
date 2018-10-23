// Copyright 2018 The LUCI Authors.
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

package frontend

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	swarming "go.chromium.org/luci/common/api/swarming/swarming/v1"
	"go.chromium.org/luci/common/data/strpair"
)

func TestBotContainsDims(t *testing.T) {
	Convey("single dimension matches", t, func() {
		So(
			botContainsDims(
				&swarming.SwarmingRpcsBotInfo{
					Dimensions: []*swarming.SwarmingRpcsStringListPair{
						{Key: "a", Value: []string{"b"}},
					},
				},
				strpair.Map{"a": []string{"b"}}),
			ShouldBeTrue,
		)
	})

	Convey("requested values are subset of available values", t, func() {
		So(
			botContainsDims(
				&swarming.SwarmingRpcsBotInfo{
					Dimensions: []*swarming.SwarmingRpcsStringListPair{
						{Key: "a", Value: []string{"b", "c"}},
					},
				},
				strpair.Map{"a": []string{"b"}}),
			ShouldBeTrue,
		)
	})

	Convey("all values match", t, func() {
		So(
			botContainsDims(
				&swarming.SwarmingRpcsBotInfo{
					Dimensions: []*swarming.SwarmingRpcsStringListPair{
						{Key: "a", Value: []string{"b", "c"}},
					},
				},
				strpair.Map{"a": []string{"b", "c"}}),
			ShouldBeTrue,
		)
	})

	Convey("requested key subset of available keys and value matches", t, func() {
		So(
			botContainsDims(
				&swarming.SwarmingRpcsBotInfo{
					Dimensions: []*swarming.SwarmingRpcsStringListPair{
						{Key: "a", Value: []string{"b"}},
						{Key: "j", Value: []string{"l"}},
					},
				},
				strpair.Map{"a": []string{"b"}}),
			ShouldBeTrue,
		)
	})

	Convey("keys do not match", t, func() {
		So(
			botContainsDims(
				&swarming.SwarmingRpcsBotInfo{
					Dimensions: []*swarming.SwarmingRpcsStringListPair{
						{Key: "j", Value: []string{"b"}},
					},
				},
				strpair.Map{"a": []string{"b"}}),
			ShouldBeFalse,
		)
	})

	Convey("values do not match", t, func() {
		So(
			botContainsDims(
				&swarming.SwarmingRpcsBotInfo{
					Dimensions: []*swarming.SwarmingRpcsStringListPair{
						{Key: "a", Value: []string{"b"}},
					},
				},
				strpair.Map{"a": []string{"c"}}),
			ShouldBeFalse,
		)
	})

	Convey("some values do not match", t, func() {
		So(
			botContainsDims(
				&swarming.SwarmingRpcsBotInfo{
					Dimensions: []*swarming.SwarmingRpcsStringListPair{
						{Key: "a", Value: []string{"b", "d"}},
					},
				},
				strpair.Map{"a": []string{"b", "c"}}),
			ShouldBeFalse,
		)
	})
}
