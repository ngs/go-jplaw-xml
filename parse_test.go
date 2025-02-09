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

	tests := []struct {
		name     string
		got      interface{}
		expected interface{}
	}{
		{
			"LawType",
			schema.LawType,
			LawTypeAct,
		},
		{
			"Era",
			schema.Era,
			EraShowa,
		},
		{
			"Year",
			schema.Year,
			27,
		},
		{
			"Num",
			schema.Num,
			231,
		},
		{
			"PromulgateMonth",
			schema.PromulgateMonth,
			7,
		},
		{
			"PromulgateDay",
			schema.PromulgateDay,
			15,
		},
		{
			"Lang",
			schema.Lang,
			LanguageJapanese,
		},
		{
			"LawNum",
			schema.LawNum,
			"昭和二十七年法律第二百三十一号",
		},
		{
			"LawTitle",
			schema.LawBody.LawTitle.CharData,
			"航空法",
		},
		{
			"ChapterCount",
			len(schema.LawBody.MainProvision.Chapter),
			13,
		},
		{
			"FirstChapterArticleCount",
			len(schema.LawBody.MainProvision.Chapter[0].Article),
			2,
		},
		{
			"FirstArticleParagraphCount",
			len(schema.LawBody.MainProvision.Chapter[0].Article[0].Paragraph),
			1,
		},
		{
			"FirstParagraphSentence",
			schema.LawBody.MainProvision.Chapter[0].Article[0].Paragraph[0].ParagraphSentence.Sentence[0].CharData,
			"この法律は、国際民間航空条約の規定並びに同条約の附属書として採択された標準、方式及び手続に準拠して、航空機の航行の安全及び航空機の航行に起因する障害の防止を図るための方法を定め、航空機を運航して営む事業の適正かつ合理的な運営を確保して輸送の安全を確保するとともにその利用者の利便の増進を図り、並びに航空の脱炭素化を推進するための措置を講じ、あわせて無人航空機の飛行における遵守事項等を定めてその飛行の安全の確保を図ることにより、航空の発達を図り、もつて公共の福祉を増進することを目的とする。",
		},
		{
			"SecondArticleEighthParagraphSentence",
			schema.LawBody.MainProvision.Chapter[0].Article[1].Paragraph[7].ParagraphSentence.Sentence[0].CharData,
			"この法律において「進入表面」とは、着陸帯の短辺に接続し、且つ、水平面に対し上方へ五十分の一以上で国土交通省令で定める<ruby>勾<rt>こう</rt></ruby>配を有する平面であつて、その投影面が進入区域と一致するものをいう。",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.expected {
				t.Errorf("Expected %s to be '%v', got '%v'", tt.name, tt.expected, tt.got)
			}
		})
	}
}
