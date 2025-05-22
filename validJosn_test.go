package main

import (
	"testing"
)

func TestValidJson(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected int
	}{
		{"Empty input", []byte(""), 1},
		{"Only open brace", []byte("{"), 1},
		{"Only close brace", []byte("}"), 1},
		{"Balanced braces", []byte("{}"), 0},
		{"Unbalanced - extra open", []byte("{{}"), 1},
		{"Unbalanced - extra close", []byte("{}{}{}{{}}}"), 1},
		{"Invalid type structure", []byte("{{}{}}{}}{"), 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsValidJson(string(test.input))
			if result != test.expected {
				t.Errorf("Test %q failed: expected %d, got %d", test.name, test.expected, result)
			}
		})
	}
}
