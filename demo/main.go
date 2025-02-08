package main

import (
	"encoding/xml"
	"io"
	"log"
	"os"

	"github.com/ngs/go-jplaw-xml"
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
	log.Println("LawTitle:", schema.LawBody.LawTitle.CharData)
	for _, chapter := range schema.LawBody.MainProvision.Chapter {
		log.Println("  Chapter:", chapter.ChapterTitle)
		for _, article := range chapter.Article {
			log.Println("    Article:", article.ArticleTitle)
			for _, paragraph := range article.Paragraph {
				log.Println("     Paragraph:", paragraph.ParagraphNum)
				for _, sentence := range paragraph.ParagraphSentence.Sentence {
					log.Println("      Sentence:", sentence.CharData)
				}
			}
		}
	}
}
