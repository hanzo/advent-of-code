package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getScore(them, me rune) int {

	me = convert(me)

	if them == me {
		return 3
	}

	if them == 'A' {
		if me == 'B' {
			return 6
		} else {
			return 0
		}
	}

	if them == 'B' {
		if me == 'C' {
			return 6
		} else {
			return 0
		}
	}

	// if them == 'C' {
	if me == 'A' {
		return 6
	} else {
		return 0
	}
}

func convert(before rune) rune {
	switch {
	case before == 'X':
		return 'A'
	case before == 'Y':
		return 'B'
	// case before == 'Z':
	// 	return 'C'
	default:
		return 'C'
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scores := make(map[rune]int)
	scores['X'] = 1
	scores['Y'] = 2
	scores['Z'] = 3

	scanner := bufio.NewScanner(file)
	curTotal := 0
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")

		them := rune(parts[0][0])
		me := rune(parts[1][0])

		curTotal += scores[me]
		curTotal += getScore(them, me)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("total: %d\n", curTotal)
}
