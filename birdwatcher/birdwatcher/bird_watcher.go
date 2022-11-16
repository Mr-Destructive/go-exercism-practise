package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var count int
	for _, c := range birdsPerDay {
		count += c
	}
	return count
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	start := (week - 1) * 7
	var count int
	for i := start; i < start+7; i++ {
		count += birdsPerDay[i]
	}
	return count
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	correct_count := birdsPerDay
	length := len(birdsPerDay)
	for i := 0; i < length-1; i += 2 {
		correct_count[i] += 1
	}
	return correct_count
}
