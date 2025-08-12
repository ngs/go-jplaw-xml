package jplaw

import (
	"encoding/xml"
	"os"
	"testing"
)

func TestUnmarshalLaw(t *testing.T) {
	// Read the test XML file
	data, err := os.ReadFile("fixtures/327AC0000000231_20240401_505AC0000000063.xml")
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}

	// Unmarshal into our Law struct
	var law Law
	err = xml.Unmarshal(data, &law)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML: %v", err)
	}

	// Test root Law element attributes
	if law.Era != EraShowa {
		t.Errorf("Expected Era to be Showa, got %v", law.Era)
	}
	if law.Lang != LanguageJapanese {
		t.Errorf("Expected Lang to be ja, got %v", law.Lang)
	}
	if law.LawType != LawTypeAct {
		t.Errorf("Expected LawType to be Act, got %v", law.LawType)
	}
	if law.Num != "231" {
		t.Errorf("Expected Num to be 231, got %v", law.Num)
	}
	if law.Year != 27 {
		t.Errorf("Expected Year to be 27, got %v", law.Year)
	}
	if law.PromulgateMonth != 7 {
		t.Errorf("Expected PromulgateMonth to be 7, got %v", law.PromulgateMonth)
	}
	if law.PromulgateDay != 15 {
		t.Errorf("Expected PromulgateDay to be 15, got %v", law.PromulgateDay)
	}

	// Test LawNum
	expectedLawNum := "昭和二十七年法律第二百三十一号"
	if law.LawNum != expectedLawNum {
		t.Errorf("Expected LawNum to be %s, got %s", expectedLawNum, law.LawNum)
	}

	// Test LawTitle
	if law.LawBody.LawTitle == nil {
		t.Fatal("Expected LawTitle to be non-nil")
	}
	expectedTitle := "航空法"
	if law.LawBody.LawTitle.Content != expectedTitle {
		t.Errorf("Expected LawTitle content to be %s, got %s", expectedTitle, law.LawBody.LawTitle.Content)
	}
	expectedKana := "こうくうほう"
	if law.LawBody.LawTitle.Kana != expectedKana {
		t.Errorf("Expected LawTitle Kana to be %s, got %s", expectedKana, law.LawBody.LawTitle.Kana)
	}

	// Test TOC existence
	if law.LawBody.TOC == nil {
		t.Fatal("Expected TOC to be non-nil")
	}
	
	// Test TOCLabel
	if law.LawBody.TOC.TOCLabel == nil {
		t.Fatal("Expected TOCLabel to be non-nil")
	}
	expectedTOCLabel := "目次"
	if law.LawBody.TOC.TOCLabel.Content != expectedTOCLabel {
		t.Errorf("Expected TOCLabel content to be %s, got %s", expectedTOCLabel, law.LawBody.TOC.TOCLabel.Content)
	}

	// Test TOCChapter count and first chapter
	if len(law.LawBody.TOC.TOCChapter) == 0 {
		t.Fatal("Expected at least one TOCChapter")
	}
	
	firstChapter := law.LawBody.TOC.TOCChapter[0]
	if firstChapter.Num != "1" {
		t.Errorf("Expected first TOCChapter Num to be 1, got %s", firstChapter.Num)
	}
	expectedChapterTitle := "第一章　総則"
	if firstChapter.ChapterTitle.Content != expectedChapterTitle {
		t.Errorf("Expected first TOCChapter title to be %s, got %s", expectedChapterTitle, firstChapter.ChapterTitle.Content)
	}
	
	// Test ArticleRange for first chapter
	if firstChapter.ArticleRange == nil {
		t.Fatal("Expected first TOCChapter ArticleRange to be non-nil")
	}
	expectedArticleRange := "（第一条・第二条）"
	if firstChapter.ArticleRange.Content != expectedArticleRange {
		t.Errorf("Expected first TOCChapter ArticleRange to be %s, got %s", expectedArticleRange, firstChapter.ArticleRange.Content)
	}

	// Test that we have sections in chapter 9
	chapter9Found := false
	for _, chapter := range law.LawBody.TOC.TOCChapter {
		if chapter.Num == "9" {
			chapter9Found = true
			if len(chapter.TOCSection) == 0 {
				t.Error("Expected Chapter 9 to have TOCSections")
			}
			if len(chapter.TOCSection) > 0 {
				firstSection := chapter.TOCSection[0]
				if firstSection.Num != "1" {
					t.Errorf("Expected first section in Chapter 9 to have Num 1, got %s", firstSection.Num)
				}
				expectedSectionTitle := "第一節　危害行為防止基本方針等"
				if firstSection.SectionTitle.Content != expectedSectionTitle {
					t.Errorf("Expected first section title to be %s, got %s", expectedSectionTitle, firstSection.SectionTitle.Content)
				}
			}
			break
		}
	}
	if !chapter9Found {
		t.Error("Expected to find Chapter 9 in TOC")
	}
}

