package jplaw

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestMixedContent_UnmarshalXML(t *testing.T) {
	tests := []struct {
		name     string
		xml      string
		wantText string
		wantHTML string
	}{
		{
			name:     "Plain text only",
			xml:      `<Sentence Num="1">これは普通のテキストです。</Sentence>`,
			wantText: "これは普通のテキストです。",
			wantHTML: "これは普通のテキストです。",
		},
		{
			name:     "Text with inline Ruby",
			xml:      `<Sentence Num="1">別表第二に掲げる測定器その他の設備であつて、次のいずれかに掲げる<Ruby>較<Rt>こう</Rt></Ruby>正又は校正</Sentence>`,
			wantText: "別表第二に掲げる測定器その他の設備であつて、次のいずれかに掲げる較正又は校正",
			wantHTML: "別表第二に掲げる測定器その他の設備であつて、次のいずれかに掲げる<ruby>較<rt>こう</rt></ruby>正又は校正",
		},
		{
			name:     "Multiple Ruby elements",
			xml:      `<Sentence><Ruby>漢<Rt>かん</Rt></Ruby><Ruby>字<Rt>じ</Rt></Ruby>の読み方</Sentence>`,
			wantText: "漢字の読み方",
			wantHTML: "<ruby>漢<rt>かん</rt></ruby><ruby>字<rt>じ</rt></ruby>の読み方",
		},
		{
			name:     "Ruby with multiple Rt elements",
			xml:      `<Sentence>これは<Ruby>振<Rt>ふ</Rt><Rt>り</Rt></Ruby>仮名です</Sentence>`,
			wantText: "これは振仮名です",
			wantHTML: "これは<ruby>振<rt>ふ</rt><rt>り</rt></ruby>仮名です",
		},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s SentenceV2
			err := xml.Unmarshal([]byte(tt.xml), &s)
			if err != nil {
				t.Fatalf("Failed to unmarshal XML: %v", err)
			}
			
			// Test plain text extraction
			gotText := s.String()
			if gotText != tt.wantText {
				t.Errorf("String() = %q, want %q", gotText, tt.wantText)
			}
			
			// Test HTML generation
			gotHTML := s.HTML()
			if gotHTML != tt.wantHTML {
				t.Errorf("HTML() = %q, want %q", gotHTML, tt.wantHTML)
			}
		})
	}
}

func TestMixedContent_ComplexStructure(t *testing.T) {
	// Test with the actual problematic XML from the law document
	xmlData := `<ItemSentence>
		<Sentence Num="1" WritingMode="vertical">別表第二に掲げる測定器その他の設備であつて、次のいずれかに掲げる<Ruby>較<Rt>こう</Rt></Ruby>正又は校正（以下この号、第三十八条の三第一項第二号及び第三十八条の八第二項において「較正等」という。）を受けたもの（その較正等を受けた日の属する月の翌月の一日から起算して一年（無線設備の点検を行うのに優れた性能を有する測定器その他の設備として総務省令で定める測定器その他の設備に該当するものにあつては、当該測定器その他の設備の区分に応じ、一年を超え三年を超えない範囲内で総務省令で定める期間）以内のものに限る。）を使用して無線設備の点検を行うものであること。</Sentence>
	</ItemSentence>`
	
	// Define a temporary struct for ItemSentence with SentenceV2
	type ItemSentenceV2 struct {
		Sentence []SentenceV2 `xml:"Sentence,omitempty"`
	}
	
	var item ItemSentenceV2
	err := xml.Unmarshal([]byte(xmlData), &item)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}
	
	if len(item.Sentence) != 1 {
		t.Fatalf("Expected 1 sentence, got %d", len(item.Sentence))
	}
	
	s := item.Sentence[0]
	
	// Check that the text contains "較"
	text := s.String()
	if !strings.Contains(text, "較") {
		t.Errorf("Plain text should contain '較', got: %q", text)
	}
	
	// Check that the HTML contains the ruby tag
	html := s.HTML()
	if !strings.Contains(html, "<ruby>較<rt>こう</rt></ruby>") {
		t.Errorf("HTML should contain ruby tag, got: %q", html)
	}
	
	// Check that Ruby appears in the correct position
	expectedPrefix := "別表第二に掲げる測定器その他の設備であつて、次のいずれかに掲げる"
	if !strings.HasPrefix(text, expectedPrefix) {
		t.Errorf("Text should start with expected prefix")
	}
	
	// Check position of 較 in the text
	idxKyou := strings.Index(text, "較")
	idxSei := strings.Index(text, "正又は校正")
	if idxKyou < 0 || idxSei < 0 || idxKyou >= idxSei {
		t.Errorf("較 should appear before 正又は校正 in the text")
	}
}

func TestMixedContent_SpecialCharacters(t *testing.T) {
	// Test HTML escaping
	xmlData := `<Sentence>これは<Ruby>&lt;タグ&gt;<Rt>たぐ</Rt></Ruby>です&amp;</Sentence>`
	
	var s SentenceV2
	err := xml.Unmarshal([]byte(xmlData), &s)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}
	
	// Plain text should have the actual characters
	wantText := "これは<タグ>です&"
	if s.String() != wantText {
		t.Errorf("String() = %q, want %q", s.String(), wantText)
	}
	
	// HTML should escape special characters
	wantHTML := "これは<ruby>&lt;タグ&gt;<rt>たぐ</rt></ruby>です&amp;"
	if s.HTML() != wantHTML {
		t.Errorf("HTML() = %q, want %q", s.HTML(), wantHTML)
	}
}