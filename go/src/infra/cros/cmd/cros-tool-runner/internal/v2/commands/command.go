// Copyright 2022 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Package commands hosts all CLI commands CTRv2 interacts with. Each command is
// a struct that represents a CLI command, holds caller-specified arguments, and
// implements an API to uniformly execute itself.
package commands

import (
	"context"
	"os/exec"
	"strings"
	"time"

	"infra/cros/cmd/cros-tool-runner/internal/common"
)

// Command is the interface of the command pattern. Only support blocking
// execution for now.
type Command interface {
	Execute(context.Context) (string, string, error)
}

// execute runs blocking command and returns stdout and stderr as strings.
func execute(ctx context.Context, name string, args []string) (string, string, error) {
	cmd := exec.CommandContext(ctx, name, args...)
	// TODO(mingkong) update RunWithTimeout since timeout is part of ctx
	return common.RunWithTimeout(ctx, cmd, time.Minute, true)
}

var utils = commandUtils{}

// commandUtils groups all utility methods for the package.
type commandUtils struct {
}

// trim removes leading and tailing whitespaces.
func (*commandUtils) trim(s string) string {
	return strings.TrimSpace(s)
}
