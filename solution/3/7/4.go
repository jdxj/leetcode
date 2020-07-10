package s374

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	lower, upper := 1, n
	for lower <= upper {
		mid := (lower + upper) >> 1
		res := guess(mid)

		switch res {
		case -1:
			lower = mid + 1
		case 0:
			return mid
		case 1:
			upper = mid - 1
		}
	}
	return -1
}

var pick int

func guess(num int) int {
	if num < pick {
		return -1
	} else if num > pick {
		return 1
	} else {
		return 0
	}
}
