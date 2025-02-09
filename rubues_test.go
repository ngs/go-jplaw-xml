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
			input:    "Hello, <Ruby>world<Rp>(</Rp><Rt>ワールド</Rt><Rp>)</Rp></Ruby>!",
			expected: "Hello, &lt;ruby&gt;world&lt;rp&gt;(&lt;/rp&gt;&lt;rt&gt;ワールド&lt;/rt&gt;&lt;rp&gt;)&lt;/rp&gt;&lt;/ruby&gt;!",
		},
		{
			name:     "Multiple ruby tags",
			input:    "Hello, <Ruby>world<Rt>! <Ruby>Good<Rt>bye!",
			expected: "Hello, &lt;ruby&gt;world&lt;rt&gt;! &lt;ruby&gt;Good&lt;rt&gt;bye!",
		},
		{
			name:     "Nested ruby tags",
			input:    "Hello, <Ruby>world <Ruby>Good<Rt>bye<Rt>!",
			expected: "Hello, &lt;ruby&gt;world &lt;ruby&gt;Good&lt;rt&gt;bye&lt;rt&gt;!",
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

func TestRestoreRubies(t *testing.T) {
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
			input:    "Hello, &lt;ruby&gt;world&lt;rp&gt;(&lt;/rp&gt;&lt;rt&gt;ワールド&lt;/rt&gt;&lt;rp&gt;)&lt;/rp&gt;&lt;/ruby&gt;!",
			expected: "Hello, <ruby>world<rp>(</rp><rt>ワールド</rt><rp>)</rp></ruby>!",
		},
		{
			name:     "Multiple ruby tags",
			input:    "Hello, &lt;ruby&gt;world&lt;rt&gt;! &lt;ruby&gt;Good&lt;rt&gt;bye!",
			expected: "Hello, <ruby>world<rt>! <ruby>Good<rt>bye!",
		},
		{
			name:     "Nested ruby tags",
			input:    "Hello, &lt;ruby&gt;world &lt;ruby&gt;Good&lt;rt&gt;bye&lt;rt&gt;!",
			expected: "Hello, <ruby>world <ruby>Good<rt>bye<rt>!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RestoreRubies(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}
