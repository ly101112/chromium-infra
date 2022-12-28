// Copyright 2022 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package templates

import (
	"fmt"
	"log"
	"os"
	"path"

	"go.chromium.org/chromiumos/config/go/test/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type crosTestProcessor struct {
	defaultPortDiscoverer portDiscoverer
	defaultServerPort     string // Default port used in cros-test
}

func newCrosTestProcessor() TemplateProcessor {
	return &crosTestProcessor{
		defaultPortDiscoverer: &defaultPortDiscoverer{},
		defaultServerPort:     "8001",
	}
}

func (p *crosTestProcessor) Process(request *api.StartTemplatedContainerRequest) (*api.StartContainerRequest, error) {
	t := request.GetTemplate().GetCrosTest()
	if t == nil {
		return nil, status.Error(codes.Internal, "unable to process")
	}

	port := portZero
	expose := make([]string, 0)
	if t.Network != hostNetworkName {
		port = p.defaultServerPort
		expose = append(expose, port)
	}
	// All non-test harness artifacts will be in <artifact_dir>/cros-test/cros-test.
	crosTestDir := path.Join(t.ArtifactDir, "cros-test", "cros-test")
	// All test result artifacts will be in <artifact_dir>/cros-test/results.
	resultDir := path.Join(t.ArtifactDir, "cros-test", "results")
	// Setting up directories. Required as podman doesn't create directories for volume mounting.
	p.createDir(crosTestDir)
	p.createDir(resultDir)
	volumes := []string{
		fmt.Sprintf("%s:%s", crosTestDir, "/tmp/test/cros-test"),
		fmt.Sprintf("%s:%s", resultDir, "/tmp/test/results"),
	}
	// Mount authorization file for gsutil if exists. See b/239855913
	gsutilAuthFile := "/home/chromeos-test/.boto"
	if _, err := os.Stat(gsutilAuthFile); err == nil {
		volumes = append(volumes, fmt.Sprintf("%s:%s", gsutilAuthFile, gsutilAuthFile))
	}
	// Mount autotest results shared folder if exists. See b/239855163
	autotestResultsFolder := "/usr/local/autotest/results/shared"
	if _, err := os.Stat(autotestResultsFolder); err == nil {
		volumes = append(volumes, fmt.Sprintf("%s:%s", autotestResultsFolder, autotestResultsFolder))
	}
	additionalOptions := &api.StartContainerRequest_Options{
		Network: t.Network,
		Expose:  expose,
		Volume:  volumes,
	}
	// It is necessary to do sudo here because /tmp/test is owned by root inside docker
	// when docker mount /tmp/test. However, the user that is running cros-test is
	// chromeos-test inside docker. Hence, the user chromeos-test does not have write
	// permission in /tmp/test. Therefore, we need to change the owner of the directory.
	cmd := fmt.Sprintf("sudo --non-interactive chown -R chromeos-test:chromeos-test %s && cros-test server -port %s", "/tmp/test", port)
	startCommand := []string{"bash", "-c", cmd}
	return &api.StartContainerRequest{Name: request.Name, ContainerImage: request.ContainerImage, AdditionalOptions: additionalOptions, StartCommand: startCommand}, nil
}

func (p *crosTestProcessor) discoverPort(request *api.StartTemplatedContainerRequest) (*api.Container_PortBinding, error) {
	t := request.GetTemplate().GetCrosTest()
	if t == nil {
		return nil, status.Error(codes.Internal, "unable to process3")
	}
	portBinding, err := p.defaultPortDiscoverer.discoverPort(request)
	if err != nil {
		return portBinding, err
	}
	if t.Network == hostNetworkName {
		portBinding.HostPort = portBinding.ContainerPort
		portBinding.HostIp = localhostIp
	}
	portBinding.Protocol = protocolTcp
	return portBinding, nil
}

// createDir creates artifact subdirectories for the given path.
func (p *crosTestProcessor) createDir(dirPath string) {
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		log.Printf("warning: cros-test template processor received error when creating directory %s: %v", dirPath, err)
	}
	log.Printf("cros-test template processor has created directory %s", dirPath)
}
