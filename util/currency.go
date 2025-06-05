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
