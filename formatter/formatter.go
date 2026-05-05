// Package formatter предоставляет функцию, которая озвращает дерево различий в заданном формате
package formatter

import (
	comparefiles "code/compareFiles"
	"code/formatter/formatJson"
	"code/formatter/plain"
	"code/formatter/stylish"
	"fmt"
)

// GetFormatter возвращает дерево различий в заданном формате
func GetFormatter(compare []comparefiles.Node, format string) (string, error) {
	switch format {
	case "stylish":
		return stylish.FormatStylish(compare), nil
	case "plain":
		return plain.FormatPlain(compare), nil
	case "json":
		data, err := formatjson.FormatJSON(compare)
		if err != nil {
			return "", err
		}
		return data, nil
	default:
		return "", fmt.Errorf("unknown format: %s", format)
	}
}
