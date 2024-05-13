package main

import (
	"fmt"
)

func checks(custards *[]byte, check byte) bool {
	for i := 0; i < 9; i++ {
		if custards[i] == check {
			return false
		}
	}
	return true
}
func check(columns *[][]byte, custards *[][]byte, row, col int, check byte) bool {

	//checkCol := checks(columns[col], check)
	//if checkCol == false {
	//	return false
	//}
	//
	//fmt.Println(columns[col])
	//
	//if row < 3 {
	//	if col < 3 {
	//		return checkCol && checks(custards[0], check)
	//	} else if col >= 3 && col < 6 {
	//		return checkCol && checks(custards[1], check)
	//	} else {
	//		return checkCol && checks(custards[2], check)
	//	}
	//
	//} else if row <= 3 && row < 6 {
	//	if col < 3 {
	//		return checkCol && checks(custards[3], check)
	//	} else if col >= 3 && col < 6 {
	//		return checkCol && checks(custards[4], check)
	//	} else {
	//		return checkCol && checks(custards[5], check)
	//	}
	//} else {
	//	if col < 3 {
	//		return checkCol && checks(custards[6], check)
	//	} else if col >= 3 && col < 6 {
	//		return checkCol && checks(custards[7], check)
	//	} else {
	//		return checkCol && checks(custards[8], check)
	//	}
	//}
	return false
}
func solveSudoku(board [][]byte) {
	columns := make([][]byte, 9)
	custards := make([][]byte, 9)

	for i := 0; i < 9; i++ {
		columns[i] = make([]byte, 9)
		custards[i] = make([]byte, 9)
	}

	needs := 0

	for i, v := range board {
		for j := 0; j < 9; j++ {
			if v[j] == 46 {
				needs++
			}
			columns[j][i] = v[j]
			if i < 3 && j <= 3 {
				custards[i][j] = v[j]
			} else if i >= 3 && i < 6 && j <= 3 {
				custards[i][j] = v[j]
			} else if i >= 6 && i < 9 && j <= 3 {
				custards[i][j] = v[j]
			}

			if i < 3 && j > 3 && j <= 6 {
				custards[i][j] = v[j]
			} else if i >= 3 && i < 6 && j > 3 && j <= 6 {
				custards[i][j] = v[j]
			} else if i >= 6 && i < 9 && j > 3 && j <= 6 {
				custards[i][j] = v[j]
			}
			if i < 3 && j > 6 {
				custards[i][j] = v[j]
			} else if i >= 3 && i < 6 && j > 6 {
				custards[i][j] = v[j]
			} else if i >= 6 && i < 9 && j > 6 {
				custards[i][j] = v[j]
			}
		}
	}

	for needs > 0 {
		for i, v := range board {
			for j := 0; j < 9; j++ {
				if v[j] == 46 {
					for k := 1; k < 10; k++ {
						if check(&columns, &custards, i, j, byte(k)) {
							board[i][j] = byte(k)
							columns[j][i] = byte(k)
							needs--
							break
						}
					}
				}
			}
		}

	}

	fmt.Println(board)

}

func main() {
	test := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'}, {'6', '.', '.', '1', '9', '5', '.', '.', '.'}, {'.', '9', '8', '.', '.', '.', '.', '6', '.'}, {'8', '.', '.', '.', '6', '.', '.', '.', '3'}, {'4', '.', '.', '8', '.', '3', '.', '.', '1'}, {'7', '.', '.', '.', '2', '.', '.', '.', '6'}, {'.', '6', '.', '.', '.', '.', '2', '8', '.'}, {'.', '.', '.', '4', '1', '9', '.', '.', '5'}, {'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	solveSudoku(test)

}
