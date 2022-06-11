package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// Execute a generic tx defined by the function fn.
// There are typically multiple CRUD ops inside fn.
func (store *Store) ExecTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	q := New(tx) // 要用tx来创建新的Queries对象，因为接下来的操作都在这个tx里面
	err = fn(q)
	if err != nil {
		rberr := tx.Rollback()
		if rberr != nil {
			return fmt.Errorf("txerr: %v, rberr: %v", err, rberr)
		}
		return err
	}
	err = tx.Commit()
	return err
}
