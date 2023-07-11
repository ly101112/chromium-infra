// Copyright 2023 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package templates

import (
	"context"
	"fmt"
	"os"
	"strings"

	"infra/cros/cmd/cros-tool-runner/internal/v2/commands"

	"go.chromium.org/chromiumos/config/go/test/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type crosVMProvisionProcessor struct {
	cmdExecutor           cmdExecutor
	defaultServerPort     string // Default port used in cros-vm-provision
	dockerArtifactDirName string // Path on the drone where service put the logs by default
}

func newCrosVMProvisionProcessor() *crosVMProvisionProcessor {
	return &crosVMProvisionProcessor{
		cmdExecutor:           &commands.ContextualExecutor{},
		defaultServerPort:     "80",
		dockerArtifactDirName: "/tmp/vm-provision",
	}
}

func (p *crosVMProvisionProcessor) Process(request *api.StartTemplatedContainerRequest) (*api.StartContainerRequest, error) {
	t := request.GetTemplate().GetCrosVmProvision()
	if t == nil {
		return nil, status.Error(codes.Internal, "unable to process")
	}

	port := portZero
	expose := make([]string, 0)
	if request.Network != hostNetworkName {
		port = p.defaultServerPort
		expose = append(expose, port)
	}

	err := generateAuthFile(request.ArtifactDir)
	if err != nil {
		return nil, err
	}

	volume := fmt.Sprintf("%s:%s", request.ArtifactDir, p.dockerArtifactDirName)
	additionalOptions := &api.StartContainerRequest_Options{
		Network: request.Network,
		Expose:  expose,
		Volume:  []string{volume},
	}

	startCommand := []string{
		"vm-provision",
		"-port", port,
		"-log", p.dockerArtifactDirName,
	}
	return &api.StartContainerRequest{Name: request.Name, ContainerImage: request.ContainerImage, AdditionalOptions: additionalOptions, StartCommand: startCommand}, nil
}

func (p *crosVMProvisionProcessor) discoverPort(request *api.StartTemplatedContainerRequest) (*api.Container_PortBinding, error) {
	// delegate to default impl.
	return defaultDiscoverPort(p.cmdExecutor, request)
}

// generateAuthFile generates file with auth token to be consumed by vm-provision
func generateAuthFile(dir string) error {

	//execute gcloud command to generate gcloud auth print-access-token
	gcloudcmd := &commands.GcloudAuthTokenPrint{}
	token, _, err := gcloudcmd.Execute(context.Background())
	if err != nil {
		return status.Error(codes.Internal, "unable to execute gcloud command for vm provision")
	}
	filepath := fmt.Sprintf("%s/%s", dir, "authToken.txt")

	// Create the file
	file, err := os.Create(filepath)
	if err != nil {
		return status.Error(codes.Internal, "unable to create token file for vm provision")
	}
	defer file.Close()

	// Write the token to the file
	_, err = file.WriteString(strings.TrimSpace(token))
	if err != nil {
		return status.Error(codes.Internal, "unable to write to token file for vm provision")
	}
	return nil
}