// Copyright 2023 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"log"
	"os"

	"github.com/maruel/subcommands"

	"go.chromium.org/luci/auth"
	"go.chromium.org/luci/auth/client/authcli"
	"go.chromium.org/luci/hardcoded/chromeinfra"
)

var (
	StdoutLog = log.New(os.Stdout, "", logFlags)
	StderrLog = log.New(os.Stderr, "", logFlags)
	logFlags  = log.LstdFlags | log.Lmicroseconds
)

func getApplication(authOpts auth.Options) *subcommands.DefaultApplication {
	return &subcommands.DefaultApplication{
		Name:  "suitemanager",
		Title: "SuiteManager CLI, go/suitemanager-dd",
		Commands: []*subcommands.Command{
			subcommands.CmdHelp,
			authcli.SubcommandInfo(authOpts, "auth-info", false),
			authcli.SubcommandLogin(authOpts, "auth-login", false),
			authcli.SubcommandLogout(authOpts, "auth-logout", false),
		},
	}
}

type suiteManagerApplication struct {
	*subcommands.DefaultApplication
	stdoutLog *log.Logger
	stderrLog *log.Logger
}

func main() {
	opts := chromeinfra.DefaultAuthOptions()
	s := &suiteManagerApplication{
		getApplication(opts),
		log.New(os.Stdout, "", logFlags),
		log.New(os.Stderr, "", logFlags)}
	os.Exit(subcommands.Run(s, nil))
}
