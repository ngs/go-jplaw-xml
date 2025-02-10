package jplaw

import "testing"

func TestConvertTags(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"<Rt>Text</Rt>", "<rt>Text</rt>"},
		{"<Rp>Text</Rp>", "<rp>Text</rp>"},
		{"<Ruby>Text</Ruby>", "<ruby>Text</ruby>"},
		{"<Sup>Text</Sup>", "<sup>Text</sup>"},
		{"<Sub>Text</Sub>", "<sub>Text</sub>"},
		{"<Rt>Text</Rt><Rp>Text</Rp>", "<rt>Text</rt><rp>Text</rp>"},
		{"<Line>Text</Line>", "<u>Text</u>"},
		{"<Line style=\"dotted\">Text</Line>", "<u style=\"border-bottom: 1px dotted;\">Text</u>"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := convertTags(test.input)
			if result != test.expected {
				t.Errorf("Expected '%s', got '%s'", test.expected, result)
			}
		})
	}
}
