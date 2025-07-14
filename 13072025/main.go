package main

import (
	"container/list"
	"fmt"
)

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	count := 0

	directions := [][]int{
		{0, 1}, {1, 0}, {0, -1}, {-1, 0}, // right, down, left, up
	}

	// BFS helper
	bfs := func(r, c int) {
		queue := list.New()
		queue.PushBack([2]int{r, c})
		grid[r][c] = '0'
		for queue.Len() > 0 {
			cell := queue.Remove(queue.Front()).([2]int)
			row, col := cell[0], cell[1]

			for _, d := range directions {
				nr := row + d[0]
				nc := col + d[1]

				if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == '1' {
					queue.PushBack([2]int{nr, nc})
					grid[nr][nc] = '0'
				}
			}
		}
	}

	// Main traversal
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				count++
				bfs(r, c)
			}
		}
	}

	return count
}

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '0', '0', '1', '0'},
		{'1', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	// Copy grid if needed, since it's mutated in place
	fmt.Println("Number of Islands:", numIslands(grid)) // Output: 3
}
