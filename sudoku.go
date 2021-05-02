package main

import (
	"fmt"
	"os"
)

// puts '0' instead of '.'

func EspaceEntre(line int, numbers string, sudoku *[9][9]int) { // enter values of numbers, "entre" in string and numbers of sudoku
	for col, char := range numbers { // assigning an index number to all symbols and assigning them the meaning of runes
		for index := 0; index <= int(char-'0'); /* this phrase is needed, because the programm displays the values of the table in ASCII */ index++ {
			sudoku[line][col] = index
		}
	}
}

func SudokuValid(sudoku *[9][9]int) bool { // this function-conditions is needed for drawing numbers in Sudoku
	// checking lines
	for line := 0; line < 9; line++ { // sets the line
		count := [10]int{}             // the condition that programme can use 10 numbers 0 -> 10
		for col := 0; col < 9; col++ { // sets the column
			count[sudoku[line][col]]++ // fill in sudoku by LINE and column
		}
		if HasDouble(count) { // call the function for check if there are identical numbers
			return false
		}
	}
	//checking columns
	for line := 0; line < 9; line++ {
		count := [10]int{}
		for col := 0; col < 9; col++ {
			count[sudoku[col][line]]++ // fill in sudoku by COLUMN and line
		}
		if HasDouble(count) { // call the function HasDouble, which looks if there are identical numbers
			return false
		}
	}
	// checking squares 3:3
	for line := 0; line < 9; line += 3 { //sets the line
		for col := 0; col < 9; col += 3 { // sets the column
			count := [10]int{}
			for lineTrois := line; lineTrois < line+3; lineTrois++ { // sets lines squares 3:3
				for colTrois := col; colTrois < col+3; colTrois++ { // sets column squares 3:3
					count[sudoku[lineTrois][colTrois]]++ // fill the SQUARES 3:3
				}
				if HasDouble(count) { // call the function for check if there are identical numbers
					return false
				}
			}
		}

	}
	return true

}

// this function looks to see if there are identical numbers
func HasDouble(count [10]int) bool {
	for index, count := range count {
		if index == 0 {
			continue
		}
		if count > 1 {
			return true
		}
	}
	return false
}
func BackTrack(sudoku *[9][9]int) bool {
	if HasNOTEmpty(sudoku) { // call the function HasNOTEmpty
		return true
	}
	for line := 0; line < 9; line++ { // sets the lines
		for col := 0; col < 9; col++ { // sets the columns
			if sudoku[line][col] == 0 { // condition, if present "0" in line and column
				for trueNbr := 1; trueNbr <= 9; trueNbr++ { // loop of numbers that can be replaced with "0"
					sudoku[line][col] = trueNbr                   // remplace zeros
					if SudokuValid(sudoku) && BackTrack(sudoku) { // checking the prescribed cinditions in the functions SudokuValid and BackTrack
						return true
					} else {
						sudoku[line][col] = 0 // condition, if present "0" in line and column
					}
				}
				return false
			}
		}
	}
	return false //to exit the loop
}
func HasNOTEmpty(sudoku *[9][9]int) bool {
	for line := 0; line < 9; line++ { // sets the lines
		for col := 0; col < 9; col++ { // sets the columns
			if sudoku[line][col] == 0 { // checking "0" in Sudoku
				return false
			}
		}
	}
	return true
}
func PrintSudoku(sudoku [9][9]int) {
	for line := 0; line < 9; line++ { // sets the lines
		for col := 0; col < 9; col++ { // sets the columns
			fmt.Printf("%d", sudoku[line][col]) // print sudoku
			if col < 8 {                        // condition for print space between all numbers in the table
				fmt.Printf(" ")
			}

		}
		fmt.Printf("\n") // transition from the one line to anther
	}
}

func main() { //this function is basic
	Args := os.Args[1:]   // removes first argument
	sudoku := [9][9]int{} // creates a table of 9 symbols by columns and lines with numerical values
	check := 0            // using to count all characters

	for index, val := range Args { // numering of all lines by indices and output of all characters of all lines in one line
		EspaceEntre(index, val, &sudoku) // function call EspaceEntre
		for range val {                  // counting all elements of table, where check will be 81
			check++
		}
	}
	// conditions with which the program wll not work and will display an error
	if check < 81 { // for condition if there are less than 81 characters
		fmt.Printf("Error\n")
	} else { // else call two funtions 1)BackTrack 2)PrintSudoku
		if BackTrack(&sudoku) {
			PrintSudoku(sudoku)
		} else { // else error
			fmt.Printf("Error\n")
		}
	}
}
