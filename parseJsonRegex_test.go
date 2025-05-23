package main

import (
	"reflect"
	"testing"
)

func TestParseJSONWithRegex(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected interface{}
		valid    bool
	}{
		{
			name:  "Valid simple object with all types",
			input: `{"key":"value","key-n":101,"key-o":{},"key-l":[]}`,
			expected: map[string]interface{}{
				"key":   "value",
				"key-n": float64(101),
				"key-o": map[string]interface{}{},
				"key-l": []interface{}{},
			},
			valid: true,
		},
		{
			name:  "Invalid missing colon",
			input: `{"key" "value"}`,
			valid: false,
		},
		{
			name:  "Invalid unbalanced braces",
			input: `{"key": "value"`,
			valid: false,
		},
		{
			name:     "Empty object",
			input:    `{}`,
			expected: map[string]interface{}{},
			valid:    true,
		},
		{
			name:     "Empty array",
			input:    `[]`,
			expected: []interface{}{},
			valid:    true,
		},
		{
			name:  "Nested array",
			input: `{"arr": [1, "two", true, null]}`,
			expected: map[string]interface{}{
				"arr": []interface{}{
					float64(1), "two", true, nil,
				},
			},
			valid: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, ok := ParseJSON(tc.input)

			if ok != tc.valid {
				t.Errorf("Expected valid: %v, got: %v", tc.valid, ok)
			}

			if tc.valid && !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result:\n%#v\nGot:\n%#v", tc.expected, actual)
			}
		})
	}
}
