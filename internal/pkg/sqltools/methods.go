package sqltools

import (
	"context"
	"database/sql"
)

func GetSimpleAttr(ctx context.Context, tx *sql.Tx, query string, values []interface{}) ([]string, error) {
	res := make([]string, 0)

	rowsAttr, err := tx.QueryContext(ctx, query, values...)
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
