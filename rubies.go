package jplaw

import (
	"regexp"
	"strings"
)

func ReplaceRubies(input string) string {
	re := regexp.MustCompile(`(?i)<(/?)(rt|rp|ruby)>`)
	return re.ReplaceAllStringFunc(input, func(match string) string {
		tag := strings.ToLower(match[1 : len(match)-1])
		return "&lt;" + tag + "&gt;"
	})
}

func RestoreRubies(input string) string {
	re := regexp.MustCompile(`&lt;(/?)(rt|rp|ruby)&gt;`)
	return re.ReplaceAllStringFunc(input, func(match string) string {
		tag := match[4 : len(match)-4]
		return "<" + tag + ">"
	})
}
