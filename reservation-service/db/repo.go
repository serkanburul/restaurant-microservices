package db

import "github.com/jackc/pgx/v4/pgxpool"

type DB struct {
	dbPool *pgxpool.Pool
	*Queries
}

func NewDB(dbPool *pgxpool.Pool) *DB {
	return &DB{
		dbPool:  dbPool,
		Queries: New(dbPool),
	}
}
