package jplaw

import "encoding/xml"

func Parse(data []byte) (Law, error) {
	var law Law
	xmlStr := string(data)
	if xmlStr == "" {
		return law, nil
	}

	xmlStr = ReplaceRubies(xmlStr)
	data = []byte(xmlStr)

	err := xml.Unmarshal(data, &law)

	if err != nil {
		return law, err
	}

	for i := range law.LawBody.MainProvision.Chapter {
		for j := range law.LawBody.MainProvision.Chapter[i].Article {
			for k := range law.LawBody.MainProvision.Chapter[i].Article[j].Paragraph {
				for l := range law.LawBody.MainProvision.Chapter[i].Article[j].Paragraph[k].ParagraphSentence.Sentence {
					law.LawBody.MainProvision.Chapter[i].Article[j].Paragraph[k].ParagraphSentence.Sentence[l].CharData = RestoreRubies(law.LawBody.MainProvision.Chapter[i].Article[j].Paragraph[k].ParagraphSentence.Sentence[l].CharData)
				}
			}
		}
	}

	return law, err
}
