package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	left := []int{}
	right := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		left_entry, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}
		left = append(left, left_entry)

		right_entry, err := strconv.Atoi(split[len(split)-1])
		if err != nil {
			panic(err)
		}
		right = append(right, right_entry)
	}

	sort.Ints(left)
	sort.Ints(right)

	numbers := make(map[int]int)

	for _, r := range right {
		v, p := numbers[r]
		if p {
			numbers[r] = v + 1
		} else {
			numbers[r] = 1
		}
	}

	var sum = 0
	for _, l := range left {
		v, p := numbers[l]
		if p {
			sum += l * v
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(sum)
}
