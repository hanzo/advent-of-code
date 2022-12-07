package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	queue := make([]rune, 0)

	for i, ch := range line {

		queue = append(queue, ch)

		if len(queue) <= 14 {
			continue
		}

		// remove oldest element
		queue = queue[1:]

		seen := make(map[rune]bool, 0)
		hasDupe := false
		for _, r := range queue {
			if _, ok := seen[r]; ok {
				hasDupe = true
				continue
			} else {
				seen[r] = true
			}
		}
		if !hasDupe {
			// fmt.Printf("%v\n", queue)
			// fmt.Printf("%v\n", seen)
			panic(i + 1)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

// queue := make([]int, 0)
// // Push to the queue
// queue = append(queue, 1)
// // Top (just get next element, don't remove it)
// x = queue[0]
// // Discard top element
// queue = queue[1:]
// // Is empty ?
// if len(queue) == 0 {
//     fmt.Println("Queue is empty !")
// }
