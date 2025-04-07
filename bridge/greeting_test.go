package greeting

import (
	"testing"
)

func TestGreet(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "basic greeting",
			input:    "John",
			expected: "Hello, John! Welcome to Gemini Bridge.",
		},
		{
			name:     "empty name",
			input:    "",
			expected: "Hello, ! Welcome to Gemini Bridge.",
		},
		{
			name:     "special characters",
			input:    "世界",
			expected: "Hello, 世界! Welcome to Gemini Bridge.",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := Greet(tc.input)
			if result != tc.expected {
				t.Errorf("Greet(%q) = %q; want %q", tc.input, result, tc.expected)
			}
		})
	}
}
