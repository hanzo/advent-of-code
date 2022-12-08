package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getKey(row, col int) string {
	return fmt.Sprintf("%d,%d", row, col)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	trees := make([][]int, 99) // hardcoded 99 based on input

	visibleTrees := make(map[string]bool)

	// iterate over each row and populate 2d array
	row := 0
	for scanner.Scan() {
		line := scanner.Text()

		trees[row] = make([]int, len(line))

		ltrCurTallest := 0
		rtlCurTallest := 0

		for col := range line {
			// todo: remove double parsing
			ltrHeight, _ := strconv.Atoi(fmt.Sprintf("%c", line[col]))
			trees[row][col] = ltrHeight

			// left to right
			if col == 0 || ltrHeight > ltrCurTallest {
				ltrCurTallest = ltrHeight
				visibleTrees[getKey(row, col)] = true
			}

			// right to left
			rtlCol := len(line) - col - 1
			rtlHeight, _ := strconv.Atoi(fmt.Sprintf("%c", line[rtlCol]))
			if rtlCol == len(line)-1 || rtlHeight > rtlCurTallest {
				rtlCurTallest = rtlHeight
				visibleTrees[getKey(row, rtlCol)] = true
			}
		}
		row += 1
	}

	// iterate over each column
	for col := 0; col < len(trees[0]); col++ {

		ttbCurHeight := 0
		bttCurHeight := 0

		for row := 0; row < len(trees); row++ {

			// top to bottom
			ttbHeight := trees[row][col]

			if row == 0 || ttbHeight > ttbCurHeight {
				ttbCurHeight = ttbHeight
				visibleTrees[getKey(row, col)] = true
			}

			// bottom to top
			bttRow := len(trees) - row - 1
			bttHeight := trees[bttRow][col]

			if bttRow == len(trees)-1 || bttHeight > bttCurHeight {
				bttCurHeight = bttHeight
				visibleTrees[getKey(bttRow, col)] = true
			}
		}
	}

	for row := range trees {
		// fmt.Println(trees[row])
		for col := 0; col < len(trees[0]); col++ {
			visChar := "_"
			if visibleTrees[getKey(row, col)] == true {
				visChar = "+"
			}
			fmt.Printf("%d%s", trees[row][col], visChar)
		}
		fmt.Println()
	}

	// fmt.Println(visibleTrees)
	fmt.Println(len(visibleTrees))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
