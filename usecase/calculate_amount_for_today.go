package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/ShreyanGoswami/interest-calculator/domain"
)

type (
	GetTodaysAmount interface {
		Execute(context.Context, GetTodaysAmountInput) (GetTodaysAmountOutput, error)
	}

	GetTodaysAmountInput struct {
		domain.InvestmentID
	}

	GetTodaysAmountOutput struct {
		Amount float64
	}

	GetTodaysAmountPresenter interface {
		Output(float64) GetTodaysAmountOutput
	}

	GetTodaysAmountInteractor struct {
		repo       domain.InvestmentRepository
		presenter  GetTodaysAmountPresenter
		ctxTimeout time.Duration
	}
)

func NewGetTodaysAmountInteractor(
	investment domain.InvestmentRepository,
	presenter GetTodaysAmountPresenter,
	t time.Duration,
) GetTodaysAmount {
	return GetTodaysAmountInteractor{
		repo:       investment,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

func (interactor GetTodaysAmountInteractor) Execute(ctx context.Context, input GetTodaysAmountInput) (GetTodaysAmountOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, interactor.ctxTimeout)
	defer cancel()

	investment, err := interactor.repo.FindInvestmentById(ctx, input.InvestmentID)

	if err != nil {
		fmt.Printf("Error while getting investment")
	}

	return interactor.presenter.Output(investment.CalculateTodaysValue()), nil
}
