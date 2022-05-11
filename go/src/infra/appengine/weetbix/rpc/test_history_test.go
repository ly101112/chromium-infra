// Copyright 2022 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package rpc

import (
	"context"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/clock/testclock"
	. "go.chromium.org/luci/common/testing/assertions"
	"go.chromium.org/luci/resultdb/rdbperms"
	"go.chromium.org/luci/server/auth"
	"go.chromium.org/luci/server/auth/authtest"
	"go.chromium.org/luci/server/span"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"

	"infra/appengine/weetbix/internal/testutil"
	"infra/appengine/weetbix/internal/testverdicts"
	"infra/appengine/weetbix/pbutil"
	pb "infra/appengine/weetbix/proto/v1"
)

func TestTestHistoryServer(t *testing.T) {
	Convey("TestHistoryServer", t, func() {
		ctx := testutil.SpannerTestContext(t)
		ctx, _ = testclock.UseTime(ctx, time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC))

		ctx = auth.WithState(ctx, &authtest.FakeState{
			Identity: "user:someone@example.com",
			IdentityPermissions: []authtest.RealmPermission{
				{
					Realm:      "project:realm",
					Permission: rdbperms.PermListTestResults,
				},
				{
					Realm:      "project:realm",
					Permission: rdbperms.PermListTestExonerations,
				},
			},
		})

		now := clock.Now(ctx)
		day := 24 * time.Hour

		var1 := pbutil.Variant("key1", "val1", "key2", "val1")
		var2 := pbutil.Variant("key1", "val2", "key2", "val1")
		var3 := pbutil.Variant("key1", "val2", "key2", "val2")
		var4 := pbutil.Variant("key1", "val1", "key2", "val2")

		_, err := span.ReadWriteTransaction(ctx, func(ctx context.Context) error {
			insertTVR := func(subRealm string, variant *pb.Variant) {
				(&testverdicts.TestVariantRealm{
					Project:     "project",
					TestID:      "test_id",
					SubRealm:    subRealm,
					Variant:     variant,
					VariantHash: pbutil.VariantHash(variant),
				}).SaveUnverified(ctx)
			}

			insertTVR("realm", var1)
			insertTVR("realm", var2)
			insertTVR("realm", var3)
			insertTVR("realm2", var4)

			insertTV := func(partitionTime time.Time, variant *pb.Variant, invId string, hasUnsubmittedChanges bool) {
				(&testverdicts.TestVerdict{
					Project:               "project",
					TestID:                "test_id",
					SubRealm:              "realm",
					PartitionTime:         partitionTime,
					VariantHash:           pbutil.VariantHash(variant),
					IngestedInvocationID:  invId,
					HasUnsubmittedChanges: hasUnsubmittedChanges,
				}).SaveUnverified(ctx)
			}

			insertTV(now.Add(-1*day), var1, "inv1", false)
			insertTV(now.Add(-1*day), var1, "inv2", false)
			insertTV(now.Add(-1*day), var2, "inv1", false)

			insertTV(now.Add(-2*day), var1, "inv1", false)
			insertTV(now.Add(-2*day), var1, "inv2", true)
			insertTV(now.Add(-2*day), var2, "inv1", true)

			insertTV(now.Add(-3*day), var3, "inv1", true)

			return nil
		})
		So(err, ShouldBeNil)

		server := NewTestHistoryServer()

		Convey("Query", func() {
			req := &pb.QueryTestHistoryRequest{
				Project: "project",
				TestId:  "test_id",
				Predicate: &pb.TestVerdictPredicate{
					SubRealm: "realm",
				},
				PageSize: 5,
			}

			Convey("unauthorised requests are rejected", func() {
				testPerm := func(ctx context.Context) {
					res, err := server.Query(ctx, req)
					So(err, ShouldErrLike, `caller does not have permission`, `in realm "project:realm"`)
					So(err, ShouldHaveGRPCStatus, codes.PermissionDenied)
					So(res, ShouldBeNil)
				}

				// No permission.
				ctx = auth.WithState(ctx, &authtest.FakeState{
					Identity: "user:someone@example.com",
				})
				testPerm(ctx)

				// testResults.list only.
				ctx = auth.WithState(ctx, &authtest.FakeState{
					Identity: "user:someone@example.com",
					IdentityPermissions: []authtest.RealmPermission{
						{
							Realm:      "project:realm",
							Permission: rdbperms.PermListTestResults,
						},
						{
							Realm:      "project:other_realm",
							Permission: rdbperms.PermListTestExonerations,
						},
					},
				})
				testPerm(ctx)

				// testExonerations.list only.
				ctx = auth.WithState(ctx, &authtest.FakeState{
					Identity: "user:someone@example.com",
					IdentityPermissions: []authtest.RealmPermission{
						{
							Realm:      "project:other_realm",
							Permission: rdbperms.PermListTestResults,
						},
						{
							Realm:      "project:realm",
							Permission: rdbperms.PermListTestExonerations,
						},
					},
				})
				testPerm(ctx)
			})

			Convey("invalid requests are rejected", func() {
				req.PageSize = -1
				res, err := server.Query(ctx, req)
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveGRPCStatus, codes.InvalidArgument)
				So(res, ShouldBeNil)
			})

			Convey("e2e", func() {
				res, err := server.Query(ctx, req)
				So(err, ShouldBeNil)
				So(res, ShouldResembleProto, &pb.QueryTestHistoryResponse{
					Verdicts: []*pb.TestVerdict{
						{
							TestId:        "test_id",
							VariantHash:   pbutil.VariantHash(var1),
							InvocationId:  "inv1",
							Status:        pb.TestVerdictStatus_EXPECTED,
							PartitionTime: timestamppb.New(now.Add(-1 * day)),
						},
						{
							TestId:        "test_id",
							VariantHash:   pbutil.VariantHash(var1),
							InvocationId:  "inv2",
							Status:        pb.TestVerdictStatus_EXPECTED,
							PartitionTime: timestamppb.New(now.Add(-1 * day)),
						},
						{
							TestId:        "test_id",
							VariantHash:   pbutil.VariantHash(var2),
							InvocationId:  "inv1",
							Status:        pb.TestVerdictStatus_EXPECTED,
							PartitionTime: timestamppb.New(now.Add(-1 * day)),
						},
						{
							TestId:        "test_id",
							VariantHash:   pbutil.VariantHash(var1),
							InvocationId:  "inv1",
							Status:        pb.TestVerdictStatus_EXPECTED,
							PartitionTime: timestamppb.New(now.Add(-2 * day)),
						},
						{
							TestId:        "test_id",
							VariantHash:   pbutil.VariantHash(var1),
							InvocationId:  "inv2",
							Status:        pb.TestVerdictStatus_EXPECTED,
							PartitionTime: timestamppb.New(now.Add(-2 * day)),
						},
					},
					NextPageToken: res.NextPageToken,
				})
				So(res.NextPageToken, ShouldNotBeEmpty)

				req.PageToken = res.NextPageToken
				res, err = server.Query(ctx, req)
				So(err, ShouldBeNil)
				So(res, ShouldResembleProto, &pb.QueryTestHistoryResponse{
					Verdicts: []*pb.TestVerdict{
						{
							TestId:        "test_id",
							VariantHash:   pbutil.VariantHash(var2),
							InvocationId:  "inv1",
							Status:        pb.TestVerdictStatus_EXPECTED,
							PartitionTime: timestamppb.New(now.Add(-2 * day)),
						},
						{
							TestId:        "test_id",
							VariantHash:   pbutil.VariantHash(var3),
							InvocationId:  "inv1",
							Status:        pb.TestVerdictStatus_EXPECTED,
							PartitionTime: timestamppb.New(now.Add(-3 * day)),
						},
					},
				})
			})
		})

		Convey("QueryStats", func() {
			req := &pb.QueryTestHistoryStatsRequest{
				Project: "project",
				TestId:  "test_id",
				Predicate: &pb.TestVerdictPredicate{
					SubRealm: "realm",
				},
				PageSize: 3,
			}

			Convey("unauthorised requests are rejected", func() {
				testPerm := func(ctx context.Context) {
					res, err := server.QueryStats(ctx, req)
					So(err, ShouldErrLike, `caller does not have permission`, `in realm "project:realm"`)
					So(err, ShouldHaveGRPCStatus, codes.PermissionDenied)
					So(res, ShouldBeNil)
				}

				// No permission.
				ctx = auth.WithState(ctx, &authtest.FakeState{
					Identity: "user:someone@example.com",
				})
				testPerm(ctx)

				// testResults.list only.
				ctx = auth.WithState(ctx, &authtest.FakeState{
					Identity: "user:someone@example.com",
					IdentityPermissions: []authtest.RealmPermission{
						{
							Realm:      "project:realm",
							Permission: rdbperms.PermListTestResults,
						},
						{
							Realm:      "project:other_realm",
							Permission: rdbperms.PermListTestExonerations,
						},
					},
				})
				testPerm(ctx)

				// testExonerations.list only.
				ctx = auth.WithState(ctx, &authtest.FakeState{
					Identity: "user:someone@example.com",
					IdentityPermissions: []authtest.RealmPermission{
						{
							Realm:      "project:other_realm",
							Permission: rdbperms.PermListTestResults,
						},
						{
							Realm:      "project:realm",
							Permission: rdbperms.PermListTestExonerations,
						},
					},
				})
				testPerm(ctx)
			})

			Convey("invalid requests are rejected", func() {
				req.PageSize = -1
				res, err := server.QueryStats(ctx, req)
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveGRPCStatus, codes.InvalidArgument)
				So(res, ShouldBeNil)
			})

			Convey("e2e", func() {
				day := 24 * time.Hour
				res, err := server.QueryStats(ctx, req)
				So(err, ShouldBeNil)
				So(res, ShouldResembleProto, &pb.QueryTestHistoryStatsResponse{
					Groups: []*pb.QueryTestHistoryStatsResponse_Group{
						{
							PartitionTime: timestamppb.New(now.Add(-1 * day)),
							VariantHash:   pbutil.VariantHash(var1),
							ExpectedCount: 2,
						},
						{
							PartitionTime: timestamppb.New(now.Add(-1 * day)),
							VariantHash:   pbutil.VariantHash(var2),
							ExpectedCount: 1,
						},
						{
							PartitionTime: timestamppb.New(now.Add(-2 * day)),
							VariantHash:   pbutil.VariantHash(var1),
							ExpectedCount: 2,
						},
					},
					NextPageToken: res.NextPageToken,
				})
				So(res.NextPageToken, ShouldNotBeEmpty)

				req.PageToken = res.NextPageToken
				res, err = server.QueryStats(ctx, req)
				So(err, ShouldBeNil)
				So(res, ShouldResembleProto, &pb.QueryTestHistoryStatsResponse{
					Groups: []*pb.QueryTestHistoryStatsResponse_Group{
						{
							PartitionTime: timestamppb.New(now.Add(-2 * day)),
							VariantHash:   pbutil.VariantHash(var2),
							ExpectedCount: 1,
						},
						{
							PartitionTime: timestamppb.New(now.Add(-3 * day)),
							VariantHash:   pbutil.VariantHash(var3),
							ExpectedCount: 1,
						},
					},
				})
			})
		})

		Convey("QueryVariants", func() {
			req := &pb.QueryVariantsRequest{
				Project:  "project",
				TestId:   "test_id",
				SubRealm: "realm",
				PageSize: 2,
			}

			Convey("unauthorised requests are rejected", func() {
				ctx = auth.WithState(ctx, &authtest.FakeState{
					Identity: "user:someone@example.com",
				})
				res, err := server.QueryVariants(ctx, req)
				So(err, ShouldErrLike, `caller does not have permission`, `in realm "project:realm"`)
				So(err, ShouldHaveGRPCStatus, codes.PermissionDenied)
				So(res, ShouldBeNil)
			})

			Convey("invalid requests are rejected", func() {
				req.PageSize = -1
				res, err := server.QueryVariants(ctx, req)
				So(err, ShouldNotBeNil)
				So(err, ShouldHaveGRPCStatus, codes.InvalidArgument)
				So(res, ShouldBeNil)
			})

			Convey("e2e", func() {
				res, err := server.QueryVariants(ctx, req)
				So(err, ShouldBeNil)
				So(res, ShouldResembleProto, &pb.QueryVariantsResponse{
					Variants: []*pb.QueryVariantsResponse_VariantInfo{
						{
							VariantHash: pbutil.VariantHash(var1),
							Variant:     var1,
						},
						{
							VariantHash: pbutil.VariantHash(var3),
							Variant:     var3,
						},
					},
					NextPageToken: res.NextPageToken,
				})
				So(res.NextPageToken, ShouldNotBeEmpty)

				req.PageToken = res.NextPageToken
				res, err = server.QueryVariants(ctx, req)
				So(err, ShouldBeNil)
				So(res, ShouldResembleProto, &pb.QueryVariantsResponse{
					Variants: []*pb.QueryVariantsResponse_VariantInfo{
						{
							VariantHash: pbutil.VariantHash(var2),
							Variant:     var2,
						},
					},
				})
			})
		})
	})
}

