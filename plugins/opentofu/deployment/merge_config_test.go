// Copyright 2025 The PipeCD Authors.
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

package deployment

import (
	"testing"

	sdk "github.com/pipe-cd/piped-plugin-sdk-go"
	"github.com/stretchr/testify/assert"

	"github.com/pipe-cd/community-plugins/plugins/opentofu/config"
)

func TestMergeConfig(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		pluginCfg *config.OpenTofuPluginConfig
		dts       []*sdk.DeployTarget[config.OpenTofuDeployTargetConfig]
		appSpec   *config.OpenTofuApplicationSpec
		expected  *config.OpenTofuDeploymentInput
	}{
		{
			name: "App spec takes highest precedence",
			pluginCfg: &config.OpenTofuPluginConfig{
				Version:           "1.6.0",
				DefaultConfig:     "main.tf",
				DefaultWorkingDir: ".",
				DefaultEnv:        []string{"TF_VAR_default=value"},
				DefaultInit:       true,
			},
			dts: []*sdk.DeployTarget[config.OpenTofuDeployTargetConfig]{
				{
					Name: "dev",
					Config: config.OpenTofuDeployTargetConfig{
						Version:    "1.8.0",
						WorkingDir: "./dev",
						Env:        []string{"TF_VAR_env=dev"},
						Init:       false,
					},
				},
			},
			appSpec: &config.OpenTofuApplicationSpec{
				Input: config.OpenTofuDeploymentInput{
					Version: "1.9.1",
					Config:  "custom.tf",
					Init:    false,
				},
			},
			expected: &config.OpenTofuDeploymentInput{
				Version:    "1.9.1",
				Config:     "custom.tf",
				WorkingDir: "./dev",
				Env:        []string{"TF_VAR_default=value", "TF_VAR_env=dev"},
				Init:       false,
			},
		},
		{
			name: "Deploy target config overrides plugin config",
			pluginCfg: &config.OpenTofuPluginConfig{
				Version:           "1.6.0",
				DefaultConfig:     "main.tf",
				DefaultWorkingDir: ".",
				DefaultEnv:        []string{"TF_VAR_default=value"},
				DefaultInit:       true,
			},
			dts: []*sdk.DeployTarget[config.OpenTofuDeployTargetConfig]{
				{
					Name: "dev",
					Config: config.OpenTofuDeployTargetConfig{
						Version:    "1.8.0",
						WorkingDir: "./dev",
						Env:        []string{"TF_VAR_env=prod"},
						Init:       false,
					},
				},
			},
			expected: &config.OpenTofuDeploymentInput{
				Version:    "1.8.0",
				Config:     "main.tf",
				WorkingDir: "./dev",
				Env:        []string{"TF_VAR_default=value", "TF_VAR_env=prod"},
				Init:       false,
			},
		},
		{
			name: "Plugin config defaults applied when no overrides",
			pluginCfg: &config.OpenTofuPluginConfig{
				Version:           "1.6.0",
				DefaultConfig:     "main.tf",
				DefaultWorkingDir: ".",
				DefaultEnv:        []string{"TF_VAR_default=value"},
				DefaultInit:       true,
			},
			expected: &config.OpenTofuDeploymentInput{
				Version:    "1.6.0",
				Config:     "main.tf",
				WorkingDir: ".",
				Env:        []string{"TF_VAR_default=value"},
				Init:       true,
			},
		},
		// TODO: Implement more test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merged := mergeConfig(tt.pluginCfg, tt.dts, tt.appSpec)
			assert.Equal(t, tt.expected.Version, merged.Version)
			assert.Equal(t, tt.expected.Config, merged.Config)
			assert.Equal(t, tt.expected.WorkingDir, merged.WorkingDir)
			assert.Equal(t, tt.expected.Env, merged.Env)
			assert.Equal(t, tt.expected.Init, merged.Init)
		})
	}
}
