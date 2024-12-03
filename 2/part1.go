package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	safe_levels := 0
	lines := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines += 1
		split := strings.Split(scanner.Text(), " ")
		level := []int{}
		for _, v := range split {
			i, _ := strconv.Atoi(v)
			level = append(level, i)
		}

		safety := safe(level)
		if safety {
			safe_levels += 1
		}
		fmt.Println(lines, safety, level)
	}

	fmt.Println(safe_levels, "from", lines)

}

func safe(level []int) bool {
	if len(level) < 2 {
		return true
	}
	first_diff := (level[0] - level[1])
	if first_diff == 0 {
		return false
	}

	should_sign := 1
	if first_diff < 0 {
		should_sign = -1
	}

	for i := 0; i+1 < len(level); i += 1 {
		v := level[i] - level[i+1]
		try := v * should_sign
		if try > 3 || try <= 0 {
			return false
		}
	}
	return true
}
