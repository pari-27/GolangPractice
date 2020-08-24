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

var ans_grid map[Cell]int = make(map[Cell]int)

func main() {
	gridSize := 9 // 4 smaller grid of 2*2
	// a map for each cell location as key and the value of the cell
	sudokuGrid := map[Cell]int{}

	//initial grid
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			sudokuGrid[Cell{i, j}] = 0
			//fmt.Println(i, j)
		}
	}
	rand.Seed(123456789123456789)
	//change it to something dynamic
	generate_grid(sudokuGrid)
	get_ansGrid(sudokuGrid)
	fmt.Println("---------------------------------")
	render_grid(sudokuGrid)

}

func generate_grid(sudokuGrid map[Cell]int) bool {
	gridSize := int(math.Sqrt(float64(len(sudokuGrid))))
	count := 0
	for count < 30 {
		row := rand.Intn(gridSize)
		col := rand.Intn(gridSize)
		number := rand.Intn(gridSize) + 1
		//fmt.Println(row, col, number)
		if safe_grid(sudokuGrid, row, col, number) {
			sudokuGrid[Cell{row, col}] = number
			//render_grid(sudokuGrid)
			if fitGrid(sudokuGrid) {

				//fmt.Println(count)
				count = count + 1
			} else {
				sudokuGrid[Cell{row, col}] = 0
			}
		}
	}
	return true
}

func getUnassignedLocation(sudokuGrid map[Cell]int) Cell {
	gridSize := int(math.Sqrt(float64(len(sudokuGrid))))
	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			if sudokuGrid[Cell{i, j}] == 0 {
				return Cell{i, j}
			}
		}

	}
	return Cell{-1, -1}

}
func get_testgrid(sudokuGrid map[Cell]int) map[Cell]int {
	test_grid := make(map[Cell]int)
	for i, v := range sudokuGrid {
		test_grid[i] = v

	}
	return test_grid
}
func get_ansGrid(sudukoGrid map[Cell]int) {

	if fitGrid(sudukoGrid) {
		render_grid(ans_grid)
	}
}
func fitGrid(sudokuGrid map[Cell]int) bool {

	test_grid := get_testgrid(sudokuGrid)

	unassignedLocation := getUnassignedLocation(test_grid)
	row, col := unassignedLocation.x, unassignedLocation.y
	if row == -1 && col == -1 {
		ans_grid = get_testgrid(test_grid)
		return true
	}
	var check bool
	gridSize := int(math.Sqrt(float64(len(sudokuGrid))))

	for k := 1; k <= gridSize; k++ {

		if safe_grid(test_grid, row, col, k) {
			test_grid[Cell{row, col}] = k
			check = fitGrid(test_grid)
			if check == true {
				return true
			}

		}
		test_grid[Cell{row, col}] = 0

	}

	return false

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
		if i != q && sudokuGrid[Cell{p, i}] == counter {
			row_flag = false
		}
		if i != p && sudokuGrid[Cell{i, q}] == counter {
			column_flag = false
		}

	}
	no_of_block := int(math.Sqrt(float64(gridSize)))
	row_index := int(p / no_of_block)
	column_index := int(q / no_of_block)
	start_index := Cell{row_index * no_of_block, column_index * no_of_block}
	end_index := Cell{(row_index * no_of_block) + (no_of_block - 1), (column_index * no_of_block) + (no_of_block - 1)}

	//check for the same integer the same box
	for i := start_index.x; i <= end_index.x; i++ {
		for j := start_index.y; j <= end_index.y; j++ {
			if sudokuGrid[Cell{i, j}] == counter {
				box_flag = false
			}

		}
	}
	//fmt.Println(empty_flag, row_flag, column_flag, box_flag)
	// render_grid(sudokuGrid)
	//check for safe grid
	if empty_flag && row_flag && column_flag && box_flag {
		return true
	} else {
		return false
	}

}

func render_grid(sudokuGrid map[Cell]int) {
	gridSize := int(math.Sqrt(float64(len(sudokuGrid))))

	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {

			str := strings.Join([]string{"[", strconv.Itoa(sudokuGrid[Cell{i, j}]), "]"}, " ")
			if (j+1)%3 == 0 && j+1 < gridSize {
				str = strings.Join([]string{"[", strconv.Itoa(sudokuGrid[Cell{i, j}]), "]", "  |  "}, " ")
			}
			fmt.Printf("%s", str)
		}
		fmt.Println("\n")
	}

}
