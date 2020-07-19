package s20

func NewUnionFind(grid [][]byte) *UnionFind {
	rows := len(grid)
	cols := len(grid[0])
	uf := &UnionFind{}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// 两种初始化对照
			// 1.
			//if grid[i][j] == '1' {
			//	uf.roots = append(uf.roots, i*cols+j)
			//	uf.count++
			//} else {
			//	uf.roots = append(uf.roots, -1)
			//}
			// 2.
			if grid[i][j] == '1' {
				uf.count++
			}
			uf.roots = append(uf.roots, i*cols+j)

			uf.rank = append(uf.rank, 0)
		}
	}
	return uf
}

type UnionFind struct {
	roots []int // 不同集合的 root
	rank  []int // 不同 root 的最大深度. 优化手段: 平衡深度 (合并时造成的不平衡)
	count int   // roots 的个数
}

func (uf *UnionFind) Find(idx int) int {
	roots := uf.roots
	// 不是 idx 不是 root
	if roots[idx] != idx {
		// 优化手段: 路径压缩 (缩短查找深度).
		// 继续向上查找.
		roots[idx] = uf.Find(roots[idx])
	}
	return roots[idx]
}

func (uf *UnionFind) Unite(x, y int) {
	// 找到它们的 root 是谁
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	// 已经在同一集合不用做任何事情
	if rootX == rootY {
		return
	}

	// 合并操作
	rank := uf.rank
	if rank[rootX] < rank[rootY] {
		rootX, rootY = rootY, rootX // 保持 rootX 的 rank 为最大
	}
	uf.roots[rootY] = rootX // 合并
	if rank[rootX] == rank[rootY] {
		rank[rootX]++
	}
	uf.count-- // 合并后 roots 个数减少
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	uf := NewUnionFind(grid)
	rows := len(grid)
	cols := len(grid[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '0' {
				continue
			}
			grid[i][j] = '0' // 可以看作是优化
			// 合并上下左右到中间
			if i-1 >= 0 && grid[i-1][j] == '1' {
				uf.Unite(i*cols+j, (i-1)*cols+j)
			}
			if i+1 < rows && grid[i+1][j] == '1' {
				uf.Unite(i*cols+j, (i+1)*cols+j)
			}
			if j-1 >= 0 && grid[i][j-1] == '1' {
				uf.Unite(i*cols+j, i*cols+(j-1))
			}
			if j+1 < cols && grid[i][j+1] == '1' {
				uf.Unite(i*cols+j, i*cols+(j+1))
			}
		}
	}
	return uf.count
}
