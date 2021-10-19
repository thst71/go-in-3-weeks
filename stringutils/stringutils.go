package stringutils

import "strings"

func Ucase(incoming string) string {
	return strings.ToUpper(incoming)
}

func Lcase(incoming string) string {
	return strings.ToLower(incoming)
}