func TestValidateQueryTestHistoryRequest(t *testing.T) {
	t.Parallel()

	Convey("validateQueryTestHistoryRequest", t, func() {
		req := &pb.QueryTestHistoryRequest{
			Project: "project",
			TestId:  "test_id",
			Predicate: &pb.TestVerdictPredicate{
				SubRealm: "realm",
			},
			PageSize: 5,
		}

		Convey("valid", func() {
			err := validateQueryTestHistoryRequest(req)
			So(err, ShouldBeNil)
		})

		Convey("no project", func() {
			req.Project = ""
			err := validateQueryTestHistoryRequest(req)
			So(err, ShouldErrLike, "project missing")
		})

		Convey("no test_id", func() {
			req.TestId = ""
			err := validateQueryTestHistoryRequest(req)
			So(err, ShouldErrLike, "test_id missing")
		})

		Convey("no predicate", func() {
			req.Predicate = nil
			err := validateQueryTestHistoryRequest(req)
			So(err, ShouldErrLike, "predicate", "unspecified")
		})

		Convey("no page size", func() {
			req.PageSize = 0
			err := validateQueryTestHistoryRequest(req)
			So(err, ShouldBeNil)
		})

		Convey("negative page size", func() {
			req.PageSize = -1
			err := validateQueryTestHistoryRequest(req)
			So(err, ShouldErrLike, "page_size", "negative")
		})
	})
}

