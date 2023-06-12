// Copyright 2023 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Responsible for the abstraction layer representing each command grouping
package cli

import (
	"errors"
	"fmt"
	"os"
)

// AbstractCommand represents a CLI grouping (e.g.: run as server, run as CLI, etc)
type AbstractCommand interface {
	// Run runs the command
	Run() error

	// Is checks if the string is representative of the command
	Is(string) bool

	// Init is an initializer for the command given the trailing args
	Init([]string) error

	// Name is the command name (for debugging)
	Name() string
}

// ParseInputs is a helper method which parses input arguments. It is
// effectively a factory method.
func ParseInputs() (AbstractCommand, error) {
	if len(os.Args) < 1 {
		return nil, errors.New("CLI arguments must be specified")
	}

	cmds := []AbstractCommand{
		NewCLICommand(),
	}

	subcommand := os.Args[1]
	options := []string{}

	for _, cmd := range cmds {
		options = append(options, cmd.Name())
		if cmd.Is(subcommand) {
			if err := cmd.Init(os.Args[2:]); err != nil {
				return nil, fmt.Errorf("failed to initialize %s command, %s", cmd.Name(), err)
			}
			return cmd, nil
		}
	}

	return nil, fmt.Errorf("unknown subcommand: %s. \nOptions are: [%s]", subcommand, options)
}
