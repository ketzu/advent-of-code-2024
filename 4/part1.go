package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check_horizontal(lines []string, x int, y int, word string) int {
	if x+len(word) > len(lines[0]) {
		return 0
	}
	for i := 1; i < len(word); i++ {
		if lines[y][x+i] != word[i] {
			return 0
		}
	}

	return 1
}

func check_vertical(lines []string, x int, y int, word string) int {
	if y+len(word) > len(lines) {
		return 0
	}
	for i := 1; i < len(word); i++ {
		if lines[y+i][x] != word[i] {
			return 0
		}
	}

	return 1
}

func check_diagonal(lines []string, x int, y int, word string) int {
	if x+len(word) > len(lines) || y+len(word) > len(lines[0]) {
		return 0
	}
	for i := 1; i < len(word); i++ {
		if lines[y+i][x+i] != word[i] {
			return 0
		}
	}

	return 1
}

func check_diagonal_backwards(lines []string, x int, y int, word string) int {
	if x-len(word)+1 < 0 || y+len(word) > len(lines[0]) {
		return 0
	}
	for i := 1; i < len(word); i++ {
		if lines[y+i][x-i] != word[i] {
			return 0
		}
	}

	return 1
}

func check(lines []string, x int, y int, word string) int {
	return check_diagonal(lines, x, y, word) + check_horizontal(lines, x, y, word) + check_vertical(lines, x, y, word) + check_diagonal_backwards(lines, x, y, word)
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

	debug := []string{}

	count := 0
	for y := 0; y < len(lines); y++ {
		var debug_line strings.Builder
		for x := 0; x < len(lines[y]); x++ {
			tmpcount := count
			switch lines[y][x] {
			case 'X':
				count += check(lines, x, y, "XMAS")
				if tmpcount != count {
					debug_line.WriteString("X")
				} else {
					debug_line.WriteString(".")
				}
			case 'S':
				count += check(lines, x, y, "SAMX")
				if tmpcount != count {
					debug_line.WriteString("S")
				} else {
					debug_line.WriteString(".")
				}
			default:
				debug_line.WriteString(".")
			}
		}
		debug = append(debug, debug_line.String())
	}
	for _, line := range debug {
		fmt.Println(line)
	}
	fmt.Println(count)
}
