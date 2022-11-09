// Copyright 2022 The ChromiumOS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// CBI corruption detection and repair logic. go/cbi-auto-recovery-dd
package cbi

import (
	"context"
	"regexp"
	"strconv"
	"strings"
	"time"

	"infra/cros/recovery/internal/components"

	labapi "go.chromium.org/chromiumos/config/go/test/lab/api"
	"go.chromium.org/luci/common/errors"
)

// CBILocation stores the port and address needed to reference CBI contents in
// EEPROM.
type CBILocation struct {
	port    string
	address string
}

const (
	// The first three initial bytes in EEPROM indicating that the chip contains
	// CBI contents. CBI contents on ALL DUTs should start with these three bytes.
	cbiMagic = "0x43 0x42 0x49"

	locateCBICommand   = "ectool locatechip"
	cbiChipType        = "0" // Maps to CBI in the `ectool locatechip` utility
	cbiIndex           = "0" // Gets the first CBI chip (there is only ever one) on the DUT.
	transferCBICommand = "ectool i2cxfer"
	cbiSize            = 256 // How many bytes of memory are stored in CBI.

	// How many bytes can be read from CBI in a single operation.
	// THIS VALUE SHOULD BE TREATED AS A HARD LIMIT. Exceeding this limit may
	// result in undefined behavior.
	readIncrement = 64
)

// How much time each CBI read and write (one call to `ectool i2cxfer`) has to complete.
// This is exceedingly generous, as each read or write should take on the order
// of milliseconds to execute. This is primarily to give some slack to DUTs experiencing connection issues.
var transferCBIDurationInSeconds = time.Second * 10

var readCBIRegex = regexp.MustCompile(`0x[[:xdigit:]]{1,2}|00`) // Match bytes printed in hex format (e.g. 00, 0x12, 0x3)
var locateCBIRegex = regexp.MustCompile(`Port:\s(\d+).*Address:\s(0x\w+)`)

// GetCBILocation uses the `ectool locatechip` utility to get the CBILocation
// from the DUT. Will return an error if the DUT doesn't support CBI or if it
// wasn't able to reach the DUT.
func GetCBILocation(ctx context.Context, run components.Runner) (*CBILocation, error) {
	locateCBIOutput, err := run(ctx, time.Second*30, locateCBICommand, cbiChipType, cbiIndex)
	if err != nil {
		return nil, errors.Annotate(err, "get CBI location: unable to determine if CBI is present on the DUT").Err()
	}

	cbiLocation, err := buildCBILocation(locateCBIOutput)
	return cbiLocation, errors.Annotate(err, "get CBI location").Err()
}

// buildCBILocation creates a CBILocation struct from the text output of an
// `ectool locatechip` call. Will return an error if the locateCBIOutput doesn't
// contain both the address and the port needed to access the CBI contents.
func buildCBILocation(locateCBIOutput string) (*CBILocation, error) {
	match := locateCBIRegex.FindStringSubmatch(locateCBIOutput)
	if len(match) < 3 {
		return nil, errors.Reason("build CBI location: no CBI contents were found on the DUT").Err()
	}
	return &CBILocation{
		port:    match[1],
		address: match[2],
	}, nil
}

// ReadCBIContents reads all <cbiSize> bytes from the CBI chip on the DUT in
// <readIncrement> sized reads using the ectool i2cxfer utility and returns a
// fully formed CBI proto.
func ReadCBIContents(ctx context.Context, run components.Runner, cbiLocation *CBILocation) (*labapi.Cbi, error) {
	hexContents := []string{}
	for offset := 0; offset < cbiSize; offset += readIncrement {
		cbiContents, err := run(ctx, time.Second*10, transferCBICommand, cbiLocation.port, cbiLocation.address, strconv.Itoa(readIncrement), strconv.Itoa(offset))
		if err != nil {
			return nil, errors.Annotate(err, "read CBI contents: unable to read CBI contents").Err()
		}

		hexBytes, err := parseBytesFromCBIContents(cbiContents, readIncrement)
		if err != nil {
			return nil, err
		}

		hexContents = append(hexContents, hexBytes...)
	}
	return &labapi.Cbi{RawContents: strings.Join(hexContents, " ")}, nil
}

// GetCBIContents uses GetCBILocation and ReadCBILocation to both locate and
// and retrieve all CBI contents from the DUT.
func GetCBIContents(ctx context.Context, run components.Runner) (*labapi.Cbi, error) {
	cbiLocation, err := GetCBILocation(ctx, run)
	if err != nil {
		return nil, errors.Annotate(err, "get CBI contents").Err()
	}

	dutCBI, err := ReadCBIContents(ctx, run, cbiLocation)
	if err != nil {
		return nil, errors.Annotate(err, "get CBI contents").Err()
	}
	return dutCBI, nil
}

// WriteCBIContents writes the provided CBI contents to the DUT in
// writeIncrement sized chunks. Propagates any errors from `ectool i2cxfer`.
func WriteCBIContents(ctx context.Context, run components.Runner, cbiLocation *CBILocation, cbi *labapi.Cbi) error {
	const (
		// How many bytes to read during our write operation. This is a quirk of the
		// ectool i2cxfer API, and is always zero when writing.
		numBytesToReadDuringWrite = "0"

		// How many bytes can be written to CBI in a single operation.
		// THIS VALUE SHOULD BE TREATED AS A HARD LIMIT. Exceeding this limit may
		// result in undefined behavior.
		writeIncrement = 8
	)
	hexBytes, err := parseBytesFromCBIContents(cbi.GetRawContents(), cbiSize)
	if err != nil {
		return errors.Annotate(err, "write CBI contents").Err()
	}
	for offset := 0; offset < cbiSize; offset += writeIncrement {
		hexByteChunk := strings.Join(hexBytes[offset:offset+writeIncrement], " ")
		// Sample command:ectool i2cxfer 0 0x50 0 0 0x43 0x42 0x49 0x98 00 00 0x2f 00
		writeResponse, err := run(ctx, transferCBIDurationInSeconds, transferCBICommand, cbiLocation.port, cbiLocation.address, numBytesToReadDuringWrite, strconv.Itoa(offset), hexByteChunk)
		if err != nil {
			return errors.Annotate(err, "write CBI contents: unable to write CBI contents: %s", writeResponse).Err()
		}
	}
	return nil
}

// parseBytesFromCBIContents reads <numBytesToRead> number of bytes from the
// raw output from a call to `ectool i2cxfer` and returns a slice of bytes
// in hex format (the same format returned from `ectool i2cxfer`).
// e.g.
// cbiContents = "Read bytes: 0x43, 0x42, 0x49"
// numBytesToRead = 2
// hexBytes = ["0x43", "0x42"]
func parseBytesFromCBIContents(cbiContents string, numBytesToRead int) ([]string, error) {
	hexBytes := readCBIRegex.FindAllString(cbiContents, numBytesToRead)
	if len(hexBytes) != numBytesToRead {
		return nil, errors.Reason("parse bytes from CBI contents: wrong amount: expected %d bytes but read %d bytes instead. CBI contents found: %s", numBytesToRead, len(hexBytes), cbiContents).Err()
	}
	return hexBytes, nil
}

// ContainsCBIMagic returns true if the rawContents of the CBI proto start with
// the CBI magic bytes.
func ContainsCBIMagic(cbi *labapi.Cbi) bool {
	return strings.HasPrefix(cbi.GetRawContents(), cbiMagic)
}
