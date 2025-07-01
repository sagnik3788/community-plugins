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

package config

// OpenTofuPluginConfig represents the plugin scope configuration.
type OpenTofuPluginConfig struct {
	// The version of the plugin configuration.
	Version string `json:"version"`
	// The default configuration file path.
	DefaultConfig string `json:"defaultConfig"`
	// The default working directory for the plugin.
	DefaultWorkingDir string `json:"defaultWorkingDir"`
	// The default environment variables for the plugin.
	DefaultEnv []string `json:"defaultEnv"`
	// Indicates whether to perform initialization by default.
	DefaultInit bool `json:"defaultInit"`
}

// OpenTofuDeployTargetConfig  represents the deployment scope configuration.
type OpenTofuDeployTargetConfig struct {
	// The version of the deployment target configuration.
	Version string `json:"version"`
	// The working directory for the deployment target.
	WorkingDir string `json:"workingDir"`
	// The environment variables for the deployment target.
	Env []string `json:"env"`
	// Indicates whether to perform initialization for the deployment target.
	Init bool `json:"init"`
}

// OpenTofuApplicationSpec  represents the application scope configuration.
type OpenTofuApplicationSpec struct {
	// The input configuration for OpenTofu deployment stages.
	Input OpenTofuDeploymentInput `json:"input"`
	// Options specific to the quick sync stage.
	QuickSync OpenTofuSyncStageOptions `json:"quickSync"`
}

func (s *OpenTofuApplicationSpec) Validate() error {
	// TODO: Validate OpenTofuApplicationSpec fields.
	return nil
}

// OpenTofuDeploymentInput is the input for OpenTofu stages.
type OpenTofuDeploymentInput struct {
	// The version of the deployment input.
	Version string `json:"version"`
	// The configuration file path for the deployment input.
	Config string `json:"config"`
	// The working directory for the deployment input.
	WorkingDir string `json:"workingDir"`
	// The environment variables for the deployment input.
	Env []string `json:"env"`
	// Indicates whether to perform initialization for the deployment input.
	Init bool `json:"init"`
}

// OpenTofuDeployStageOptions holds options specific to the quick sync stage.
type OpenTofuSyncStageOptions struct {
	// Indicates whether to automatically approve changes during the quick sync stage.
	AutoApprove bool `json:"autoApprove"`
}
