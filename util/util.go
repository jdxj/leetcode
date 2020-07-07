package util

func Min2(p1, p2 int) int {
	if p1 < p2 {
		return p1
	}
	return p2
}

func Min3(p1, p2, p3 int) int {
	if p2 > p3 {
		p2, p3 = p3, p2
	}
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
