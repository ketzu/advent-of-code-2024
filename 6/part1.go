package main

import (
	"bufio"
	"fmt"
	"os"
)

func read_input(filename string) [][]byte {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	input := [][]byte{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bytes := scanner.Bytes()
		tmp := make([]byte, len(bytes))
		copy(tmp, bytes)
		input = append(input, tmp)
	}
	return input
}

func find_start(world [][]byte) (int, int) {
	for y, row := range world {
		for x, cell := range row {
			if cell == '^' {
				return x, y
			}
		}
	}
	return -1, -1
}

func turn_right(x, y int) (int, int) {
	if x == 0 && y == -1 {
		return 1, 0 // right
	}
	if x == 1 && y == 0 {
		return 0, 1 // down
	}
	if x == 0 && y == 1 {
		return -1, 0 // left
	}
	if x == -1 && y == 0 {
		return 0, -1 // upo
	}
	return 0, 0
}

func to_direction(x, y int) string {
	if x == 0 && y == -1 {
		return "up"
	}
	if x == 1 && y == 0 {
		return "right"
	}
	if x == 0 && y == 1 {
		return "down"
	}
	if x == -1 && y == 0 {
		return "left"
	}
	return "unknown"
}

func main() {
	world := read_input("input.txt")
	x, y := find_start(world)
	d_x, d_y := 0, -1

	fmt.Println(x, y, to_direction(d_x, d_y))
	fmt.Println(len(world), len(world[0]))

	counter := 1
	world[y][x] = 'X'

	for true {
		t_x := x + d_x
		t_y := y + d_y

		if t_y < 0 || t_y >= len(world) || t_x < 0 || t_x >= len(world[t_y]) {
			break
		}

		switch world[t_y][t_x] {
		case '.':
			counter += 1
			world[t_y][t_x] = 'X'
			fallthrough
		case 'X':
			x = t_x
			y = t_y
		case '#':
			d_x, d_y = turn_right(d_x, d_y)
		}
	}

	fmt.Println(counter)
}
