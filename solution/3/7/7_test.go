package s374

import (
	"fmt"
	"testing"
)

func TestGuessNumber(t *testing.T) {
	pick = 6
	res := guessNumber(10)
	fmt.Printf("%d\n", res)
}
