// Copyright 2017 The LUCI Authors.
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

package bugs

import (
	"bytes"
	"fmt"
	"infra/monorail"
	"strings"
	"text/template"

	"golang.org/x/net/context"

	"go.chromium.org/gae/service/info"
	"go.chromium.org/luci/common/errors"

	"infra/appengine/luci-migration/common"
	"infra/appengine/luci-migration/storage"
)

var statusChangedTmpl = template.Must(template.New("").
	Funcs(template.FuncMap{
		"pathEscape": common.PathEscape,
		"percent": func(f float64) string {
			return fmt.Sprintf("%d%%", int(100.0*f))
		},
	}).
	Parse(`
Status changed to "{{.Builder.Migration.Status}}"
{{- if not .Migrated}} (correctness {{.Builder.Migration.Correctness | percent}}, speed {{.Builder.Migration.Speed | percent}})
For the latest status, see https://{{.Hostname}}/masters/{{.Builder.ID.Master|pathEscape}}/builders/{{.Builder.ID.Builder|pathEscape}}
{{end}}
`))

// PostComment posts a comment on the builder bug about the current status.
func PostComment(c context.Context, client ClientFactory, builder *storage.Builder) error {
	tmplArgs := map[string]interface{}{
		"Builder":  builder,
		"Hostname": info.DefaultVersionHostname(c),
		"Migrated": builder.Migration.Status == storage.StatusMigrated,
	}
	contentBuf := &bytes.Buffer{}
	if err := statusChangedTmpl.Execute(contentBuf, tmplArgs); err != nil {
		return errors.Annotate(err, "could not execute comment template").Err()
	}

	var label string
	switch builder.Migration.Status {
	case storage.StatusLUCIWAI, storage.StatusMigrated:
		label = "MigrationStatus-WAI"
	default:
		label = "-MigrationStatus-WAI"
	}

	var status string
	if builder.Migration.Status == storage.StatusMigrated {
		status = monorail.StatusFixed
	}

	req := &monorail.InsertCommentRequest{
		Issue: &monorail.IssueRef{
			ProjectId: builder.IssueID.Project,
			IssueId:   int32(builder.IssueID.ID),
		},
		SendEmail: true,
		Comment: &monorail.InsertCommentRequest_Comment{
			Content: strings.TrimSpace(contentBuf.String()),
			Updates: &monorail.Update{
				Status: status,
				Labels: []string{label},
			},
		},
	}

	_, err := client(builder.IssueID.Hostname).InsertComment(c, req)
	if err != nil {
		return errors.Annotate(err, "InsertComment RPC failed").Err()
	}
	return nil
}
