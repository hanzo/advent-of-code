package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

		firstHalf := line[:len(line)/2]
		secondHalf := line[len(line)/2:]

		var commonChar rune

		for _, char := range firstHalf {
			if strings.ContainsRune(secondHalf, char) {
				commonChar = char
				break
			}

		}

		var charVal int
		if commonChar >= 65 && commonChar <= 90 {
			charVal = int(commonChar) - 38
		} else {
			charVal = int(commonChar) - 96
		}

		curTotal += charVal

		// fmt.Printf("line: %s, first: %s, second: %s, common: %c (%v)\n", line, firstHalf, secondHalf, commonChar, charVal)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("total: %d\n", curTotal)
}
