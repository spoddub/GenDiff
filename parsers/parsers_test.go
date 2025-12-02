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
		assert.Contains(t, data, "common")
		assert.Contains(t, data, "group1")
		assert.Contains(t, data, "group2")
	}

	common, ok := data["common"].(map[string]any)
	if assert.True(t, ok, "common should be an object") {
		assert.Equal(t, "Value 1", common["setting1"])
		assert.EqualValues(t, 200, common["setting2"])
		assert.Equal(t, true, common["setting3"])
	}

	group2, ok := data["group2"].(map[string]any)
	if assert.True(t, ok, "group2 should be an object") {
		deep, ok := group2["deep"].(map[string]any)
		if assert.True(t, ok, "group2.deep should be an object") {
			assert.EqualValues(t, 45, deep["id"])
		}
	}
}

func TestParseYAML(t *testing.T) {
	path := fixturePath("file1.yml")

	data, err := Parse(path)
	if !assert.NoError(t, err, "Parse should not return error for YAML") {
		return
	}

	if assert.NotNil(t, data) {
		assert.Contains(t, data, "common")
		assert.Contains(t, data, "group1")
		assert.Contains(t, data, "group2")
	}

	common, ok := data["common"].(map[string]any)
	if assert.True(t, ok, "common should be an object") {
		assert.Equal(t, "Value 1", common["setting1"])
		assert.EqualValues(t, 200, common["setting2"])
		assert.Equal(t, true, common["setting3"])
	}

	group2, ok := data["group2"].(map[string]any)
	if assert.True(t, ok, "group2 should be an object") {
		deep, ok := group2["deep"].(map[string]any)
		if assert.True(t, ok, "group2.deep should be an object") {
			assert.EqualValues(t, 45, deep["id"])
		}
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
