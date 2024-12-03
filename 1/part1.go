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

	fmt.Println(left)

	var sum = 0
	for i := range left {
		tmp := left[i] - right[i]
		if tmp > 0 {
			sum += tmp
		} else {
			sum -= tmp
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(sum)
}
