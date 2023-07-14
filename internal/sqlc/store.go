package internal

import "github.com/jackc/pgx/v5/pgxpool"

type TxStore interface {
	Querier
}

type Store struct {
	*Queries
	pool *pgxpool.Pool
}

func NewStore(pool *pgxpool.Pool) TxStore {
	return &Store{
		pool:    pool,
		Queries: New(pool),
	}
}
