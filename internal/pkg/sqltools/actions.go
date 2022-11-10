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

func RunTxOnConn(ctx context.Context, options *sql.TxOptions, db *sql.DB, action func(ctx context.Context, tx *sql.Tx) error) error {
	conn, _ := db.Conn(ctx)
	defer conn.Close()

	tx, err := conn.BeginTx(ctx, options)
	if err != nil {
		return err
	}
	defer func() {
		err = tx.Rollback()
		if err != nil {
			return
		}
	}()

	err = action(ctx, tx)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func RunQuery(ctx context.Context, db *sql.DB, action func(ctx context.Context, conn *sql.Conn) error) error {
	conn, _ := db.Conn(ctx)
	defer conn.Close()

	err := action(ctx, conn)
	if err != nil {
		return err
	}

	return nil
}
