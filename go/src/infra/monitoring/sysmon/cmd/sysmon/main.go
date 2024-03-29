// Copyright (c) 2016 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/logging/gologger"
	"go.chromium.org/luci/common/tsmon"

	"infra/monitoring/sysmon/android"
	"infra/monitoring/sysmon/cipd"
	"infra/monitoring/sysmon/cros"
	"infra/monitoring/sysmon/docker"
	"infra/monitoring/sysmon/puppet"
	"infra/monitoring/sysmon/system"

	"infra/cmdsupport/service"
)

func main() {
	fs := flag.NewFlagSet("", flag.ExitOnError)

	tsmonFlags := tsmon.NewFlags()
	tsmonFlags.Flush = "auto"
	tsmonFlags.Register(fs)

	loggingConfig := logging.Config{Level: logging.Info}
	loggingConfig.AddFlags(fs)

	fs.Parse(os.Args[1:])

	c := context.Background()
	c = gologger.StdConfig.Use(c)
	c = loggingConfig.Set(c)

	if err := tsmon.InitializeFromFlags(c, &tsmonFlags); err != nil {
		panic(fmt.Sprintf("failed to initialize tsmon: %s", err))
	}

	// Register metric callbacks.
	android.Register()
	cipd.Register()
	cros.Register()
	docker.Register()
	puppet.Register()
	system.Register() // Should be registered last.

	if tsmonFlags.Flush == "auto" {
		stop := make(chan struct{})

		// Support running as a Windows Service, or as a regular background process.
		s := &service.Service{
			Start: func() int {
				// tsmon's auto-flusher goroutine will call the metric callbacks and flush
				// the metrics every minute.
				<-stop
				return 0
			},
			Stop: func() {
				close(stop)
			},
		}
		os.Exit(service.Run(s))
	} else {
		// Flush once and exit.
		tsmon.Flush(c)
	}
}
