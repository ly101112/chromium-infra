// Copyright 2022 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package templates

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"go.chromium.org/chromiumos/config/go/test/api"
	labApi "go.chromium.org/chromiumos/config/go/test/lab/api"
	"go.chromium.org/luci/common/errors"
	"infra/cros/cmd/cros-tool-runner/internal/v2/commands"
)

// ContainerLookuper provides interface to lookup information for a container
type ContainerLookuper interface {
	LookupContainerPortBindings(name string) ([]*api.Container_PortBinding, error)
}

// templateUtils implements ContainerLookuper
type templateUtils struct {
	ContainerLookuper
}

var TemplateUtils = templateUtils{}

// parsePortBindingString parses the output from `docker container port` command
// The input string example: `81/tcp -> 0.0.0.0:42223`
// Unsupported binding format (e.g. IPv6) is ignored, and both binding and error
// are returned as nil. Other unexpected errors will return error.
func (*templateUtils) parsePortBindingString(input string) (*api.Container_PortBinding, error) {
	r := regexp.MustCompile(`(?P<ContainerPort>\d+)/(?P<Protocol>\w+) -> (?P<HostIp>[\d\\.]+):(?P<HostPort>\d+)`)
	match := r.FindStringSubmatch(input)
	if match == nil {
		log.Printf("warning: ignore unrecognized port binding input %s", input)
		return nil, nil
	}
	containerPort, err := strconv.Atoi(match[1])
	if err != nil {
		return nil, err
	}
	hostPort, err := strconv.Atoi(match[4])
	if err != nil {
		return nil, err
	}
	return &api.Container_PortBinding{
		ContainerPort: int32(containerPort),
		Protocol:      match[2],
		HostIp:        match[3],
		HostPort:      int32(hostPort),
	}, nil
}

// parseMultilinePortBindings parses multiline output from `docker container
// port` command since Docker allows multiple ports to be published in one
// container. However, the CTRv2 server only allows one port to be published.
// Unsupported binding format (e.g. IPv6) is ignored.
func (u *templateUtils) parseMultilinePortBindings(multiline string) ([]*api.Container_PortBinding, error) {
	result := make([]*api.Container_PortBinding, 0)
	for _, line := range strings.Split(multiline, "\n") {
		if line == "" {
			continue
		}
		binding, err := u.parsePortBindingString(line)
		if err != nil {
			return result, err
		}
		if binding != nil {
			result = append(result, binding)
		}
	}
	return result, nil
}

func (*templateUtils) retrieveContainerPortOutputFromCommand(name string) (string, error) {
	cmd := commands.ContainerPort{Name: name}
	stdout, _, err := cmd.Execute(context.Background())
	if err != nil {
		return "", nil
	}
	return strings.TrimSpace(stdout), err
}

// LookupContainerPortBindings is the API to get port bindings for a container
func (u *templateUtils) LookupContainerPortBindings(name string) ([]*api.Container_PortBinding, error) {
	output, err := u.retrieveContainerPortOutputFromCommand(name)
	if err != nil {
		return nil, err
	}
	return u.parseMultilinePortBindings(output)
}

// endpointToAddress converts an endpoint to an address string
func (*templateUtils) endpointToAddress(endpoint *labApi.IpEndpoint) string {
	return fmt.Sprintf("%s:%d", endpoint.Address, endpoint.Port)
}

// writeToFile writes proto message to a file
func (*templateUtils) writeToFile(file string, content proto.Message) error {
	f, err := os.Create(file)
	if err != nil {
		return errors.Annotate(err, "fail to create file %v", file).Err()
	}
	m := jsonpb.Marshaler{}
	if err := m.Marshal(f, content); err != nil {
		return errors.Annotate(err, "fail to marshal request to file %v", file).Err()
	}
	return nil
}
