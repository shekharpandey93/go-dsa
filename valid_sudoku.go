package main


func main() {
	var arr [][]byte
	arr = append(arr, []byte{"5","3",".",".","7",".",".",".","."})
	arr := [][]{
	[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]
	}
}

func isValidSudoku(board [][]byte) bool {

	var rows, columns, squares [9][9]bool
	for i, row := range board{
		for j, column := range row{
			if column != '.' {
				k := column-49
				if rows[j][k] || columns[i][k] || squares[3*(i/3)+j/3][k] {
					return false
				}
				rows[j][k] = true
				columns[i][k] = true
				squares[3*(i/3)+j/3][k] = true
			}
		}
	}
	return true
	// for i := 0; i < 9; i++ {
	//     if !isValidRow(board[i]) {
	//         return false
	//     }
	// }
	// for i := 0; i < 9; i++ {
	//     if !isValidColumn(board, i) {
	//         return false
	//     }
	// }
	// for i := 0; i < 9; i += 3 {
	//     for j := 0; j < 9; j += 3 {
	//         if !isValidCell(board, i, j) {
	//             return false
	//         }
	//     }
	// }
	// return true
}

func isValidRow(board []byte) bool {
	checker, num := 0, 0
	for i := 0; i < 9; i++ {
		if board[i] != '.' {
			num = int(board[i] - '0')
			if checker & (1 << num) != 0 {
				return false
			}
			checker = checker | (1 << num)
		}
	}
	return true
}
func isValidColumn(board [][]byte, j int) bool {
	checker, num := 0, 0
	for i := 0; i < 9; i++ {
		if board[i][j] != '.' {
			num = int(board[i][j] - '0')
			if checker & (1 << num) != 0 {
				return false
			}
			checker = checker | (1 << num)
		}
	}
	return true
}
func isValidCell(board [][]byte, n, m int) bool {
	checker, num := 0, 0
	for i := n; i < n + 3; i++ {
		for j := m; j < m + 3; j++ {
			if board[i][j] != '.' {
				num = int(board[i][j] - '0')
				if checker & (1 << num) != 0 {
					return false
				}
				checker = checker | (1 << num)
			}
		}
	}
	return true
}



