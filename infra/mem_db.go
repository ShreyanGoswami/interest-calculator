package infra

import (
	"github.com/ShreyanGoswami/interest-calculator/domain"
	"github.com/hashicorp/go-memdb"
)

type InMemoryInvestmentRepo struct {
	db *memdb.MemDB
}

func NewInMemoryInvestmentRepo(schema *memdb.DBSchema) (*InMemoryInvestmentRepo, error) {
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}
	return &InMemoryInvestmentRepo{db}, nil
}

func (a InMemoryInvestmentRepo) FillInvestmentInMemoryDatabase(tableName string, investments []domain.Investment) {
	txn := a.db.Txn(true)
	for _, investment := range investments {
		if err := txn.Insert(tableName, investment); err != nil {
			panic(err)
		}
	}
	txn.Commit()
}
