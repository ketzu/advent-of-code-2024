package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var verbose = false

func is_correct(prints []int, rules map[int][]int) bool {
	seen := make(map[int]struct{})

	if verbose {
		fmt.Println(prints)
		fmt.Println(rules)
	}

	for i := len(prints) - 1; i >= 0; i-- {
		value := prints[i]
		if verbose {
			fmt.Println("checking", value, "seen", seen)
		}
		if candidates, ok := rules[value]; ok {
			for _, candidate := range candidates {
				if _, ok := seen[candidate]; ok {
					if verbose {
						fmt.Println("seen", candidate, "violating", value, candidates)
					}
					return false
				}
			}
		}
		seen[value] = struct{}{}
	}
	return true
}

func add_rule(rules map[int][]int, a int, b int) {
	if _, ok := rules[b]; !ok {
		rules[b] = []int{}
	}
	rules[b] = append(rules[b], a)
}

func read_rules(filename string) map[int][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	rules := make(map[int][]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "|")
		a, _ := strconv.Atoi(split[0])
		b, _ := strconv.Atoi(split[1])
		add_rule(rules, a, b)
	}
	return rules
}

func read_prints(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	prints := [][]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ",")
		print := []int{}
		for _, s := range split {
			i, _ := strconv.Atoi(s)
			print = append(print, i)
		}
		prints = append(prints, print)
	}
	return prints
}

func main() {
	rules := read_rules("rules.txt")
	prints := read_prints("prints.txt")

	sum := 0
	for _, print := range prints {
		if is_correct(print, rules) {
			middle := print[len(print)/2]
			if verbose {
				fmt.Println(print, "passed", middle)
			}
			sum += middle
		} else {
			if verbose {
				fmt.Println(print, "failed")
			}
		}
	}
	fmt.Println(sum)
}
