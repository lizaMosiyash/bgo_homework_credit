package credit
import "math"

func Calculate(sumOfCredit, countOfMonth, percentOfRate float64) (monthlyPayment1 int,totalSumOfCredit int, sumOfPercent int) {
	monthlyPercent := percentOfRate/12/100
	partOfAnnCoef := math.Pow((1+monthlyPercent), countOfMonth)
	AnnCoef := monthlyPercent*partOfAnnCoef/(partOfAnnCoef-1)
	monthlyPayment := AnnCoef*sumOfCredit
	monthlyPayment1 = int(monthlyPayment)
	totalSumOfCredit = int(monthlyPayment) * int(countOfMonth)
	sumOfPercent = totalSumOfCredit - int(sumOfCredit)
	return monthlyPayment1, sumOfPercent, totalSumOfCredit
}
