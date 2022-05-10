// Copyright 2020 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package controller

import (
	"context"
	"fmt"
	"time"

	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/server/auth"
	"google.golang.org/grpc/codes"
	grpcStatus "google.golang.org/grpc/status"

	ufspb "infra/unifiedfleet/api/v1/models"
	api "infra/unifiedfleet/api/v1/rpc"
	"infra/unifiedfleet/app/model/configuration"
)

const (
	// LUCI Auth group which is used to verify if a service account has permissions to run public Chromium tests in ChromeOS lab
	PublicUsersToChromeOSAuthGroup = "public-chromium-in-chromeos-builders"

	// Date Format to parse launch date for Device info read from DLM
	DateFormat = "2006-01-02"
)

// InvalidBoardError is the error raised when a private board is specified for a public test
type InvalidBoardError struct {
	Board string
}

func (e *InvalidBoardError) Error() string {
	return fmt.Sprintf("Cannnot run public tests on a private board : %s", e.Board)
}

// InvalidModelError is the error raised when a private model is specified for a public test
type InvalidModelError struct {
	Model string
}

func (e *InvalidModelError) Error() string {
	return fmt.Sprintf("Cannot run public tests on a private model : %s", e.Model)
}

// InvalidImageError is the error raised when an invalid image is specified for a public test
type InvalidImageError struct {
	Image string
}

func (e *InvalidImageError) Error() string {
	return fmt.Sprintf("Cannot run public tests on an image which is not allowlisted : %s", e.Image)
}

// InvalidTestError is the error raised when an invalid image is specified for a public test
type InvalidTestError struct {
	TestName string
}

func (e *InvalidTestError) Error() string {
	return fmt.Sprintf("Public user cannnot run the not allowlisted test : %s", e.TestName)
}

func IsValidTest(ctx context.Context, req *api.CheckFleetTestsPolicyRequest) error {
	logging.Infof(ctx, "Request to check from crosfleet: %s", req)
	logging.Infof(ctx, "Service account being validated: %s", auth.CurrentIdentity(ctx).Email())
	isMemberInPublicGroup, err := auth.IsMember(ctx, PublicUsersToChromeOSAuthGroup)
	if err != nil {
		// Ignoring error for now till we validate the service account membership check is correct
		logging.Errorf(ctx, "Request to check public chrome auth group membership failed: %s", err)
		return nil
	}

	if !isMemberInPublicGroup {
		return nil
	}

	// Validate if the board and model are public
	if req.Board == "" {
		return grpcStatus.Errorf(codes.InvalidArgument, "Invalid input - Board cannot be empty for public tests.")
	}
	if !contains(getValidPublicBoards(), req.Board) {
		return &InvalidBoardError{Board: req.Board}
	}
	if req.Model == "" {
		return grpcStatus.Errorf(codes.InvalidArgument, "Invalid input - Model cannot be empty for public tests.")
	}
	if !contains(getValidPublictModels(), req.Model) {
		return &InvalidModelError{Model: req.Model}
	}

	// Validate Test Name
	if req.TestName == "" {
		return grpcStatus.Errorf(codes.InvalidArgument, "Invalid input - Test name cannot be empty for public tests.")
	}
	if !contains(getValidPublicTestNames(), req.TestName) {
		return &InvalidTestError{TestName: req.TestName}
	}

	// Validate Image
	if req.Image == "" {
		return grpcStatus.Errorf(codes.InvalidArgument, "Invalid input - Image cannot be empty for public tests.")
	}
	if !contains(getValidPublicImages(), req.Image) {
		return &InvalidImageError{Image: req.Image}
	}

	return nil
}

func ImportPublicBoardsAndModels(ctx context.Context, goldenEyeDevices *ufspb.GoldenEyeDevices) error {
	for _, device := range goldenEyeDevices.Devices {
		if device.LaunchDate == "" {
			continue
		}
		launchDate, err := time.Parse(DateFormat, device.LaunchDate)
		if err != nil {
			// Ignore and process the rest of the data
			logging.Infof(ctx, "Failed to parse Launch Date from Golden Eye Device data %s", device.LaunchDate)
			continue
		}
		if launchDate.Before(time.Now()) {
			// Already launched board and model, can be added to allowed list
			for _, board := range device.Boards {
				logging.Infof(ctx, "Launched Board from Golden Eye Device data %s", board.PublicCodename)
				var modelNames []string
				for _, model := range board.Models {
					modelNames = append(modelNames, model.Name)
				}
				configuration.AddPublicBoardModelData(ctx, board.PublicCodename, modelNames)
			}
		}
	}
	return nil
}

func getValidPublicTestNames() []string {
	return []string{"tast.lacros"}
}

func getValidPublicBoards() []string {
	return []string{"eve", "octopus"}
}

func getValidPublictModels() []string {
	return []string{"eve", "octopus"}
}

func getValidPublicImages() []string {
	return []string{"eve-full/R100-14495.0.0-rc1", "octopus-full/R101-14543.0.0-rc1"}
}

func contains(listItems []string, name string) bool {
	for _, item := range listItems {
		if name == item {
			return true
		}
	}
	return false
}
