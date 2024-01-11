package services

import (
	"encoding/json"
	"log/slog"
	"os"
	"strings"
	"fmt"

	"github.com/google/uuid"
	"github.com/nginx/agent/v3/internal/client"
	"github.com/nginx/agent/v3/internal/data-sources/nginx"
	file "github.com/nginx/agent/v3/internal/data-sources/os"
	"github.com/nginx/agent/v3/internal/models/instances"
)

type InstanceService struct {
	instances []*instances.Instance
}

func NewInstanceService() *InstanceService {
	return &InstanceService{}
}

func (is *InstanceService) GetInstances() ([]*instances.Instance, error) {
	processes, err := file.GetProcesses()
	if err != nil {
		is.instances = []*instances.Instance{}
	} else {
		is.instances, err = nginx.GetInstances(processes)
	}

	return is.instances, err
}

func (is *InstanceService) UpdateInstanceConfig(tenantID uuid.UUID, instanceID uuid.UUID, filesUrl string) error {
	cachePath := fmt.Sprintf("/var/lib/nginx-agent/config/%v/cache.json", instanceID.String())
	lastConfigApply := make(map[string]*instances.File)
	currentConfigApply := make(map[string]*instances.File)

	cacheData, err := os.ReadFile(cachePath)
	if err != nil {
		slog.Error("Unable to read file cache.json", "error", err)
	}

	err = json.Unmarshal(cacheData, &lastConfigApply)
	if err != nil {
		slog.Error("Unable to unmarshal data from cache.json", "error", err)
	}

	hcd := client.NewHttpConfigDownloader()

	filesMetaData, err := hcd.GetFilesMetadata(filesUrl, tenantID)
	if err != nil {
		slog.Error("Error getting files metadata", "error", err)
		return err
	}

filesLoop:
	for _, fileData := range filesMetaData.Files {
		if fileData.Path != "" && !strings.HasSuffix(fileData.Path, "/") {
			if lastConfigApply != nil && len(lastConfigApply) > 0 {
				fileOnSystem, ok := lastConfigApply[fileData.Path]
				if ok && !fileData.LastModified.AsTime().After(fileOnSystem.LastModified.AsTime()) {
					slog.Debug("Skipping file as latest version is already on disk", "filePath", fileData.Path)
					currentConfigApply[fileData.Path] = lastConfigApply[fileData.Path]
					continue filesLoop
				}

			}

			fileDownloadResponse, err := hcd.GetFile(fileData, filesUrl, tenantID)
			if err != nil {
				slog.Error("Error getting file data", "err", err)
			}

			slog.Info("============== \n Downloading File:", "path", fileData.Path)

			err = file.WriteFile(fileDownloadResponse.FileContent, fileDownloadResponse.FilePath)
			if err != nil {
				slog.Error("Error writing to file", "err", err)
			}

			currentConfigApply[fileData.Path] = &instances.File{
				Version:      fileData.Version,
				Path:         fileData.Path,
				LastModified: fileData.LastModified,
			}
		}
	}

	cacheJson, err := json.MarshalIndent(currentConfigApply, "", "  ")
	if err != nil {
		slog.Error("Unable marshal cache data", "error", err)
	}

	err = file.WriteFile(cacheJson, cachePath)
	if err != nil {
		slog.Error("Unable to write cache", "error", err)
	}

	return err
}
