// Copyright (c) F5, Inc.
//
// This source code is licensed under the Apache License, Version 2.0 license found in the
// LICENSE file in the root directory of this source tree.

package config

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
)

type (
	//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6@v6.7.0 -generate
	//counterfeiter:generate . CacheInterface
	CacheInterface interface{}

	Cache struct {
		previousFileCache FileCache
		currentFileCache  FileCache
		cachePath         string
	}
)

func NewCache(cachePath, instanceID string) *Cache {
	if cachePath == "" {
		cachePath = fmt.Sprintf(cacheLocation, instanceID)
	}

	previousFileCache, err := readInstanceCache(cachePath)
	if err != nil {
		slog.Warn("", err)
		previousFileCache = make(FileCache)
	}

	return &Cache{
		cachePath:         cachePath,
		previousFileCache: previousFileCache,
	}
}

func readInstanceCache(cachePath string) (previousFileCache FileCache, err error) {
	previousFileCache = make(FileCache)

	if _, statErr := os.Stat(cachePath); os.IsNotExist(statErr) {
		return previousFileCache, fmt.Errorf("cache.json does not exist %s: %w", cachePath, statErr)
	}

	cacheData, err := os.ReadFile(cachePath)
	if err != nil {
		return previousFileCache, fmt.Errorf("error reading file cache.json %s: %w", cachePath, err)
	}
	err = json.Unmarshal(cacheData, &previousFileCache)
	if err != nil {
		return previousFileCache, fmt.Errorf("error unmarshalling data from cache.json %s: %w", cachePath, err)
	}

	return previousFileCache, err
}
