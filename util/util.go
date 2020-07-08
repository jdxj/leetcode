package util

import "fmt"

func Min2(p1, p2 int) int {
	if p1 < p2 {
		return p1
	}
	return p2
}

func Min3(p1, p2, p3 int) int {
	p2 = Min2(p2, p3)
	return Min2(p1, p2)
}

func Min4(p1, p2, p3, p4 int) int {
	p2 = Min3(p2, p3, p4)
	return Min2(p1, p2)
}

func Max2(p1, p2 int) int {
	if p1 > p2 {
		return p1
	}
	return p2
}

func Max3(p1, p2, p3 int) int {
	if p2 < p3 {
		p2, p3 = p3, p2
	}
	return Max2(p1, p2)
}

func IsEqualSliceInt(p1, p2 []int) bool {
	if len(p1) != len(p2) {
		return false
	}
	for i := range p1 {
		if p1[i] != p2[i] {
			return false
		}
	}
	return true
}

func PrintTwoDimensionalArray(p [][]int) {
	for _, v := range p {
		fmt.Printf("%v\n", v)
	}
}
