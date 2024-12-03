package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	r, _ := regexp.Compile(`mul\((\d+),(\d+)\)`)

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matches := r.FindAllStringSubmatch(scanner.Text(), -1)
		for _, v := range matches {
			o1, _ := strconv.Atoi(v[1])
			o2, _ := strconv.Atoi(v[2])
			sum += o1 * o2
			fmt.Println(o1, o2, o1*o2, sum)
		}
	}
	fmt.Println(sum)
}
