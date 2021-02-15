package infra

import (
	"testing"
	"time"

	"github.com/ShreyanGoswami/interest-calculator/domain"
	"github.com/hashicorp/go-memdb"
)

func TestMemDB_AddSchema(t *testing.T) {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"investment": &memdb.TableSchema{
				Name: "investment",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "InvestmentID"},
					},
				},
			},
		},
	}
	repo, _ := NewInMemoryInvestmentRepo(schema)

	investments := []domain.Investment{
		{
			InvestmentID:        "id-1",
			StartDate:           time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC),
			EndDate:             time.Date(2021, time.July, 1, 0, 0, 0, 0, time.UTC),
			PrincipalAmount:     1000,
			CurrentAmount:       1000,
			InterestRatePerDay:  0.55,
			InterestRateOverall: 100,
		},
	}

	tests := []struct {
		name      string
		db        *InMemoryInvestmentRepo
		data      []domain.Investment
		tableName string
	}{
		{
			name:      "Successfully added data to database",
			db:        repo,
			data:      investments,
			tableName: "investment",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.db.FillInvestmentInMemoryDatabase(tc.tableName, tc.data)
			// Query database
			txn := tc.db.db.Txn(false)
			defer txn.Abort()
			raw, err := txn.First("investment", "id", "id-1")
			if err != nil {
				panic(err)
			}
			investment := raw.(domain.Investment)
			if !validateInvestment(investment, investments[0]) {
				t.Errorf("[TestCase '%s'] Expected %v| Result %v",
					tc.name,
					investments[0],
					investment)
			}
		})
	}
}

func validateInvestment(actualInvestment, expectedInvestement domain.Investment) bool {
	if actualInvestment.InvestmentID != expectedInvestement.InvestmentID {
		return false
	}
	// TODO add more assertions
	return true
}
