package main

import (
	"encoding/xml"
	"io"
	"log"
	"os"

	"go.ngs.io/jp-law-xml"
)

func main() {
	file, err := os.Open("../fixtures/327AC0000000231_20240401_505AC0000000063.xml")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	var schema jplaw.Law
	err = xml.Unmarshal(data, &schema)

	if err != nil {
		log.Fatalf("Failed to parse schema: %v", err)
	}

	log.Println("LawNum:", schema.LawNum)
	if schema.LawBody.LawTitle != nil {
		log.Println("LawTitle:", schema.LawBody.LawTitle.Content)
	}
	for _, chapter := range schema.LawBody.MainProvision.Chapter {
		log.Println("  Chapter:", chapter.ChapterTitle.Content)
		for _, article := range chapter.Article {
			if article.ArticleTitle != nil {
				log.Println("    Article:", article.ArticleTitle.Content)
			}
			for _, paragraph := range article.Paragraph {
				log.Println("     Paragraph:", paragraph.ParagraphNum.Content)
				for _, sentence := range paragraph.ParagraphSentence.Sentence {
					log.Println("      Sentence:", sentence.Content)
				}
			}
		}
	}
}
