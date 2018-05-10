// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package gerrit

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	ds "go.chromium.org/gae/service/datastore"

	"infra/tricium/api/admin/v1"
	"infra/tricium/appengine/common/track"
	"infra/tricium/appengine/common/triciumtest"
)

func TestReportResultsRequest(t *testing.T) {
	Convey("Test Environment", t, func() {
		tt := &triciumtest.Testing{}
		ctx := tt.Context()

		// Add request -> run -> function -> worker to datastore.
		functionName := "MyLinter"
		request := &track.AnalyzeRequest{
			GitRepo: "https://chromium-review.googlesource.com",
			GitRef:  "refs/changes/88/508788/1",
		}
		So(ds.Put(ctx, request), ShouldBeNil)
		requestKey := ds.KeyForObj(ctx, request)
		run := &track.WorkflowRun{ID: 1, Parent: requestKey}
		So(ds.Put(ctx, run), ShouldBeNil)
		runKey := ds.KeyForObj(ctx, run)
		So(ds.Put(ctx, &track.FunctionRun{
			ID:     "MyLinter",
			Parent: runKey,
		}), ShouldBeNil)
		analyzerKey := ds.NewKey(ctx, "FunctionRun", "MyLinter", 0, runKey)
		So(ds.Put(ctx, &track.FunctionRunResult{
			ID:          1,
			Parent:      analyzerKey,
			Name:        "MyLinter",
			NumComments: 2,
		}), ShouldBeNil)
		workerName := "MyLinter_UBUNTU"
		So(ds.Put(ctx, &track.WorkerRun{
			ID:     workerName,
			Parent: analyzerKey,
		}), ShouldBeNil)

		// Add example Comment and associated CommentSelection entities.
		workerKey := ds.NewKey(ctx, "WorkerRun", workerName, 0, analyzerKey)
		json1, err := json.Marshal("{\"category\":\"L\",\"message\":\"Line too long\"}")
		json2, err := json.Marshal("{\"category\":\"L\",\"message\":\"Line too short\"}")
		So(err, ShouldBeNil)
		comments := []*track.Comment{
			{Parent: workerKey, Comment: json1},
			{Parent: workerKey, Comment: json2},
			{Parent: workerKey, Comment: json2},
		}
		So(ds.Put(ctx, comments), ShouldBeNil)
		So(ds.Put(ctx, &track.CommentSelection{
			ID:       1,
			Parent:   ds.KeyForObj(ctx, comments[0]),
			Included: true,
		}), ShouldBeNil)
		So(ds.Put(ctx, &track.CommentSelection{
			ID: 1, Parent: ds.KeyForObj(ctx, comments[1]),
			Included: true,
		}), ShouldBeNil)
		// The third comment added is not "included" when reporting
		// comments.
		So(ds.Put(ctx, &track.CommentSelection{
			ID:       1,
			Parent:   ds.KeyForObj(ctx, comments[2]),
			Included: false,
		}), ShouldBeNil)

		Convey("Reports only included comments", func() {
			mock := &mockRestAPI{}
			err := reportResults(ctx, &admin.ReportResultsRequest{
				RunId:    run.ID,
				Analyzer: "MyLinter",
			}, mock)
			So(err, ShouldBeNil)
			// This only includes the two selected comments.
			So(len(mock.LastComments), ShouldEqual, len(comments)-1)
		})

		Convey("Does not report results when reporting is disabled", func() {
			request.GerritReportingDisabled = true
			So(ds.Put(ctx, request), ShouldBeNil)
			mock := &mockRestAPI{}
			err := reportResults(ctx, &admin.ReportResultsRequest{
				RunId:    run.ID,
				Analyzer: functionName,
			}, mock)
			So(err, ShouldBeNil)
			So(len(mock.LastComments), ShouldEqual, 0)
		})

		// Put more comments in until the number of included comments
		// has reached the maximum comment number threshold.
		for len(comments) < maxComments+1 {
			comment := &track.Comment{Parent: workerKey, Comment: json1}
			comments = append(comments, comment)
			So(ds.Put(ctx, comment), ShouldBeNil)
			So(ds.Put(ctx, &track.CommentSelection{
				ID:       1,
				Parent:   ds.KeyForObj(ctx, comment),
				Included: true,
			}), ShouldBeNil)
		}
		So(len(comments), ShouldEqual, maxComments+1)

		Convey("Reports when number of comments is at maximum", func() {
			mock := &mockRestAPI{}
			err := reportResults(ctx, &admin.ReportResultsRequest{
				RunId:    run.ID,
				Analyzer: functionName,
			}, mock)
			So(err, ShouldBeNil)
			So(len(mock.LastComments), ShouldEqual, maxComments)
		})

		// Put one more comment in;
		comment := &track.Comment{Parent: workerKey, Comment: json1}
		comments = append(comments, comment)
		So(ds.Put(ctx, comment), ShouldBeNil)
		So(ds.Put(ctx, &track.CommentSelection{
			ID:       1,
			Parent:   ds.KeyForObj(ctx, comment),
			Included: true,
		}), ShouldBeNil)
		So(len(comments), ShouldEqual, maxComments+2)

		Convey("Does not report when number of comments exceeds maximum", func() {
			mock := &mockRestAPI{}
			err := reportResults(ctx, &admin.ReportResultsRequest{
				RunId:    run.ID,
				Analyzer: functionName,
			}, mock)
			So(err, ShouldBeNil)
			So(len(mock.LastComments), ShouldEqual, 0)
		})
	})
}
