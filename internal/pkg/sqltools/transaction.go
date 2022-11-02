package sqltools

import (
	"context"
	"database/sql"
)

func RunTx(ctx context.Context, options *sql.TxOptions, db *sql.DB, action func(tx *sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, options)
	if err != nil {
		return err
	}
	defer func() {
		err = tx.Rollback()
		if err != nil {
			return
		}
	}()

	err = action(tx)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
