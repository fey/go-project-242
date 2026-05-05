// Package stylish предоставляет функцию для вывода данных в формате stylish
package stylish

import (
	comparefiles "code/compareFiles"
	"fmt"
	"sort"
	"strings"
)

const defaultIndent = 4

// makeIndent создает отступ для текущего уровня
func makeIndent(depth, spaceCount int) string {
	return strings.Repeat(" ", depth*spaceCount-2)
}

// makeBackIndent создает закрывающий отступ для вложенных объектов
func makeBackIndent(depth, spaceCount int) string {
	return strings.Repeat(" ", (depth-1)*spaceCount)
}

// Stringify преобразует значение в строку
func Stringify(value any, depth int) string {
	if value == nil {
		return "null"
	}

	if !isMap(value) {
		return fmt.Sprintf("%v", value)
	}

	valueMap, ok := value.(map[string]any)
	if !ok {
		return fmt.Sprintf("%v", value)
	}

	keys := getSortedKeysFromMap(valueMap)

	lines := make([]string, 0, len(keys))
	indent := strings.Repeat(" ", depth*defaultIndent)
	for _, key := range keys {
		val := valueMap[key]
		lines = append(lines, fmt.Sprintf("%s%s: %s", indent, key, Stringify(val, depth+1)))
	}

	result := strings.Join(lines, "\n")
	backIndent := makeBackIndent(depth, defaultIndent)

	return fmt.Sprintf("{\n%s\n%s}", result, backIndent)
}

// isMap проверяет, является ли значение map
func isMap(value any) bool {
	_, ok := value.(map[string]any)
	return ok
}

// getSortedKeysFromMap возвращает отсортированные ключи из map
func getSortedKeysFromMap(m map[string]any) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	return keys
}

// FormatStylish форматирует дерево различий в стиле stylish
func FormatStylish(tree []comparefiles.Node) string {
	var iter func(nodes []comparefiles.Node, depth int) string
	iter = func(nodes []comparefiles.Node, depth int) string {
		var builder strings.Builder
		builder.WriteString("{\n")
		for _, node := range nodes {
			switch node.Type {
			case comparefiles.Nested:
				childStr := iter(node.Children, depth+1)
				fmt.Fprintf(&builder, "%s  %s: %s\n", makeIndent(depth, defaultIndent), node.Key, childStr)
			case comparefiles.Unchanged:
				fmt.Fprintf(&builder, "%s  %s: %s\n", makeIndent(depth, defaultIndent), node.Key, Stringify(node.OldValue, depth+1))
			case comparefiles.Deleted:
				fmt.Fprintf(&builder, "%s- %s: %s\n", makeIndent(depth, defaultIndent), node.Key, Stringify(node.OldValue, depth+1))
			case comparefiles.Added:
				fmt.Fprintf(&builder, "%s+ %s: %s\n", makeIndent(depth, defaultIndent), node.Key, Stringify(node.NewValue, depth+1))
			case comparefiles.Changed:
				fmt.Fprintf(&builder, "%s- %s: %s\n", makeIndent(depth, defaultIndent), node.Key, Stringify(node.OldValue, depth+1))
				fmt.Fprintf(&builder, "%s+ %s: %s\n", makeIndent(depth, defaultIndent), node.Key, Stringify(node.NewValue, depth+1))
			}
		}
		builder.WriteString(makeBackIndent(depth, defaultIndent) + "}")
		return builder.String()
	}

	return iter(tree, 1)
}
