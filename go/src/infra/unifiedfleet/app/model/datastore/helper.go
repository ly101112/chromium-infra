// Copyright 2020 The Chromium OS Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package datastore

import (
	"fmt"
	"strings"
)

// GetServoID returns the servo_id for searching
func GetServoID(servoHostname string, servoPort int32) string {
	if servoHostname != "" && servoPort != 0 {
		return fmt.Sprintf("%s%d", servoHostname, servoPort)
	}
	return ""
}

// GetOSIndex returns a slics of strings for a given string
func GetOSIndex(osversion string) []string {
	str := strings.TrimSpace(osversion)
	str = strings.ToLower(str)
	str = strings.Replace(str, " ", "_", -1)
	str = strings.Replace(str, ",", "_", -1)
	return strings.Split(str, "_")
}
