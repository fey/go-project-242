package code

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenDiff(t *testing.T) {
	tests := []struct {
		name        string
		filepath1   string
		filepath2   string
		format      string
		expected    string
		expectError bool
		errorMsg    string
	}{
		{
			name:      "compare json",
			filepath1: "./testdata/file1.json",
			filepath2: "./testdata/file2.json",
			format:    "stylish",
			expected:  "./testdata/stylish.txt",
		},
		{
			name:      "emptyFormat",
			filepath1: "./testdata/file1.json",
			filepath2: "./testdata/file2.json",
			format:    "",
			expected:  "./testdata/stylish.txt",
		},
		{
			name:        "unknownFormat",
			filepath1:   "./testdata/file1.json",
			filepath2:   "./testdata/file2.json",
			format:      "txt",
			expectError: true,
		},
		{
			name:      "compare yml",
			filepath1: "./testdata/file1.yml",
			filepath2: "./testdata/file2.yml",
			format:    "stylish",
			expected:  "./testdata/stylish.txt",
		},
		{
			name:      "plainFormat",
			filepath1: "./testdata/file1.yml",
			filepath2: "./testdata/file2.yml",
			format:    "plain",
			expected:  "./testdata/plain.txt",
		},
		{
			name:      "jsonFormat",
			filepath1: "./testdata/file1.yml",
			filepath2: "./testdata/file2.yml",
			format:    "json",
			expected:  "./testdata/json.txt",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got, err := GenDiff(tc.filepath1, tc.filepath2, tc.format)

			if tc.expectError {
				assert.Error(t, err, "Expected error but got none")
				assert.Empty(t, got)
			} else {
				assert.NoError(t, err)
				data, err := os.ReadFile(tc.expected)
				assert.NoError(t, err)
				want := string(data)
				assert.Equal(t, want, got)
			}
		})
	}
}
