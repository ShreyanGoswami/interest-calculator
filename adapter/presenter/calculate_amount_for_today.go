package presenter

import "github.com/ShreyanGoswami/interest-calculator/usecase"

type calculateAmountForTodayPresenter struct{}

func NewCalculateAmountForTodayPresenter() usecase.GetTodaysAmountPresenter {
	return calculateAmountForTodayPresenter{}
}

func (a calculateAmountForTodayPresenter) Output(amount float64) usecase.GetTodaysAmountOutput {
	return usecase.GetTodaysAmountOutput{Amount: amount}
}
