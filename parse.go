package jplaw

import "encoding/xml"

func Parse(data []byte) (Law, error) {
	xmlStr := string(data)
	if xmlStr == "" {
		return Law{}, nil
	}
	xmlStr = convertTags(xmlStr)
	data = []byte(xmlStr)
	var schema Law
	err := xml.Unmarshal(data, &schema)
	return schema, err
}
