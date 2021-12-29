package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	var result float32
	switch {
	case balance < 0:
		result = 3.213
	case balance >= 0 && balance < 1000:
		result = 0.5
	case balance >= 1000 && balance < 5000:
		result = 1.621
	case balance >= 5000:
		result = 2.475
	}
	return result
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance)/100)
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var years int = 0
	var resultBalance = balance

	for i := 0; resultBalance < targetBalance; i++ {
		resultBalance = AnnualBalanceUpdate(resultBalance)
		years = i
	}
	return years + 1
}
