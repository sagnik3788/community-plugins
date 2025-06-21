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
	Version           string   `json:"version"`
	DefaultConfig     string   `json:"defaultConfig"`
	DefaultWorkingDir string   `json:"defaultWorkingDir"`
	DefaultEnv        []string `json:"defaultEnv"`
	DefaultInit       bool     `json:"defaultInit"`
}

// OpenTofuDeployTargetConfig  represents the deployment scope configuration.
type OpenTofuDeployTargetConfig struct {
	Name   string       `json:"name"`
	Config DeployConfig `json:"config"`
}

type DeployConfig struct {
	Version    string   `json:"version"`
	WorkingDir string   `json:"workingDir"`
	Env        []string `json:"env"`
	Init       bool     `json:"init"`
}

// OpenTofuApplicationSpec  represents the application scope configuration.
type OpenTofuApplicationSpec struct {
	Input     OpenTofuDeploymentInput    `json:"input"`
	QuickSync OpenTofuDeployStageOptions `json:"quickSync"`
}

func (s *OpenTofuApplicationSpec) Validate() error {
	// TODO: Validate OpenTofuApplicationSpec fields.
	return nil
}

// OpenTofuDeploymentInput is the input for OpenTofu stages.
type OpenTofuDeploymentInput struct {
	Version    string   `json:"version"`
	Config     string   `json:"config"`
	WorkingDir string   `json:"workingDir"`
	Env        []string `json:"env"`
	Init       bool     `json:"init"`
}

// OpenTofuDeployStageOptions holds options specific to the quick sync stage.
type OpenTofuDeployStageOptions struct {
	AutoApprove bool `json:"autoApprove"`
}
