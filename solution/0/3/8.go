// 38
package s03

var (
	cache = make(map[int]string, 30)
	buf   = make([]byte, 0, 1000)
)

func countAndSay(n int) string {
	if v, ok := cache[n]; ok {
		return v
	}
	// 找不到的话就计算
	var res string
	for i := 0; i < n; i++ {
		res = next(i, res)
	}
	return res
}

func next(idx int, pre string) string {
	nextIdx := idx + 1
	if v, ok := cache[nextIdx]; ok {
		return v
	}
	if idx == 0 {
		cache[nextIdx] = "1"
		return cache[nextIdx]
	}

	// 清空
	buf = buf[:0]
	preLen := len(pre)
	for start, end := 0, 0; start < preLen; {
		if end < preLen && pre[end] == pre[start] {
			end++
			continue
		}
		count := end - start
		buf = append(buf, '0'+byte(count), pre[start])
		start = end
	}
	cache[nextIdx] = string(buf)
	return cache[nextIdx]
}
