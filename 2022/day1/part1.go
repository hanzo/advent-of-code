package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	curMaxTotal := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			if curTotal > curMaxTotal {
				curMaxTotal = curTotal
			}
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

	fmt.Printf("max: %d\n", curMaxTotal)
}
