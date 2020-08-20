package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

//for the key of the grid
type Cell struct {
	x int
	y int
}

func main() {
	//for the full grid size
	gridSize := 4 // 4 smaller grid of 2*2
	// a map for each cell location as key and the value of the cell
	sudokuGrid := map[Cell]int{}

	//initial grid
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			str := strings.Join([]string{"[", strconv.Itoa(sudokuGrid[Cell{i, j}]), "]"}, "")
			sudokuGrid[Cell{i, j}] = 0
			fmt.Printf("%s", str)

		}
		fmt.Println()
	}

	//fill the grid with random number

	counter := 1

	already_placed_value := int(math.Pow(float64(gridSize), 2) / 2)

	for already_placed_value < 0 {
		p, q := rand.Intn(4), rand.Intn(4)
		fmt.Println(p, q)
		sudokuGrid[Cell{p, q}] = counter
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
			already_placed_value -= 1
		}
	}

	fmt.Println(sudokuGrid)

}
