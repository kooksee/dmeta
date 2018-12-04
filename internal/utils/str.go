package utils

import (
	"fmt"
	"strings"
)

func Str(a interface{}) string {
	if a == nil {
		return ""
	}
	return fmt.Sprintf("%s", a)
}

func StrJoin(sep string, s ... string) string {
	return strings.Join(s, sep)
}
func StrOf(s ... string) []string {
	return s
}
