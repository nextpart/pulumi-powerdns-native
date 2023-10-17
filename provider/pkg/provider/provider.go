/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 *
 * This file incorporates work covered by the following copyright and
 * permission notice:
 *
 *   Copyright 2016-2023, Pulumi Corporation.
 *
 *   Licensed under the Apache License, Version 2.0 (the "License");
 *   you may not use this file except in compliance with the License.
 *   You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *   Unless required by applicable law or agreed to in writing, software
 *   distributed under the License is distributed on an "AS IS" BASIS,
 *   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *   See the License for the specific language governing permissions and
 *   limitations under the License.
 */

package provider

import (
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
	"github.com/pulumi/pulumi-go-provider/middleware/schema"

	pdns "github.com/nextpart/pulumi-powerdns-native/provider/pkg/provider/powerdns"
)

// Version is initialized by the Go linker to contain the semver of this build.
var Version string

const Name string = "powerdns"

func Provider() p.Provider {
	// We tell the provider what resources it needs to support.
	// In this case, a single custom resource.
	return infer.Provider(infer.Options{
		Metadata: schema.Metadata{
			DisplayName: "PowerDNS",
			Description: "The Pulumi PowerDNS provider provides resources to interact with a PowerDNS Authoritative DNS Server.",
			Keywords: []string{
				"pulumi",
				"powerdns",
				"dns",
			},
			Homepage:          "https://pulumi.com",
			License:           "MPL-2.0",
			Repository:        "https://github.com/nextpart/pulumi-powerdns-native",
			PluginDownloadURL: "github://api.github.com/nextpart/pulumi-powerdns-native",
			Publisher:         "Nextpart",
			LogoURL:           "https://raw.githubusercontent.com/nextpart/pulumi-powerdns-native/master/assets/logo.png",
			// This contains language specific details for generating the provider's SDKs
			LanguageMap: map[string]any{
				"csharp": map[string]any{
					"packageReferences": map[string]string{
						"Pulumi": "3.*",
					},
				},
				"go": map[string]any{
					"generateResourceContainerTypes": true,
					"importBasePath":                 "github.com/nextpart/pulumi-powerdns-native/sdk/go/powerdns",
				},
				"nodejs": map[string]any{
					"packageName": "@nextpart/powerdns",
					"dependencies": map[string]string{
						"@pulumi/pulumi": "^3.0.0",
					},
				},
				"python": map[string]any{
					"requires": map[string]string{
						"pulumi": ">=3.0.0,<4.0.0",
					},
				},
				"java": map[string]any{
					"buildFiles":                      "gradle",
					"gradleNexusPublishPluginVersion": "1.1.0",
					"dependencies": map[string]any{
						"com.pulumi:pulumi":               "0.6.0",
						"com.google.code.gson:gson":       "2.8.9",
						"com.google.code.findbugs:jsr305": "3.0.2",
					},
				},
			},
		},
		Resources: []infer.InferredResource{
			infer.Resource[*pdns.Zone, pdns.ZoneArgs, pdns.ZoneState](),
			infer.Resource[*pdns.Record, pdns.RecordArgs, pdns.RecordState](),
		},
		ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
			"powerdns": "index",
		},
		Config: infer.Config[*pdns.Config](),
	})
}
