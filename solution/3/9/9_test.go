package s394

import "testing"

func TestDecodeString(t *testing.T) {
	s := "3[a2[c]]"
	correct := "accaccacc"
	result := decodeString(s)
	if result != correct {
		t.Fatalf("failed, correct: %s, your: %s\n", correct, result)
	}

	s = "abc3[cd]xyz"
	correct = "abccdcdcdxyz"
	result = decodeString(s)
	if result != correct {
		t.Fatalf("failed, correct: %s, your: %s\n", correct, result)
	}

	s = "2[abc]3[cd]ef"
	correct = "abcabccdcdcdef"
	result = decodeString(s)
	if result != correct {
		t.Fatalf("failed, correct: %s, your: %s\n", correct, result)
	}
}
