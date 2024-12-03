// (mul\([0-9]+,[0-9]+\))
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Part1(path string, numbers *regexp.Regexp) int {
	input, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	instructions, err := regexp.Compile("mul\\([0-9]+,[0-9]+\\)")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(input)
	var total int = 0
	for scanner.Scan() {
		is := instructions.FindAllString(scanner.Text(), len(scanner.Text()))
		for _, i := range is {
			nums := numbers.FindAllString(i, len(i))
			n1, err := strconv.Atoi(nums[0])
			if err != nil {
				panic(err)
			}
			n2, err := strconv.Atoi(nums[1])
			if err != nil {
				panic(err)
			}
			total += n1 * n2
		}
	}
	return total
}

func Part2(path string, numbers *regexp.Regexp) int {
	instructions, err := regexp.Compile("do\\(\\)|don't\\(\\)|mul\\([0-9]+,[0-9]+\\)")
	if err != nil {
		panic(err)
	}
	input, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)
	var total int = 0
	var enabled bool = true
	for scanner.Scan() {
		is := instructions.FindAllString(scanner.Text(), len(scanner.Text()))
		for _, i := range is {
			if i == "do()" {
				enabled = true
				continue
			} else if i == "don't()" {
				enabled = false
				continue
			} else if !enabled {
				continue
			}
			nums := numbers.FindAllString(i, len(i))
			n1, err := strconv.Atoi(nums[0])
			if err != nil {
				panic(err)
			}
			n2, err := strconv.Atoi(nums[1])
			if err != nil {
				panic(err)
			}
			total += n1 * n2
		}
	}
	return total
}

func main() {
	if len(os.Args[1:]) != 2 {
		panic("Usage: ./main [1|2] file")
	}
	numbers, err := regexp.Compile("[0-9]+")
	if err != nil {
		panic(err)
	}
	var total int
	switch os.Args[1] {
	default:
	case "1":
		total = Part1(os.Args[2], numbers)
	case "2":
		total = Part2(os.Args[2], numbers)
	}
	fmt.Println("Total:", total)
}