func TestValidateQueryTestHistoryStatsRequest(t *testing.T) {
	t.Parallel()

	Convey("validateQueryTestHistoryStatsRequest", t, func() {
		req := &pb.QueryTestHistoryStatsRequest{
			Project: "project",
			TestId:  "test_id",
			Predicate: &pb.TestVerdictPredicate{
				SubRealm: "realm",
			},
			PageSize: 5,
		}

		Convey("valid", func() {
			err := validateQueryTestHistoryStatsRequest(req)
			So(err, ShouldBeNil)
		})

		Convey("no project", func() {
			req.Project = ""
			err := validateQueryTestHistoryStatsRequest(req)
			So(err, ShouldErrLike, "project missing")
		})

		Convey("no test_id", func() {
			req.TestId = ""
			err := validateQueryTestHistoryStatsRequest(req)
			So(err, ShouldErrLike, "test_id missing")
		})

		Convey("no predicate", func() {
			req.Predicate = nil
			err := validateQueryTestHistoryStatsRequest(req)
			So(err, ShouldErrLike, "predicate", "unspecified")
		})

		Convey("no page size", func() {
			req.PageSize = 0
			err := validateQueryTestHistoryStatsRequest(req)
			So(err, ShouldBeNil)
		})

		Convey("negative page size", func() {
			req.PageSize = -1
			err := validateQueryTestHistoryStatsRequest(req)
			So(err, ShouldErrLike, "page_size", "negative")
		})
	})
}

func TestValidateQueryVariantsRequest(t *testing.T) {
	t.Parallel()

	Convey("validateQueryVariantsRequest", t, func() {
		req := &pb.QueryVariantsRequest{
			Project:  "project",
			TestId:   "test_id",
			PageSize: 5,
		}

		Convey("valid", func() {
			err := validateQueryVariantsRequest(req)
			So(err, ShouldBeNil)
		})

		Convey("no project", func() {
			req.Project = ""
			err := validateQueryVariantsRequest(req)
			So(err, ShouldErrLike, "project missing")
		})

		Convey("no test_id", func() {
			req.TestId = ""
			err := validateQueryVariantsRequest(req)
			So(err, ShouldErrLike, "test_id missing")
		})

		Convey("no page size", func() {
			req.PageSize = 0
			err := validateQueryVariantsRequest(req)
			So(err, ShouldBeNil)
		})

		Convey("negative page size", func() {
			req.PageSize = -1
			err := validateQueryVariantsRequest(req)
			So(err, ShouldErrLike, "page_size", "negative")
		})
	})
}
