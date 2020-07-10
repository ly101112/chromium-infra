// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package common

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/golang/protobuf/jsonpb"
	isolateservice "go.chromium.org/luci/common/api/isolate/isolateservice/v1"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/isolated"
	"go.chromium.org/luci/common/isolatedclient"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/sync/parallel"
	"go.chromium.org/luci/server/auth"

	tricium "infra/tricium/api/v1"
)

const (
	isolateDevServerURL  = "https://isolateserver-dev.appspot.com"
	isolateProdServerURL = "https://isolateserver.appspot.com"
	isolateNamespace     = "default-gzip" // The namespace to use when isolating new inputs.
)

// IsolateAPI defines the interface to the isolate server.
//
// The interface is tuned to the needs of Tricium and Tricium data.
type IsolateAPI interface {
	// IsolateGitFileDetails isolates Git file details based on the
	// corresponding Tricium data type definition.
	//
	// The Git file details data type should be isolated with the following
	// path tricium/data/git_file_details.json and the data format as specified
	// in tricium/api/v1/data.proto.
	//
	// Note that this isolate has no command and includes no other isolates.
	IsolateGitFileDetails(c context.Context, serverURL string, d *tricium.Data_GitFileDetails) (string, error)

	// FetchIsolatedResult fetches isolated Tricium result output as a JSON
	// string.
	//
	// The output is assumed to be on the form of a Tricium result and
	// located in tricium/data/results.json in the isolated output.
	FetchIsolatedResults(c context.Context, serverURL, namespace, isolatedOutput string) (string, error)
}

// IsolateServer implements the IsolateAPI interface.
var IsolateServer isolateServer

type isolateServer struct {
}

// IsolateGitFileDetails implements the IsolateAPI interface.
func (s isolateServer) IsolateGitFileDetails(c context.Context, serverURL string, d *tricium.Data_GitFileDetails) (string, error) {
	chunks := make([]*isoChunk, 2)
	mode := 0444

	// Create Git file details chunk.
	gitDetailsData, err := (&jsonpb.Marshaler{}).MarshalToString(d)
	if err != nil {
		return "", errors.Annotate(err, "failed to marshal git file details to JSON").Err()
	}
	gitDetailsSize := int64(len(gitDetailsData))
	chunks[0] = &isoChunk{
		data:  []byte(gitDetailsData),
		isIso: false,
	}
	h := isolated.GetHash(isolateNamespace)
	chunks[0].file = &isolated.File{
		Digest: isolated.HashBytes(h, chunks[0].data),
		Mode:   &mode,
		Size:   &gitDetailsSize,
	}

	// Create isolate chunk.
	iso := isolated.New(h)
	path, err := tricium.GetPathForDataType(d)
	if err != nil {
		return "", errors.Reason("failed to get data file path, data: %v", d).Err()
	}
	iso.Files[path] = *chunks[0].file
	isoData, err := json.Marshal(iso)
	if err != nil {
		return "", errors.Annotate(err, "failed to marshal git file details isolate").Err()
	}
	isoSize := int64(len(isoData))
	chunks[1] = &isoChunk{
		data:  []byte(isoData),
		isIso: true,
	}
	chunks[1].file = &isolated.File{
		Digest: isolated.HashBytes(h, chunks[1].data),
		Mode:   &mode,
		Size:   &isoSize,
	}

	// Isolate chunks.
	if err := s.isolateChunks(c, serverURL, chunks); err != nil {
		return "", errors.Annotate(err, "failed to isolate chunks").Err()
	}

	// Return isolate hash.
	return string(chunks[1].file.Digest), nil
}

// FetchIsolatedResults implements the IsolateAPI interface.
func (s isolateServer) FetchIsolatedResults(c context.Context, serverURL, namespace, digest string) (string, error) {
	outIso, err := s.fetchIsolated(c, serverURL, namespace, digest)
	if err != nil {
		return "", errors.Annotate(err, "failed to fetch output isolate").Err()
	}
	resultsFile, ok := outIso.Files["tricium/data/results.json"]
	if !ok {
		return "", errors.Reason("missing results file in isolated output, isolated output: (%s, %s)", namespace, digest).Err()
	}
	buf := &buffer{}
	if err := s.fetch(c, serverURL, namespace, string(resultsFile.Digest), buf); err != nil {
		return "", errors.Annotate(err, "failed to fetch result file").Err()
	}
	// Note: if there's an issue with keeping all output in memory, this could
	// potentially be changed to use io.Reader.
	return string(buf.Bytes()), nil
}

