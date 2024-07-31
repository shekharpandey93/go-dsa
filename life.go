package main

import "fmt"

func main()  {
	rows:= 15
	cols := 15
	board := [15][15]int{ { 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
		{ 0, 0, 0, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0 },
		{ 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
		{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
		{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
		{ 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
		{ 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
		{ 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
		{ 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 },
		{ 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0 }}
	generationWeWantToTest(board, rows,cols, 2)
}

func generationWeWantToTest(board [15][15]int, rows,cols, generationWeWantToTest int )  {
	future := [15][15]int{}
	//generation := 1;
	//for generation <= generationWeWantToTest {
		for i := 1; i < rows -1; i++ {
			for j := 1; j < cols -1; j++ {
				aliveNeighbours := 0;
				for d1 := -1; d1 <= 1; d1++ {
					for d2 := -1; d2 <= 1; d2++ {
						if (d1 == 0 && d2 == 0) {
							continue
						}
						aliveNeighbours  += board[i + d1][j + d2]
					}
				}
				//aliveNeighbours -= board[i][j]

				future[i][j] = board[i][j] // assign value
				if ((board[i][j] == 1) && (aliveNeighbours < 2)) { //dies only single
					future[i][j] = 0
				} else if ((board[i][j] == 1) && (aliveNeighbours > 3)) {  //dies only over population
					future[i][j] = 0
				} else if ((board[i][j] == 0) && (aliveNeighbours > 3)) { // active
					future[i][j] = 1
				}
			}
		}
		//generation++
	//}
	fmt.Println("=====future===",future)
}
