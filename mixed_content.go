package jplaw

import (
	"encoding/xml"
	"fmt"
	"html"
	"strings"
)

// ContentNode represents a node in mixed content (text or Ruby element)
type ContentNode interface {
	String() string     // Returns plain text representation
	HTML() string       // Returns HTML representation with Ruby tags
	IsRuby() bool       // Returns true if this is a Ruby node
}

// TextNode represents plain text in mixed content
type TextNode struct {
	Text string
}

// String returns the plain text
func (t TextNode) String() string {
	return t.Text
}

// HTML returns HTML-escaped text
func (t TextNode) HTML() string {
	return html.EscapeString(t.Text)
}

// IsRuby returns false for text nodes
func (t TextNode) IsRuby() bool {
	return false
}

// RubyNode represents a Ruby annotation in mixed content
type RubyNode struct {
	Base string   // The base text (e.g., "較")
	Rt   []string // The ruby text annotations (e.g., ["こう"])
}

// String returns the base text only
func (r RubyNode) String() string {
	return r.Base
}

// HTML returns HTML ruby markup
func (r RubyNode) HTML() string {
	if len(r.Rt) == 0 {
		return html.EscapeString(r.Base)
	}
	
	var sb strings.Builder
	sb.WriteString("<ruby>")
	sb.WriteString(html.EscapeString(r.Base))
	for _, rt := range r.Rt {
		sb.WriteString("<rt>")
		sb.WriteString(html.EscapeString(rt))
		sb.WriteString("</rt>")
	}
	sb.WriteString("</ruby>")
	return sb.String()
}

// IsRuby returns true for Ruby nodes
func (r RubyNode) IsRuby() bool {
	return true
}

// MixedContent holds ordered content nodes (text and Ruby elements)
type MixedContent struct {
	Nodes []ContentNode
}

// String returns the plain text representation of all nodes
func (m MixedContent) String() string {
	var sb strings.Builder
	for _, node := range m.Nodes {
		sb.WriteString(node.String())
	}
	return sb.String()
}

// HTML returns the HTML representation with Ruby tags
func (m MixedContent) HTML() string {
	var sb strings.Builder
	for _, node := range m.Nodes {
		sb.WriteString(node.HTML())
	}
	return sb.String()
}

// UnmarshalXML implements custom XML unmarshaling to preserve mixed content order
func (m *MixedContent) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	m.Nodes = []ContentNode{}
	
	for {
		token, err := d.Token()
		if err != nil {
			return err
		}
		
		switch t := token.(type) {
		case xml.EndElement:
			// End of the mixed content element
			return nil
			
		case xml.CharData:
			// Plain text content
			text := string(t)
			if text != "" {
				m.Nodes = append(m.Nodes, TextNode{Text: text})
			}
			
		case xml.StartElement:
			// Check if it's a Ruby element
			if t.Name.Local == "Ruby" {
				var ruby Ruby
				if err := d.DecodeElement(&ruby, &t); err != nil {
					return err
				}
				
				// Convert Ruby to RubyNode
				rubyNode := RubyNode{
					Base: ruby.Content,
					Rt:   make([]string, len(ruby.Rt)),
				}
				for i, rt := range ruby.Rt {
					rubyNode.Rt[i] = rt.Content
				}
				m.Nodes = append(m.Nodes, rubyNode)
			} else {
				// Skip unknown elements for now
				if err := d.Skip(); err != nil {
					return err
				}
			}
		}
	}
}

// SentenceV2 is an improved version of Sentence that properly handles mixed content
type SentenceV2 struct {
	Num         int          `xml:"Num,attr,omitempty"`
	Function    string       `xml:"Function,attr,omitempty"`
	Indent      string       `xml:"Indent,attr,omitempty"`
	WritingMode WritingMode  `xml:"WritingMode,attr,omitempty"`
	MixedContent MixedContent // Custom unmarshaling for mixed content
}

// UnmarshalXML implements custom unmarshaling for SentenceV2
func (s *SentenceV2) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
	return s.MixedContent.UnmarshalXML(d, start)
}

// String returns the plain text content
func (s *SentenceV2) String() string {
	return s.MixedContent.String()
}

// HTML returns the HTML representation with Ruby tags
func (s *SentenceV2) HTML() string {
	return s.MixedContent.HTML()
}