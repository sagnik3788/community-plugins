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
	sdk "github.com/pipe-cd/piped-plugin-sdk-go"

	"github.com/pipe-cd/community-plugins/plugins/opentofu/config"
)

// mergeConfig merges configuration from plugin scope, deploy target, and application spec
// following the precedence order: application spec > deploy target > plugin scope
func mergeConfig(pluginCfg *config.OpenTofuPluginConfig, dts []*sdk.DeployTarget[config.OpenTofuDeployTargetConfig], appSpec *config.OpenTofuApplicationSpec) *config.OpenTofuDeploymentInput {
	// Start with plugin scope defaults
	merged := &config.OpenTofuDeploymentInput{
		Version:    pluginCfg.Version,
		Config:     pluginCfg.DefaultConfig,
		WorkingDir: pluginCfg.DefaultWorkingDir,
		Env:        append([]string{}, pluginCfg.DefaultEnv...),
		Init:       pluginCfg.DefaultInit,
	}

	// Override with deploy target config if available
	if len(dts) > 0 && dts[0].Name != "" {
		deployCfg := dts[0].Config.Config
		if deployCfg.Version != "" {
			merged.Version = deployCfg.Version
		}
		if deployCfg.WorkingDir != "" {
			merged.WorkingDir = deployCfg.WorkingDir
		}
		if len(deployCfg.Env) > 0 {
			merged.Env = append(merged.Env, deployCfg.Env...)
		}
		// Init from deploy target takes precedence over plugin default
		merged.Init = deployCfg.Init
	}

	// Override with application spec (highest precedence)
	if appSpec != nil {
		appInput := appSpec.Input
		if appInput.Version != "" {
			merged.Version = appInput.Version
		}
		if appInput.Config != "" {
			merged.Config = appInput.Config
		}
		if appInput.WorkingDir != "" {
			merged.WorkingDir = appInput.WorkingDir
		}
		if len(appInput.Env) > 0 {
			merged.Env = append(merged.Env, appInput.Env...)
		}
		// Init from application spec takes highest precedence
		merged.Init = appInput.Init
	}

	return merged
}
