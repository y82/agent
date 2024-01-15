package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/google/uuid"
	"github.com/nginx/agent/v3/internal/apis/http"
	file "github.com/nginx/agent/v3/internal/data-sources/os"
)

func main() {
	exampleTenantID := "7332d596-d2e6-4d1e-9e75-70f91ef9bd0e"
	exampleInstanceID := "aecea348-62c1-4e3d-b848-6d6cdeb1cb9c"

	handler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})
	logger := slog.New(handler)
	slog.SetDefault(logger)

	// TODO: remove when example is no longer needed
	tenantID, err := uuid.Parse(exampleTenantID)
	if err != nil {
		logger.Error("Failed to create UUID for TenantID", "err", err)
	}

	instanceID, err := uuid.Parse(exampleInstanceID)
	if err != nil {
		logger.Error("Failed to create UUID for InstanceID", "err", err)
	}

	managementServer := "0.0.0.0:8090"
	filesUrl := fmt.Sprintf("http://%v/instance/%s/files/", managementServer, instanceID)

	dataplaneServer := http.NewDataplaneServer("0.0.0.0:8091", logger)
	go dataplaneServer.Run(context.Background())

	err = file.UpdateNginxConfig(tenantID, instanceID, filesUrl)

	if err != nil {
		logger.Error("Error updating Instance Config", "err", err)
	}
}
