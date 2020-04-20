package _5_to_21

//Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.
//
//Note: You can only move either down or right at any point in time.
//
//Example:
//
//Input:
//[
//[1,3,1],
//[1,5,1],
//[4,2,1]
//]
//Output: 7
//Explanation: Because the path 1→3→1→1→1 minimizes the sum.


func minPathSum(grid [][]int) int {
	R := len(grid)
	if R == 0 {
		return 0
	}
	C := len(grid[0])
	if C == 0 {
		return 0
	}
	for c := 1; c != C; c++ { // first row
		grid[0][c] += grid[0][c-1]
	}
	for r := 1; r != R; r++ {
		grid[r][0] += grid[r-1][0]
		for c := 1; c != C; c++ {
			grid[r][c] += min(grid[r-1][c], grid[r][c-1])
		}
	}
	return grid[R-1][C-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b

}