func TestMarshalLaw(t *testing.T) {
	// Create a simple Law struct for testing marshaling
	law := Law{
		Era:             EraHeisei,
		Year:            31,
		Num:             "1",
		PromulgateMonth: 4,
		PromulgateDay:   1,
		LawType:         LawTypeAct,
		Lang:            LanguageJapanese,
		LawNum:          "平成三十一年法律第一号",
		LawBody: LawBody{
			LawTitle: &LawTitle{
				Content: "テスト法",
				Kana:    "てすとほう",
			},
			MainProvision: MainProvision{
				Article: []Article{
					{
						Num: "1",
						ArticleTitle: &ArticleTitle{
							Content: "目的",
						},
						Paragraph: []Paragraph{
							{
								Num: 1,
								ParagraphNum: ParagraphNum{
									Content: "1",
								},
								ParagraphSentence: ParagraphSentence{
									Sentence: []Sentence{
										{
											Num:     1,
											Content: "この法律は、テストを目的とする。",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	// Marshal to XML
	data, err := xml.MarshalIndent(law, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal Law to XML: %v", err)
	}

	// Unmarshal back to verify structure
	var unmarshaled Law
	err = xml.Unmarshal(data, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal marshaled XML: %v", err)
	}

	// Verify some key fields
	if unmarshaled.Era != law.Era {
		t.Errorf("Era mismatch after round-trip: expected %v, got %v", law.Era, unmarshaled.Era)
	}
	if unmarshaled.LawNum != law.LawNum {
		t.Errorf("LawNum mismatch after round-trip: expected %s, got %s", law.LawNum, unmarshaled.LawNum)
	}
	if unmarshaled.LawBody.LawTitle.Content != law.LawBody.LawTitle.Content {
		t.Errorf("LawTitle content mismatch after round-trip: expected %s, got %s", 
			law.LawBody.LawTitle.Content, unmarshaled.LawBody.LawTitle.Content)
	}

	t.Logf("Successfully marshaled and unmarshaled Law struct")
}

func TestWritingModeConstants(t *testing.T) {
	if WritingModeVertical != "vertical" {
		t.Errorf("Expected WritingModeVertical to be 'vertical', got %s", WritingModeVertical)
	}
	if WritingModeHorizontal != "horizontal" {
		t.Errorf("Expected WritingModeHorizontal to be 'horizontal', got %s", WritingModeHorizontal)
	}
}

func TestEmptyOptionalFields(t *testing.T) {
	// Test that omitempty works correctly for optional fields
	law := Law{
		Era:     EraReiwa,
		Year:    5,
		Num:     "1",
		LawType: LawTypeAct,
		Lang:    LanguageJapanese,
		LawNum:  "令和五年法律第一号",
		LawBody: LawBody{
			MainProvision: MainProvision{
				Article: []Article{
					{
						Num: "1",
						Paragraph: []Paragraph{
							{
								Num: 1,
								ParagraphNum: ParagraphNum{
									Content: "1",
								},
								ParagraphSentence: ParagraphSentence{
									Sentence: []Sentence{
										{
											Content: "テスト",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	data, err := xml.MarshalIndent(law, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal Law: %v", err)
	}

	xmlStr := string(data)
	
	// Check that optional attributes are omitted when zero/empty
	if containsString(xmlStr, "PromulgateMonth=\"0\"") {
		t.Error("Expected zero PromulgateMonth to be omitted")
	}
	if containsString(xmlStr, "PromulgateDay=\"0\"") {
		t.Error("Expected zero PromulgateDay to be omitted")
	}
	
	t.Logf("Generated XML (partial): %s", xmlStr[:min(500, len(xmlStr))])
}

func containsString(haystack, needle string) bool {
	return len(haystack) >= len(needle) && findString(haystack, needle) >= 0
}

func findString(haystack, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
