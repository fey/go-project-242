// Package code предоставляет функцию парсинга файлов
package code

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// Parser функция парсинга файла
func Parser(path string) (map[string]any, error) {
	var result map[string]any

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("file reading error %s: %w", path, err)
	}

	ext := strings.ToLower(filepath.Ext(path))

	switch ext {
	case ".json":
		if err := json.Unmarshal(data, &result); err != nil {
			return nil, fmt.Errorf("error parsing JSON file %s: %w", path, err)
		}

	case ".yaml", ".yml":
		if err := yaml.Unmarshal(data, &result); err != nil {
			return nil, fmt.Errorf("error parsing YAML file %s: %w", path, err)
		}

	default:
		return nil, fmt.Errorf("unsupported file format %s. Supported formats: .json, .yaml, .yml", ext)
	}

	return result, nil
}
