package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Job struct {
	weight int
	length int
}

// This file describes a set of jobs with positive and integral weights and lengths.  It has the format
//
// [number_of_jobs]
//
// [job_1_weight] [job_1_length]
//
// [job_2_weight] [job_2_length]
//
// ...
//
// For example, the third line of the file is "74 59", indicating that the second job has weight 74 and length 59.
//
// You should NOT assume that edge weights or lengths are distinct.
//
// Your task in this problem is to run the greedy algorithm that schedules jobs in decreasing order of the difference (weight - length).  Recall from lecture that this algorithm is not always optimal.  IMPORTANT: if two jobs have equal difference (weight - length), you should schedule the job with higher weight first.  Beware: if you break ties in a different way, you are likely to get the wrong answer.  You should report the sum of weighted completion times of the resulting schedule --- a positive integer --- in the box below.
//
// ADVICE: If you get the wrong answer, try out some small test cases to debug your algorithm (and post your test cases to the discussion forum).
func weightedCompletionTime(jobs *[]Job) int {
	sort.Slice(*jobs, func(i int, j int) bool {
		c1 := (*jobs)[i].weight - (*jobs)[i].length
		c2 := (*jobs)[j].weight - (*jobs)[j].length
		if c1 != c2 {
			return c1 > c2
		}

		return (*jobs)[i].weight > (*jobs)[j].weight
	})

	weightedSum := 0
	processedLength := 0
	for _, j := range *jobs {
		processedLength += j.length
		weightedSum += j.weight * processedLength
	}

	return weightedSum
}

func weightedCompletionTime2(jobs *[]Job) int {
	sort.Slice(*jobs, func(i int, j int) bool {
		c1 := float64((*jobs)[i].weight) / float64((*jobs)[i].length)
		c2 := float64((*jobs)[j].weight) / float64((*jobs)[j].length)
		if c1 != c2 {
			return c1 > c2
		}

		return (*jobs)[i].weight > (*jobs)[j].weight
	})

	weightedSum := 0
	processedLength := 0
	for _, j := range *jobs {
		processedLength += j.length
		weightedSum += j.weight * processedLength
	}

	return weightedSum
}

func main() {
	data, err := os.ReadFile("jobs.txt")
	if err != nil {
		return
	}

	d := strings.Split(strings.TrimSpace(string(data)), "\n")

	totalLength, _ := strconv.Atoi(strings.TrimSpace(d[0]))
	jobs := make([]Job, 0, totalLength)
	for _, a := range d[1:] {
		dd := strings.Fields(a)
		w, _ := strconv.Atoi(strings.TrimSpace(dd[0]))
		l, _ := strconv.Atoi(strings.TrimSpace(dd[1]))
		jobs = append(jobs, Job{w, l})
	}

	fmt.Println(weightedCompletionTime(&jobs))
	fmt.Printf("By Ratio w/d %d", weightedCompletionTime2(&jobs))
}
