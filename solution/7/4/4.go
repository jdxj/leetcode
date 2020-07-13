// 744
package s74

import "sort"

func nextGreatestLetter(letters []byte, target byte) byte {
	idx := sort.Search(len(letters), func(i int) bool {
		return letters[i] > target
	})
	if idx == len(letters) {
		idx = 0
	}
	return letters[idx]
}
