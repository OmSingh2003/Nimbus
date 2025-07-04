package util 

const (
	USD = "USD"
	INR = "INR"
	EUR = "EUR"
)

// IsSupportedCurrency checks if the given currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD,INR,EUR :
	return true 
	}
	return false
}

// GetWelcomeCreditAmount returns the welcome credit amount in cents for different currencies
// Base amount is $100 USD
func GetWelcomeCreditAmount(currency string) int64 {
	switch currency {
	case USD:
		return 10000 // $100.00 in cents
	case EUR:
		return 9500  // €95.00 in cents (approximate exchange rate)
	case INR:
		return 830000 // ₹8300.00 in paisa (approximate exchange rate)
	default:
		return 10000 // Default to USD equivalent
	}
}
