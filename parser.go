package code

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func ParseFile(path string) (map[string]any, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		return nil, err
	}
	extension := filepath.Ext(absPath)

	switch extension {
	case ".json":
		var result map[string]any
		if err := json.Unmarshal(data, &result); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, fmt.Errorf("%s extension is not supported", extension)
	}
}
