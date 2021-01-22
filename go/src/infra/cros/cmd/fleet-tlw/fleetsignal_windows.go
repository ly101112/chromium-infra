// Copyright 2021 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// +build windows

package main

import (
	"os"
)

func setupSignalHandler() <-chan os.Signal {
	panic("windows not supported")
}
