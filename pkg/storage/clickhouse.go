package storage

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/pkg/errors"
)

type Clickhouse interface {
	Open() (err error)
	Close() []error
	GetDB() *sql.DB
	Insert(
		ctx context.Context,
		cb func(ctx context.Context, tx *sql.Tx) error,
	) error
}

type clickhouse struct {
	Params *Params
	DB     *sql.DB
}

func NewClickhouse(p *Params) (*clickhouse, error) {
	ch := &clickhouse{
		Params: p,
	}

	if err := ch.Open(); err != nil {
		return nil, errors.Wrap(err, "ch.Open")
	}

	return ch, nil
}

func (t *clickhouse) Open() (err error) {
	t.DB, err = sql.Open("clickhouse", t.Params.Dsn)
	if err != nil {
		return
	}

	err = t.DB.Ping()

	return errors.Wrap(err, "DB.Ping")
}

func (t *clickhouse) Close() []error {
	if err := t.DB.Close(); err != nil {
		return []error{err}
	}

	return nil
}

func (t *clickhouse) GetDB() *sql.DB {
	return t.DB
}

func (t *clickhouse) Insert(
	ctx context.Context,
	cb func(ctx context.Context, tx *sql.Tx) error,
) error {
	tx, err := t.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err = cb(ctx, tx); err != nil {
		if errR := tx.Rollback(); errR != nil {
			log.Println(errors.Wrap(err, "tx.Rollback"))
		}
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}
