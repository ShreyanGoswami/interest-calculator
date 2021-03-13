package repository

import (
	"context"

	"github.com/ShreyanGoswami/interest-calculator/domain"
	"github.com/ShreyanGoswami/interest-calculator/infra/database"
)

type InMemoryInvestmentTableDB struct {
	tableName string
	db        *database.InMemoryMemDB
}

func NewInMemoryInvestmentTableDB(db *database.InMemoryMemDB) InMemoryInvestmentTableDB {
	return InMemoryInvestmentTableDB{
		tableName: "investment",
		db:        db,
	}
}

func (a InMemoryInvestmentTableDB) AddInvestment(context context.Context, investment domain.Investment) (domain.Investment, error) {
	err := a.db.InsertData(a.tableName, investment)
	if err != nil {
		panic(err)
	}
	return investment, nil
}

func (a InMemoryInvestmentTableDB) FindInvestmentById(context context.Context, id domain.InvestmentID) (domain.Investment, error) {
	return a.db.FindOne(a.tableName, id)
}

func (a InMemoryInvestmentTableDB) CloseInvestment(context context.Context, id domain.InvestmentID) (domain.Investment, error) {
	// TODO needs to be moved to mem_db
	// investment, errInFindInvestment := a.FindInvestmentById(context, id)
	// if errInFindInvestment != nil {
	// 	panic(errInFindInvestment)
	// }
	// investment.IsClosed = true
	// updatedInvestment, errInUpdate := a.AddInvestment(context, investment)
	// if errInUpdate != nil {
	// 	panic(errInUpdate)
	// }
	// return updatedInvestment, nil
	return domain.Investment{}, nil
}

func (a InMemoryInvestmentTableDB) FillInvestmentInMemoryDatabase(investments []domain.Investment) {
	// TODO needs to be moved to mem_db
	// txn := a.db.Txn(true)
	// for _, investment := range investments {
	// 	if err := txn.Insert(a.tableName, investment); err != nil {
	// 		panic(err)
	// 	}
	// }
	// txn.Commit()
}
