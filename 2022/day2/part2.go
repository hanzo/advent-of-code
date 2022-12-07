package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// func getPlay(them, outcome rune) rune {
// 	switch {
// 		case them == 'A'

// }

func getScore(them, me rune) int {

	// me = convert(me)

	if them == me {
		return 3
	}

	if them == 'A' {
		if me == 'X' {
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
	scores['A'] = 1
	scores['B'] = 2
	scores['C'] = 3

	scanner := bufio.NewScanner(file)
	curTotal := 0
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")

		them := rune(parts[0][0])
		me := rune(parts[1][0])

		if me == 'X' { // lose
			curTotal += 0
			switch {
			case them == 'A':
				curTotal += scores['C']
			case them == 'B':
				curTotal += scores['A']
			default:
				curTotal += scores['B']
			}
		} else if me == 'Y' { // draw
			curTotal += 3
			curTotal += scores[them]
		} else { // win
			curTotal += 6
			switch {
			case them == 'A':
				curTotal += scores['B']
			case them == 'B':
				curTotal += scores['C']
			default:
				curTotal += scores['A']
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("total: %d\n", curTotal)
}
