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

type Config struct{}

// OpenTofuDeployTargetConfig represents the deployment scope configuration.
type OpenTofuDeployTargetConfig struct {
	// The version of OpenTofu to use.
	// Empty means the pre-installed version will be used.
	Version string `json:"version,omitempty"`
	// The working directory for the deployment target.
	WorkingDir string `json:"workingDir,omitempty"`
	// The environment variables for the deployment target.
	Env []string `json:"env,omitempty"`
	// Indicates whether to perform initialization for the deployment target.
	Init bool `json:"init"`
	// The configuration file path.
	Config string `json:"config,omitempty"`
}

// OpenTofuApplicationSpec represents the application scope configuration.
type OpenTofuApplicationSpec struct {
	// The input configuration for OpenTofu deployment stages.
	Input OpenTofuDeploymentInput `json:"input"`
}

func (s *OpenTofuApplicationSpec) Validate() error {
	// TODO: Validate OpenTofuApplicationSpec fields.
	return nil
}

// OpenTofuDeploymentInput is the input for OpenTofu stages.
type OpenTofuDeploymentInput struct {
	// The version of OpenTofu to use.
	Version string `json:"version,omitempty"`
	// The configuration file path for the deployment input.
	Config string `json:"config,omitempty"`
	// The working directory for the deployment input.
	WorkingDir string `json:"workingDir,omitempty"`
	// The environment variables for the deployment input.
	Env []string `json:"env,omitempty"`
	// Indicates whether to perform initialization for the deployment input.
	Init bool `json:"init"`
}

// OpenTofuPlanStageOptions contains all configurable values for a OPENTOFU_PLAN stage.
type OpenTofuPlanStageOptions struct {
	// TODO: Add options for plan stage.
}

// OpenTofuApplyStageOptions contains all configurable values for a OPENTOFU_APPLY stage.
type OpenTofuApplyStageOptions struct {
	// Whether to automatically approve changes during the apply stage.
	// Default: false
	AutoApprove bool `json:"autoApprove"`
}
