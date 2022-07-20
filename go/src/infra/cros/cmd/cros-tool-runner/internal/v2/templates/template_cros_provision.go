package templates

// Copyright 2022 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

import (
	"fmt"
	"log"
	"path"

	"go.chromium.org/chromiumos/config/go/test/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type crosProvisionProcessor struct {
	TemplateProcessor
	placeholderPopulator  placeholderPopulator
	serverPort            string // Default port used in cros-provision
	dockerArtifactDirName string // Path on the drone where service put the logs by default
	inputFileName         string // File in artifact dir to be passed to cros-provision
}

func newCrosProvisionProcessor() *crosProvisionProcessor {
	return &crosProvisionProcessor{
		placeholderPopulator:  newPopulatorRouter(),
		serverPort:            "80",
		dockerArtifactDirName: "/tmp/provisionservice",
		inputFileName:         "in.json",
	}
}

func (p *crosProvisionProcessor) Process(request *api.StartTemplatedContainerRequest) (*api.StartContainerRequest, error) {
	t := request.Template.GetCrosProvision()
	if t == nil {
		return nil, status.Error(codes.Internal, "unable to process")
	}

	// constants TODO(mingkong): define constants with namespacing to avoid typos
	volume := fmt.Sprintf("%s:%s", t.ArtifactDir, p.dockerArtifactDirName)
	additionalOptions := &api.StartContainerRequest_Options{
		Network: t.Network,
		Expose:  []string{p.serverPort},
		Volume:  []string{volume},
	}
	startCommand := []string{
		"cros-provision",
		"server",
		"-metadata", path.Join(p.dockerArtifactDirName, p.inputFileName), // input file flag for cros-provision v2 is metadata
	}
	p.processPlaceholders(request)
	err := p.writeInputFile(request)
	if err != nil {
		return nil, err
	}
	return &api.StartContainerRequest{Name: request.Name, ContainerImage: request.ContainerImage, AdditionalOptions: additionalOptions, StartCommand: startCommand}, nil
}

func (p *crosProvisionProcessor) processPlaceholders(request *api.StartTemplatedContainerRequest) {
	t := request.Template.GetCrosProvision()
	if t.InputRequest.DutServer == nil {
		return
	}
	populatedDutServer, err := p.placeholderPopulator.populate(*t.InputRequest.DutServer)
	if err != nil {
		log.Printf("warning: error %v when processing dut server placeholder %v"+
			" in cros-provision input request, skipping to process template as is",
			err, t.InputRequest.DutServer)
		return
	}
	t.InputRequest.DutServer = &populatedDutServer
}

func (p *crosProvisionProcessor) writeInputFile(request *api.StartTemplatedContainerRequest) error {
	t := request.Template.GetCrosProvision()
	filePath := path.Join(t.ArtifactDir, p.inputFileName)
	return TemplateUtils.writeToFile(filePath, t.InputRequest)
}
