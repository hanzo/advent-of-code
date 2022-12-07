package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Directory struct {
	name          string
	parentDir     *Directory
	subDirs       []*Directory
	totalFileSize int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	totalSize := 0

	scanner := bufio.NewScanner(file)

	rootNode := Directory{
		name: "/",
	}

	var curNode *Directory

	for scanner.Scan() {
		line := scanner.Text()

		cmdParts := strings.Split(line, " ")

		if cmdParts[0] == "$" {

			cmd := cmdParts[1]

			if cmd == "cd" {
				newDir := cmdParts[2]

				if newDir == "/" {
					curNode = &rootNode
				} else if newDir == ".." {
					dirSize := curNode.totalFileSize
					if curNode.totalFileSize <= 100000 {
						totalSize += dirSize
					}
					curNode = curNode.parentDir
					curNode.totalFileSize += dirSize // add subdir's size to parent

				} else { // cd foobar
					newNode := Directory{
						name:      newDir,
						parentDir: curNode,
					}
					curNode = &newNode
				}

			} else { // $ ls
				continue
			}

		} else if cmdParts[0] == "dir" { // dir foobar
			newNode := Directory{
				name:      cmdParts[1],
				parentDir: curNode,
			}
			curNode.subDirs = append(curNode.subDirs, &newNode)

		} else { // 1234 file.txt
			size, _ := strconv.Atoi(cmdParts[0])
			curNode.totalFileSize += size
		}
	}

	for curNode.parentDir != nil {
		// fmt.Printf("in dir: %s\n", curNode.name)
		dirSize := curNode.totalFileSize
		if curNode.totalFileSize <= 100000 {
			totalSize += dirSize
		}
		curNode = curNode.parentDir
		curNode.totalFileSize += dirSize // add subdir's size to parent
	}

	fmt.Println(totalSize)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
