// Copyright 2017 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package gerrit

import (
	"fmt"
	"sort"
	"sync"
	"testing"
	"time"

	"github.com/golang/protobuf/proto"
	. "github.com/smartystreets/goconvey/convey"
	ds "go.chromium.org/gae/service/datastore"
	tq "go.chromium.org/gae/service/taskqueue"
	"go.chromium.org/luci/auth/identity"
	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/clock/testclock"
	"go.chromium.org/luci/server/auth"
	"go.chromium.org/luci/server/auth/authtest"

	gr "golang.org/x/build/gerrit"
	"golang.org/x/net/context"

	"infra/tricium/api/admin/v1"
	"infra/tricium/api/v1"
	"infra/tricium/appengine/common"
	"infra/tricium/appengine/common/track"
	"infra/tricium/appengine/common/triciumtest"
)

const (
	host = "https://chromium-review.googlesource.com"
)

// mockPollRestAPI allows for modification of change state returned by QueryChanges.
type mockPollRestAPI struct {
	sync.Mutex
	changes map[string][]gr.ChangeInfo
}

func (m *mockPollRestAPI) QueryChanges(c context.Context, host, project string, ts time.Time) ([]gr.ChangeInfo, bool, error) {
	m.Lock()
	defer m.Unlock()

	if m.changes == nil {
		m.changes = make(map[string][]gr.ChangeInfo)
	}
	id := gerritProjectID(host, project)
	changes, _ := m.changes[id]
	return changes, false, nil
}

func (*mockPollRestAPI) PostRobotComments(c context.Context, host, change, revision string, runID int64, comments []*track.Comment) error {
	// not used by the poller
	return nil
}

func (m *mockPollRestAPI) GetChangedLines(c context.Context, host, change, revision string) (ChangedLinesInfo, error) {
	// not used by the poller
	return ChangedLinesInfo{}, nil
}

func (m *mockPollRestAPI) addChanges(host, project string, c []gr.ChangeInfo) {
	m.Lock()
	defer m.Unlock()

	if m.changes == nil {
		m.changes = make(map[string][]gr.ChangeInfo)
	}
	id := gerritProjectID(host, project)
	changes, _ := m.changes[id]
	changes = append(changes, c...)
	m.changes[id] = changes
}

type mockConfigProvider struct {
	Projects map[string]*tricium.ProjectConfig
}

func (*mockConfigProvider) GetServiceConfig(c context.Context) (*tricium.ServiceConfig, error) {
	return nil, nil // not used by the poller
}

func (m *mockConfigProvider) GetProjectConfig(c context.Context, p string) (*tricium.ProjectConfig, error) {
	pc, ok := m.Projects[p]
	if !ok {
		return nil, fmt.Errorf("nonexistent project")
	}
	return pc, nil
}

// GetAllProjectConfigs implements the ProviderAPI.
func (m *mockConfigProvider) GetAllProjectConfigs(c context.Context) (map[string]*tricium.ProjectConfig, error) {
	return m.Projects, nil
}

func numEnqueuedAnalyzeRequests(ctx context.Context) int {
	return len(tq.GetTestable(ctx).GetScheduledTasks()[common.AnalyzeQueue])
}

func TestPollAllProjectsBehavior(t *testing.T) {
	Convey("Test Environment", t, func() {
		ctx := triciumtest.Context()

		cp := &mockConfigProvider{
			Projects: map[string]*tricium.ProjectConfig{
				"a-project": {
					Repos: []*tricium.RepoDetails{
						{
							Source: &tricium.RepoDetails_GerritProject{
								GerritProject: &tricium.GerritProject{
									Host:    host,
									Project: "infra",
									GitUrl:  "https://repo-host.com/infra",
								},
							},
						},
					},
				},
				"b-project": {
					Repos: []*tricium.RepoDetails{
						{
							Source: &tricium.RepoDetails_GerritProject{
								GerritProject: &tricium.GerritProject{
									Host:    host,
									Project: "project/tricium-gerrit",
									GitUrl:  "https://repo-host.com/playground",
								},
							},
						},
					},
				},
			},
		}

		Convey("Poll puts a poll project request in the task queue for each project", func() {
			So(poll(ctx, cp), ShouldBeNil)
			tasks := tq.GetTestable(ctx).GetScheduledTasks()[common.PollProjectQueue]
			var projects []string
			for _, task := range tasks {
				request := &admin.PollProjectRequest{}
				So(proto.Unmarshal(task.Payload, request), ShouldBeNil)
				projects = append(projects, request.Project)
			}
			sort.Strings(projects)
			So(projects, ShouldResemble, []string{"a-project", "b-project"})
		})
	})
}

