package s200

import "testing"

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
