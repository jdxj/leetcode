// 151
package s15

import (
	"strings"
)

func reverseWords(s string) string {
	fields := strings.Fields(s)
	if len(fields) < 1 {
		return ""
	}

	for i, j := 0, len(fields)-1; i < j; i, j = i+1, j-1 {
		fields[i], fields[j] = fields[j], fields[i]
	}
	return strings.Join(fields, " ")
}
