/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	for i := 1; i < n; i++ {
		grid[0][i] = grid[0][i-1] + grid[0][i]
	}
	for i := 1; i < m; i++ {
		grid[i][0] = grid[i-1][0] + grid[i][0]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
		}
	}

	return grid[m-1][n-1]
}

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}

// @lc code=end

