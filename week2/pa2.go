package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var swapCount = 0

func merge(lines []int) []int {
	if len(lines) <= 1 {
		return lines
	}

	mid := len(lines) / 2
	left := merge(lines[:mid])
	right := merge(lines[mid:])

	result := make([]int, 0, len(lines))
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			swapCount += len(left) - i
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

func main() {
	// read file
	data, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	nums := make([]int, 0, len(lines))

	for _, line := range lines {
		line = strings.TrimSpace(line)
		n, _ := strconv.Atoi(line)
		nums = append(nums, n)
	}

	sorted := merge(nums)
	for _, line := range sorted {
		fmt.Println(line)
	}

	fmt.Printf("Swap count: %d\n", swapCount)
}
