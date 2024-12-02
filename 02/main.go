package main

import (
	"strconv"
	"fmt"
	"os"
	"bufio"
	"strings"
	"math"
)

func Part1(path string) int {
	input, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	var safe int = 0
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		t := strings.Split(scanner.Text(), " ")
		nums := make([]int, len(t))
		for i, l := range t {
			nums[i], err = strconv.Atoi(l)
			if err != nil {
				panic(err)
			}
		}
		if Analyze(nums) {
			safe++
		}
	}
	return safe
}

func Analyze(nums []int) bool {
	prevSign := 0
	for i := range nums[1:] {
		var sign int
		if sign = 1; (nums[i] - nums[i + 1]) > 0 {
			sign = -1
		}
		d := int(math.Abs(float64(nums[i] - nums[i + 1])))
		if d < 1 || d > 3 || (prevSign != 0 && sign != prevSign) {
			return false
		}
		prevSign = sign
	}
	return true
}

func Part2(path string) int {
	input, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	var safe int = 0
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		t := strings.Split(scanner.Text(), " ")
		nums := make([]int, len(t))
		for i, l := range t {
			nums[i], err = strconv.Atoi(l)
			if err != nil {
				panic(err)
			}
		}
		if Analyze(nums) {
			safe++
		} else {
			for i := range nums {
				l := make([]int, 0)
				for j := range nums {
					if j != i {
						l = append(l, nums[j])
					}
				}
				if Analyze(l) {
					safe++
					break
				}
			}
		}
	}
	return safe
}

func main() {
	if (len(os.Args[1:]) != 2) {
		panic("Usage: ./main [1|2] file")
	}
	var answer int
	switch (os.Args[1]) {
	default:
	case "1":
		answer = Part1(os.Args[2])
	case "2":
		answer = Part2(os.Args[2])
	}
	fmt.Println("Total Safe Reports:", answer)
}
