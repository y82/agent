// Copyright (c) F5, Inc.
//
// This source code is licensed under the Apache License, Version 2.0 license found in the
// LICENSE file in the root directory of this source tree.

package service

import (
	"context"
	"fmt"

	"github.com/nginx/agent/v3/api/grpc/instances"
	agentconfig "github.com/nginx/agent/v3/internal/config"
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
	dataplaneConfigServices map[instances.Type]config.DataplaneConfig // key is instance type
	// instanceConfigWriter    map[string]writer.ConfigWriterInterface   // key is instanceID
	agentConfig *agentconfig.Config
}

func NewConfigService(agentConfig *agentconfig.Config) *ConfigService {
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

func (cs *ConfigService) UpdateInstanceConfiguration(_ context.Context, correlationID, _ string,
	instance *instances.Instance,
) *instances.ConfigurationStatus {
	// remove when tenantID is being set
	// tenantID, err := uuid.Parse("7332d596-d2e6-4d1e-9e75-70f91ef9bd0e")

	// configClient := client.NewHTTPConfigClient(cs.agentConfig.Client.Timeout)

	// configWriter := cs.instanceConfigWriter[instance.GetInstanceId()]
	// configWriter :=

	// if configWriter != nil {
	//	cs.instanceConfigWriter[instance.GetInstanceId()] = writer.NewConfigWriter(configClient,
	//	cs.agentConfig, instance.GetInstanceId())
	// }

	// skippedFiles, _ := configWriter.Write(context.Background(), "", tenantID)
	// if err != nil {
	//	return &instances.ConfigurationStatus{
	//		InstanceId:    instance.GetInstanceId(),
	//		CorrelationId: correlationID,
	//		Status:        instances.Status_FAILED,
	//		Message:       fmt.Sprintf("%s", err),
	//	}
	//}

	// if skippedFiles != nil {
	//	slog.Debug("skip")
	//}

	// configWriter.SetDataplaneConfig(cs.dataplaneConfigServices[instance.GetType()])
	// err = configWriter.Validate(instance)
	//
	// if err != nil {
	//	// Add Rollback
	//	return &instances.ConfigurationStatus{
	//		InstanceId:    instance.GetInstanceId(),
	//		CorrelationId: correlationID,
	//		Status:        instances.Status_FAILED,
	//		Message:       fmt.Sprintf("%s", err),
	//	}
	//}
	//
	// err = configWriter.Reload(instance)
	// if err != nil {
	//	// Add Rollback
	//	return &instances.ConfigurationStatus{
	//		InstanceId:    instance.GetInstanceId(),
	//		CorrelationId: correlationID,
	//		Status:        instances.Status_FAILED,
	//		Message:       fmt.Sprintf("%s", err),
	//	}
	//}
	//
	// err = configWriter.Complete()
	//
	// if err != nil {
	//	slog.Error("Error: ", err)
	//}

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
