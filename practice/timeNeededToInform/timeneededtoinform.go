// /1376. Time Needed to Inform All Employees
// Medium
// Topics
// premium lock iconCompanies
// Hint
//
// A company has n employees with a unique ID for each employee from 0 to n - 1. The head of the company is the one with headID.
//
// Each employee has one direct manager given in the manager array where manager[i] is the direct manager of the i-th employee, manager[headID] = -1. Also, it is guaranteed that the subordination relationships have a tree structure.
//
// The head of the company wants to inform all the company employees of an urgent piece of news. He will inform his direct subordinates, and they will inform their subordinates, and so on until all employees know about the urgent news.
//
// The i-th employee needs informTime[i] minutes to inform all of his direct subordinates (i.e., After informTime[i] minutes, all his direct subordinates can start spreading the news).
//
// Return the number of minutes needed to inform all the employees about the urgent news.
//
// Example 1:
//
// Input: n = 1, headID = 0, manager = [-1], informTime = [0]
// Output: 0
// Explanation: The head of the company is the only employee in the company.
package main

import "fmt"

func max(nums []int) int {
	maxVal := nums[0]
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	nm := make([]int, n)

	for emp, mgr := range manager {
		for mgr != -1 {
			nm[emp] = nm[emp] + informTime[mgr]

			emp = mgr
			mgr = manager[mgr]
		}
	}

	return max(nm)
}

func main() {
	output := numOfMinutes(6, 2, []int{2, 2, -1, 2, 2, 2}, []int{0, 0, 1, 0, 0, 0})
	fmt.Println(output)
}
