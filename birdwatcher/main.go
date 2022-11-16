package main

import (
	"birdwatcher/birdwatcher"
	"fmt"
)

func main() {
	weeks := []int{2, 3, 15, 6, 7, 12, 5, 13, 14, 6, 7, 8, 8, 10}
	total_bird_count := birdwatcher.TotalBirdCount([]int{3, 4, 5})
	birds_in_week := birdwatcher.BirdsInWeek(weeks, 2)
	fixed_count := birdwatcher.FixBirdCountLog(weeks)

	fmt.Println(total_bird_count)
	fmt.Println(birds_in_week)
	fmt.Println(fixed_count)
}
