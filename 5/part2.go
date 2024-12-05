package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func is_correct(prints []int, rules map[int][]int) bool {
	seen := make(map[int]struct{})

	for i := len(prints) - 1; i >= 0; i-- {
		value := prints[i]
		if candidates, ok := rules[value]; ok {
			for _, candidate := range candidates {
				if _, ok := seen[candidate]; ok {
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
		if !is_correct(print, rules) {
			Sort(print, rules)
			middle := print[len(print)/2]
			sum += middle
		}
	}
	fmt.Println(sum)
}

func contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func rules_contains(rules map[int][]int, index int, value int) bool {
	rule, ok := rules[index]
	return ok && contains(rule, value)
}

func Sort(print []int, rules map[int][]int) {
	sort.Slice(print, func(i, j int) bool {
		a := print[i]
		b := print[j]
		if rules_contains(rules, a, b) {
			return true
		} else if rules_contains(rules, b, a) {
			return false
		} else {
			return a < b
		}
	})
}