func TestPollProjectBasicBehavior(t *testing.T) {

	Convey("Test Environment", t, func() {
		ctx := triciumtest.Context()

		now := time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC)
		ctx, tc := testclock.UseTime(ctx, now)
		ctx = auth.WithState(ctx, &authtest.FakeState{
			Identity: identity.AnonymousIdentity,
		})

		cp := &mockConfigProvider{
			Projects: map[string]*tricium.ProjectConfig{
				"infra": {
					Repos: []*tricium.RepoDetails{
						{
							Source: &tricium.RepoDetails_GerritProject{
								GerritProject: &tricium.GerritProject{
									Host:    host,
									Project: "infra",
									GitUrl:  "https://repo-host.com/infra",
								},
							},
						},
						{
							Source: &tricium.RepoDetails_GerritProject{
								GerritProject: &tricium.GerritProject{
									Host:    host,
									Project: "project/tricium-gerrit",
									GitUrl:  "https://repo-host.com/playground",
								},
							},
						},
					},
				},
			},
		}
		projects, err := cp.GetAllProjectConfigs(ctx)
		So(err, ShouldBeNil)

		gerritProjects := []*tricium.GerritProject{
			projects["infra"].Repos[0].GetGerritProject(),
			projects["infra"].Repos[1].GetGerritProject(),
		}

		Convey("First poll, no changes", func() {
			api := &mockPollRestAPI{}
			So(pollProject(ctx, "infra", api, cp), ShouldBeNil)
			Convey("Creates tracking entries for Gerrit projects", func() {
				for _, gd := range gerritProjects {
					p := &Project{ID: gerritProjectID(gd.Host, gd.Project)}
					So(ds.Get(ctx, p), ShouldBeNil)
				}
			})
			Convey("Does not enqueue analyze requests", func() {
				So(numEnqueuedAnalyzeRequests(ctx), ShouldEqual, 0)
			})
		})

		Convey("Second poll, no changes", func() {
			api := &mockPollRestAPI{}
			So(pollProject(ctx, "infra", api, cp), ShouldBeNil)
			// Store last poll timestamps from first poll.
			lastPolls := make(map[string]time.Time)
			for _, gd := range gerritProjects {
				p := &Project{ID: gerritProjectID(gd.Host, gd.Project)}
				So(ds.Get(ctx, p), ShouldBeNil)
				lastPolls[p.ID] = p.LastPoll
			}
			So(pollProject(ctx, "infra", api, cp), ShouldBeNil)
			Convey("Does not update timestamp of last poll", func() {
				for _, gd := range gerritProjects {
					p := &Project{ID: gerritProjectID(gd.Host, gd.Project)}
					So(ds.Get(ctx, p), ShouldBeNil)
					t, _ := lastPolls[p.ID]
					So(t.Equal(p.LastPoll), ShouldBeTrue)
				}
			})
			Convey("Does not enqueue analyze requests", func() {
				So(numEnqueuedAnalyzeRequests(ctx), ShouldEqual, 0)
			})
		})

		Convey("First poll, with changes", func() {
			api := &mockPollRestAPI{}
			lastChangeTs := clock.Now(ctx)
			// Fill up with one change per project.
			for _, gd := range gerritProjects {
				api.addChanges(gd.Host, gd.Project, []gr.ChangeInfo{
					{
						Project: gd.Project,
						Updated: gr.TimeStamp(lastChangeTs),
						Owner:   &gr.AccountInfo{Email: "user@example.com"},
					},
				})
			}
			So(pollProject(ctx, "infra", api, cp), ShouldBeNil)
			Convey("Does not enqueue analyze requests", func() {
				So(numEnqueuedAnalyzeRequests(ctx), ShouldEqual, 0)
			})
		})

		Convey("Second poll, with new changes adding files", func() {
			api := &mockPollRestAPI{}
			lastChangeTs := tc.Now().UTC()
			// Fill up with one change per project.
			for _, gd := range gerritProjects {
				revisions := map[string]gr.RevisionInfo{
					"abcdef": {Files: map[string]*gr.FileInfo{"README.md": {}}},
				}
				api.addChanges(gd.Host, gd.Project, []gr.ChangeInfo{
					{
						ID:              "project~branch~Ideadc0de",
						Project:         gd.Project,
						Status:          "NEW",
						CurrentRevision: "abcdef",
						Updated:         gr.TimeStamp(lastChangeTs),
						Revisions:       revisions,
						Owner:           &gr.AccountInfo{Email: "user@example.com"},
					},
				})
			}
			So(pollProject(ctx, "infra", api, cp), ShouldBeNil)
			tc.Add(time.Second)
			So(pollProject(ctx, "infra", api, cp), ShouldBeNil)
			Convey("Updates last poll timestamp to last change timestamp", func() {
				for _, gd := range gerritProjects {
					p := &Project{ID: gerritProjectID(gd.Host, gd.Project)}
					So(ds.Get(ctx, p), ShouldBeNil)
					So(lastChangeTs.Equal(p.LastPoll), ShouldBeTrue)
				}
			})
			Convey("Enqueues analyze requests for each repo in the project", func() {
				So(numEnqueuedAnalyzeRequests(ctx), ShouldEqual, len(gerritProjects))
				tasks := tq.GetTestable(ctx).GetScheduledTasks()[common.AnalyzeQueue]
				var projects []string
				var repos []string
				for _, task := range tasks {
					ar := &tricium.AnalyzeRequest{}
					So(proto.Unmarshal(task.Payload, ar), ShouldBeNil)
					projects = append(projects, ar.Project)
					repos = append(repos, ar.GetGerritRevision().GitUrl)
				}
				So(projects, ShouldResemble, []string{"infra", "infra"})
				sort.Strings(repos)
				So(repos, ShouldResemble, []string{
					"https://repo-host.com/infra",
					"https://repo-host.com/playground",
				})

			})
			Convey("Adds change tracking entities", func() {
				for _, gd := range gerritProjects {
					So(ds.Get(ctx, &Change{
						ID:     "project~branch~Ideadc0de",
						Parent: ds.NewKey(ctx, "GerritProject", gerritProjectID(gd.Host, gd.Project), 0, nil),
					}), ShouldBeNil)
				}
			})
		})

		Convey("Poll with changes that include deleted and binary files", func() {
			api := &mockPollRestAPI{}
			lastChangeTs := tc.Now().UTC()
			// Fill up with one change per project.
			for _, gd := range gerritProjects {
				revisions := map[string]gr.RevisionInfo{
					"abcdef": {
						Files: map[string]*gr.FileInfo{
							"changed.txt": {},
							"deleted.txt": {Status: "D"},
							"binary.png":  {Binary: true},
						},
					},
				}
				api.addChanges(gd.Host, gd.Project, []gr.ChangeInfo{
					{
						ID:              "project~branch~Ideadc0de",
						Project:         gd.Project,
						Status:          "NEW",
						CurrentRevision: "abcdef",
						Updated:         gr.TimeStamp(lastChangeTs),
						Revisions:       revisions,
						Owner:           &gr.AccountInfo{Email: "user@example.com"},
					},
				})
			}
			So(pollProject(ctx, "infra", api, cp), ShouldBeNil)
			tc.Add(time.Second)
			So(pollProject(ctx, "infra", api, cp), ShouldBeNil)
			Convey("Enqueued analyze requests do not include deleted files", func() {
				tasks := tq.GetTestable(ctx).GetScheduledTasks()[common.AnalyzeQueue]
				So(len(tasks), ShouldEqual, len(gerritProjects))
				for _, task := range tasks {
					ar := &tricium.AnalyzeRequest{}
					err := proto.Unmarshal(task.Payload, ar)

					// Sorting files according to their paths to account for random
					// enumeration in go maps.
					sort.Slice(ar.Files, func(i, j int) bool {
						return ar.Files[i].Path < ar.Files[j].Path
					})
					So(err, ShouldBeNil)
					So(ar.Files, ShouldResemble, []*tricium.Data_File{
						{
							Path:     "binary.png",
							IsBinary: true,
							Status:   tricium.Data_MODIFIED,
						},
						{
							Path:     "changed.txt",
							IsBinary: false,
							Status:   tricium.Data_MODIFIED,
						},
					})
				}
			})
		})

		Convey("Poll when there is a change with no files", func() {
			api := &mockPollRestAPI{}
			lastChangeTs := tc.Now().UTC()
			// Fill up with one change per project.
			for _, gd := range gerritProjects {
				revisions := map[string]gr.RevisionInfo{
					"abcdef": {Files: make(map[string]*gr.FileInfo)},
				}
				api.addChanges(gd.Host, gd.Project, []gr.ChangeInfo{
					{
						ID:              "project~branch~Ideadc0de",
						Project:         gd.Project,
						Status:          "NEW",
						CurrentRevision: "abcdef",
						Updated:         gr.TimeStamp(lastChangeTs),
						Revisions:       revisions,
						Owner:           &gr.AccountInfo{Email: "user@example.com"},
					},
				})
			}
			So(pollProject(ctx, "infra", api, cp), ShouldBeNil)
			tc.Add(time.Second)
			So(pollProject(ctx, "infra", api, cp), ShouldBeNil)
			Convey("Does not enqueue analyze requests", func() {
				So(numEnqueuedAnalyzeRequests(ctx), ShouldEqual, 0)
			})
		})

		Convey("Poll with many changes, so paging is used", func() {
			api := &mockPollRestAPI{}
			// The first poll stores the timestamp.
			So(pollProject(ctx, "infra", api, cp), ShouldBeNil)
			tc.Add(time.Second)

			// Fill up each project with multiple changes.
			numChanges := 6
			revBase := "abcdef"
			branch := "master"
			changeIDFooter := "Ideadc0de"
			for _, gd := range gerritProjects {
				var changes []gr.ChangeInfo
				for i := 0; i < numChanges; i++ {
					tc.Add(time.Second)
					changeID := fmt.Sprintf("%s~%s~%s%d", gd.Project, branch, changeIDFooter, i)
					rev := fmt.Sprintf("%s%d", revBase, i)
					files := map[string]*gr.FileInfo{"README.md": {}}
					revisions := make(map[string]gr.RevisionInfo)
					revisions[rev] = gr.RevisionInfo{Files: files}
					changes = append(changes, gr.ChangeInfo{
						ID:              changeID,
						Project:         gd.Project,
						Status:          "NEW",
						CurrentRevision: rev,
						Updated:         gr.TimeStamp(tc.Now().UTC()),
						Revisions:       revisions,
						Owner:           &gr.AccountInfo{Email: "user@example.com"},
					})
				}
				api.addChanges(gd.Host, gd.Project, changes)

			}
			So(pollProject(ctx, "infra", api, cp), ShouldBeNil)
			Convey("Enqueues analyze requests", func() {
				So(numEnqueuedAnalyzeRequests(ctx), ShouldEqual, len(gerritProjects)*numChanges)
			})
			Convey("Adds change tracking entities", func() {
				for _, gd := range gerritProjects {
					for i := 0; i < numChanges; i++ {
						So(ds.Get(ctx, &Change{
							ID:     fmt.Sprintf("%s~%s~%s%d", gd.Project, branch, changeIDFooter, i),
							Parent: ds.NewKey(ctx, "GerritProject", gerritProjectID(gd.Host, gd.Project), 0, nil),
						}), ShouldBeNil)
					}
				}
			})
		})
	})
}

