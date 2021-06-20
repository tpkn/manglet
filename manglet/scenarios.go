package manglet

import (
	"regexp"
	"strings"
)

var digits_rule = regexp.MustCompile("[0-9]")
var letters_rule = regexp.MustCompile("(?i)[a-z]")
var symbols_rule = regexp.MustCompile("[" + regexp.QuoteMeta("!$%^&*()_+|~=`{}[]:\";'<>?./\\") + "]")

//	Scenarios applies mangle functions, set by user
func Scenarios(line string, mangle_groups []string) string {
	for _, s := range mangle_groups {
		switch s {
		case "l": line = letters_rule.ReplaceAllString(line, "x")
		case "d": line = digits_rule.ReplaceAllString(line, "0")
		case "s": line = symbols_rule.ReplaceAllString(line, "-")
		}
	}
	
	return line
}

func QuickReplace(text, dictionary, replacement string) string {
	if len(text) > 0 {
		for _, v := range  strings.Split(dictionary, "") {
			result := strings.ReplaceAll(text, v, replacement)
			if result != text {
				text = result
			}
		}
	}

	return text
}