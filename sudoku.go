package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	readTable()
}

// 3: Check some base cases. If the index is at the end of the matrix,
//  i.e. i=N-1 and j=N then check if the grid is safe or not, if safe
//  print the grid and return true else return false. The other base case
//  is when the value of column is N, i.e j = N,
//  then move to next row, i.e. i++ and j = 0.
func isSafe(grid []int, row, col, num int) bool {
	return true
}

func printTable(table [][]int) {
	for j, row := range table {
		for i, d := range row {
			if i%3 == 0 {
				fmt.Print("|") //		|0 9 6|0 4 0|0 0 1|
			} // 				 		|1 0 0|0 6 0|0 0 4|
			fmt.Print(d)  //	 		|5 0 4|8 1 0|3 9 0|
			if i%3 != 2 { //     		-------------------
				fmt.Print(" ") // 		|0 0 7|9 5 0|0 4 3|
			} // 			  	  		|0 3 0|0 8 0|0 0 0|
			if i == 8 { // 		  		|4 0 5|0 2 3|0 1 8|
				fmt.Print("|") // 		-------------------
			} // 				  		|0 1 0|6 3 0|0 5 9|
		} //					  		|0 5 9|0 7 0|8 3 0|
		fmt.Println() // 		 		|0 0 3|5 9 0|0 0 7|
		if j%3 == 2 && j != len(table) {
			fmt.Println("-------------------")
		}
	}
}

func readTable() [][]int { // go run . "blablabla" >> grid
	N := 9
	var table [][]int
	for i := 0; i < N; i++ {
		table = append(table, make([]int, N))
	}

	rows := os.Args[1:]
	for i, row := range rows {
		for j, symb := range row {
			if symb != '.' {
				table[i][j], _ = strconv.Atoi(string(symb))
			} else {
				table[i][j] = 0
			}
		}
	}
	return table
}
