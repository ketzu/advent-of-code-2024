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

	r, _ := regexp.Compile(`(?:(mul)\((\d+),(\d+)\)|(do)\(\)|(don't)\(\))`)

	sum := 0
	do := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		matches := r.FindAllStringSubmatch(scanner.Text(), -1)
		for _, v := range matches {
			switch len(v[1]) + len(v[4]) + len(v[5]) {
			case 2: // do
				do = true
			case 5: // don't
				do = false
			case 3: // mul
				if do {
					o1, _ := strconv.Atoi(v[2])
					o2, _ := strconv.Atoi(v[3])
					sum += o1 * o2
				}
			}
		}
	}
	fmt.Println(sum)
}
