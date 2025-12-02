package code

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func fixturePath(filename string) string {
	return filepath.Join("testdata", "fixture", filename)
}

func TestGenDiffFlatJSONStylish(t *testing.T) {
	file1 := fixturePath("file1.json")
	file2 := fixturePath("file2.json")
	expectedPath := fixturePath("flat_stylish.txt")

	got, err := GenDiff(file1, file2, "stylish")
	if !assert.NoError(t, err, "GenDiff returned error") {
		return
	}

	expectedBytes, err := os.ReadFile(expectedPath)
	if !assert.NoError(t, err, "cannot read expected fixture") {
		return
	}

	expected := strings.TrimSpace(string(expectedBytes))
	got = strings.TrimSpace(got)

	assert.Equal(t, expected, got, "diff result does not match expected output")
}
