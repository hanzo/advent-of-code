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
	// [N]     [Q]         [N]
	// [R]     [F] [Q]     [G] [M]
	// [J]     [Z] [T]     [R] [H] [J]
	// [T] [H] [G] [R]     [B] [N] [T]
	// [Z] [J] [J] [G] [F] [Z] [S] [M]
	// [B] [N] [N] [N] [Q] [W] [L] [Q] [S]
	// [D] [S] [R] [V] [T] [C] [C] [N] [G]
	// [F] [R] [C] [F] [L] [Q] [F] [D] [P]
	//  1   2   3   4   5   6   7   8   9

	stacks := make([]Stack, 10)
	stacks[1].Push('F')
	stacks[1].Push('D')
	stacks[1].Push('B')
	stacks[1].Push('Z')
	stacks[1].Push('T')
	stacks[1].Push('J')
	stacks[1].Push('R')
	stacks[1].Push('N')

	stacks[2].Push('R')
	stacks[2].Push('S')
	stacks[2].Push('N')
	stacks[2].Push('J')
	stacks[2].Push('H')

	stacks[3].Push('C')
	stacks[3].Push('R')
	stacks[3].Push('N')
	stacks[3].Push('J')
	stacks[3].Push('G')
	stacks[3].Push('Z')
	stacks[3].Push('F')
	stacks[3].Push('Q')

	stacks[4].Push('F')
	stacks[4].Push('V')
	stacks[4].Push('N')
	stacks[4].Push('G')
	stacks[4].Push('R')
	stacks[4].Push('T')
	stacks[4].Push('Q')

	stacks[5].Push('L')
	stacks[5].Push('T')
	stacks[5].Push('Q')
	stacks[5].Push('F')

	stacks[6].Push('Q')
	stacks[6].Push('C')
	stacks[6].Push('W')
	stacks[6].Push('Z')
	stacks[6].Push('B')
	stacks[6].Push('R')
	stacks[6].Push('G')
	stacks[6].Push('N')

	stacks[7].Push('F')
	stacks[7].Push('C')
	stacks[7].Push('L')
	stacks[7].Push('S')
	stacks[7].Push('N')
	stacks[7].Push('H')
	stacks[7].Push('M')

	stacks[8].Push('D')
	stacks[8].Push('N')
	stacks[8].Push('Q')
	stacks[8].Push('M')
	stacks[8].Push('T')
	stacks[8].Push('J')

	stacks[9].Push('P')
	stacks[9].Push('G')
	stacks[9].Push('S')

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)

		inputParts := strings.Split(line, " ")

		moveCount, _ := strconv.Atoi(inputParts[1])
		fromStack, _ := strconv.Atoi(inputParts[3])
		toStack, _ := strconv.Atoi(inputParts[5])

		fmt.Printf("line: %s, move %d from %d to %d\n", line, moveCount, fromStack, toStack)

		for i := 0; i < moveCount; i++ {
			crate := stacks[fromStack].Pop()

			stacks[toStack].Push(crate)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 1; i <= 9; i++ {
		// fmt.Printf("stack %d has top element %c\n", i, stacks[i].Peek())
		fmt.Printf("%c", stacks[i].Peek())
	}

}

// package stack

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

// Create a new stack
func New() *Stack {
	return &Stack{nil, 0}
}

// Return the number of items in the stack
func (this *Stack) Len() int {
	return this.length
}

// View the top item on the stack
func (this *Stack) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}

// Pop the top item of the stack and return it
func (this *Stack) Pop() interface{} {
	if this.length == 0 {
		return nil
	}

	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

// Push a value onto the top of the stack
func (this *Stack) Push(value interface{}) {
	n := &node{value, this.top}
	this.top = n
	this.length++
}
