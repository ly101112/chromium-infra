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

package app

const botSummaryKind = "fleetBotSummary"

// fleetBotSummaryEntity is a datastore entity that stores the fleet.BotSummary in
// JSON format. In effect, this is a cache of task and bot history information
// from the Swarming server over a fixed time period.
type fleetBotSummaryEntity struct {
	_kind string `gae:"$kind,fleetBotSummary"`
	// Swarming bot's dut_id dimension. This dimension is an opaque reference
	// to the managed DUT's uuid in skylab inventory data.
	DutID string `gae:"$id"`
	// Data is the fleet.BotSummary object serialized to protobuf binary format.
	Data []byte `gae:",noindex"`
}
