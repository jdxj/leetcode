package s27

// 278
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	lower, upper := 1, n
	for lower <= upper {
		mid := (lower + upper) >> 1
		if isBadVersion(mid) {
			// 判断当前是否是最小 bad version
			if mid-1 >= 0 && !isBadVersion(mid-1) {
				return mid
			}
			upper = mid - 1
		} else {
			lower = mid + 1
		}
	}
	return -1
}

var firstBV int

func isBadVersion(version int) bool {
	return version >= firstBV
}
