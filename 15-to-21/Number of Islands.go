package _5_to_21

//Given a 2d grid map of '1's (land) and '0's (water), count the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.
//
//Example 1:
//
//Input:
//11110
//11010
//11000
//00000
//
//Output: 1
//
//Example 2:
//
//Input:
//11000
//11000
//00100
//00011
//
//Output: 3

func numIslands(grid [][]byte) int {
	l := len(grid)
	var numberOfLand int

	for i:=0; i< l; i++{
		row := grid[i]
		for j:=0; j<len(row); j++{
			if grid[i][j] == '1'{
				numberOfLand += dfs(grid, i, j)
			}

		}
	}

	return numberOfLand
}

func dfs(grid[][]byte, i , j int) int{
	if i<0  || i >= len(grid) || j < 0 || j >= len(grid[i]) || grid[i][j] == '0'{
		return 0
	}

	grid[i][j] = '0'
	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
	return 1
}

