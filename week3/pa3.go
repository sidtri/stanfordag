package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var comparisions = 0

func main() {
	data, err := os.ReadFile("qsdata.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	nums := make([]int, 0, len(lines))
	for _, line := range lines {
		n, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			fmt.Println("Error", err)
			return
		}
		nums = append(nums, n)
	}

	qs(nums, 0, len(nums)-1)
	fmt.Println(nums)
	fmt.Println("Comparisions", comparisions)
}

func qs(nums []int, lb int, rb int) {
	if lb >= rb {
		return
	}

	mid := partition(nums, lb, rb)

	qs(nums, lb, mid-1)
	qs(nums, mid+1, rb)
}

func partition(nums []int, lb int, rb int) (mid int) {
	pivot := nums[lb]
	i := lb + 1

	for j := lb + 1; j <= rb; j++ {
		comparisions += 1
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i += 1
		}
	}

	nums[lb], nums[i-1] = nums[i-1], nums[lb]
	return i - 1
}

func partitionRight(nums []int, lb int, rb int) (mid int) {
	pi := rb
	pivot := nums[pi]

	rb -= 1

	for lb <= rb {
		comparisions += 1
		if nums[lb] <= pivot {
			lb += 1
			continue
		}

		nums[lb], nums[rb] = nums[rb], nums[lb]
		rb -= 1
	}

	rb += 1
	nums[pi], nums[rb] = nums[rb], nums[pi]
	return rb
}
