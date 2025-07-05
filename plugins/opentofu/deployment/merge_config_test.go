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
		name     string
		dts      []*sdk.DeployTarget[config.OpenTofuDeployTargetConfig]
		appSpec  *config.OpenTofuApplicationSpec
		expected *config.OpenTofuDeploymentInput
	}{
		{
			name: "App spec takes highest precedence",
			dts: []*sdk.DeployTarget[config.OpenTofuDeployTargetConfig]{
				{
					Name: "dev",
					Config: config.OpenTofuDeployTargetConfig{
						Version:    "1.8.0",
						Config:     "main.tf",
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
					Init:    true,
				},
			},
			expected: &config.OpenTofuDeploymentInput{
				Version:    "1.9.1",
				Config:     "custom.tf",
				WorkingDir: "./dev",
				Env:        []string{"TF_VAR_env=dev"},
				Init:       true,
			},
		},
		{
			name: "Deploy target config provides defaults",
			dts: []*sdk.DeployTarget[config.OpenTofuDeployTargetConfig]{
				{
					Name: "prod",
					Config: config.OpenTofuDeployTargetConfig{
						Version:    "1.8.0",
						Config:     "main.tf",
						WorkingDir: "./prod",
						Env:        []string{"TF_VAR_env=prod"},
						Init:       true,
					},
				},
			},
			expected: &config.OpenTofuDeploymentInput{
				Version:    "1.8.0",
				Config:     "main.tf",
				WorkingDir: "./prod",
				Env:        []string{"TF_VAR_env=prod"},
				Init:       true,
			},
		},
		{
			name: "Empty deploy targets returns zero values",
			dts:  []*sdk.DeployTarget[config.OpenTofuDeployTargetConfig]{},
			expected: &config.OpenTofuDeploymentInput{
				Version:    "",
				Config:     "",
				WorkingDir: "",
				Env:        nil,
				Init:       false,
			},
		},
		{
			name: "Nil app spec uses only deploy target config",
			dts: []*sdk.DeployTarget[config.OpenTofuDeployTargetConfig]{
				{
					Name: "staging",
					Config: config.OpenTofuDeployTargetConfig{
						Version:    "1.7.0",
						Config:     "staging.tf",
						WorkingDir: "./staging",
						Env:        []string{"TF_VAR_env=staging"},
						Init:       false,
					},
				},
			},
			appSpec: nil,
			expected: &config.OpenTofuDeploymentInput{
				Version:    "1.7.0",
				Config:     "staging.tf",
				WorkingDir: "./staging",
				Env:        []string{"TF_VAR_env=staging"},
				Init:       false,
			},
		},
		{
			name: "App spec env appends to deploy target env",
			dts: []*sdk.DeployTarget[config.OpenTofuDeployTargetConfig]{
				{
					Name: "dev",
					Config: config.OpenTofuDeployTargetConfig{
						Version: "1.8.0",
						Config:  "main.tf",
						Env:     []string{"TF_VAR_env=dev"},
					},
				},
			},
			appSpec: &config.OpenTofuApplicationSpec{
				Input: config.OpenTofuDeploymentInput{
					Env: []string{"TF_VAR_app=value"},
				},
			},
			expected: &config.OpenTofuDeploymentInput{
				Version: "1.8.0",
				Config:  "main.tf",
				Env:     []string{"TF_VAR_env=dev", "TF_VAR_app=value"},
				Init:    false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			merged := mergeConfig(tt.dts, tt.appSpec)
			assert.Equal(t, *tt.expected, *merged)
		})
	}
}
