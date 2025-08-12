package jplaw

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestSentence_BackwardCompatibility(t *testing.T) {
	// Test that the improved Sentence still works with existing code
	xmlData := `<Sentence Num="1">別表第二に掲げる測定器その他の設備であつて、次のいずれかに掲げる<Ruby>較<Rt>こう</Rt></Ruby>正又は校正</Sentence>`
	
	var s Sentence
	err := xml.Unmarshal([]byte(xmlData), &s)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}
	
	// Test backward compatibility: Content field should have plain text
	if s.Content == "" {
		t.Error("Content field should not be empty for backward compatibility")
	}
	
	// Content should contain the Ruby base text
	if !strings.Contains(s.Content, "較") {
		t.Errorf("Content should contain '較', got: %q", s.Content)
	}
	
	// Ruby field should have the Ruby elements for backward compatibility
	if len(s.Ruby) != 1 {
		t.Errorf("Expected 1 Ruby element for backward compatibility, got %d", len(s.Ruby))
	}
	
	if len(s.Ruby) > 0 {
		if s.Ruby[0].Content != "較" {
			t.Errorf("Ruby[0].Content = %q, want %q", s.Ruby[0].Content, "較")
		}
		if len(s.Ruby[0].Rt) != 1 || s.Ruby[0].Rt[0].Content != "こう" {
			t.Errorf("Ruby[0].Rt[0].Content = %q, want %q", s.Ruby[0].Rt[0].Content, "こう")
		}
	}
	
	// Test new HTML() method
	html := s.HTML()
	if !strings.Contains(html, "<ruby>較<rt>こう</rt></ruby>") {
		t.Errorf("HTML() should contain proper ruby tag, got: %q", html)
	}
	
	// Verify the Ruby appears in the correct position in HTML
	expectedHTMLPrefix := "別表第二に掲げる測定器その他の設備であつて、次のいずれかに掲げる<ruby>"
	if !strings.Contains(html, expectedHTMLPrefix) {
		t.Errorf("Ruby should appear in correct position in HTML")
	}
}

func TestSentence_MultipleRubyElements(t *testing.T) {
	xmlData := `<Sentence>最初の<Ruby>漢<Rt>かん</Rt></Ruby><Ruby>字<Rt>じ</Rt></Ruby>と次の<Ruby>単<Rt>たん</Rt></Ruby><Ruby>語<Rt>ご</Rt></Ruby>です</Sentence>`
	
	var s Sentence
	err := xml.Unmarshal([]byte(xmlData), &s)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}
	
	// Should have all Ruby elements
	if len(s.Ruby) != 4 {
		t.Errorf("Expected 4 Ruby elements, got %d", len(s.Ruby))
	}
	
	// HTML should have Ruby tags in correct positions
	html := s.HTML()
	expected := "最初の<ruby>漢<rt>かん</rt></ruby><ruby>字<rt>じ</rt></ruby>と次の<ruby>単<rt>たん</rt></ruby><ruby>語<rt>ご</rt></ruby>です"
	if html != expected {
		t.Errorf("HTML() = %q, want %q", html, expected)
	}
	
	// Plain text should have all kanji
	text := s.Content
	if text != "最初の漢字と次の単語です" {
		t.Errorf("Content = %q, want %q", text, "最初の漢字と次の単語です")
	}
}