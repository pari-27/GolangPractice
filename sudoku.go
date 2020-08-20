package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

//for the key of the grid
type Cell struct {
	x int
	y int
}

func main() {
	gridSize := 4 // 4 smaller grid of 2*2
	// a map for each cell location as key and the value of the cell
	sudokuGrid := map[Cell]int{}

	//initial grid
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			sudokuGrid[Cell{i, j}] = 0
			fmt.Println(i, j)
		}
	}

	//change it to something dynamic
	fmt.Println(safe_grid(sudokuGrid, 0, 1, 1))
	fmt.Println(safe_grid(sudokuGrid, 0, 3, 4))
	fmt.Println(safe_grid(sudokuGrid, 1, 0, 2))
	fmt.Println(safe_grid(sudokuGrid, 1, 2, 1))
	fmt.Println(safe_grid(sudokuGrid, 2, 0, 3))
	fmt.Println(safe_grid(sudokuGrid, 2, 2, 2))
	fmt.Println(safe_grid(sudokuGrid, 3, 1, 4))
	fmt.Println(safe_grid(sudokuGrid, 3, 3, 3))

	render_grid(sudokuGrid)

}

func safe_grid(sudokuGrid map[Cell]int, p int, q int, counter int) bool {

	//for the full grid size
	gridSize := int(math.Sqrt(float64(len(sudokuGrid))))

	empty_flag, row_flag, column_flag, box_flag := true, true, true, true
	//check if the cell is blank
	if (sudokuGrid[Cell{p, q}] != 0) {
		//set flag empty_flag to false
		empty_flag = false
	}

	//check for the same integer in the row and column
	for i := 0; i < gridSize; i++ {
		if sudokuGrid[Cell{p, i}] == counter {
			row_flag = false
		}
		if sudokuGrid[Cell{i, q}] == counter {
			column_flag = false
		}

	}
	no_of_block := int(math.Sqrt(float64(gridSize)))
	row_index := int(p / no_of_block)
	column_index := int(q / no_of_block)
	start_index := Cell{row_index * no_of_block, column_index * no_of_block}
	end_index := Cell{(row_index * no_of_block) + 2, (column_index * no_of_block) + 2}

	//check for the same integer the same box
	for i := start_index.x; i <= end_index.x; i++ {
		for j := start_index.y; j <= end_index.y; j++ {
			if sudokuGrid[Cell{i, j}] == counter {
				box_flag = false
			}

		}
	}
	//check for safe grid
	if empty_flag && row_flag && column_flag && box_flag {
		sudokuGrid[Cell{p, q}] = counter
		return true
	} else {
		return false
	}
	//fmt.Println(empty_flag, row_flag, column_flag, box_flag)

}

func render_grid(sudokuGrid map[Cell]int) {
	gridSize := int(math.Sqrt(float64(len(sudokuGrid))))

	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			str := strings.Join([]string{"[", strconv.Itoa(sudokuGrid[Cell{i, j}]), "]"}, "")
			fmt.Printf("%s", str)

		}
		fmt.Println()
	}

}
