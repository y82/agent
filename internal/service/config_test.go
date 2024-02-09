// Copyright (c) F5, Inc.
//
// This source code is licensed under the Apache License, Version 2.0 license found in the
// LICENSE file in the root directory of this source tree.

package service

import (
	"context"
	"fmt"
	"testing"

	configfakes2 "github.com/nginx/agent/v3/internal/datasource/config/configfakes"

	config2 "github.com/nginx/agent/v3/internal/config"

	"github.com/stretchr/testify/require"

	"github.com/nginx/agent/v3/api/grpc/instances"
	"github.com/nginx/agent/v3/internal/model"
	"github.com/nginx/agent/v3/internal/service/config/configfakes"
	"github.com/stretchr/testify/assert"
)

func TestConfigService_SetConfigContext(t *testing.T) {
	expectedConfigContext := &model.NginxConfigContext{
		AccessLogs: []*model.AccessLog{{Name: "access.logs"}},
	}

	configService := NewConfigService(&config2.Config{})
	configService.SetConfigContext(expectedConfigContext)

	assert.Equal(t, expectedConfigContext, configService.configContext)
}

func TestConfigService_ParseInstanceConfiguration(t *testing.T) {
	expectedConfigContext := &model.NginxConfigContext{
		AccessLogs: []*model.AccessLog{{Name: "access.logs"}},
	}

	configService := NewConfigService(&config2.Config{})

	fakeDataplaneConfig := &configfakes.FakeDataplaneConfig{}
	fakeDataplaneConfig.ParseConfigReturns(expectedConfigContext, nil)

	configService.dataplaneConfigServices[instances.Type_NGINX] = fakeDataplaneConfig

	result, err := configService.ParseInstanceConfiguration("123", &instances.Instance{Type: instances.Type_NGINX})

	require.NoError(t, err)
	assert.Equal(t, expectedConfigContext, result)

	_, err = configService.ParseInstanceConfiguration("123", &instances.Instance{Type: instances.Type_UNKNOWN})

	require.Error(t, err)
}

func TestUpdateInstanceConfiguration(t *testing.T) {
	instanceID := "ae6c58c1-bc92-30c1-a9c9-85591422068e"
	correlationID := "dfsbhj6-bc92-30c1-a9c9-85591422068e"
	ctx := context.TODO()
	instance := instances.Instance{
		InstanceId: instanceID,
		Type:       instances.Type_NGINX,
	}
	cs := NewConfigService(&config2.Config{})
	tests := []struct {
		name        string
		writeErr    error
		validateErr error
		reloadErr   error
		expected    *instances.ConfigurationStatus
	}{
		{
			name:        "write fails",
			writeErr:    fmt.Errorf("error writing config"),
			validateErr: nil,
			reloadErr:   nil,
			expected: &instances.ConfigurationStatus{
				InstanceId:    instanceID,
				CorrelationId: correlationID,
				Status:        instances.Status_FAILED,
				Message:       "error writing config",
			},
		},
		{
			name:        "validate fails",
			writeErr:    nil,
			validateErr: fmt.Errorf("error validating config"),
			reloadErr:   nil,
			expected: &instances.ConfigurationStatus{
				InstanceId:    instanceID,
				CorrelationId: correlationID,
				Status:        instances.Status_FAILED,
				Message:       "error validating config",
			},
		},
		{
			name:        "reload fails",
			writeErr:    nil,
			validateErr: nil,
			reloadErr:   fmt.Errorf("error reloading config"),
			expected: &instances.ConfigurationStatus{
				InstanceId:    instanceID,
				CorrelationId: correlationID,
				Status:        instances.Status_FAILED,
				Message:       "error reloading config",
			},
		},
		{
			name:        "success",
			writeErr:    nil,
			validateErr: nil,
			reloadErr:   nil,
			expected: &instances.ConfigurationStatus{
				InstanceId:    instanceID,
				CorrelationId: correlationID,
				Status:        instances.Status_SUCCESS,
				Message:       "Config applied successfully",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mockConfigWriter := configfakes2.FakeConfigWriterInterface{}
			mockConfigWriter.WriteReturns(test.writeErr)
			mockConfigWriter.ReloadReturns(test.reloadErr)
			mockConfigWriter.ValidateReturns(test.validateErr)

			filesURL := fmt.Sprintf("/instance/%s/files/", instanceID)

			result := cs.UpdateInstanceConfiguration(ctx, correlationID, filesURL, &instance)
			assert.Equal(t, test.expected, result)
		})
	}
}
