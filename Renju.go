package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

type Coordinate struct {
	row    int
	column int
}

func main() {
	const ROW, COLUMN = 15, 15
	var count uint = 1
	var turn string
	var piece string
	var finished bool = false

	var wg sync.WaitGroup

	grid := [ROW + 1][COLUMN + 1]string{} // 2d map grid for the main game

	blackBoard := []Coordinate{}
	whiteBoard := []Coordinate{}

	for row := 0; row <= ROW; row++ {
		for column := 0; column <= COLUMN; column++ {
			if row == 0 {
				if column < 10 {
					grid[row][column] = strconv.Itoa(column) + " "
				} else {
					grid[row][column] = strconv.Itoa(column)
				}
			} else if column == 0 {
				if row < 10 {
					grid[row][column] = strconv.Itoa(row) + " "
				} else {
					grid[row][column] = strconv.Itoa(row)
				}
			} else {
				grid[row][column] = "  "
			}
		}
	}

	for !finished {

		wg.Add(4)

		for row := 0; row <= ROW; row++ {
			fmt.Println(grid[row])
		}

		if numType(count) {
			turn = "Black"
			piece = "B "
		} else {
			turn = "White"
			piece = "W "
		}

		fmt.Printf("It is %v's turn\n", turn)

		a, b := input(ROW, COLUMN)
		for grid[a][b] != "  " {
			fmt.Println("!space already occupied!")
			a, b = input(ROW, COLUMN)
		}

		board := []Coordinate{}

		if turn == "Black" {
			blackBoard = append(blackBoard, Coordinate{a, b})
			board = blackBoard
		} else {
			whiteBoard = append(whiteBoard, Coordinate{a, b})
			board = whiteBoard
		}

		sort(board)

		grid[a][b] = piece

		go func() { // hor check
			defer wg.Done()
			indx := 1
			for _, element := range board {
				for element.column <= 11 && indx < len(board) {
					if containsHor(element.column, board) && containsHor(element.column+1, board) && containsHor(element.column+2, board) && containsHor(element.column+3, board) && containsHor(element.column+4, board) {
						for row := 0; row <= ROW; row++ {
							fmt.Println(grid[row])
						}
						fmt.Printf("%v wins!\n", turn)
						os.Exit(1)
					}
					indx++
				}
			}
		}()

		go func() { // vert check
			defer wg.Done()
			indy := 1
			for _, element := range board {
				for element.row <= 11 && indy < len(board) {
					if containsVert(element.row, board) && containsVert(element.row+1, board) && containsVert(element.row+2, board) && containsVert(element.row+3, board) && containsVert(element.row+4, board) {
						for row := 0; row <= ROW; row++ {
							fmt.Println(grid[row])
						}
						fmt.Printf("%v wins!\n", turn)
						os.Exit(1)
					}
					indy++
				}
			}
		}()

		go func() { // diagonal 1 check
			defer wg.Done()
			indd1 := 1
			for _, element := range board {
				for element.row <= 11 && element.column <= 11 && indd1 < len(board) {
					if containsDiag(element.row, element.column, board) && containsDiag(element.row+1, element.column+1, board) && containsDiag(element.row+2, element.column+2, board) && containsDiag(element.row+3, element.column+3, board) && containsDiag(element.row+4, element.column+4, board) {
						for row := 0; row <= ROW; row++ {
							fmt.Println(grid[row])
						}
						fmt.Printf("%v wins!\n", turn)
						os.Exit(1)
					}
					indd1++
				}
			}
		}()

		go func() { // diagonal 2 check
			defer wg.Done()
			indd2 := 1
			for _, element := range board {
				for element.row <= 11 && element.column >= 5 && indd2 < len(board) {
					if containsDiag(element.row, element.column, board) && containsDiag(element.row+1, element.column-1, board) && containsDiag(element.row+2, element.column-2, board) && containsDiag(element.row+3, element.column-3, board) && containsDiag(element.row+4, element.column-4, board) {
						for row := 0; row <= ROW; row++ {
							fmt.Println(grid[row])
						}
						fmt.Printf("%v wins!\n", turn)
						os.Exit(1)
					}
					indd2++
				}
			}
		}()
		wg.Wait()

		count++
	}
}

func numType(a uint) bool { // false if even, true if odd
	if a%2 == 0 {
		return false
	}
	return true
}

func containsHor(val int, brd []Coordinate) bool {
	for _, element := range brd {
		if element.column == val {
			return true
		}
	}
	return false
}

func containsVert(val int, brd []Coordinate) bool {
	for _, element := range brd {
		if element.row == val {
			return true
		}
	}
	return false
}

func containsDiag(val1, val2 int, brd []Coordinate) bool {
	for _, element := range brd {
		if element.row == val1 && element.column == val2 {
			return true
		}
	}
	return false
}

func input(i, j int) (int, int) { // row and column input control
	x := bufio.NewScanner(os.Stdin)
	y := bufio.NewScanner(os.Stdin)

	fmt.Printf("Enter row number, press ENTER, then enter column number: \n")
	x.Scan()
	y.Scan()

	if x.Text() == "x" && y.Text() == "x" { // type two x's to exit the game
		exit()
	}

	inputx, errx := strconv.ParseInt(x.Text(), 10, 64)
	inputy, erry := strconv.ParseInt(y.Text(), 10, 64)

	intInputx := int(inputx)
	intInputy := int(inputy)

	for errx != nil || erry != nil || intInputx <= 0 || intInputx > i || intInputy <= 0 || intInputy > j {
		x = bufio.NewScanner(os.Stdin)
		y = bufio.NewScanner(os.Stdin)
		fmt.Printf("Invalid coordinates for row or column. Re-enter inputs row and column: \n")
		x.Scan()
		y.Scan()

		if x.Text() == "x" && y.Text() == "x" { // type two x's to exit the game
			exit()
		}
		inputx, errx = strconv.ParseInt(x.Text(), 10, 64)
		inputy, erry = strconv.ParseInt(y.Text(), 10, 64)

		intInputx = int(inputx)
		intInputy = int(inputy)
	}

	return intInputx, intInputy
}

func sort(brd []Coordinate) {
	for i := 1; i < len(brd); i++ {
		key := brd[i]
		var j int = i - 1
		for j >= 0 && brd[j].row > key.row {
			brd[j+1] = brd[j]
			brd[j] = key
			j--
		}

		if j >= 0 && brd[j].row == key.row {
			for j >= 0 && brd[j].column > key.column {
				brd[j+1] = brd[j]
				brd[j] = key
				j--
			}
		}
	}
}

func exit() {
	fmt.Println("Exiting...")

	os.Exit(1)
}
