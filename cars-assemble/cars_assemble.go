package cars

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {
	//panic("SuccessRate not implemented")
	var successRate float64 = 0
	if speed == 0 {
		successRate = 0.0
	}
	if speed > 0 {
		successRate = 1
	}
	if speed > 4 {
		successRate = 0.9
	}
	if speed > 8 {
		successRate = 0.77
	}

	return successRate
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	//panic("CalculateProductionRatePerHour not implemented")
	var productionRate float64 = float64(221 * speed)
	return productionRate * SuccessRate(speed)
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	//panic("CalculateProductionRatePerMinute not implemented")
	return int(CalculateProductionRatePerHour(speed)) / 60
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	//panic("CalculateLimitedProductionRatePerHour not implemented")
	var rate float64 = CalculateProductionRatePerHour(speed)
	if rate > limit {
		return limit
	}
	return rate
}
