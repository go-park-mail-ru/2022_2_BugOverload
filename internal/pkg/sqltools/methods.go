package sqltools

import (
	"context"
	"database/sql"
	stdErrors "github.com/pkg/errors"
	"go-park-mail-ru/2022_2_BugOverload/internal/pkg/errors"
)

func GetSimpleAttr(ctx context.Context, conn *sql.Conn, query string, args ...any) ([]string, error) {
	res := make([]string, 0)

	rowsAttr, err := conn.QueryContext(ctx, query, args...)
	if err != nil {
		return []string{}, err
	}
	defer rowsAttr.Close()

	for rowsAttr.Next() {
		var value sql.NullString

		err = rowsAttr.Scan(&value)
		if err != nil {
			return []string{}, err
		}

		res = append(res, value.String)
	}

	return res, nil
}

func GetSimpleAttrOnConn(ctx context.Context, db *sql.DB, query string, args ...any) ([]string, error) {
	res := make([]string, 0)

	err := RunQuery(ctx, db, func(ctx context.Context, conn *sql.Conn) error {
		var err error

		res, err = GetSimpleAttr(ctx, conn, query, args...)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil && !stdErrors.Is(err, sql.ErrNoRows) {
		return []string{}, errors.ErrPostgresRequest
	}

	return res, nil
}
