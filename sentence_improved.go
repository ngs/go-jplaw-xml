package jplaw

import (
	"encoding/xml"
	"fmt"
)

// SentenceImproved is an improved version of Sentence that handles mixed content properly
// while maintaining backward compatibility
type SentenceImproved struct {
	Num          int             `xml:"Num,attr,omitempty"`
	Function     string          `xml:"Function,attr,omitempty"` // main or proviso
	Indent       string          `xml:"Indent,attr,omitempty"`
	WritingMode  WritingMode     `xml:"WritingMode,attr,omitempty"`
	Content      string          // Plain text content for backward compatibility
	Line         []Line          `xml:"Line,omitempty"`
	QuoteStruct  []QuoteStruct   `xml:"QuoteStruct,omitempty"`
	ArithFormula []ArithFormula  `xml:"ArithFormula,omitempty"`
	Ruby         []Ruby          // Ruby elements for backward compatibility
	Sup          []Sup           `xml:"Sup,omitempty"`
	Sub          []Sub           `xml:"Sub,omitempty"`
	MixedContent MixedContent    // The properly parsed mixed content
}

// UnmarshalXML implements custom unmarshaling to handle mixed content while maintaining backward compatibility
func (s *SentenceImproved) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// Parse attributes
	for _, attr := range start.Attr {
		switch attr.Name.Local {
		case "Num":
			fmt.Sscanf(attr.Value, "%d", &s.Num)
		case "Function":
			s.Function = attr.Value
		case "Indent":
			s.Indent = attr.Value
		case "WritingMode":
			s.WritingMode = WritingMode(attr.Value)
		}
	}
	
	// Parse mixed content
	if err := s.MixedContent.UnmarshalXML(d, start); err != nil {
		return err
	}
	
	// Set backward compatibility fields
	s.Content = s.MixedContent.String()
	
	// Extract Ruby elements for backward compatibility
	s.Ruby = []Ruby{}
	for _, node := range s.MixedContent.Nodes {
		if rubyNode, ok := node.(RubyNode); ok && rubyNode.IsRuby() {
			ruby := Ruby{
				Content: rubyNode.Base,
				Rt:      make([]Rt, len(rubyNode.Rt)),
			}
			for i, rt := range rubyNode.Rt {
				ruby.Rt[i] = Rt{Content: rt}
			}
			s.Ruby = append(s.Ruby, ruby)
		}
	}
	
	return nil
}

// String returns the plain text content
func (s *SentenceImproved) String() string {
	return s.MixedContent.String()
}

// HTML returns the HTML representation with Ruby tags
func (s *SentenceImproved) HTML() string {
	return s.MixedContent.HTML()
}

// GetContent returns the plain text content (for backward compatibility)
func (s *SentenceImproved) GetContent() string {
	if s.Content != "" {
		return s.Content
	}
	return s.MixedContent.String()
}

// GetRubyHTML returns HTML with Ruby annotations
func (s *SentenceImproved) GetRubyHTML() string {
	return s.MixedContent.HTML()
}