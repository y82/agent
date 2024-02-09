// Copyright (c) F5, Inc.
//
// This source code is licensed under the Apache License, Version 2.0 license found in the
// LICENSE file in the root directory of this source tree.

package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nginx/agent/v3/api/grpc/instances"
	"github.com/nginx/agent/v3/internal/client"
	config2 "github.com/nginx/agent/v3/internal/config"
	config3 "github.com/nginx/agent/v3/internal/datasource/config"
	"github.com/nginx/agent/v3/internal/service/config"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6@v6.7.0 -generate
//counterfeiter:generate . ConfigServiceInterface
type ConfigServiceInterface interface {
	SetConfigContext(instanceConfigContext any)
	UpdateInstanceConfiguration(
		ctx context.Context,
		correlationID, location string,
		instance *instances.Instance,
	) *instances.ConfigurationStatus
	ParseInstanceConfiguration(
		correlationID string,
		instance *instances.Instance,
	) (instanceConfigContext any, err error)
}

type ConfigService struct {
	configContext           any
	dataplaneConfigServices map[instances.Type]config.DataplaneConfig
	// TODO have configwriter per instacne Id,
	// instanceConfigWriter    map[string]config3.ConfigWriter // key is instanceID
	agentConfig *config2.Config
}

func NewConfigService(agentConfig *config2.Config) *ConfigService {
	nginxConfigService := config.NewNginx()

	return &ConfigService{
		dataplaneConfigServices: map[instances.Type]config.DataplaneConfig{
			instances.Type_NGINX:                nginxConfigService,
			instances.Type_NGINX_PLUS:           nginxConfigService,
			instances.Type_NGINX_GATEWAY_FABRIC: config.NewNginxGatewayFabric(),
		},
		agentConfig: agentConfig,
	}
}

func (cs *ConfigService) SetConfigContext(instanceConfigContext any) {
	cs.configContext = instanceConfigContext
}

func (cs *ConfigService) UpdateInstanceConfiguration(ctx context.Context, correlationID, location string,
	instance *instances.Instance,
) *instances.ConfigurationStatus {
	// TODO: remove when tenantID is being set
	tenantID, _ := uuid.Parse("7332d596-d2e6-4d1e-9e75-70f91ef9bd0e")

	configClient := client.NewHTTPConfigClient(cs.agentConfig.Client.Timeout)

	// TODO: check if instanceID has a config writer in the map if it doesnt create new one and add to map
	// change needed for testing so we can use the fake config writer and not read the previous cache each time
	configWriter := config3.NewConfigWriter(configClient, cs.agentConfig, instance.GetInstanceId())

	_, err := configWriter.Write(ctx, location, tenantID)
	if err != nil {
		return &instances.ConfigurationStatus{
			InstanceId:    instance.GetInstanceId(),
			CorrelationId: correlationID,
			Status:        instances.Status_FAILED,
			Message:       fmt.Sprintf("%s", err),
		}
	}

	configWriter.SetDataplaneConfig(cs.dataplaneConfigServices[instance.GetType()])
	err = configWriter.Validate(instance)

	if err != nil {
		// TODO: Rollback
		return &instances.ConfigurationStatus{
			InstanceId:    instance.GetInstanceId(),
			CorrelationId: correlationID,
			Status:        instances.Status_FAILED,
			Message:       fmt.Sprintf("%s", err),
		}
	}

	err = configWriter.Reload(instance)
	if err != nil {
		// TODO: Rollback
		return &instances.ConfigurationStatus{
			InstanceId:    instance.GetInstanceId(),
			CorrelationId: correlationID,
			Status:        instances.Status_FAILED,
			Message:       fmt.Sprintf("%s", err),
		}
	}

	// TODO: Complete
	return &instances.ConfigurationStatus{
		InstanceId:    instance.GetInstanceId(),
		CorrelationId: correlationID,
		Status:        instances.Status_SUCCESS,
		Message:       "Config applied successfully",
	}
}

func (cs *ConfigService) ParseInstanceConfiguration(
	_ string,
	instance *instances.Instance,
) (instanceConfigContext any, err error) {
	conf, ok := cs.dataplaneConfigServices[instance.GetType()]

	if !ok {
		return nil, fmt.Errorf("unknown instance type %s", instance.GetType())
	}

	return conf.ParseConfig(instance)
}
