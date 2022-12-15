package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// InMemory...
type Postgres struct {
	db *sqlx.DB
}

// InitDb...
func InitDb(psqlConfig string) (*Postgres, error) {
	var err error
	tempDB, err := sqlx.Connect("postgres", psqlConfig)
	if err != nil {
		return nil, err
	}

	return &Postgres{
		db: tempDB,
	}, nil
}
