package main

import (
	"bufio"
	"fmt"
	"os"
)

func check_same(lines []string, x int, y int) bool {
	return (lines[y+1][x+1] == 'M' && lines[y-1][x-1] == 'S') || (lines[y+1][x+1] == 'S' && lines[y-1][x-1] == 'M')
}

func check_different(lines []string, x int, y int) bool {
	return (lines[y-1][x+1] == 'M' && lines[y+1][x-1] == 'S') || (lines[y-1][x+1] == 'S' && lines[y+1][x-1] == 'M')
}

func check_x(lines []string, x int, y int) bool {
	return check_same(lines, x, y) && check_different(lines, x, y)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	count := 0
	for y := 1; y < len(lines)-1; y++ {
		for x := 1; x < len(lines[y])-1; x++ {
			if lines[y][x] == 'A' && check_x(lines, x, y) {
				count += 1
			}
		}
	}
	fmt.Println(count)
}
