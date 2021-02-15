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

func (_ InMemoryInvestmentRepo) FillInvestmentInMemoryDatabase(db *memdb.MemDB, investments []domain.Investment) {
	txn := db.Txn(true)
	for _, investment := range investments {
		if err := txn.Insert("investments", investment); err != nil {
			panic(err)
		}
	}
	txn.Commit()
}
