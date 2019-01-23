package common

const (
    Coin       int64 = 1e8
	Amount100  int64 = 100 * Coin
	Amount1000 int64 = 1000 * Coin
	MaxSendCount     = 100
)

func DivisionAmount(account int64) []int64 {
	count1000 := account/Amount1000
	count100 := (account - count1000 * Amount1000)/Amount100
	lastAmount :=  account - count1000 * Amount1000 - count100 * Amount100
	var amounts []int64
	for i := 0; i < int(count1000); i++  {
		amounts = append(amounts, Amount1000)
	}
	for i := 0; i < int(count100); i++  {
		amounts = append(amounts, Amount100)
	}
	if lastAmount > 0 {
		amounts = append(amounts, lastAmount)
	}
	return amounts
}
