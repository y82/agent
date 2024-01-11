/**
 * Copyright (c) F5, Inc.
 *
 * This source code is licensed under the Apache License, Version 2.0 license found in the
 * LICENSE file in the root directory of this source tree.
 */

package os

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteFile(t *testing.T) {
	filePath := "/tmp/test.conf"
	fileContent := []byte("location /test {\n    return 200 \"Test location\\n\";\n}")

	err := WriteFile(fileContent, filePath)
	assert.NoError(t, err)
	assert.FileExists(t, filePath)

	data, err := os.ReadFile(filePath)
	assert.NoError(t, err)
	assert.Equal(t, fileContent, data)

	err = os.Remove(filePath)
	assert.NoError(t, err)
	assert.NoFileExists(t, filePath)
}
