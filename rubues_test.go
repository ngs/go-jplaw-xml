package jplaw

import "testing"

func TestReplaceRubies(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "No ruby tag",
			input:    "Hello, world!",
			expected: "Hello, world!",
		},
		{
			name:     "Single ruby tag",
			input:    "Hello, <Ruby>world<Rt>ワールド</Rt></Ruby>!",
			expected: "Hello, &lt;Ruby&gt;world&lt;Rt&gt;ワールド&lt;/Rt&gt;&lt;/Ruby&gt;!",
		},
		{
			name:     "Multiple ruby tags",
			input:    "Hello, <Ruby>world<Rt>! <Ruby>Good<Rt>bye!",
			expected: "Hello, &lt;world&gt;! &lt;Good&gt;bye!",
		},
		{
			name:     "Nested ruby tags",
			input:    "Hello, <Ruby>world <Ruby>Good<Rt>bye<Rt>!",
			expected: "Hello, &lt;world &lt;Good&gt;bye&gt;!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReplaceRubies(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}
