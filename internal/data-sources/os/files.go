/**
 * Copyright (c) F5, Inc.
 *
 * This source code is licensed under the Apache License, Version 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package os

import (
	"log/slog"
	"os"
	"path"
)

func WriteFile(fileContent []byte, filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		slog.Debug("File does not exist, creating")
		err = os.MkdirAll(path.Dir(filePath), 0o750)
		if err != nil {
			slog.Error("Error creating directory", "error", err)
			return err
		}
	}

	err := os.WriteFile(filePath, fileContent, 0o644)
	if err != nil {
		slog.Error("Error writing to file", "error", err)
		return err
	}
	slog.Debug("Content written to file", "filePath", filePath)

	return nil
}
