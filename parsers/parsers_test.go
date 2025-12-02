package parsers

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func fixturePath(filename string) string {
	return filepath.Join("..", "testdata", "fixture", filename)
}

func TestParseJSON(t *testing.T) {
	path := fixturePath("file1.json")

	data, err := Parse(path)
	if !assert.NoError(t, err, "Parse should not return error for JSON") {
		return
	}

	if assert.NotNil(t, data) {
		assert.Equal(t, "hexlet.io", data["host"])
		assert.EqualValues(t, 50, data["timeout"])
		assert.Equal(t, "123.234.53.22", data["proxy"])
		assert.Equal(t, false, data["follow"])
	}
}

func TestParseYAML(t *testing.T) {
	path := fixturePath("file1.yml")

	data, err := Parse(path)
	if !assert.NoError(t, err, "Parse should not return error for YAML") {
		return
	}

	if assert.NotNil(t, data) {
		assert.Equal(t, "hexlet.io", data["host"])
		assert.EqualValues(t, 50, data["timeout"])
		assert.Equal(t, "123.234.53.22", data["proxy"])
		assert.Equal(t, false, data["follow"])
	}
}

func TestParseUnsupportedExtension(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.txt")

	err := os.WriteFile(path, []byte("some content"), 0o644)
	if !assert.NoError(t, err, "failed to create temp file") {
		return
	}

	_, err = Parse(path)
	if assert.Error(t, err, "expected error for unsupported extension") {
		assert.Contains(t, err.Error(), "unsupported file format: .txt")
	}
}
