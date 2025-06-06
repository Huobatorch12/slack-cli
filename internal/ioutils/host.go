// Copyright 2022-2025 Salesforce, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ioutils

import (
	"os"

	"github.com/toughtackle/slack-cli/internal/goutils"
)

func GetHostname() string {
	// Storing hostname provides context about where errors are coming from.
	// It is also useful for rate limiting.
	// Regardless of its utility, since it may be PII, we should hash it.
	var hostname, err = os.Hostname()
	if err != nil {
		hostname = "unknown"
	} else {
		// hash the string
		hostname, err = goutils.HashString(hostname)
		if err != nil {
			hostname = "unknown"
		}
	}
	return hostname
}
