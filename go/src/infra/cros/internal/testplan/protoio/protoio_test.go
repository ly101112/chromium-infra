// Copyright 2024 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package protoio_test

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/testing/protocmp"

	test_api_v1 "go.chromium.org/chromiumos/config/go/test/api/v1"
	bbpb "go.chromium.org/luci/buildbucket/proto"

	"infra/cros/internal/testplan/protoio"
)

func TestReadBinaryOrJSONPb(t *testing.T) {
	ctx := context.Background()

	testPlan := &test_api_v1.HWTestPlan{
		Id: &test_api_v1.HWTestPlan_TestPlanId{Value: "testplan"},
	}

	tempdir := t.TempDir()
	binPbPath := filepath.Join(tempdir, "test1.binpb")
	jsonPbPath := filepath.Join(tempdir, "test2.jsonpb")
	misnamedBinPbPath := filepath.Join(tempdir, "test3.jsonpb")
	misnamedJSONPbPath := filepath.Join(tempdir, "test4.binpb")

	protoBytes, err := proto.Marshal(testPlan)
	if err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile(binPbPath, protoBytes, os.ModePerm); err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile(jsonPbPath, []byte(`{"id": {"value": "testplan"}}`), os.ModePerm); err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile(misnamedBinPbPath, protoBytes, os.ModePerm); err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile(misnamedJSONPbPath, []byte(`{"id": {"value": "testplan"}}`), os.ModePerm); err != nil {
		t.Fatal(err)
	}

	for _, path := range []string{binPbPath, jsonPbPath, misnamedBinPbPath, misnamedJSONPbPath} {
		readMessage := &test_api_v1.HWTestPlan{}
		err := protoio.ReadBinaryOrJSONPb(ctx, path, readMessage)
		if err != nil {
			t.Errorf("error reading path %s: %s", path, err)
		}

		if diff := cmp.Diff(testPlan, readMessage, protocmp.Transform()); diff != "" {
			t.Errorf("returned unexpected diff in read message (-want +got):\n%s", diff)
		}
	}
}

func TestWriteJsonl(t *testing.T) {
	outPath := filepath.Join(t.TempDir(), "out.jsonl")
	if err := protoio.WriteJsonl([]*test_api_v1.HWTestPlan{
		{Id: &test_api_v1.HWTestPlan_TestPlanId{
			Value: "testplan1",
		}},
		{Id: &test_api_v1.HWTestPlan_TestPlanId{Value: "testid2"}},
	}, outPath); err != nil {
		t.Fatal(err)
	}

	readBytes, err := os.ReadFile(outPath)
	if err != nil {
		t.Fatal(err)
	}

	expectedString := "{\"id\":{\"value\":\"testplan1\"}}\n{\"id\":{\"value\":\"testid2\"}}\n"
	if diff := cmp.Diff(expectedString, string(readBytes), protocmp.Transform()); diff != "" {
		t.Errorf("returned unexpected diff in read message (-want +got):\n%s", diff)
	}
}

func TestReadJsonl(t *testing.T) {
	inPath := filepath.Join(t.TempDir(), "in.jsonl")
	buildsContent := `{"builder": {"builder": "builderA"}}
{"builder": {"builder": "builderB"}}`
	if err := os.WriteFile(inPath, []byte(buildsContent), os.ModePerm); err != nil {
		t.Fatal(err)
	}

	parsedBuilds, err := protoio.ReadJsonl(inPath, func() *bbpb.Build { return &bbpb.Build{} })
	if err != nil {
		t.Fatal(err)
	}

	expectedBuilds := []*bbpb.Build{
		{
			Builder: &bbpb.BuilderID{Builder: "builderA"},
		},
		{
			Builder: &bbpb.BuilderID{Builder: "builderB"},
		},
	}
	if diff := cmp.Diff(expectedBuilds, parsedBuilds, protocmp.Transform()); diff != "" {
		t.Errorf("returned unexpected diff in read message (-want +got):\n%s", diff)
	}
}

func TestFilepathAsJsonpb(t *testing.T) {
	tests := []struct {
		name, input, expected string
	}{
		{
			name:     "binary proto",
			input:    "a/b/test.binpb",
			expected: "a/b/test.jsonpb",
		},
		{
			name:     "text file",
			input:    "a/b/test.txt",
			expected: "a/b/test.jsonpb",
		},
		{
			name:     "json proto",
			input:    "a/b/test.jsonpb",
			expected: "a/b/test.jsonpb",
		},
		{
			name:     "empty",
			input:    "",
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := protoio.FilepathAsJsonpb(test.input); got != test.expected {
				t.Errorf("FilepathAsJsonpb(%s) = %s, want %s", test.input, got, test.expected)
			}
		})
	}
}