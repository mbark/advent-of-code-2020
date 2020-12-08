package util

import "strings"


func ReadInput(in, splitBy string) []string {
	trimmed := strings.Trim(in, "\n")
	return strings.Split(trimmed, splitBy)
}
