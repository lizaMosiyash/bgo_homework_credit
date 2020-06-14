package credit
import "math"

func Calculate(sumOfCredit, countOfMonth, percentOfRate float64) (monthlyPayment, totalSumOfCredit, sumOfPercent int) {
	monthlyPercent := percentOfRate/12/100
	partOfAnnCoef := math.Pow((1+monthlyPercent), countOfMonth)
	AnnCoef := monthlyPercent*partOfAnnCoef/(partOfAnnCoef-1)
	monthlyPayment = int(AnnCoef*sumOfCredit)
	totalSumOfCredit = monthlyPayment * int(countOfMonth)
	sumOfPercent = totalSumOfCredit - int(sumOfCredit)
	return
}
