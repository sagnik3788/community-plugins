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

	pluginCfg := &config.OpenTofuPluginConfig{
		Version:           "1.6.0",
		DefaultConfig:     "main.tf",
		DefaultWorkingDir: ".",
		DefaultEnv:        []string{"TF_VAR_default=value"},
		DefaultInit:       true,
	}

	deployTarget := &sdk.DeployTarget[config.OpenTofuDeployTargetConfig]{
		Name: "dev",
		Config: config.OpenTofuDeployTargetConfig{
			Config: config.DeployConfig{
				Version:    "1.8.0",
				WorkingDir: "./dev",
				Env:        []string{"TF_VAR_env=dev"},
				Init:       false,
			},
		},
	}

	appSpec := &config.OpenTofuApplicationSpec{
		Input: config.OpenTofuDeploymentInput{
			Version: "1.9.1",
			Config:  "custom.tf",
			Init:    false,
		},
	}

	merged := mergeConfig(pluginCfg, []*sdk.DeployTarget[config.OpenTofuDeployTargetConfig]{deployTarget}, appSpec)

	assert.Equal(t, "1.9.1", merged.Version)                                        // App spec takes highest precedence
	assert.Equal(t, "custom.tf", merged.Config)                                     // App spec takes highest precedence
	assert.Equal(t, "./dev", merged.WorkingDir)                                     // Deploy target takes precedence over plugin
	assert.Equal(t, false, merged.Init)                                             // Deploy target takes precedence over plugin
	assert.Equal(t, []string{"TF_VAR_default=value", "TF_VAR_env=dev"}, merged.Env) // Both envs combined
}
