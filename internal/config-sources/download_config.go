/**
 * Copyright (c) F5, Inc.
 *
 * This source code is licensed under the Apache License, Version 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package configsources

import (
	// "encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/nginx/agent/v3/internal/models/instances"
	"google.golang.org/protobuf/encoding/protojson"
)

type httpConfigDownloader struct {
	logger     *slog.Logger
	httpClient http.Client
}

func NewHttpConfigDownloader(logger *slog.Logger) *httpConfigDownloader {
	httpClient := http.Client{
		Timeout: time.Second * 10,
	}

	return &httpConfigDownloader{
		httpClient: httpClient,
		logger:     logger,
	}
}

func (hcd *httpConfigDownloader) GetFilesMetadata(filesUrl string, tenantID uuid.UUID) (*instances.Files, error) {
	files := instances.Files{}

	req, err := http.NewRequest(http.MethodGet, filesUrl, nil)
	req.Header.Set("tenantId", tenantID.String())
	if err != nil {
		hcd.logger.Error("Error making request", "err", err)
		return nil, err
	}

	resp, err := hcd.httpClient.Do(req)
	if err != nil {
		hcd.logger.Error("Error response from request", "err", err)
		return nil, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		hcd.logger.Error("Error reading response body", "err", err)
		return nil, err
	}

	hcd.logger.Info("Files download", "files", data)

	// TODO: look into why version is an unknown field and why this is needed
	pb := protojson.UnmarshalOptions{DiscardUnknown: true}
	err = pb.Unmarshal(data, &files)

	if err != nil {
		hcd.logger.Error("Error unmarshal data", "err", err)
		return nil, err
	}

	return &files, nil
}

func (hcd *httpConfigDownloader) GetFile(file *instances.File, filesUrl string, tenantID uuid.UUID) (*instances.FileDownloadResponse, error) {
	response := instances.FileDownloadResponse{}
	params := url.Values{}

	params.Add("version", file.Version)
	params.Add("encoded", "true")

	filePath := url.QueryEscape(file.Path)

	fileUrl := fmt.Sprintf("%v%v?%v", filesUrl, filePath, params.Encode())
	hcd.logger.Info(fileUrl)

	req, err := http.NewRequest(http.MethodGet, fileUrl, nil)
	req.Header.Set("tenantId", tenantID.String())
	if err != nil {
		hcd.logger.Error("Error making request", "err", err)
		return nil, err
	}

	resp, err := hcd.httpClient.Do(req)
	if err != nil {
		hcd.logger.Error("Error response from request", "err", err)
		return nil, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		hcd.logger.Error("Error reading response body", "err", err)
		return nil, err
	}

	// TODO: look into why type is an unknown field and why this is needed
	pb := protojson.UnmarshalOptions{DiscardUnknown: true}
	err = pb.Unmarshal(data, &response)

	if err != nil {
		hcd.logger.Error("Error unmarshal data", "err", err)
		return nil, err
	}

	return &response, err
}
