package main

import (
	"fmt"
	"math/rand"
	"slices"
)

func main() {
	birthdayRandoms := []int{}
	sameBirthday := false
	for range 23 {
		birthday := rand.Intn(365)
		if slices.Contains(birthdayRandoms, birthday) {
			sameBirthday = true
		}
		birthdayRandoms = append(birthdayRandoms, birthday)
	}

	fmt.Println("Same birthday exists", sameBirthday)
}
