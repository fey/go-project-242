// Package code предоставляет функцию сравнивания файлов
package code

import (
	comparefiles "code/compareFiles"
	"code/formatter"
	code "code/parser"
	"fmt"
)

// GenDiff Функция сравнивания файлов.
func GenDiff(filepath1, filepath2, format string) (string, error) {
	if filepath1 == "" || filepath2 == "" {
		return "", fmt.Errorf("file paths cannot be empty: %q, %q", filepath1, filepath2)
	}

	const defaultFormat = "stylish"

	if format == "" {
		format = defaultFormat
	}
	data1, err1 := code.Parser(filepath1)
	if err1 != nil {
		return "", err1

	}

	data2, err2 := code.Parser(filepath2)
	if err2 != nil {
		return "", err2
	}

	diff := comparefiles.CompareFiles(data1, data2)

	result, err := formatter.GetFormatter(diff, format)

	if err != nil {
		return "", err
	}

	return result, nil
}
