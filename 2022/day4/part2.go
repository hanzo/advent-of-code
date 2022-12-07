package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	curTotal := 0

	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)

		ranges := strings.Split(line, ",")

		range1 := strings.Split(ranges[0], "-")
		range1Start, _ := strconv.Atoi(range1[0])
		range1End, _ := strconv.Atoi(range1[1])

		range2 := strings.Split(ranges[1], "-")
		range2Start, _ := strconv.Atoi(range2[0])
		range2End, _ := strconv.Atoi(range2[1])

		fmt.Printf("line: %s, %d-%d, %d-%d\n", line, range1Start, range1End, range2Start, range2End)

		if (range1Start >= range2Start && range1Start <= range2End) ||
			(range1End >= range2Start && range1End <= range2End) ||
			(range2Start >= range1Start && range2Start <= range1End) ||
			(range2End >= range1Start && range2End <= range1End) {
			fmt.Printf("range in contained\n")
			curTotal += 1
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("total: %d\n", curTotal)
}
