// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package acceptance_test

import (
	"testing"

	"github.com/GoogleCloudPlatform/buildpacks/internal/acceptance"
)

func TestGCFAcceptanceDotNet(t *testing.T) {
	// TODO(mattrobertson): Remove when stack images are published.
	if acceptance.PullImages() {
		t.Skip("Disabled for continuous builds")
	}
	builder, cleanup := acceptance.CreateBuilder(t)
	t.Cleanup(cleanup)

	testCases := []acceptance.Test{
		{
			Name: "cs single target",
			App:  "cs_single_target",
			Env:  []string{"GOOGLE_FUNCTION_TARGET=TestFunction.Function"},
			Path: "/function",
		},
		{
			Name: "cs multiple targets",
			App:  "cs_multiple_targets",
			Env:  []string{"GOOGLE_FUNCTION_TARGET=TestFunction.Function"},
			Path: "/function",
		},
		{
			Name: "cs nested configuration",
			App:  "cs_nested_configuration",
			Env:  []string{"GOOGLE_FUNCTION_TARGET=TestFunction.Function"},
			Path: "/function",
		},
		{
			Name: "fs function",
			App:  "fs_function",
			Env:  []string{"GOOGLE_FUNCTION_TARGET=fs_function.Function"},
			Path: "/function",
		},
		{
			Name: "vb function",
			App:  "vb_function",
			Env:  []string{"GOOGLE_FUNCTION_TARGET=vb_function.CloudFunction"},
			Path: "/function",
		},
	}

	for _, tc := range testCases {
		tc.Env = append(tc.Env, "GOOGLE_RUNTIME=dotnet3", "X_GOOGLE_TARGET_PLATFORM=gcf")
		tc.FilesMustExist = append(tc.FilesMustExist,
			"/layers/google.utils.archive-source/src/source-code.tar.gz",
			"/workspace/.googlebuild/source-code.tar.gz",
		)
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()
			acceptance.TestApp(t, builder, tc)
		})
	}
}
