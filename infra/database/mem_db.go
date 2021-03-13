package database

import (
	"github.com/ShreyanGoswami/interest-calculator/domain"
	"github.com/hashicorp/go-memdb"
)

type InMemoryMemDB struct {
	db                  *memdb.MemDB
	investmentTableName string
}

const idFieldName = "id"

func NewInMemoryDatabase() (*InMemoryMemDB, error) {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"investment": {
				Name: "investment",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "InvestmentID"},
					},
				},
			},
		},
	}
	db, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}
	tableName := "investment"
	return &InMemoryMemDB{db, tableName}, nil
}

func (a InMemoryMemDB) InsertData(tableName string, data interface{}) error {
	txn := a.db.Txn(true)
	defer txn.Abort()
	if err := txn.Insert(tableName, data); err != nil {
		return err
	}
	txn.Commit()
	return nil
}

func (a InMemoryMemDB) FindOne(tableName string, query interface{}) (domain.Investment, error) {
	txn := a.db.Txn(false)
	defer txn.Abort()
	raw, err := txn.First(tableName, idFieldName, query)
	if err != nil {
		return domain.Investment{}, err
	}
	return raw.(domain.Investment), nil
}
