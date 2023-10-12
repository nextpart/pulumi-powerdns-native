// Copyright 2016-2023, Pulumi Corporation.
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

package provider

import (
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"

	pdns "github.com/nextpart/pulumi-powerdns/provider/pkg/provider/powerdns"
)

// Version is initialized by the Go linker to contain the semver of this build.
var Version string

const Name string = "powerdns"

func Provider() p.Provider {
	// We tell the provider what resources it needs to support.
	// In this case, a single custom resource.
	return infer.Provider(infer.Options{
		Resources: []infer.InferredResource{
			infer.Resource[pdns.Random, pdns.RandomArgs, pdns.RandomState](),
			infer.Resource[*pdns.Zone, pdns.ZoneArgs, pdns.ZoneState](),
			infer.Resource[*pdns.Record, pdns.RecordArgs, pdns.RecordState](),
		},
		ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
			"powerdns": "index",
		},
		Config: infer.Config[*pdns.Config](),
	})
}



