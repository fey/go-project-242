// Package jsonfmt предоставляет функцию для вывода данных в формате json
package formatjson

import (
	comparefiles "code/compareFiles"
	"encoding/json"
	"fmt"
)

// FormatJSON форматирует дерево различий в стиле json
func FormatJSON(nodes []comparefiles.Node) (string, error) {
	wrapper := map[string]any{
		"diff": nodes,
	}

	data, err := json.MarshalIndent(wrapper, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal JSON: %w", err)
	}
	return string(data), nil
}
