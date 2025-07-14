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
)

func Test_FetchDefinedStages(t *testing.T) {
	plugin := &Plugin{}
	desiredStages := []string{"OPENTOFU_PLAN", "OPENTOFU_APPLY", "OPENTOFU_ROLLBACK"}
	expectedstages := plugin.FetchDefinedStages()

	assert.Equal(t, desiredStages, expectedstages, "Defined stages should match the expected stages")
}

func Test_BuildPipelineSyncStages(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input *sdk.BuildPipelineSyncStagesInput
		want  *sdk.BuildPipelineSyncStagesResponse
	}{
		{
			name: "single stage without rollback",
			input: &sdk.BuildPipelineSyncStagesInput{
				Request: sdk.BuildPipelineSyncStagesRequest{
					Stages: []sdk.StageConfig{
						{
							Name:  stageApply,
							Index: 1,
						},
					},
					Rollback: false,
				},
			},
			want: &sdk.BuildPipelineSyncStagesResponse{
				Stages: []sdk.PipelineStage{
					{
						Name:               "OPENTOFU_APPLY",
						Index:              1,
						Rollback:           false,
						Metadata:           map[string]string{},
						AvailableOperation: sdk.ManualOperationNone,
					},
				},
			},
		},
		{
			name: "multiple stages without rollback",
			input: &sdk.BuildPipelineSyncStagesInput{
				Request: sdk.BuildPipelineSyncStagesRequest{
					Stages: []sdk.StageConfig{
						{
							Name:  stagePlan,
							Index: 1,
						},
						{
							Name:  stageApply,
							Index: 3,
						},
					},
					Rollback: false,
				},
			},
			want: &sdk.BuildPipelineSyncStagesResponse{
				Stages: []sdk.PipelineStage{
					{
						Name:               "OPENTOFU_PLAN",
						Index:              1,
						Rollback:           false,
						Metadata:           map[string]string{},
						AvailableOperation: sdk.ManualOperationNone,
					},
					{
						Name:               "OPENTOFU_APPLY",
						Index:              3,
						Rollback:           false,
						Metadata:           map[string]string{},
						AvailableOperation: sdk.ManualOperationNone,
					},
				},
			},
		},
		{
			name: "multiple stages with rollback",
			input: &sdk.BuildPipelineSyncStagesInput{
				Request: sdk.BuildPipelineSyncStagesRequest{
					Stages: []sdk.StageConfig{
						{
							Name:  stagePlan,
							Index: 2,
						},
						{
							Name:  stageApply,
							Index: 3,
						},
					},
					Rollback: true,
				},
			},
			want: &sdk.BuildPipelineSyncStagesResponse{
				Stages: []sdk.PipelineStage{
					{
						Name:               "OPENTOFU_PLAN",
						Index:              2,
						Rollback:           false,
						Metadata:           map[string]string{},
						AvailableOperation: sdk.ManualOperationNone,
					},
					{
						Name:               "OPENTOFU_APPLY",
						Index:              3,
						Rollback:           false,
						Metadata:           map[string]string{},
						AvailableOperation: sdk.ManualOperationNone,
					},
					{
						Name:               "OPENTOFU_ROLLBACK",
						Index:              2,
						Rollback:           true,
						Metadata:           map[string]string{},
						AvailableOperation: sdk.ManualOperationNone,
					},
				},
			},
		},
	}

	p := &Plugin{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := p.BuildPipelineSyncStages(t.Context(), nil, tt.input)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestPlugin_BuildQuickSyncStages(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input *sdk.BuildQuickSyncStagesInput
		want  *sdk.BuildQuickSyncStagesResponse
	}{
		{
			name: "no rollback",
			input: &sdk.BuildQuickSyncStagesInput{
				Request: sdk.BuildQuickSyncStagesRequest{
					Rollback: false,
				},
			},
			want: &sdk.BuildQuickSyncStagesResponse{
				Stages: []sdk.QuickSyncStage{
					{
						Name:               stageApply,
						Description:        "Sync by applying any detected changes",
						Rollback:           false,
						Metadata:           map[string]string{},
						AvailableOperation: sdk.ManualOperationNone,
					},
				},
			},
		},
		{
			name: "with rollback",
			input: &sdk.BuildQuickSyncStagesInput{
				Request: sdk.BuildQuickSyncStagesRequest{
					Rollback: true,
				},
			},
			want: &sdk.BuildQuickSyncStagesResponse{
				Stages: []sdk.QuickSyncStage{
					{
						Name:               "OPENTOFU_APPLY",
						Description:        "Sync by applying any detected changes",
						Rollback:           false,
						Metadata:           map[string]string{},
						AvailableOperation: sdk.ManualOperationNone,
					},
					{
						Name:               "OPENTOFU_ROLLBACK",
						Description:        "Rollback by applying the previous OpenTofu files",
						Rollback:           true,
						Metadata:           map[string]string{},
						AvailableOperation: sdk.ManualOperationNone,
					},
				},
			},
		},
	}

	p := &Plugin{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := p.BuildQuickSyncStages(t.Context(), nil, tt.input)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
