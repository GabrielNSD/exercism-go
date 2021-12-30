package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var total int = 0
	for i := 0; i < len(birdsPerDay); i++ {
		total = total + birdsPerDay[i]
	}
	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var initIndex int = (week - 1) * 7
	var birdsCount int = 0
	for i := initIndex; i < initIndex+7; i++ {
		birdsCount = birdsCount + birdsPerDay[i]
	}
	return birdsCount
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i = i + 2 {
		birdsPerDay[i] = birdsPerDay[i] + 1
	}
	return birdsPerDay
}
