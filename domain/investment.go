package domain

import (
	"context"
	"time"
)

type (
	InvestmentRepository interface {
		AddInvestment(context.Context, Investment) (Investment, error)
		UpdateInvestment(context.Context, Investment) (Investment, error)
		FindInvestmentById(context.Context, InvestmentID) (Investment, error)
		CloseInvestment(context.Context, InvestmentID) (Investment, error)
	}

	InvestmentID string

	Investment struct {
		InvestmentID
		CurrentAmount       float64
		InterestRatePerDay  float64
		StartDate           time.Time
		EndDate             time.Time
		Duration            time.Time
		PrincipalAmount     float64
		InterestRateOverall float64
	}

	InvestmentReport struct {
		amount float64
		date   time.Time
	}
)

func (investment *Investment) calculateInterest() float64 {
	return investment.CurrentAmount * (investment.InterestRatePerDay / 100)
}

func (investment Investment) CalculateTodaysValue() float64 {
	return investment.CurrentAmount + investment.calculateInterest()
}

func (investment Investment) GenerateHistory(endDateForReport time.Time) ([]InvestmentReport, error) {
	investmentReport := []InvestmentReport{}
	numberOfDays := endDateForReport.Sub(investment.StartDate).Hours() / 24
	prevDate := investment.StartDate
	currAmount := investment.PrincipalAmount
	for i := 0; i < int(numberOfDays); i++ {
		currVal := InvestmentReport{}
		currVal.date = prevDate
		currVal.amount = currAmount
		currAmount = calculateAmountAfterInterest(currAmount, investment.InterestRatePerDay)
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
		PrincipalAmount:    initialAmount,
		CurrentAmount:      currAmount,
		InterestRatePerDay: interestPerDay,
	}
}
