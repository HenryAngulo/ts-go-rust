package main

import (
	"fmt"
	"strings"
)

func getInput() string{
  return `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`
}

func main() {
	treeChar := "#"
	treeCount := 0
	for r, line := range strings.Split(getInput(), "\n"){
		currentPos := (r * 3) % len(line)
		if string(line[currentPos]) == treeChar{
			treeCount++
		}
	}
	fmt.Printf("Tree count: %d\n", treeCount)
}