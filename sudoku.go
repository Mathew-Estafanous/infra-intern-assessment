package main

// SolveSudoku solves a 9x9 sudoku puzzle through the means of a Constraint Satisfaction Problem
// using backtracking to find the solution. To optimize backtracking we use interleaving to validate
// the choice before continuing to the next solution. Doing so eliminates unnecessary traversals of
// invalid solutions - thereby reducing the number of backtracking steps.
func SolveSudoku(grid [][]int) [][]int {
	// deep copy grid to prevent mutation of original grid.
	result := make([][]int, len(grid))
	for i := range grid {
		result[i] = make([]int, len(grid[i]))
		copy(result[i], grid[i])
	}

	solveWithBacktracking(result)
	return result
}

func solveWithBacktracking(grid [][]int) bool {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			// continue when cell is already filled.
			if grid[r][c] != 0 {
				continue
			}

			for choice := 1; choice <= 9; choice++ {
				if !isValid(grid, r, c, choice) {
					// choice is not valid, continue to next choice.
					continue
				}

				grid[r][c] = choice
				// continue down the solution tree using this choice for grid[r][c].
				if solveWithBacktracking(grid) {
					// a solution was found with this choice.
					return true
				}
				// the choice did not lead to a solution, reset the cell and try the next choice.
				grid[r][c] = 0
			}
			return false
		}
	}
	return true
}

// isValid ensures that the choice is valid for the given row, column and 3x3 grid.
// Because implementation is in constant time, multiple calls can be made with limited impact on performance.
func isValid(grid [][]int, row, col, choice int) bool {
	// check all elements in row are not equal to choice
	for x := 0; x < 9; x++ {
		if grid[row][x] == choice {
			return false
		}
	}

	// validate all elements in column are not equal to choice
	for x := 0; x < 9; x++ {
		if grid[x][col] == choice {
			return false
		}
	}

	// identify the 3x3 grid that this cell belongs to.
	cellRow := row - (row % 3)
	cellCol := col - (col % 3)

	// validate all elements in 3x3 grid are not equal to choice
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if grid[r+cellRow][c+cellCol] == choice {
				return false
			}
		}
	}
	return true
}
