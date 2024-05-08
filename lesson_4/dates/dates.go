package dates

const DaysInWeek = 7

func weekToDays(weeks int) int {
	return weeks * DaysInWeek
}

func daysToWeeks(days int) float64 {
	return float64(days) / DaysInWeek
}
