package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// testing there
func main() {
	databytes, err := os.ReadFile("2sum.txt")
	if err != nil {
		fmt.Println("Issue reading  2sum.txt file")
		return
	}
	numbersArr := []int{}

	data := strings.Split(string(databytes), "\n")

	hm := make(map[int]int)
	for _, d := range data {
		if strings.TrimSpace(d) == "" {
			continue
		}
		number, err := strconv.Atoi(strings.TrimSpace(d))
		if err != nil {
			fmt.Println("failed to parse number;", d)
			continue
		}
		numbersArr = append(numbersArr, number)
		hm[number]++
	}

	sort.Ints(numbersArr)

	found := make(map[int]bool)
	for _, x := range numbersArr {
		lo := sort.SearchInts(numbersArr, -10000-x)
		hi := sort.SearchInts(numbersArr, 10000-x+1)
		for _, y := range numbersArr[lo:hi] {
			if x == y && hm[x] < 2 {
				continue
			}
			found[x+y] = true
		}

	}
	fmt.Println("Total count: ", len(found))
}
