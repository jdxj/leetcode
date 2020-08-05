// 412
package s41

import (
	"strconv"
)

var (
	convert = map[int]string{
		3: "Fizz",
		5: "Buzz",
	}
	order = []int{
		3,
		5,
	}

	buf = make([]string, 50)
)

func fizzBuzz(n int) []string {
	buf = buf[:0]
	for i := 0; i < n; i++ {
		tmp := ""
		for j := 0; j < len(order); j++ {
			if (i+1)%order[j] == 0 {
				tmp += convert[order[j]]
			}
		}
		if tmp == "" {
			tmp = strconv.Itoa(i + 1)
		}
		buf = append(buf, tmp)
	}
	return buf
}
