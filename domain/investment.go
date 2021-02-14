package domain

import "time"

type (
	InvestmentRepository interface {
		CalculateTodaysValue() (float64, error)
		GenerateHistory(time.Time) ([]InvestmentReport, error)
	}

	InvestmentID string

	Investment struct {
		InvestmentID
		currentAmount       float64
		interestRatePerDay  float64
		startDate           time.Time
		endDate             time.Time
		duration            time.Time
		principalAmount     float64
		interestRateOverall float64
	}

	InvestmentReport struct {
		amount float64
		date   time.Time
	}
)

func (investment *Investment) calculateInterest() float64 {
	return investment.currentAmount * (investment.interestRatePerDay / 100)
}

func (investment Investment) CalculateTodaysValue() (float64, error) {
	return investment.currentAmount + investment.calculateInterest(), nil
}

func (investment Investment) GenerateHistory(endDateForReport time.Time) ([]InvestmentReport, error) {
	investmentReport := []InvestmentReport{}
	numberOfDays := endDateForReport.Sub(investment.startDate).Hours() / 24
	prevDate := investment.startDate
	currAmount := investment.principalAmount
	for i := 0; i < int(numberOfDays); i++ {
		currVal := InvestmentReport{}
		currVal.date = prevDate
		currVal.amount = currAmount
		currAmount = calculateAmountAfterInterest(currAmount, investment.interestRatePerDay)
		investmentReport = append(investmentReport, currVal)
		prevDate = prevDate.Add(time.Hour * time.Duration(24))
	}
	return investmentReport, nil
}

func calculateAmountAfterInterest(amount, interest float64) float64 {
	return amount * (1 + interest/100)
}

func NewInvestment(initialAmount, currAmount, interestPerDay float64) Investment {
	return Investment{
		principalAmount:    initialAmount,
		currentAmount:      currAmount,
		interestRatePerDay: interestPerDay,
	}
}
