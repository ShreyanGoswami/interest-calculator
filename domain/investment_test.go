package domain

import (
	"testing"
	"time"
)

func TestInvestment_CalculateTodaysInterestAmount(t *testing.T) {

	tests := []struct {
		name       string
		investment Investment
		expected   float64
	}{
		{
			name:       "Successfully calculate interest earned today",
			investment: NewInvestment(1.0, 10.0, 1.0),
			expected:   10.1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.investment.CalculateTodaysValue()
			if actual != tc.expected {
				t.Errorf("[TestCase '%s'] Result: '%v' | Expected: '%v'",
					tc.name,
					actual,
					tc.expected)
			}
		})
	}
}

func TestInvestment_GenerateReport(t *testing.T) {
	tests := []struct {
		name       string
		investment Investment
		expected   []InvestmentReport
		args       time.Time
	}{
		{
			name: "Succesfully generate two days of report",
			investment: Investment{
				startDate:           time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
				principalAmount:     1000,
				currentAmount:       1000,
				interestRatePerDay:  1.0,
				interestRateOverall: 100,
			},
			args: time.Date(2021, time.January, 3, 0, 0, 0, 0, time.UTC),
			expected: []InvestmentReport{
				{
					amount: 1000,
					date:   time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					amount: 1000,
					date:   time.Date(2021, time.January, 2, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			name: "Successfully generate report for two days at the boundary of two months",
			investment: Investment{
				startDate:           time.Date(2021, time.January, 31, 0, 0, 0, 0, time.UTC),
				principalAmount:     1000,
				currentAmount:       1000,
				interestRatePerDay:  1.0,
				interestRateOverall: 100,
			},
			args: time.Date(2021, time.February, 2, 0, 0, 0, 0, time.UTC),
			expected: []InvestmentReport{
				{
					amount: 1000,
					date:   time.Date(2021, time.January, 31, 0, 0, 0, 0, time.UTC),
				},
				{
					amount: 1000,
					date:   time.Date(2021, time.February, 1, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			name: "Successfully generate report for two days at the boundary of two years",
			investment: Investment{
				startDate:           time.Date(2021, time.December, 31, 0, 0, 0, 0, time.UTC),
				principalAmount:     1000,
				currentAmount:       1000,
				interestRatePerDay:  1.0,
				interestRateOverall: 100,
			},
			args: time.Date(2022, time.January, 2, 0, 0, 0, 0, time.UTC),
			expected: []InvestmentReport{
				{
					amount: 1000,
					date:   time.Date(2021, time.January, 31, 0, 0, 0, 0, time.UTC),
				},
				{
					amount: 1000,
					date:   time.Date(2021, time.February, 1, 0, 0, 0, 0, time.UTC),
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, _ := tc.investment.GenerateHistory(tc.args)
			if compareInvestmentReports(actual, tc.expected) {
				t.Errorf("[TestCase '%s'] Result: '%v' | Expected: '%v'",
					tc.name,
					actual,
					tc.expected)
			}
		})
	}
}

func compareInvestmentReports(report1, report2 []InvestmentReport) bool {
	if len(report1) != len(report2) {
		return false
	}

	for _, itemInReport1 := range report1 {
		isItemFoundInReport2 := false
		for _, itemInreport2 := range report2 {
			if itemInreport2.amount == itemInReport1.amount && compareDates(itemInReport1.date, itemInreport2.date) {
				isItemFoundInReport2 = true
			}
		}
		if isItemFoundInReport2 == false {
			return false
		}
	}
	return true
}

func compareDates(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
