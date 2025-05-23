package main

import (
	"reflect"
	"testing"
)

func TestValidateAndExtractTyped(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]interface{}
		valid    bool
	}{
		{
			name: "Valid full types",
			input: `{
				"active": true,
				"name": "Alice",
				"score": 42.5,
				"deleted": null,
				"visits": 12
			}`,
			expected: map[string]interface{}{
				"active":  true,
				"name":    "Alice",
				"score":   42.5,
				"deleted": nil,
				"visits":  float64(12),
			},
			valid: true,
		},
		{
			name:     "Empty JSON",
			input:    `{}`,
			expected: map[string]interface{}{},
			valid:    true,
		},
		{
			name:  "Invalid unquoted key",
			input: `{active: true}`,
			valid: false,
		},
		{
			name:  "Invalid string without quotes",
			input: `{"name": Alice}`,
			valid: false,
		},
		{
			name:  "Trailing comma",
			input: `{"name": "Alice",}`,
			valid: false,
		},
		{
			name:  "Number as float with exponent",
			input: `{"value": 1.23e4}`,
			expected: map[string]interface{}{
				"value": 1.23e4,
			},
			valid: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, valid := ValidateAndExtractTyped(tt.input)
			if valid != tt.valid {
				t.Errorf("Expected valid=%v, got %v", tt.valid, valid)
			}
			if valid && !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected result=%v, got %v", tt.expected, result)
			}
		})
	}
}
