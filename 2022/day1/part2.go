package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	curTotal := 0
	elfTotals := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			elfTotals = append(elfTotals, curTotal)
			curTotal = 0
		} else {
			intVal, _ := strconv.Atoi(line)
			curTotal += int(intVal)
		}

		// fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(elfTotals)

	fmt.Printf("totals: %v\n", elfTotals)

	topThree := elfTotals[len(elfTotals)-1] + elfTotals[len(elfTotals)-2] + elfTotals[len(elfTotals)-3]

	fmt.Printf("top 3: %v\n", topThree)
}
