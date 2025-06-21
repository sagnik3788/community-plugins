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
	"context"

	sdk "github.com/pipe-cd/piped-plugin-sdk-go"

	"github.com/pipe-cd/community-plugins/plugins/opentofu/config"
)

type Plugin struct{}

// ensure the type Plugin implements sdk.DeploymentPlugin.
var _ sdk.DeploymentPlugin[config.OpenTofuPluginConfig, config.OpenTofuDeployTargetConfig, config.OpenTofuApplicationSpec] = (*Plugin)(nil)

func (*Plugin) FetchDefinedStages() []string {
	return []string{}
}

// BuildPipelineSyncStages builds the stages that will be executed by the plugin.
func (p *Plugin) BuildPipelineSyncStages(ctx context.Context, cfg *config.OpenTofuPluginConfig, input *sdk.BuildPipelineSyncStagesInput) (*sdk.BuildPipelineSyncStagesResponse, error) {
	return &sdk.BuildPipelineSyncStagesResponse{
		Stages: []sdk.PipelineStage{},
	}, nil
}

// ExecuteStage executes the given stage.
func (p *Plugin) ExecuteStage(ctx context.Context, cfg *config.OpenTofuPluginConfig, dts []*sdk.DeployTarget[config.OpenTofuDeployTargetConfig], input *sdk.ExecuteStageInput[config.OpenTofuApplicationSpec]) (*sdk.ExecuteStageResponse, error) {
	return &sdk.ExecuteStageResponse{
		Status: sdk.StageStatusSuccess,
	}, nil
}

// DetermineVersions determines the versions of artifacts for the deployment.
func (p *Plugin) DetermineVersions(ctx context.Context, cfg *config.OpenTofuPluginConfig, input *sdk.DetermineVersionsInput[config.OpenTofuApplicationSpec]) (*sdk.DetermineVersionsResponse, error) {
	return &sdk.DetermineVersionsResponse{
		Versions: []sdk.ArtifactVersion{},
	}, nil
}

// DetermineStrategy determines the sync strategy for the deployment.
func (p *Plugin) DetermineStrategy(ctx context.Context, cfg *config.OpenTofuPluginConfig, input *sdk.DetermineStrategyInput[config.OpenTofuApplicationSpec]) (*sdk.DetermineStrategyResponse, error) {
	return &sdk.DetermineStrategyResponse{
		Strategy: sdk.SyncStrategyQuickSync,
	}, nil
}

// BuildQuickSyncStages builds the stages for quick sync.
func (p *Plugin) BuildQuickSyncStages(ctx context.Context, cfg *config.OpenTofuPluginConfig, input *sdk.BuildQuickSyncStagesInput) (*sdk.BuildQuickSyncStagesResponse, error) {
	return &sdk.BuildQuickSyncStagesResponse{
		Stages: []sdk.QuickSyncStage{},
	}, nil
}
