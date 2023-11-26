package eight_queen

import "fmt"

func CanPlace(board Per, row int, column int) bool {
	for i := 0; i < row; i++ {
		// 列相同
		if board[i] == column {
			return false
		}

		// 正对角线
		if i+board[i] == row+column {
			return false
		}

		// 反对角线
		if board[i]-i == column-row {
			return false
		}
	}

	return true
}

func PrintBoard(board Per) {
	for row := 0; row < len(board); row++ {
		for column := 0; column < len(board); column++ {
			if board[row] == column {
				fmt.Print("X")
			} else {
				fmt.Print("-")
			}
		}
		fmt.Println()
	}

	fmt.Println()
}

func GetQueen(N int) int {
	count := 0

	pers := Permutation(N)
	for _, board := range pers {

		var validBoard = true
		for row, column := range board {
			if !CanPlace(board, row, column) {
				validBoard = false
				break
			}
		}

		if validBoard {
			PrintBoard(board)
			count += 1
		}
	}

	return count
}
