package main

import (
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type minHeap []int

func (s minHeap) Len() int {
	return len(s)
}

func (s minHeap) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s minHeap) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *minHeap) Push(x any) {
	*s = append(*s, x.(int))
}

func (s *minHeap) Pop() any {
	old := *s
	length := len(old)
	popped := old[length-1]
	*s = old[0 : length-1]
	return popped
}

func (s *minHeap) Display() int {
	return (*s)[0]
}

type maxHeap []int

func (s maxHeap) Len() int {
	return len(s)
}

func (s maxHeap) Empty() bool {
	return s.Len() == 0
}

func (s maxHeap) Less(i, j int) bool {
	return s[i] > s[j]
}

func (s maxHeap) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *maxHeap) Push(x any) {
	*s = append(*s, x.(int))
}

func (s *maxHeap) Pop() any {
	old := *s
	length := len(old)
	popped := old[length-1]
	*s = old[0 : length-1]
	return popped
}

func (s *maxHeap) Display() int {
	return (*s)[0]
}

func median(numbers <-chan int) int {
	min := minHeap{}
	max := maxHeap{}
	heap.Init(&min)
	heap.Init(&max)
	var medianScore int = 0
	for num := range numbers {
		if max.Empty() || max.Display() >= num {
			heap.Push(&max, num)
		} else {
			heap.Push(&min, num)
		}

		// rebalance
		if max.Len() > (min.Len() + 1) {
			data := heap.Pop(&max).(int)
			heap.Push(&min, data)
		}

		if min.Len() > max.Len() {
			data := heap.Pop(&min).(int)
			heap.Push(&max, data)
		}

		// calculate scores
		medianScore += max.Display()
	}

	return medianScore
}

func main() {
	data, err := os.ReadFile("median.txt")
	if err != nil {
		fmt.Println("error happen, ", err)
		return
	}

	d := strings.Split(strings.TrimSpace(string(data)), "\n")

	numbers := make(chan int)
	result := make(chan int)
	go func() {
		result <- median(numbers)
	}()
	for _, strnum := range d {
		val, err := strconv.Atoi(strings.TrimSpace(strnum))
		if err != nil {
			fmt.Println("Some issues with parsing integers")
			continue
		}
		numbers <- val
	}
	close(numbers)

	fmt.Println(<-result % 10000)
}
