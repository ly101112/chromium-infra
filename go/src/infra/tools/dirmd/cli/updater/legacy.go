// Copyright 2020 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package updater

import (
	"encoding/json"
	"fmt"

	"infra/tools/dirmd"
	dirmdpb "infra/tools/dirmd/proto"
)

func toLegacyFormat(m *dirmd.Mapping) []byte {
	dirToComponent := make(map[string]string, len(m.Dirs))
	dirToTeam := make(map[string]string, len(m.Dirs))

	for dir, meta := range m.Dirs {
		if meta.TeamEmail != "" {
			dirToTeam[dir] = meta.TeamEmail
		}
		if comp := legacyComponent(meta); comp != "" {
			dirToComponent[dir] = comp
		}
	}

	marshalled, err := json.MarshalIndent(map[string]interface{}{
		"AAA-README": []string{
			"",
			"This file is generated by infra.git/go/src/tools/dirmd/cli/updater",
			"by parsing the contents of OWNERS and DIR_METADATA files throughout the chromium source code.",
			"",
			"Manual edits of this file will be overwritten by an automated process.",
		},
		"dir-to-component": dirToComponent,
		"dir-to-team":      dirToTeam,

		// The existence of these keys is important to some legacy applications,
		// but correctness is not.
		"components-to-team":  map[string]interface{}{},
		"teams-per-component": map[string]interface{}{},
	}, "", "  ")
	if err != nil {
		panic(err)
	}
	return marshalled
}

func legacyComponent(meta *dirmdpb.Metadata) string {
	ret := meta.GetMonorail().GetComponent()
	if ret == "" {
		return ret
	}

	if os := legacyOS(meta.Os); os != "" {
		ret += fmt.Sprintf("(%s)", os)
	}
	return ret
}

func legacyOS(os dirmdpb.OS) string {
	switch os {
	case dirmdpb.OS_LINUX:
		return "Linux"
	case dirmdpb.OS_WINDOWS:
		return "Windows"
	case dirmdpb.OS_MAC:
		return "Mac"
	case dirmdpb.OS_ANDROID:
		return "Android"
	case dirmdpb.OS_IOS:
		return "iOS"
	case dirmdpb.OS_CHROME_OS:
		return "ChromeOS"
	case dirmdpb.OS_FUCHSIA:
		return "Fuchsia"
	default:
		return ""
	}
}
