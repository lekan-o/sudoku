package main
import (
	"fmt"
	"os"
)
func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}
	var board [9][9]byte
	for i := 0; i < 9; i++ {
		row := os.Args[i+1]
		if len(row) != 9 {
			fmt.Println("Error")
			return
		}
		for j := 0; j < 9; j++ {
			ch := row[j]
			if (ch < '1' || ch > '9') && ch != '.' {
				fmt.Println("Error")
				return
			}
			board[i][j] = ch
		}
	}
	if !isValidSudoku(board) {
		fmt.Println("Error")
		return
	}
	if solveSudoku(&board) {
		printBoard(board)
	} else {
		fmt.Println("Error")
	}
}
func canPlace(board *[9][9]byte, row, col int, num byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == num {
			return false
		}
		if board[i][col] == num {
			return false
		}		
		subRow := 3*(row/3) + i/3
		subCol := 3*(col/3) + i%3
		if board[subRow][subCol] == num {
			return false
		}
	}
	return true
}
func solveSudoku(board *[9][9]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				for num := byte('1'); num <= '9'; num++ {
					if canPlace(board, i, j, num) {
						board[i][j] = num
						if solveSudoku(board) {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}
func printBoard(board [9][9]byte) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%c", board[i][j])
			if j < 8 {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}
func isValidSudoku(board [9][9]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			ch := board[i][j]
			if ch != '.' {
				board[i][j] = '.'
				if !canPlace(&board, i, j, ch) {
					return false
				}
				board[i][j] = ch
			}
		}
	}
	return true
}
