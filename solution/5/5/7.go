// 557
package s55

import "bytes"

func reverseWords(s string) string {
	sb := []byte(s)
	for start := 0; start < len(sb); {
		offset := bytes.Index(sb[start:], []byte(" "))
		if offset < 0 {
			offset = len(sb[start:])
		}
		end := start + offset
		reverse(sb, start, end-1)
		start = end + 1
	}
	return string(sb)
}

func reverse(sb []byte, start, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		sb[i], sb[j] = sb[j], sb[i]
	}
}
