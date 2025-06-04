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

package main

import (
	"context"
	"encoding/json"
	"testing"

	sdk "github.com/pipe-cd/piped-plugin-sdk-go"
	"github.com/stretchr/testify/assert"
)

// TODO: add tests for executeHello(), executeGoodbye()

func TestFetchDefinedStages(t *testing.T) {
	t.Parallel()
	plugin := &plugin{}
	got := plugin.FetchDefinedStages()
	want := []string{
		"EXAMPLE_HELLO",
		"EXAMPLE_GOODBYE",
	}

	assert.Equal(t, want, got)
}

func TestBuildPipelineSyncStages(t *testing.T) {
	t.Parallel()

	plugin := &plugin{}
	got, err := plugin.BuildPipelineSyncStages(context.Background(), &pluginConfig{}, &sdk.BuildPipelineSyncStagesInput{
		Request: sdk.BuildPipelineSyncStagesRequest{
			Stages: []sdk.StageConfig{
				{
					Index: 0,
					Name:  "EXAMPLE_HELLO",
					Config: mustMarshal(helloStageOptions{
						Name: "world",
					}),
				},
				{
					Index: 1,
					Name:  "EXAMPLE_GOODBYE",
					Config: mustMarshal(goodbyeStageOptions{
						Message: "world",
					}),
				},
			},
			Rollback: true,
		},
	})

	want := &sdk.BuildPipelineSyncStagesResponse{
		Stages: []sdk.PipelineStage{
			{
				Index:              0,
				Name:               "EXAMPLE_HELLO",
				Rollback:           false,
				Metadata:           map[string]string{},
				AvailableOperation: sdk.ManualOperationNone,
			},
			{
				Index:              1,
				Name:               "EXAMPLE_GOODBYE",
				Rollback:           false,
				Metadata:           map[string]string{},
				AvailableOperation: sdk.ManualOperationNone,
			},
		},
	}

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func mustMarshal(v any) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}
