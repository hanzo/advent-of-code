package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getScenicScore(trees [][]int, startRow, startCol int) int {

	// any edge tree will have at least one direction with a scenic score of zero, so the total scenic score will be zero
	if startRow == 0 || startRow == len(trees)-1 || startCol == 0 || startCol == len(trees[0])-1 {
		return 0
	}

	height := trees[startRow][startCol]

	// left
	leftCount := 1
	for col := startCol - 1; col > 0; col-- {
		if trees[startRow][col] >= height {
			break
		}
		leftCount++
	}

	// right
	rightCount := 1
	for col := startCol + 1; col < len(trees[0])-1; col++ {
		if trees[startRow][col] >= height {
			break
		}
		rightCount++
	}

	// up
	upCount := 1
	for row := startRow - 1; row > 0; row-- {
		if trees[row][startCol] >= height {
			break
		}
		upCount++
	}

	// down
	downCount := 1
	for row := startRow + 1; row < len(trees)-1; row++ {
		if trees[row][startCol] >= height {
			break
		}
		downCount++
	}

	return leftCount * rightCount * upCount * downCount
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	trees := make([][]int, 99) // hardcoded 99 based on input

	row := 0
	for scanner.Scan() {
		line := scanner.Text()

		trees[row] = make([]int, len(line))

		for col := range line {
			height, _ := strconv.Atoi(fmt.Sprintf("%c", line[col]))
			trees[row][col] = height
		}
		row += 1
	}

	bestScenicScore := 0

	for row := range trees {
		for col := 0; col < len(trees[0]); col++ {

			scenicScore := getScenicScore(trees, row, col)
			if scenicScore > bestScenicScore {
				bestScenicScore = scenicScore
			}

			fmt.Printf("%d-%d,", trees[row][col], scenicScore)
		}
		fmt.Println()
	}

	fmt.Println(bestScenicScore)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
