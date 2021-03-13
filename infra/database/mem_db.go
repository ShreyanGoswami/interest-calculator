package database

import (
	"context"

	"github.com/ShreyanGoswami/interest-calculator/domain"
	"github.com/hashicorp/go-memdb"
)

type InMemoryInvestmentRepo struct {
	db                  *memdb.MemDB
	investmentTableName string
}

const idFieldName = "id" // TODO this needs to go somewhere else because it is common information for all implementations of the database

func NewInMemoryInvestmentRepo(schema *memdb.DBSchema) (*InMemoryInvestmentRepo, error) {
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}
	tableName := "investment"
	return &InMemoryInvestmentRepo{db, tableName}, nil
}

func (a InMemoryInvestmentRepo) AddInvestment(context context.Context, investment domain.Investment) (domain.Investment, error) {
	txn := a.db.Txn(true)
	defer txn.Abort()
	if err := txn.Insert(a.investmentTableName, investment); err != nil {
		panic(err)
	}
	txn.Commit()
	return investment, nil
}

func (a InMemoryInvestmentRepo) FindInvestmentById(context context.Context, id domain.InvestmentID) (domain.Investment, error) {
	txn := a.db.Txn(false)
	defer txn.Abort()
	queryID := string(id)
	raw, err := txn.First(a.investmentTableName, idFieldName, queryID)
	if err != nil {
		return domain.Investment{}, err
	}
	return raw.(domain.Investment), nil
}

func (a InMemoryInvestmentRepo) CloseInvestment(context context.Context, id domain.InvestmentID) (domain.Investment, error) {
	investment, errInFindInvestment := a.FindInvestmentById(context, id)
	if errInFindInvestment != nil {
		panic(errInFindInvestment)
	}
	investment.IsClosed = true
	updatedInvestment, errInUpdate := a.AddInvestment(context, investment)
	if errInUpdate != nil {
		panic(errInUpdate)
	}
	return updatedInvestment, nil
}

func (a InMemoryInvestmentRepo) FillInvestmentInMemoryDatabase(investments []domain.Investment) {
	txn := a.db.Txn(true)
	for _, investment := range investments {
		if err := txn.Insert(a.investmentTableName, investment); err != nil {
			panic(err)
		}
	}
	txn.Commit()
}
