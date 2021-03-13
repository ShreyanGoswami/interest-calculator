package action

import (
	"net/http"

	"github.com/ShreyanGoswami/interest-calculator/infra/logger"
	"github.com/ShreyanGoswami/interest-calculator/usecase"
)

type CalculateTodaysAmountAction struct {
	uc  usecase.GetTodaysAmount
	log logger.Logger
}

func NewCalculateTodaysAmountAction(uc usecase.GetTodaysAmount, log logger.Logger) CalculateTodaysAmountAction {
	return CalculateTodaysAmountAction{
		uc:  uc,
		log: log,
	}
}

func (a CalculateTodaysAmountAction) Execute(res http.ResponseWriter, req *http.Request) {
	inputInvestmentId := req.URL.Query().Get("investment_id")
	input := usecase.NewGetTodaysAmountInput(inputInvestmentId)
	a.uc.Execute(req.Context(), input)
}