func (s isolateServer) isolateChunks(c context.Context, serverURL string, chunks []*isoChunk) error {
	// Check presence of isolated files.
	dgsts := make([]*isolateservice.HandlersEndpointsV1Digest, len(chunks))
	for i, chnk := range chunks {
		dgsts[i] = &isolateservice.HandlersEndpointsV1Digest{
			Digest:     string(chnk.file.Digest),
			Size:       *chnk.file.Size,
			IsIsolated: chnk.isIso,
		}
	}
	client, err := s.createIsolateClient(c, serverURL, isolateNamespace)
	if err != nil {
		return err
	}
	states, err := client.Contains(c, dgsts)
	if err != nil {
		return errors.Annotate(err, "failed to check isolate contains").Err()
	}
	// Push chunks not already present in parallel.
	return parallel.FanOutIn(func(ch chan<- func() error) {
		for i, st := range states {
			if st != nil {
				i, st := i, st
				ch <- func() error {
					return client.Push(c, st, isolatedclient.NewBytesSource(chunks[i].data))
				}
			}
		}
	})
}

func (s isolateServer) fetch(c context.Context, serverURL, namespace, digest string, buf *buffer) error {
	client, err := s.createIsolateClient(c, serverURL, namespace)
	if err != nil {
		return err
	}
	if err := client.Fetch(c, isolated.HexDigest(digest), buf); err != nil {
		return errors.Annotate(err, "failed to fetch").Err()
	}
	return nil
}

func (s isolateServer) fetchIsolated(c context.Context, serverURL, namespace, digest string) (*isolated.Isolated, error) {
	buf := &buffer{}
	if err := s.fetch(c, serverURL, namespace, digest, buf); err != nil {
		return nil, errors.Annotate(err, "failed to fetch isolated").Err()
	}
	iso := &isolated.Isolated{}
	json.Unmarshal(buf.Bytes(), iso)
	logging.Fields{
		"isolatedContents": string(buf.Bytes()),
		"isolated":         iso,
	}.Infof(c, "Fetched isolated.")
	return iso, nil
}

func (s isolateServer) createIsolateClient(c context.Context, serverURL, namespace string) (*isolatedclient.Client, error) {
	authTransport, err := auth.GetRPCTransport(c, auth.AsSelf)
	if err != nil {
		return nil, errors.Annotate(err, "failed to setup auth transport for isolate client").Err()
	}
	anonTransport, err := auth.GetRPCTransport(c, auth.NoAuth)
	if err != nil {
		return nil, errors.Annotate(err, "failed to setup anonymous transport for isolate client").Err()
	}
	return isolatedclient.NewClient(
		serverURL,
		isolatedclient.WithAnonymousClient(&http.Client{Transport: anonTransport}),
		isolatedclient.WithAuthClient(&http.Client{Transport: authTransport}),
		isolatedclient.WithNamespace(namespace)), nil
}

type isoChunk struct {
	data  []byte
	isIso bool
	file  *isolated.File
}

type buffer struct {
	bytes.Buffer
}

func (f *buffer) Seek(a int64, b int) (int64, error) {
	if a != 0 || b != 0 {
		return 0, errors.Reason("non-zero value given to buffer.Seek").Err()
	}
	f.Reset()
	return 0, nil
}

// MockIsolator mocks the IsolateAPI interface for testing.
var MockIsolator mockIsolator

type mockIsolator struct{}

// IsolateGitFileDetails is a mock function for MockIsolator.
//
// For any testing actually using the return values, create a new mock.
func (mockIsolator) IsolateGitFileDetails(c context.Context, serverURL string, d *tricium.Data_GitFileDetails) (string, error) {
	return "mockmockmock", nil
}

// FetchIsolatedResults is a mock function for MockIsolator.
//
// For any testing using the return value, create a new mock.
func (mockIsolator) FetchIsolatedResults(c context.Context, serverURL, namespace, isolatedOutput string) (string, error) {
	return "mockmockmock", nil
}
