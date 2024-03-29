// Copyright 2018 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

//go:build !windows
// +build !windows

package atutil

import (
	"syscall"
)

// Signaled returns true if autoserv exited from a signal.
func (r *Result) Signaled() bool {
	ws := syscall.WaitStatus(r.Exit)
	return ws.Signaled()
}
