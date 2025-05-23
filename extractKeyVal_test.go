package main

import (
	"reflect"
	"testing"
)

func TestExtractKeyValuePairs(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]string
	}{
		{
			"Single pair",
			`{"name": "John"}`,
			map[string]string{"name": "John"},
		},
		{
			"Multiple pairs",
			`{"a": "1", "b": "2"}`,
			map[string]string{"a": "1", "b": "2"},
		},
		{
			"Escaped characters",
			`{"quote": "She said \"hi\"", "path": "C:\\\\folder"}`,
			map[string]string{
				"quote": `She said "hi"`,
				"path":  `C:\\folder`,
			},
		},
		{
			"Extra spaces",
			`{ "x" : "y" , "z" : "w" }`,
			map[string]string{"x": "y", "z": "w"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractKeyValuePairs(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ExtractKeyValuePairs(%q) = %v; expected %v", tt.input, result, tt.expected)
			}
		})
	}
}
