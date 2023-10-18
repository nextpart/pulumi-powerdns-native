// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	powerdns "github.com/nextpart/pulumi-powerdns-native/provider/pkg/provider"
)

// copied from encoding/json for use with JSONMarshal above
func MarshalIndent(v any) ([]byte, error) {
	// json.Marshal normally escapes HTML. This one doesn't
	// https://stackoverflow.com/questions/28595664/how-to-stop-json-marshal-from-escaping-and
	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "    ")
	err := encoder.Encode(v)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func main() {
	flag.Usage = func() {
		const usageFormat = "Usage: %s <schema-file>"
		fmt.Fprintf(flag.CommandLine.Output(), usageFormat, os.Args[0])
		flag.PrintDefaults()
	}

	var version string
	flag.StringVar(&version, "version", powerdns.Version, "the provider version to record in the generated code")

	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		return
	}
	s, err := powerdns.Schema(version)
	if err != nil {
		panic(err)
	}

	// sort keys
	var arg map[string]any
	err = json.Unmarshal([]byte(s), &arg)
	if err != nil {
		panic(err)
	}

	// remove version key
	delete(arg, "version")

	// use custom marshal indent to skip html escaping
	out, err := MarshalIndent(arg)
	if err != nil {
		panic(err)
	}

	schemaPath := args[0]
	err = os.WriteFile(schemaPath, out, 0600)
	if err != nil {
		panic(err)
	}
}