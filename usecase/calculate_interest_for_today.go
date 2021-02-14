package usecase

import "context"

type (
	GetTodaysAmount interface {
		Execute(context.Context, GetTodaysAmountInput) (GetTodaysAmountOutput, error)
	}

	GetTodaysAmountInput struct {
	}

	GetTodaysAmountOutput struct {
	}
)
