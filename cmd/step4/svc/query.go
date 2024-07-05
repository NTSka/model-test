package svc

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"log"
	"test-model/pkg/proto/event"
	"time"
)

const query = `INSERT INTO test (date, assets, event_host, event_type,
                  action, importance, object, meta1, meta2,
                  meta3, meta4, meta5)`

func (t *svc) Insert(ctx context.Context, data []*event.EventStep3) error {
	return t.clickhouse.Insert(ctx, func(ctx context.Context, tx *sql.Tx) error {
		stmt, err := tx.PrepareContext(ctx, query)
		if err != nil {
			return err
		}

		defer func() {
			if errS := stmt.Close(); errS != nil {
				log.Print(errors.Wrap(err, "stmt.Close"))
			}
		}()

		for _, e := range data {
			_, err = stmt.ExecContext(
				ctx, time.Unix(e.Event.Event.Timestamp, 0),
				e.Event.Event.Event.Assets,
				e.Event.Event.Event.EventSrc.Host,
				e.Event.Event.Event.EventSrc.Type,
				e.Event.Event.Event.Action,
				int8(e.Event.Event.Event.Importance),
				e.Event.Event.Event.Object,
				e.Event.Event.Meta1,
				e.Event.Meta2,
				e.Event.Meta3,
				e.Meta4,
				e.Meta5)

			if err != nil {
				return err
			}
		}

		return nil
	})
}
