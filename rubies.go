package jplaw

import (
	"regexp"
)

// RubyTagReplacer is a regular expression to replace <Ruby>, <Rt> tags with &lt; and &gt;

var rubyTagReplacer = regexp.MustCompile(`(?s)<Ruby>(.*?)<Rt>(.*?)</Ruby>`)

// Replace replaces <Ruby>, <Rt> tags with &lt; and &gt;
func ReplaceRubies(s string) string {
	return rubyTagReplacer.ReplaceAllString(s, "&lt;$1&gt;$2")
}
