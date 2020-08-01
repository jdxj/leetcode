package s55

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	s := "123456789"
	sb := []byte(s)
	reverse(sb, 0, len(sb)-1)
	fmt.Printf("%s\n", sb)
}

func TestReverseWords(t *testing.T) {
	s := "Let's take LeetCode contest"
	result := reverseWords(s)
	fmt.Printf("%s\n", result)
}
