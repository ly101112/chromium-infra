// Copyright 2022 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"infra/tools/vpython_ng/pkg/application"
	"infra/tools/vpython_ng/pkg/python"
	"infra/tools/vpython_ng/pkg/wheels"

	"go.chromium.org/luci/common/errors"
)

type PythonRuntime struct {
	Version     string
	Executable  string
	CIPDName    string
	SpecPattern string
}

func GetPythonRuntime(ver string) *PythonRuntime {
	return map[string]*PythonRuntime{
		"2.7": {
			Version:     "2.7",
			Executable:  "python",
			CIPDName:    "cpython",
			SpecPattern: ".vpython",
		},
		"3.8": {
			Version:     "3.8",
			Executable:  "python3",
			CIPDName:    "cpython3",
			SpecPattern: ".vpython3",
		},
	}[ver]
}

func DefaultPythonVersion() string {
	switch filepath.Base(os.Args[0]) {
	case "vpython":
		return "2.7"
	default:
		return "3.8"
	}
}

func main() {
	rt := GetPythonRuntime(DefaultPythonVersion())

	app := application.Application{
		PruneThreshold:    7 * 24 * time.Hour, // One week.
		MaxPrunesPerSweep: 3,

		Environments: os.Environ(),
		Arguments:    os.Args[1:],
	}
	app.Initialize()

	app.Must(app.ParseEnvs())
	app.Must(app.ParseArgs())

	if app.Bypass {
		// no-op for tool mode if we are bypassing vpython
		if app.ToolMode != "" {
			return
		}
		app.PythonExecutable = rt.Executable
		app.Must(app.ExecutePython())
		return
	}

	app.Must(app.LoadSpec())

	if v := app.VpythonSpec.PythonVersion; v != "" {
		rt = GetPythonRuntime(v)
	}
	cpython, err := python.CPythonFromRelativePath(rt.Version, rt.CIPDName)
	if err != nil {
		app.Fatal(err)
	}
	app.PythonExecutable = rt.Executable

	env := python.Environment{
		Executable: rt.Executable,
		CPython:    cpython,
		Virtualenv: python.VirtualenvFromCIPD("version:2@16.7.10.chromium.7"),
	}
	wheel, err := wheels.FromSpec(app.VpythonSpec, env.Pep425Tags())
	if err != nil {
		app.Fatal(err)
	}
	venv := env.WithWheels(wheel)

	if !app.Help && app.ToolMode != "" {
		app.Must(func() error {
			switch app.ToolMode {
			case "install":
				app.PruneThreshold = 0
				return app.BuildVENV(venv)
			case "verify":
				return wheels.Verify(app.VpythonSpec)
			default:
				return errors.Reason("unknown -vpython-tool command: %s", app.ToolMode).Err()
			}
		}())
		return
	}

	if app.Help {
		// Continue to execute python to print its help message after vpython's.
		fmt.Println(app.Usage)
	}
	app.Must(app.BuildVENV(venv))
	app.Must(app.ExecutePython())
}
