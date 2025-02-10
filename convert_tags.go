package jplaw

import (
	"regexp"
	"strings"
)

func convertTags(source string) string {
	tags := []string{"Rt", "Rp", "Ruby", "Sup", "Sub"}
	for _, tag := range tags {
		re := regexp.MustCompile(`(?i)<(/?)` + tag + `>`)
		source = re.ReplaceAllStringFunc(source, strings.ToLower)
	}

	lineRe := regexp.MustCompile(`(?i)<Line(?: style="([^"]*)")?>`)
	source = lineRe.ReplaceAllStringFunc(source, func(match string) string {
		submatches := lineRe.FindStringSubmatch(match)
		if len(submatches) > 1 && submatches[1] != "" {
			style := ""
			switch submatches[1] {
			case "dotted":
				style = "border-bottom: 1px dotted;"
			case "double":
				style = "border-bottom: 3px double;"
			case "none":
				style = "border-bottom: none;"
			case "solid":
				style = "border-bottom: 1px solid;"
			default:
				return `<u>`
			}
			return `<u style="` + style + `">`
		}
		return `<u>`
	})

	source = strings.ReplaceAll(source, "</Line>", "</u>")

	return source
}
