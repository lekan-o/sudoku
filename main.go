package main

import (
	"fmt"
	"os"
)

// Usage: go run main.go <9 rows of 9 chars each, digits or '.'>
func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}

	var board [9][9]rune
	// Read and validate 9 rows of input from command line arguments.
	for row := 0; row < 9; row++ {
		rowStr := os.Args[row+1]
		if len(rowStr) != 9 {
			fmt.Println("Error")
			return
		}
		for col, num := range rowStr {
			// each character must be '1'..'9' or '.' for empty
			if (num < '1' || num > '9') && num != '.' {
				fmt.Println("Error")
				return
			}
			board[row][col] = num
		}
	}

	// quick validity check (no duplicates in rows/cols/blocks)
	if !isValid(&board) || !solve(&board) {
		fmt.Println("Error")
		return
	}

	for row := 0; row < 9; row++ {
		for column := 0; column < 9; column++ {
			if column > 0 {
				fmt.Print(" ")
			}
			fmt.Printf("%c", board[row][column])
		}
		fmt.Println()
	}
}

// canPlace checks whether placing num at (row,col) is legal under Sudoku rules.
// num is the rune (character) for the board cell — a digit '1'..'9' or '.'.
// canPlace returns true if `num` can be placed at (row,col)
// without breaking Sudoku rules (no same number in row, col, or 3x3 block).
func canPlace(board *[9][9]rune, row, col int, num rune) bool {
	for i := 0; i < 9; i++ {
		// check same row and same column
		if board[row][i] == num || board[i][col] == num {
			return false
		}
		// compute the indices inside the 3x3 sub-block
		subRow := 3*(row/3) + i/3
		subCol := 3*(col/3) + i%3
		if board[subRow][subCol] == num {
			return false
		}
	}
	return true
}

// isValid ensures pre-filled cells don't violate rules.
func isValid(board *[9][9]rune) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			num := board[row][col]
			if num != '.' {
				// temporarily remove the value and check if it could be placed
				board[row][col] = '.'
				ok := canPlace(board, row, col, num)
				board[row][col] = num
				if !ok {
					return false
				}
			}
		}
	}
	return true
}

// solve uses simple backtracking.
func solve(board *[9][9]rune) bool {
	// backtracking solver: find an empty cell and try numbers 1..9
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				for num := '1'; num <= '9'; num++ {
					if canPlace(board, row, col, num) {
						board[row][col] = num
						if solve(board) {
							return true
						}
						// undo and try next
						board[row][col] = '.'
					}
				}
				// no valid number found for this empty cell
				return false
			}
		}
	}
	return true
}
