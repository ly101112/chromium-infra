// Copyright 2016 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

//go:build unix
// +build unix

package infraenv

var (
	// Paths on the system where credentials are stored.
	//
	// This path is provisioned by Puppet.
	systemCredentialDirs = []string{
		"/creds/service_accounts",
	}
)
