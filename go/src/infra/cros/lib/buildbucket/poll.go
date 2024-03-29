package buildbucket

import (
	"context"
	"fmt"
	"time"

	bbpb "go.chromium.org/luci/buildbucket/proto"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/retry"
	"go.chromium.org/luci/grpc/grpcutil"
	"go.chromium.org/luci/grpc/prpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/fieldmaskpb"
)

// Default fields returned by Buildbucket. The builds requested in this function
// will eventually be returned to the caller (usually a recipe), so this set
// should be kept in sync with the default fields in
// https://source.chromium.org/chromium/infra/infra/+/main:recipes-py/recipe_modules/buildbucket/api.py.
var defaultFields = []string{
	"builder",
	"create_time",
	"created_by",
	"critical",
	"end_time",
	"id",
	"input",
	"number",
	"output",
	"start_time",
	"status",
	"update_time",
}

// DefaultPRPCOpts returns a set of Options that work well with Buildbucket for
// most use cases.
func DefaultPRPCOpts() *prpc.Options {
	opts := prpc.DefaultOptions()
	opts.Retry = func() retry.Iterator {
		return &retry.ExponentialBackoff{
			Limited: retry.Limited{
				Delay:   time.Second,
				Retries: 10,
			},
			Multiplier: 2.0,
			MaxDelay:   5 * time.Minute,
		}
	}
	opts.Debug = true
	return opts
}

// PollForOutputProp polls until all of buildIds are completed or have set
// outputProp. Between each call to Buildbucket, the fn. sleeps for interval.
// This functionality is similar to the `bb collect` command, but allows polling
// for an output prop being set, instead of only build completion; this is
// useful for things like polling until a builder has published images.
//
// On completion, returns a map from build id -> Build for all of buildIds.
func PollForOutputProp(
	ctx context.Context,
	client bbpb.BuildsClient,
	buildIds []int64,
	outputProp string,
	interval time.Duration,
) (map[int64]*bbpb.Build, error) {
	// Keep two sets of builds, one with builds that are completed, one with
	// builds that have the output property set, but are not completed. Builds
	// that are completed will not be requested in later iterations.
	//
	// Another potential optimization is to only request propSetBuilds after
	// polling is complete, to get a final status (this would also allow
	// requesting fewer fields during the polling requests). The downside is
	// that there will be an additional request at the end.
	completedBuilds := make(map[int64]*bbpb.Build)
	propSetBuilds := make(map[int64]*bbpb.Build)

	for {
		batchReq := &bbpb.BatchRequest{}
		for _, id := range buildIds {
			if _, found := completedBuilds[id]; found {
				logging.Debugf(ctx, "build %d already completed, not requesting", id)
				continue
			}

			batchReq.Requests = append(batchReq.Requests, &bbpb.BatchRequest_Request{
				Request: &bbpb.BatchRequest_Request_GetBuild{
					GetBuild: &bbpb.GetBuildRequest{
						Id: id,
						// Also request tags, to match the behavior in
						// https://source.chromium.org/chromiumos/chromiumos/codesearch/+/main:infra/recipes/recipe_modules/orch_menu/api.py.
						Fields: &fieldmaskpb.FieldMask{Paths: append(defaultFields, "tags")},
					},
				},
			})
		}

		batchResp, err := client.Batch(ctx, batchReq)
		if err != nil {
			return nil, err
		}

		for _, resp := range batchResp.GetResponses() {
			var build *bbpb.Build
			switch resp.GetResponse().(type) {
			case *bbpb.BatchResponse_Response_GetBuild:
				build = resp.GetGetBuild()
			case *bbpb.BatchResponse_Response_Error:
				// One of the responses in the batch is an error. If the error
				// is transient (we also consider DeadlineExceeded transient),
				// just skip doing any updates to the build on this iteration;
				// the build status will be requested again next iteration. If
				// the error is non-transient, return an error to exit the loop.
				code := status.FromProto(resp.GetError()).Code()
				if grpcutil.IsTransientCode(code) || code == codes.DeadlineExceeded {
					logging.Warningf(ctx, "got transient error in BatchResponse, will retry next iteration: %q", resp.GetError())
					continue
				}

				return nil, fmt.Errorf("got error in BatchResponse: %q", resp.GetError())
			default:
				return nil, fmt.Errorf("got unexpected response type: %q", resp)
			}

			// If the build has completed add it to completedBuilds and remove
			// it from propSetBuilds (if it was there). If the build has set the
			// property, add it to propSetBuilds.
			if (build.GetStatus() & bbpb.Status_ENDED_MASK) != 0 {
				logging.Infof(ctx, "build %d completed", build.Id)
				completedBuilds[build.Id] = build

				if _, found := propSetBuilds[build.Id]; found {
					logging.Debugf(ctx, "build %d in propSetBuilds, moved to completedBuilds", build.Id)
					delete(propSetBuilds, build.Id)
				}
			} else if _, containsOutputProp := build.GetOutput().GetProperties().GetFields()[outputProp]; containsOutputProp {
				logging.Infof(ctx, "build %d has output prop (%q)", build.Id, outputProp)
				propSetBuilds[build.Id] = build
			} else {
				logging.Infof(ctx, "still waiting for build %d to complete or set output property", build.Id)
			}
		}

		// Check that a given build is not in both of the sets, this indicates
		// an invalid state.
		for _, id := range buildIds {
			_, buildInCompleted := completedBuilds[id]
			_, buildInPropSet := propSetBuilds[id]

			if buildInCompleted && buildInPropSet {
				panic(fmt.Sprintf("build %d in both completedBuilds and propSetBuilds, invalid state", id))
			}
		}

		// Check that the sum of the size of the sets combined is not greater
		// than the length of buildIds, this indicates an invalid state.
		if len(completedBuilds)+len(propSetBuilds) > len(buildIds) {
			panic("sum of builds in completedBuilds and propSetBuilds > buildIds, invalid state")
		}

		// Once all requested builds are completed or set the output prop,
		// break. Otherwise sleep for interval.
		if len(completedBuilds)+len(propSetBuilds) == len(buildIds) {
			break
		}

		logging.Infof(ctx, "sleeping for %s", interval)
		time.Sleep(interval)
	}

	allBuilds := make(map[int64]*bbpb.Build)
	for id, build := range completedBuilds {
		allBuilds[id] = build
	}
	for id, build := range propSetBuilds {
		allBuilds[id] = build
	}
	return allBuilds, nil
}
