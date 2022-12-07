package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func unique(runeSlice string) []rune {
	keys := make(map[rune]bool)
	list := []rune{}
	for _, entry := range runeSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func getBadgePriority(elves []string) int {

	charMap := make(map[rune]int)

	for _, line := range elves {

		unique := unique(line)

		for _, char := range unique {
			charMap[char] += 1
		}
	}

	fmt.Printf("%v\n", charMap)

	var commonChar rune
	for char, count := range charMap {
		if count == 3 {
			commonChar = char
			break
		}
	}

	fmt.Println(commonChar)

	var charVal int
	if commonChar >= 65 && commonChar <= 90 {
		charVal = int(commonChar) - 38
	} else {
		charVal = int(commonChar) - 96
	}

	// fmt.Printf("line: %s, first: %s, second: %s, common: %c (%v)\n", line, firstHalf, secondHalf, commonChar, charVal)

	return charVal
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	curTotal := 0

	var curElves []string
	for scanner.Scan() {
		line := scanner.Text()

		curElves = append(curElves, line)

		if len(curElves) == 3 {
			curTotal += getBadgePriority(curElves)
			curElves = nil
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("total: %d\n", curTotal)
}
