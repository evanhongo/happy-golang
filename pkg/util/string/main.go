package string

import (
	"regexp"
	"strings"
)

var spaceReg = regexp.MustCompile(`\s+`)

func CleanStr(str string) string {
	newStr := spaceReg.ReplaceAllString(str, " ")
	return strings.TrimSpace(newStr)
}
