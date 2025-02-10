package jplaw

import (
	"io"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	file, err := os.Open("fixtures/327AC0000000231_20240401_505AC0000000063.xml")
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	schema, err := Parse(data)

	if err != nil {
		t.Fatalf("Failed to parse schema: %v", err)
	}

	tests := [][]interface{}{
		{schema.LawType, LawTypeAct, "LawType"},
		{schema.Era, EraShowa, "Era"},
		{schema.Year, 27, "Year"},
		{schema.Num, 231, "Num"},
		{schema.PromulgateMonth, 7, "PromulgateMonth"},
		{schema.PromulgateDay, 15, "PromulgateDay"},
		{schema.Lang, LanguageJapanese, "Lang"},
		{string(schema.LawNum), "昭和二十七年法律第二百三十一号", "LawNum"},
		{schema.LawBody.LawTitle.Content, "航空法", "LawTitle"},
		{len(schema.LawBody.MainProvision.Chapter), 13, "Chapter count"},
		{len(schema.LawBody.MainProvision.Chapter[0].Article), 2, "Article count in first Chapter"},
		{len(schema.LawBody.MainProvision.Chapter[0].Article[0].Paragraph), 1, "Paragraph count in first Article"},
		{schema.LawBody.MainProvision.Chapter[0].Article[0].Paragraph[0].ParagraphSentence.Sentence[0].Content, "この法律は、国際民間航空条約の規定並びに同条約の附属書として採択された標準、方式及び手続に準拠して、航空機の航行の安全及び航空機の航行に起因する障害の防止を図るための方法を定め、航空機を運航して営む事業の適正かつ合理的な運営を確保して輸送の安全を確保するとともにその利用者の利便の増進を図り、並びに航空の脱炭素化を推進するための措置を講じ、あわせて無人航空機の飛行における遵守事項等を定めてその飛行の安全の確保を図ることにより、航空の発達を図り、もつて公共の福祉を増進することを目的とする。", "ParagraphSentence"},
		{schema.LawBody.MainProvision.Chapter[0].Article[1].Paragraph[7].ParagraphSentence.Sentence[0].Content, "この法律において「進入表面」とは、着陸帯の短辺に接続し、且つ、水平面に対し上方へ五十分の一以上で国土交通省令で定める<ruby>勾<rt>こう</rt></ruby>配を有する平面であつて、その投影面が進入区域と一致するものをいう。", "ParagraphSentenceWithRuby"},
	}

	for _, test := range tests {
		t.Run(test[2].(string), func(t *testing.T) {
			if test[0] != test[1] {
				t.Errorf("Expected %s to be '%v', got '%v'", test[2], test[1], test[0])
			}
		})
	}
}