func TestPollProjectWhitelistBehavior(t *testing.T) {

	Convey("Test Environment", t, func() {
		ctx := triciumtest.Context()

		var (
			noWhitelistProject = "no-whitelist-project"
			whitelistProject   = "whitelist-group-project"
			whitelistGroup     = "whitelist-group-name"
		)

		now := time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC)
		ctx, tc := testclock.UseTime(ctx, now)
		ctx = auth.WithState(ctx, &authtest.FakeState{
			Identity: identity.AnonymousIdentity,
			FakeDB: authtest.FakeDB{
				"user:whitelisteduser@example.com": []string{whitelistGroup},
			},
		})

		cp := &mockConfigProvider{
			Projects: map[string]*tricium.ProjectConfig{
				noWhitelistProject: {
					Repos: []*tricium.RepoDetails{
						{
							Source: &tricium.RepoDetails_GerritProject{
								GerritProject: &tricium.GerritProject{
									Host:    host,
									Project: noWhitelistProject,
									GitUrl:  "https://repo-host.com/no-whitelist",
								},
							},
						},
					},
				},
				whitelistProject: {
					Repos: []*tricium.RepoDetails{
						{
							Source: &tricium.RepoDetails_GerritProject{
								GerritProject: &tricium.GerritProject{
									Host:    host,
									Project: whitelistProject,
									GitUrl:  "https://repo-host.com/whitelist",
								},
							},
							WhitelistedGroup: []string{whitelistGroup},
						},
					},
				},
			},
		}

		projects, err := cp.GetAllProjectConfigs(ctx)
		So(err, ShouldBeNil)

		var gerritProjects []*tricium.GerritProject
		for _, pc := range projects {
			for _, repo := range pc.Repos {
				if gd := repo.GetGerritProject(); gd != nil {
					gerritProjects = append(gerritProjects, gd)
				}
			}
		}

		Convey("No whitelisted groups means all changes are analyzed", func() {
			api := &mockPollRestAPI{}
			lastChangeTs := tc.Now().UTC()
			revisions := map[string]gr.RevisionInfo{
				"abcdef": {Files: map[string]*gr.FileInfo{"README.md": {}}},
			}
			api.addChanges(host, noWhitelistProject, []gr.ChangeInfo{
				{
					ID:              "project~branch~Ideadc0de",
					Project:         noWhitelistProject,
					Status:          "NEW",
					CurrentRevision: "abcdef",
					Updated:         gr.TimeStamp(lastChangeTs),
					Revisions:       revisions,
					Owner:           &gr.AccountInfo{Email: "whitelisteduser@example.com"},
				},
			})
			So(pollProject(ctx, noWhitelistProject, api, cp), ShouldBeNil)
			tc.Add(time.Second)
			So(pollProject(ctx, noWhitelistProject, api, cp), ShouldBeNil)
			Convey("Enqueues an analyze request", func() {
				So(numEnqueuedAnalyzeRequests(ctx), ShouldEqual, 1)
			})
		})

		Convey("Poll with a change by a whitelisted user", func() {
			api := &mockPollRestAPI{}
			lastChangeTs := tc.Now().UTC()
			revisions := map[string]gr.RevisionInfo{
				"abcdef": {Files: map[string]*gr.FileInfo{"README.md": {}}},
			}
			api.addChanges(host, whitelistProject, []gr.ChangeInfo{
				{
					ID:              "project~branch~Ideadc0de",
					Project:         whitelistProject,
					Status:          "NEW",
					CurrentRevision: "abcdef",
					Updated:         gr.TimeStamp(lastChangeTs),
					Revisions:       revisions,
					Owner:           &gr.AccountInfo{Email: "whitelisteduser@example.com"},
				},
			})
			So(pollProject(ctx, whitelistProject, api, cp), ShouldBeNil)
			tc.Add(time.Second)
			So(pollProject(ctx, whitelistProject, api, cp), ShouldBeNil)
			Convey("Does not enqueue analyze requests", func() {
				So(numEnqueuedAnalyzeRequests(ctx), ShouldEqual, 1)
			})
		})

		Convey("Poll with a change by an unwhitelisted user", func() {
			api := &mockPollRestAPI{}
			lastChangeTs := tc.Now().UTC()
			revisions := map[string]gr.RevisionInfo{
				"abcdef": {Files: map[string]*gr.FileInfo{"README.md": {}}},
			}
			api.addChanges(host, whitelistProject, []gr.ChangeInfo{
				{
					ID:              "project~branch~Ideadc0de",
					Project:         whitelistProject,
					Status:          "NEW",
					CurrentRevision: "abcdef",
					Updated:         gr.TimeStamp(lastChangeTs),
					Revisions:       revisions,
					Owner:           &gr.AccountInfo{Email: "somebody-else@example.com"},
				},
			})
			So(pollProject(ctx, whitelistProject, api, cp), ShouldBeNil)
			tc.Add(time.Second)
			So(pollProject(ctx, whitelistProject, api, cp), ShouldBeNil)
			Convey("Does not enqueue analyze requests", func() {
				So(numEnqueuedAnalyzeRequests(ctx), ShouldEqual, 0)
			})
		})
	})
}

func TestStatusCode(t *testing.T) {
	ctx := triciumtest.Context()

	Convey("Valid codes", t, func() {
		So(statusFromCode(ctx, "A"), ShouldEqual, tricium.Data_ADDED)
		So(statusFromCode(ctx, "D"), ShouldEqual, tricium.Data_DELETED)
		So(statusFromCode(ctx, "R"), ShouldEqual, tricium.Data_RENAMED)
		So(statusFromCode(ctx, "C"), ShouldEqual, tricium.Data_COPIED)
		So(statusFromCode(ctx, "W"), ShouldEqual, tricium.Data_REWRITTEN)
		So(statusFromCode(ctx, "M"), ShouldEqual, tricium.Data_MODIFIED)
	})

	Convey("Unknown status means modified", t, func() {
		So(statusFromCode(ctx, "X"), ShouldEqual, tricium.Data_MODIFIED)
	})
}
