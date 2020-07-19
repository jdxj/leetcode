package s20

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewUnionFind(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	count := numIslands(grid)
	if count != 3 {
		t.Fatalf("failed, corrent: %d, count: %d\n", 3, count)
	}
}

func TestDiv(t *testing.T) {
	n := 19
	correct := 82
	if next(n) != correct {
		t.Fatalf("n: %d, correct: %d\n", n, correct)
	}
	n = 999
	correct = 243
	if next(n) != correct {
		t.Fatalf("n: %d, correct: %d\n", n, correct)
	}
	n = 9999999999999
	correct = 1053
	if next(n) != correct {
		t.Fatalf("n: %d, correct: %d\n", n, correct)
	}
}

func TestCharInString(t *testing.T) {
	n := "abc"
	fmt.Printf("%s\n", reflect.TypeOf(n[0]).Kind())
	for i, v := range n {
		fmt.Printf("%d: %s\n", i, reflect.TypeOf(v).Kind())
	}
}
