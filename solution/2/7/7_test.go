package s27

import "testing"

func TestNumSquares(t *testing.T) {
	correct := 3
	count := numSquares(12)
	if count != correct {
		t.Fatalf("failed, correct: %d, your: %d\n", correct, count)
	}
}
