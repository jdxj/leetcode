package s739

import (
	"fmt"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	T := []int{73, 74, 75, 71, 69, 72, 76, 73}
	result := dailyTemperatures(T)
	fmt.Printf("%v\n", result)
}

func TestFloodFill(t *testing.T) {
	image := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	floodFill(image, 1, 1, 2)
	for _, v := range image {
		fmt.Printf("%v\n", v)
	}
}
