package dataresource

import (
	"STACK-GO/adapter"
	"context"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

// Implement
func New(dialect, dsn string, idleConn, maxConn int) (adapter.RepoAdapter, error) {
	db, err := sqlx.Open(dialect, dsn)
	if err != nil {
		fmt.Printf("Database Error - Connection - Cannot Create Connection to Database Err: %s", err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("Database Error - Ping Connection - Bad Ping to Database Err: %s", err)
	}

	db.SetMaxIdleConns(idleConn)
	db.SetMaxOpenConns(maxConn)

	return &Repository{db}, nil
}

// func (r *Repository) Close() {
// 	r.db.Close()
// }

func (r *Repository) FindByID(id string, data interface{}, query string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.db.GetContext(ctx, data, query, id)
	// defer r.Close()

	if err != nil {
		fmt.Printf("Database Error - FIND BY ID - Cannot Pull Data. Err: %s", err)
		return data, err
	}

	return data, nil
}

func (r *Repository) List(data interface{}, query string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.db.SelectContext(ctx, data, query)
	if err != nil {
		fmt.Printf("Database Error - List - Cannot Pull All Data. Err: %s", err)
		return nil, err
	}

	return data, nil
}

func (r *Repository) Create(data interface{}, query string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row, err := r.db.NamedExecContext(ctx, query, data)
	if err != nil {
		fmt.Printf("Database Error - Create - Cannot Insert Data. Err: %s", err)
		return 0, err
	}

	return row.LastInsertId()
}

func (r *Repository) Update(data interface{}, id, query string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row, err := r.db.NamedExecContext(ctx, query, data)
	if err != nil {
		fmt.Printf("Database Error - Update - Cannot Update Data. Err: %s", err)
		return 0, err
	}

	return row.RowsAffected()
}

func (r *Repository) Delete(id, query string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	row, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		fmt.Printf("Database Error - Delete - Cannot Delete Data. Err: %s", err)
		return 0, err
	}

	return row.RowsAffected()
}

func (r *Repository) Auth(uname, pass, query string, data interface{}) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.db.GetContext(ctx, data, query, uname, pass)
	// defer r.Close()

	if err != nil {
		fmt.Printf("Database Error - FIND BY ID - Cannot Pull Data. Err: %s", err)
		return nil, err
	}

	return data, nil
}
