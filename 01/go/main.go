package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ReadInts(path string) ([]int, []int) {
	input, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	left := make([]int, 0)
	right := make([]int, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), "   ")
		num1, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}
		num2, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}
		left = append(left, num1)
		right = append(right, num2)
	}
	sort.Slice(left, func(a, b int) bool {
		return left[a] < left[b]
	})
	sort.Slice(right, func(a, b int) bool {
		return right[a] < right[b]
	})
	return left, right
}

func Part1(path string) int {
	left, right := ReadInts(path)
	if len(left) != len(right) {
		panic("Different Length Slices")
	}
	var total int = 0
	for i, l := range left {
		total += int(math.Abs(float64(l - right[i])))
	}
	return total
}

func Part2(path string) int {
	left, right := ReadInts(path)
	if len(left) != len(right) {
		panic("Different Length Slices")
	}
	totals := make([]int, right[len(right)-1]+1)
	for _, r := range right {
		totals[r] += 1
	}
	total := 0
	for _, l := range left {
		total += l * totals[l]
	}
	return total
}

func main() {
	if len(os.Args) != 3 {
		panic("Usage: ./main [1|2] file")
	}
	switch os.Args[1] {
	default:
	case "1":
		fmt.Println("Total Distance:", Part1(os.Args[2]))
	case "2":
		fmt.Println("Similarity Score:", Part2(os.Args[2]))
	}
}
