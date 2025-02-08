package jplaw

import (
	"encoding/xml"
	"io"
	"os"
	"testing"
)

func TestParseSchema(t *testing.T) {
	file, err := os.Open("fixtures/327AC0000000231_20240401_505AC0000000063.xml")
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	var schema Law
	err = xml.Unmarshal(data, &schema)

	if err != nil {
		t.Fatalf("Failed to parse schema: %v", err)
	}

	if schema.LawType != LawTypeAct {
		t.Errorf("Expected LawType to be '%s', got '%s'", LawTypeAct, schema.LawType)
	}

	if schema.Era != EraShowa {
		t.Errorf("Expected Era to be '%s', got '%s'", EraShowa, schema.Era)
	}

	if schema.Year != 27 {
		t.Errorf("Expected Year to be 27, got %d", schema.Year)
	}

	if schema.Num != 231 {
		t.Errorf("Expected Num to be 231, got %d", schema.Num)
	}

	if schema.PromulgateMonth != 7 {
		t.Errorf("Expected PromulgateMonth to be 7, got %d", schema.PromulgateMonth)
	}

	if schema.PromulgateDay != 15 {
		t.Errorf("Expected PromulgateDay to be 15, got %d", schema.PromulgateDay)
	}

	if schema.Lang != LanguageJapanese {
		t.Errorf("Expected Lang to be 'ja', got '%s'", schema.Lang)
	}

	if schema.LawNum == "" {
		t.Error("Expected LawNum to be non-empty")
	}

	if schema.LawBody.LawTitle.CharData == "" {
		t.Error("Expected LawTitle to be non-empty")
	}

	if len(schema.LawBody.MainProvision.Chapter) == 0 {
		t.Error("Expected at least one Chapter in MainProvision")
	}

	if len(schema.LawBody.MainProvision.Chapter[0].Article) == 0 {
		t.Error("Expected at least one Article in the first Chapter")
	}

	if len(schema.LawBody.MainProvision.Chapter[0].Article[0].Paragraph) == 0 {
		t.Error("Expected at least one Paragraph in the first Article")
	}

	if schema.LawBody.MainProvision.Chapter[0].Article[0].Paragraph[0].ParagraphSentence.Sentence[0].CharData == "" {
		t.Error("Expected ParagraphSentence to have non-empty CharData")
	}

	if schema.LawNum != "昭和二十七年法律第二百三十一号" {
		t.Errorf("Expected LawNum to be '昭和二十七年法律第二百三十一号', got '%s'", schema.LawNum)
	}

	if schema.LawBody.LawTitle.CharData != "航空法" {
		t.Errorf("Expected LawTitle to be '航空法', got '%s'", schema.LawBody.LawTitle.CharData)
	}

	if len(schema.LawBody.MainProvision.Chapter) != 13 {
		t.Errorf("Expected 13 Chapter in MainProvision, got %d", len(schema.LawBody.MainProvision.Chapter))
	}

	if len(schema.LawBody.MainProvision.Chapter[0].Article) != 2 {
		t.Errorf("Expected 2 Article in the first Chapter, got %d", len(schema.LawBody.MainProvision.Chapter[0].Article))
	}

	if len(schema.LawBody.MainProvision.Chapter[0].Article[0].Paragraph) != 1 {
		t.Errorf("Expected 1 Paragraph in the first Article, got %d", len(schema.LawBody.MainProvision.Chapter[0].Article[0].Paragraph))
	}

	if schema.LawBody.MainProvision.Chapter[0].Article[0].Paragraph[0].ParagraphSentence.Sentence[0].CharData != "この法律は、国際民間航空条約の規定並びに同条約の附属書として採択された標準、方式及び手続に準拠して、航空機の航行の安全及び航空機の航行に起因する障害の防止を図るための方法を定め、航空機を運航して営む事業の適正かつ合理的な運営を確保して輸送の安全を確保するとともにその利用者の利便の増進を図り、並びに航空の脱炭素化を推進するための措置を講じ、あわせて無人航空機の飛行における遵守事項等を定めてその飛行の安全の確保を図ることにより、航空の発達を図り、もつて公共の福祉を増進することを目的とする。" {
		t.Errorf("Expected ParagraphSentence to have 'この法律は、(snip)', got '%s'", schema.LawBody.MainProvision.Chapter[0].Article[0].Paragraph[0].ParagraphSentence.Sentence[0].CharData)
	}

	// Add more assertions based on the expected structure of the schema
}
