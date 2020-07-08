package s739

// DFS
// 1 1 1
// 1 1 0
// 1 0 1
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	oldColor := image[sr][sc]
	// 不用改
	if oldColor == newColor {
		return image
	}
	image[sr][sc] = newColor

	// 修改 "上" 为 newColor
	if sr-1 >= 0 && image[sr-1][sc] == oldColor {
		floodFill(image, sr-1, sc, newColor)
	}
	// 修改 "下" 为 newColor
	if sr+1 < len(image) && image[sr+1][sc] == oldColor {
		floodFill(image, sr+1, sc, newColor)
	}
	// 修改 "左" 为 newColor
	if sc-1 >= 0 && image[sr][sc-1] == oldColor {
		floodFill(image, sr, sc-1, newColor)
	}
	// 修改 "右" 为 newColor
	if sc+1 < len(image[0]) && image[sr][sc+1] == oldColor {
		floodFill(image, sr, sc+1, newColor)
	}
	return image
}
